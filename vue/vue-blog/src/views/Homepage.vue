<template >
    <div class="conainer">
        <div class="nav">
            <div @click="homepage">首页</div>
            <n-popselect @update:value="serachcate" v-model:value="selectedCategory" :options="cateorytyoptions"
                trigger="click">
                <div>分类<span>{{ categyryName }}</span></div>
            </n-popselect>
            <div @click="dashborad">后台</div>
        </div>
        <n-divider />
        <div class="post">
            <n-space class="serach">
                <n-input v-model:value="pageinfo.keyword" :style="{ width: `500px` }" placeholder="请输入关键字" />
                <n-button type="primary" ghost @click="loadBlogs(0)">搜索</n-button>
            </n-space>

            <div v-for="(blog, index) in blogListInfo" style="margin-bottom: 15px;cursor: pointer;">
                <n-card :title="blog.title" @click="toDetail(blog)">
                    {{ blog.content }}
                    <template #footer>
                        <n-space align="center">
                            <div>
                                发布时间：{{ blog.createtime }}
                            </div>
                        </n-space>
                    </template>
                </n-card>
            </div>

            <n-pagination @update:page="loadBlogs" v-model:page="pageinfo.p" :page-count="pageinfo.pagetotal" />
        </div>
        <n-divider />
        <div class="footer">
            <div>Power by xxx</div>
            <div>XICP备xxxxx号-1</div>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, inject, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'


const router = useRouter()
const route = useRoute()


const message = inject("message")
const axios = inject("axios")
const dialog = inject("dialog")

const selectedCategory = ref(0)
const cateorytyoptions = ref([])
const blogListInfo = ref([])

const pageinfo = reactive({
    p: 1,
    pnum: 10,
    pagetotal: 0,
    total: 0,
    keyword: "",
    categoryid:0,
})


onMounted(() => {
    loadCategorys();
    loadBlogs();
})

const categyryName = computed(() => {
    //寻找符合的条件列
    let selectOption = cateorytyoptions.value.find((option) => { return option.value == selectedCategory.value })
    return selectOption ? selectOption.label : ""
})


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

//文章列表读取
const loadBlogs = async (page = 0) => {
    if (page != 0) {
        pageinfo.p = page
    }
    let res = await axios.get(`/bk/bok?keyword=${pageinfo.keyword}&p=${pageinfo.p}&pnum=${pageinfo.pnum}&category=${pageinfo.categoryid}`)
    let temp = res.data.date
    for (let date of temp) {
        date.content += "..."
    }
    blogListInfo.value = temp
    pageinfo.total = res.data.total
    pageinfo.pagetotal = parseInt(pageinfo.total / pageinfo.pnum) + (pageinfo.total % pageinfo.pnum > 0 ? 1 : 0)
    console.log(res)
}

const serachcate = (categoryid) => {
    pageinfo.categoryid = categoryid
    loadBlogs()
}

const toDetail = (blog) => {
    router.push({ path: "/detail", query: { id: blog.id } })
}

const homepage = () => {
    router.push("/")
}

const dashborad = () => {
    router.push("/login")
}


</script>

<style lang="scss" scoped>
.serach {
    margin-bottom: 15px;
}

.container {

    widows: 1200px;
    margin: 20 auto;
}

.post {
    margin-left: 6%;
}

.nav {
    display: flex;
    font-size: 20px;
    padding-top: 40px;
    color: #64676a;
    margin-left: 10%;

    div {
        cursor: pointer;
        margin-right: 20px;

        &:hover {
            color: #f60;
        }
    }

    span {
        font-size: 13px;
    }
}

.footer {
    text-align: center;
    line-height: 25px;
    color: #64676a;
}</style>