import { createApp } from 'vue'
import App from './App.vue'
// import '@/css/reset.css';
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import '@/css/style.css';
import '@/css/common.css';
import '@/assets/js/iconfont'
import router from '@/router/router'
import { createPinia } from 'pinia'


const app = createApp(App)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.use(router).use(createPinia()).mount('#app')