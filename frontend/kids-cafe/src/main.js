import { createApp } from 'vue'
import { createPinia } from "pinia";  // ✅ Pinia 가져오기
import './style.css'
import App from './App.vue'
import router from './router'  // router 설정을 import
import '@/assets/css/global.css';  // 글로벌 CSS 연결\

//const pinia = createPinia();  // ✅ Pinia 인스턴스 생성

createApp(App)
    .use(router)  // router를 Vue 앱에 연결
   // .use(pinia)
    .mount('#app')
