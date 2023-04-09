import { RouteRecordRaw } from 'vue-router'

import User from "./pages/User.vue"

const api: Array<RouteRecordRaw> = [
    {
      path: '/usuario',
      name: 'Usuario',
      component: User
    }
]
  
export default api