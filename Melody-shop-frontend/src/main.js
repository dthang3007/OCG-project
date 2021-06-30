import { createApp } from 'vue'
import App from './App.vue'
import routes from "./routes"
import store from './store'
import moshaToast from 'mosha-vue-toastify'
import 'mosha-vue-toastify/dist/style.css'

const app=createApp(App)
app.use(routes)
app.use(store)
app.use(moshaToast)

app.mount('#app')
