// vue
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from './router'
// varlet
import Varlet from '@varlet/ui'

// local
import App from './App.vue'

// varlet
import '@varlet/ui/es/style'
import './assets/css/varlet.css'
import '@varlet/touch-emulator'
// local
import './assets/css/main.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Varlet)

app.mount('#app')
