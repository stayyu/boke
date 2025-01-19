<template >
    <div class="container">
        <n-button @click="back">返回</n-button>
        <!-- 标题 -->
        <n-h1>{{ bloginfo.title }}</n-h1>
        <!-- 文章内容 -->
        <div class="blog-content">     
              <div v-html="bloginfo.content"></div>
            </div>
 
    </div>
</template>

<script setup>
import { ref, reactive, inject, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'


const router = useRouter()
const route = useRoute()
const bloginfo=ref({})
const axios = inject("axios")

onMounted(()=>{
    loadBlog()
})

const loadBlog= async()=>{
	console.log(route.query.id)
    let res = await axios.get(`/bk/bok?p=&id=${route.query.id}`)
    bloginfo.value=res.data
}


const back=()=>{
    router.push("/")
}


</script>

<style>
.blog-content img{
    max-width: 100%!important;
}
</style>
<style lang="scss" scoped>
.container {
    width: 1200px;
    margin: 0 auto;
}

</style>