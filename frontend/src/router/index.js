import { createRouter, createWebHistory } from 'vue-router'
import Register from '../views/Register.vue' // 假设你的注册页面组件位于 src/views/Register.vue

const routes = [
    {
        path: '/register',
        name: 'Register',
        component: Register
    },
    // 其他路由...
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
})

export default router