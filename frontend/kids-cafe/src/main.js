import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'  // router 설정을 import
import '@/assets/css/global.css';  // 글로벌 CSS 연결\


createApp(App)
    .use(router)  // router를 Vue 앱에 연결
    .mount('#app')
