// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
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

const routes = [
  { path: '/', component: Home },
  // { path: '/book_a_party', component: Booking },
  { path: '/book_a_party', component: SimplyBook },
  { path: '/book_a_party/quickbook', component: QuickBook },
  { path: '/book_a_party/confirm_detail', component: ConfirmDetail },
  { path: '/book_a_party/select_room/:branchID', component: Partyroom },
  { path: '/book_a_party/select_branch', component: SelectBranch },
  { path: '/book_a_party/select-time/:roomID', component: SelectTIme },
  { path: '/book_a_party/foodopitions', component: Foodoptions },
  { path: '/book_a_party/client-info', component: ClientInfo },
  { path: '/book_a_party/booking-confirm', component: BookingConfirm },
  { path: '/faq', component: FAQ },
  { path: '/menu', component: Menu },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
