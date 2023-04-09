import './style.css'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
// http
import http from './http/index'
import { AxiosKey } from './http/symbols'

const app = createApp(App)

app.use(router)

app.provide(AxiosKey, http)

app.mount('#app')