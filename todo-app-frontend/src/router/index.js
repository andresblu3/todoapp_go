import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../views/HomePage.vue'

const routes = [
    {
        path: '/',
        name: 'HomePage',
        component: HomePage
    },
]

const router = createRouter({
    history: createWebHistory(process.env.VUE_APP_BASE_URL),
    routes
})

export default router