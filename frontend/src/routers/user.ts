import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '../components/Dashboard.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
]

const router = createRouter({
  history: createWebHistory("/usuario"),
  routes
})

export default router