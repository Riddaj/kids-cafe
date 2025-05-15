// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import PickBranch from '../views/PickBranch.vue';
import Booking from '../views/Booking.vue';
import SimplyBook from '../components/SimplyBook.vue';
import ConfirmDetail from '../components/ConfirmDetail.vue'
import Partyroom from '../components/Partyroom.vue';
import SelectBranch from '../components/SelectBranch.vue';
import FAQ from '../components/FAQ.vue';
import SelectTIme from '../components/SelectTIme.vue';
import QuickBook from '../components/QuickBook.vue';
import Foodoptions from '../components/Foodoptions.vue';
import ClientInfo from '../components/ClientInfo.vue';
import BookingConfirm from '../components/BookingConfirm.vue';
import Menu from '../components/Menu.vue';
import AddMenu from '../components/AddMenu.vue';
import Price from '../components/Price.vue';
import PartiesEvents from '../components/PartiesEvents.vue';
import AboutUs from '../components/AboutUs.vue';
import AdminDashboard from '../components/AdminDashboard.vue';
import Entryrules from '../components/Entryrules.vue';
import Contact from '../components/Contact.vue';
import Login from '../components/Login.vue';
import AddParty from '../components/AddParty.vue';



const routes = [
  { path: '/', component: PickBranch },
  { path: '/home/:branchID', component: Home },
  // { path: '/book_a_party', component: Booking },
  { path: '/book_a_party/:branchID', component: SimplyBook },
  { path: '/book_a_party/quickbook', component: QuickBook },
  { path: '/book_a_party/confirm_detail', component: ConfirmDetail },
  { path: '/book_a_party/select_room/:branchID', component: Partyroom },
  { path: '/book_a_party/select_branch', component: SelectBranch },
  { path: '/book_a_party/select-time/:roomID', component: SelectTIme },
  { path: '/book_a_party/foodopitions', component: Foodoptions },
  { path: '/book_a_party/client-info', component: ClientInfo },
  { path: '/book_a_party/booking-confirm', name: 'booking-confirm', component: BookingConfirm },
  { path: '/faq/:branchID', component: FAQ },
  { path: '/entryrules/:branchID', component: Entryrules },
  { path: '/menu/:branchID', name: 'menu', component: Menu },
  { path: '/price/:branchID', component: Price },
  { path: '/admin/menu/:branchID', component: AddMenu },
  { path: '/parties-events/:branchID', component: PartiesEvents },
  { path: '/about-us/:branchID', component: AboutUs },
  { path: '/admin', component: AdminDashboard, name: 'admin'},
  { path: '/contact/:branchID', component: Contact },
  { path: '/login', component: Login },
  { path: '/add_party', component: AddParty },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// 여기에 추가!
router.beforeEach((to, from, next) => {
  const defaultTitle = 'Twinkle Kids Cafe';
  document.title = to.meta.title || defaultTitle;

   // ✅ 2. 어드민 경로 보호
   const requiresAuth = to.path.startsWith("/admin");
   const token = localStorage.getItem("authToken");
 
   if (requiresAuth && !token) {
     next('/login/');
   } else {
     next();
   }

  // next();
});


export default router;
