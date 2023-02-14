import './style.css'
import { createApp } from 'vue'
import User from '../apps/user.vue'
import router from '../routers/user'

createApp(User)
  .use(router)
  .mount('#user')