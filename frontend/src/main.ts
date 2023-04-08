import './style.css'
import { createApp } from 'vue'
import { globalVariables, GlobalVariables } from './globals/variables'
import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(router)

app.provide('globalVariables', globalVariables as GlobalVariables)

app.mount('#app')