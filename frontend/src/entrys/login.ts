import './style.css'
import { createApp } from 'vue'
import Login from '../apps/login.vue'
import router from '../routers/login'

createApp(Login)
  .use(router)
  .mount('#login')