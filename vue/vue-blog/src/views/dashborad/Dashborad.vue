<template>
    <div class="main-panel">
        <div class="menus">
            <div v-for="(menu, index) in menus" @click="toPage(menu)">
                {{ menu.name }}
            </div>
        </div>
            <div style="padding:20px;width:100%;">
                <router-view></router-view>
            </div>
        </div>
  
    <div class="title">后台管理系统</div>
</template>

<script setup>
import { Adminstore } from '../../stores/Adminstore';
import { ref, reactive, inject } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()


const message = inject("message")
const axios = inject("axios")

const adminstore = Adminstore()
 
let menus = [
    { name: "文章管理", href: "/dashborad/article" },
    { name: "分类管理", href: "/dashboard/category" },
    { name: "退出", href: "logout" },
];

const toPage = (menu) => {
    if (menu.href=='logout') {
        router.push("/login")
    } else {
        router.push(menu.href)
    }
}

</script>

<style lang="scss" scoped>
.main-panel {
    display: flex;
    color: dimgray;
    max-width: 1500px;
    margin: 0 auto;
}

.menus {
    padding: 20px 0;
    box-sizing: border-box;
    line-height: 55px;
    text-align: center;
    width: 180px;
    height: 95vh;
    border-right: 1px solid #dadada;

    div {
        cursor: pointer;

        &:hover {
            color: #fd760f;
        }
    }
}

.title{
    font-size: 65px;
    font-weight: bold;
    text-align: right;
    position: fixed;
    color: black;
    right:calc((100vw - 1500px)/2);
    bottom:20px;
}
</style>
