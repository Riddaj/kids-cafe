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
    <div class="book-a-party-section">
      <p class="party-description">
        Celebrate magical moments at Twinkle Kids Cafe — <br>where little giggles, colorful balloons, and joyful memories fill the air.<br/>
        Create a party your child will remember forever in a space designed for laughter, play, and love.
      </p>
      <div class="book-a-party">
        <a class="btn-book-outline" :href="`/book_a_party/select_room/${this.$route.params.branchID}`" title="Book a party">Book a party</a>
      </div>
    </div>
    <div class="hiring">
      <img :src="getPartyroomImage(branchID)" alt="party room" />
    </div>
    <div class="hiring">
      <img :src="getPrivateRoomImage(branchID)" alt="private room" />
    </div>
    <Footer/>
  </div>
</template>

<script>
import NavBar from './NavBar.vue';
import Footer from './Footer.vue';
import { Swiper, SwiperSlide } from 'swiper/vue'; // 이 부분만 수정
import { Autoplay, Pagination, Navigation } from 'swiper/modules';
import 'swiper/swiper-bundle.css';  // Swiper 스타일 추가

export default {
  name: 'App',
  components: {
    NavBar,
    Footer,
    Swiper,
    SwiperSlide,
  },
  data() {
    return {
      branchID: this.$route.params.branchID,
      images: [],
    };
  },
    // Swiper 옵션은 템플릿 내에서 처리하므로 mounted에서 new Swiper()를 호출하지 않음
    setup() {
    return {
      modules: [Autoplay, Pagination, Navigation],
    };
  },
  methods: {
    getPartyroomImage(branchID) {
      const images = {
        burwood: require('@/assets/B_hiring3.png'),
        hornsby: require('@/assets/H_hiring3.png'),
        // 필요 시 다른 지점 추가
      };
      return images[branchID];
    },
    getPrivateRoomImage(branchID) {
      const images = {
        burwood: require('@/assets/B_private.png'),
        hornsby: require('@/assets/H_private.png'),
        // 필요 시 다른 지점 추가
      };
      return images[branchID];
    },
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

.book-a-party-section {
  background-color: #f6f2ef; /* 이미지 배경과 유사한 톤 */
  padding: 60px 20px; /* 넉넉한 여백 */
  text-align: center;
}

.btn-book-outline {
  border: 1px solid black;
  background: transparent;
  color: black;
  text-transform: uppercase;
  font-weight: 500;
  letter-spacing: 2px;
  padding: 12px 28px;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.3s ease, color 0.3s ease;
  text-decoration: none;
  display: inline-block;
}

.btn-book-outline:hover {
  background-color: black;
  color: white;
}

.party-description {
  max-width: 700px;           /* 👉 문단 폭 제한 */
  margin: 0 auto 40px;        /* 👉 가운데 정렬 */
  font-size: 1.3rem;
  line-height: 1.7;
  color: #222;
  text-align: justify;        /* 👉 양쪽 정렬 느낌 추가 */
  text-align-last: center;    /* ✅ 마지막 줄은 가운데로 (시각적으로 더 자연스러움) */
}

/* ✅ 모바일 대응 */
@media (max-width: 768px) {
  .party-description {
    font-size: 1.1rem;
    line-height: 1.6;
    padding: 0 16px;
    text-align: left; /* 작은 화면에서는 justify가 오히려 불편할 수 있음 */
  }

  .btn-book-outline {
    padding: 10px 20px;
    font-size: 13px;
  }
}
</style>