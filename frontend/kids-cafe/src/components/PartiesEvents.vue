<template>
  <div id="app">
    <NavBar/>
    <swiper
      ref="mySwiper"
      :spaceBetween="30"
      :centeredSlides="true"
      :modules="modules"
      :autoplay="{
        delay: 2500,
        disableOnInteraction: false
      }"
      :pagination="{
        clickable: true
      }"
      :navigation="true"
      class="mySwiper"
    >
      <swiper-slide v-for="(img, index) in images" :key="index">
        <img :src="img" alt="gallery" class="w-full h-auto" />
      </swiper-slide>
    </swiper>
    <div class="hiring">
      <img src="@/assets/partyroom.png">
    </div>
    <div class="hiring">
      <img src="@/assets/private.png">
    </div>
  </div>
</template>

<script>
import NavBar from './NavBar.vue';
import { Swiper, SwiperSlide } from 'swiper/vue'; // 이 부분만 수정
import { Autoplay, Pagination, Navigation } from 'swiper/modules';
import 'swiper/swiper-bundle.css';  // Swiper 스타일 추가

export default {
  name: 'App',
  components: {
    NavBar,
    Swiper,
    SwiperSlide,
  },
  data() {
    return {
      images: [],
    };
  },
    // Swiper 옵션은 템플릿 내에서 처리하므로 mounted에서 new Swiper()를 호출하지 않음
    setup() {
    return {
      modules: [Autoplay, Pagination, Navigation],
    };
  },
  mounted() {

    const branchID = this.$route.params.branchID;
    console.log("branchID ****** = ", branchID)

    if (branchID === 'burwood') {
      this.images = [
        '/images/burwood/IMG-20250415-WA0079.jpg',
        '/images/burwood/IMG-20250415-WA0093.jpg',
        '/images/burwood/IMG-20250415-WA0085.jpg',
        '/images/burwood/IMG-20250415-WA0080.jpg',
      ];
    } else if (branchID === 'hornsby') {
      this.images = [
        '/images/hornsby/IMG-20250415-WA0005.jpg',
        '/images/hornsby/IMG-20250415-WA0023.jpg',
        '/images/hornsby/IMG-20250415-WA0026.jpg',
        '/images/hornsby/IMG-20250415-WA0002.jpg',
        '/images/hornsby/IMG-20250415-WA0005.jpg',
        '/images/hornsby/IMG-20250415-WA0040.jpg',
      ];
    } 

    this.$nextTick(() => {
      // Swiper 인스턴스를 가져와서 출력
      console.log("왜 사진 안넘어가냐고;; == ", this.$refs.mySwiper.swiper);
    });
    // 만약 swiper 인스턴스가 제대로 작동하지 않으면 확인을 위해 출력해볼 수 있음
    //console.log("왜 사진 안넘어가냐고;; == ", this.$refs.mySwiper.swiper);
  },
};
</script>

<style>
/* 스타일 설정 */
#app {
  height: 100%;
}

html,
body {
  position: relative;
  height: 100%;
}

.hiring img {
  width: 100%;
  height: auto;
  display: block;
  object-fit: cover;
}

.swiper {
  width: 100%;
  height: 100%;
}

.swiper-slide {
  width: 100%;
  height: 800px; /* 원하는 높이 */
  text-align: center;
  font-size: 18px;
  background: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
}

.swiper-slide img {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover; /* 이미지가 잘리지 않게 비율 맞춰 자르기 */
}

</style>