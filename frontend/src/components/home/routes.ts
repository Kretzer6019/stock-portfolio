import { RouteRecordRaw } from 'vue-router'

import Index from "./pages/Index.vue"

const home: Array<RouteRecordRaw> = [
    {
        path: '/login',
        name: 'Login',
        component: Index
    },
]
  
export default home