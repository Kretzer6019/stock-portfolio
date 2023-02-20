import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Index from '../components/pages/Index.vue'
import User from '../components/pages/User.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/login',
    name: 'Login',
    component: Index
  },
  {
    path: '/usuario',
    name: 'Usuario',
    component: User
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router