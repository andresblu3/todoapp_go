import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import vuetify from './plugins/vuetify'
import 'animate.css'
import '@mdi/font/css/materialdesignicons.css' // Importar los iconos de Material Design

const app = createApp(App)

app.config.globalProperties.$http = axios
app.config.globalProperties.$http.defaults.baseURL = 'http://localhost:8000'

app.use(router)
app.use(vuetify)

app.mount('#app')