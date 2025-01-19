import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import {router} from './router'
import naive from 'naive-ui'
import axios from 'axios'
import { createDiscreteApi } from 'naive-ui'
import { Adminstore} from './stores/Adminstore'

const app = createApp(App)


axios.defaults.baseURL="http://127.0.0.1:8081"
const { message, notification, dialog,}=createDiscreteApi(["message", "dialog", "notification"])



app.provide("axios",axios)
app.provide("message",message)
app.provide("notification",notification)
app.provide("dialog",dialog)
app.provide("server_url",axios.defaults.baseURL)

app.use(createPinia())
const adminstore=Adminstore()

//拦截器
axios.interceptors.request.use((config)=>{
    config.headers.token=adminstore.token
    return config
})




app.use(router)
app.use(naive)
app.mount('#app')
