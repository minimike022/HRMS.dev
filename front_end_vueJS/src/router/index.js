import {createRouter, createWebHistory}from 'vue-router'
import Login from "/src/pages/Login.vue"
const Application_Forms = () => import("/src/pages/Application_Forms.vue")

const routes = [
    {
        path:"/user/login",
        component: Login
    },
    {
        path:"/application_forms",
        component: Application_Forms
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router;