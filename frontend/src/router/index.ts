import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Login from '../components/auth/SignIn.vue'
import Home from '../components/Dashboard.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router