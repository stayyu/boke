<template>
  <div>
    <n-button @click="showADDModal = true">添加</n-button>
    <n-table :bordered="false" :single-line="false">
      <thead>
        <tr>
          <th>编号</th>
          <th>名称</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(category, index) in categoryList">
          <td>{{ category.id }}</td>
          <td>{{ category.name }}</td>
          <td>
            <n-space>
              <n-button @click="toUpdata(category)">修改</n-button>
              <n-button @click="deleteCatecory(category)">删除</n-button>
            </n-space>
          </td>
        </tr>
      </tbody>
    </n-table>
    <n-modal v-model:show="showADDModal" preset="dialog" title="Dialog">
      <template #header>
        <div>添加分类</div>
      </template>
      <div>
        <n-input v-model:value="addCategory.id" type="text" placeholder="请输入编号"></n-input>
        <n-input v-model:value="addCategory.name" type="text" placeholder="请输入名称"></n-input>
      </div>
      <template #action>
        <div>
          <n-button @click="add">提交</n-button>
        </div>
      </template>
    </n-modal>
    <n-modal v-model:show="showUpadteModal" preset="dialog" title="Dialog">
      <template #header>
        <div>修改分类</div>
      </template>
      <div>
        <n-input v-model:value="updataCategory.name" type="text" placeholder="请输入名称"></n-input>
      </div>
      <template #action>
        <div>
          <n-button @click="updata">提交</n-button>
        </div>
      </template>
    </n-modal>
  </div>
</template>

<script setup>
import { Adminstore } from '../../stores/Adminstore';
import { ref, reactive, inject, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()


const message = inject("message")
const dialog = inject("dialog")
const axios = inject("axios")

const adminstore = Adminstore()

const showADDModal = ref(false)
const showUpadteModal = ref(false)
const categoryList = ref([{}])
const addCategory = reactive({
  name: "",
  id:""
})
const updataCategory = reactive({
  id: 0,
  name: ""
})

onMounted(() => {
  loadDatas()
})

const loadDatas = async () => {
  let res = await axios.get("/bk/cate")
  categoryList.value=res.data.date
}

//添加分类
const add = async () => {
  let res = await axios.post("/bk/cate/create",
  { name: addCategory.name ,
      id:addCategory.id})
  if (res.status == 200) {
    loadDatas()
    message.info("修改成功")
  } else {
    message.error("修改失败");
  }
  showADDModal.value = false;
}
//修改分类
const toUpdata = async (category) => {
  showUpadteModal.value = true
  showUpadteModal.id = category.id
  showUpadteModal.name = category.name
}

const updata = async () => {
  let res = await axios.put("/bk/cate/update",{
     name:updataCategory.name,
     id:showUpadteModal.id
    })
  console.log(res)
  if (res.status== 200) {
    loadDatas()
    message.info(res.data.msg)
  } else {
    message.error(res.data.msg);
  }
  showUpadteModal.value = false;
}



const deleteCatecory = async (category) => {
  dialog.warning({
    title: '警告',
    content: '是否要删除',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: async () => {
      console.log(category.id)
      let res = await axios.delete(`/bk/cate/del/${category.id}`)
      if (res.status == 200) {
        loadDatas()
        message.info("删除成功")
      } else {
        message.error("删除失败");
      }
    },
    onNegativeClick: () => { }
  })



}




</script>

<style lang="scss" scoped></style>