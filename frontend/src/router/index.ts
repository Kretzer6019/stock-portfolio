import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

import pageRoutes from '../components/router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/login'
  },
  ...pageRoutes
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router