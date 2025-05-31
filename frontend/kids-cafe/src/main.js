import { createApp } from 'vue'
//import { createPinia } from "pinia";  // ✅ Pinia 가져오기
import './style.css'
import App from './App.vue'
import router from './router'  // router 설정을 import
import '@/assets/css/global.css';  // 글로벌 CSS 연결
import PrimeVue from 'primevue/config';  // PrimeVue 설정 import
import Tree from 'primevue/tree';        // Tree 컴포넌트 import
import '@fortawesome/fontawesome-free/css/all.min.css';
import { auth } from './firebase' // Firebase 초기화

//const pinia = createPinia();  // ✅ Pinia 인스턴스 생성

let app;
auth.onAuthStateChanged(() => {
  if (!app) {
    app = createApp(App);
    app.use(router);
    app.mount('#app');
  }
});

/*
createApp(App)
.use(router)  // router를 Vue 앱에 연결
// .use(pinia)
.component('Tree', Tree)
.use(PrimeVue)
.mount('#app')

*/