import { RouteRecordRaw } from 'vue-router'
import api from './api/routes'
import home from './home/routes'

const routes: Array<RouteRecordRaw> = [
    ...api,
    ...home,
]

export default routes