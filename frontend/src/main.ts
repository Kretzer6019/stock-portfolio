import './style.css'

import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
import axios from './http/axios'

const app = createApp(App)
app.use(router)
app.use(axios, {
    baseUrl: "http://localhost:8080",
})
app.mount('#app')