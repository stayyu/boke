<template>
    <div>
        <n-tabs v-model:value="tabvalue" justify-content="start" type="line">
            <n-tab-pane name="list" tab="文章列表">
                <div v-for="(blog, index) in blogListInfo">
                    <n-card :title="blog.title">
                        {{ blog.content }}
                        <template #footer>
                            <n-space align="center">
                                <div>
                                    发布时间：{{ blog.createtime }}
                                    <br>
                                    <n-button @click="toupdata(blog)">修改</n-button>
                                    <n-button @click="deldata(blog)">删除</n-button>
                                </div>
                            </n-space>
                        </template>
                    </n-card>
                </div>
                <n-space>
                    <div @click="toPage(pagenum)" v-for="pagenum in pageinfo.pagetotal">
                        <div :style="`color:` + (pagenum == pageinfo.p ? `blue` : ``)"> {{ pagenum }}</div>
                    </div>
                </n-space>
            </n-tab-pane>
            <n-tab-pane name="add" tab="添加文章">
                <n-form-item label="标题">
                    <n-input v-model:value="addAtricle.title" placeholder="请输入标题" />
                </n-form-item>
                <n-form-item label="分类">
                    <n-select v-model:value="addAtricle.categoryid" :options="cateorytyoptions" />
                </n-form-item>

                <n-form-item label="内容">
                    <rich-text-editor v-model="addAtricle.content"></rich-text-editor>
                </n-form-item>

                <n-form-item label="">
                    <n-button @click="add">提交</n-button>
                </n-form-item>


            </n-tab-pane>
            <n-tab-pane name="updata" tab="修改">
                <n-form-item label="标题">
                    <n-input v-model:value="updataAtricle.title" placeholder="请输入标题" />
                </n-form-item>
                <n-form-item label="分类">
                    <n-select v-model:value="updataAtricle.categoryid" :options="cateorytyoptions" />
                </n-form-item>

                <n-form-item label="内容">
                    <rich-text-editor v-model="updataAtricle.content"></rich-text-editor>
                </n-form-item>

                <n-form-item label="">
               
                    <n-button @click="updata">提交</n-button>
                </n-form-item>

            </n-tab-pane>
        </n-tabs>
    </div>
</template>

<script setup>
import { Adminstore } from '../../stores/Adminstore';
import { ref, reactive, inject, onMounted, onBeforeUnmount } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import RichTextEditor from '../../components/RichTextEitor.vue'


const router = useRouter()
const route = useRoute()


const message = inject("message")
const axios = inject("axios")
const dialog = inject("dialog")

const adminstore = Adminstore()
const cateorytyoptions = ref([])
const blogListInfo = ref([])
const tabvalue = ref("list")

const addAtricle = reactive({
    categoryid: 0,
    title: "",
    content: "",
	id: ""
})

const updataAtricle = reactive({
    categoryid: 0,
    title: "",
    content: "",
    id: ""
})

const pageinfo = reactive({
    p: 1,
    pnum: 10,
    pagetotal: 0,
    total: 0,
})


onMounted(() => {
    loadCategorys()
    loadBlogs()
})

//文章列表读取
const loadBlogs = async () => {
    let res = await axios.get(`/bk/bok?p=${pageinfo.p}&pnum=${pageinfo.pnum}`)
    let temp = res.data.date
    for (let date of temp) {
        date.content += "..."
    }
    blogListInfo.value = temp
    pageinfo.total = res.data.total
    pageinfo.pagetotal = parseInt(pageinfo.total / pageinfo.pnum) + (pageinfo.total % pageinfo.pnum > 0 ? 1 : 0)
}


//加载分类选项
const loadCategorys = async () => {
    let res = await axios.get("/bk/cate")
    cateorytyoptions.value = res.data.date.map((item) => {
        return {
            label: item.name,
            value: item.id
        }
    })
}

const add = async () => {
    let res = await axios.post("/bk/bok/create", addAtricle)
    if (res.status == 200) {
        message.info("添加成功")
        addAtricle.title = ""
        addAtricle.content = null
    } else {
        message.error("添加失败");
    }
    loadBlogs()
}

const updata = async () => {
    let res = await axios.put("/bk/bok/updata", updataAtricle)
    if (res.status == 200) {
        tabvalue.value = "list"
        loadBlogs()
        message.info("修改成功")
    } else {
        message.error("添加失败");
    }
}

const deldata= async (blog) => {
  dialog.warning({
    title: '警告',
    content: '是否要删除',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: async () => {
      console.log(blog.id)
      let res = await axios.delete(`/bk/bok/del/${blog.id}`)
      if (res.status == 200) {
        loadBlogs()
        message.info("删除成功")
      } else {
        message.error("删除失败");
      }
    },
    onNegativeClick: () => { }
  })



}

const toPage = async (pagenum) => {
    pageinfo.p = pagenum
    loadBlogs()
}

const toupdata = async (blog) => {
    tabvalue.value = "updata"
    let res = await axios.get("/bk/bok/updata?cid" + blog.categoryid)
    console.log(blog.id)
    updataAtricle.id = blog.id
    updataAtricle.title = res.data.title
    updataAtricle.content = res.data.content
    updataAtricle.categoryid = res.data.categoryid
}


</script>

<style lang="scss" scoped></style>