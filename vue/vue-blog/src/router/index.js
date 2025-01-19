import { createRouter, createWebHistory } from 'vue-router'


let routes = [

  { path: "/login", component: () => import("../views/Login.vue") },
  { path: "/", component: () => import("../views/Homepage.vue") },
  { path: "/detail", component: () => import("../views/detail.vue") },
  {
    path: "/dashborad", component: () => import("../views/dashborad/Dashborad.vue"), children: [
      { path: "/dashboard/category", component: () => import("../views/dashborad/Category.vue") },
      { path: "/dashborad/article", component: () => import("../views/dashborad/Article.vue") },
    ]
  },
]


const router = createRouter({
  history: createWebHistory(),
  routes,
})



export { router, routes }
