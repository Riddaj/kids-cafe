<template>
    <div id="app" style="background-color: ivory;">
        <NavBar/>
        <!-- Swiper + 텍스트 래퍼 -->
      <div class="swiper-wrapper-container">
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
        <!-- ✅ 텍스트 오버레이 -->
        <div class="swiper-overlay-text">
          <h1>Welcome to Twinkle Kids Cafe!</h1>
          <p style="color: white;">Let’s play, learn, and celebrate!</p>
        </div>
      </div>
        <div>
            <section class="slider">
                    <div class="main-text">
                        <div>
                            <img src="@/assets/toys.png" alt="toys" class="toys" />
                        </div>
                        <h1>Welcome to <br>
                            {{ branchName }}<br>
                            Twinkle <br>
                            Kids Cafe<br>
                        </h1>
                    </div>
                    <div class="main-text">
                        <p>Welcome to the center where your children are able <br> 
                            to play, learn, excercise and have the best Birthday 
                            <br>parties!&nbsp;</p>
                    </div>
                    <!-- 퀵 예약 버튼 -->
                    <!-- <div class="quickbtn-container">
                            <a class="quickbtn-book" href="/book_a_party/quickbook" title="Quick Party Room Booking">
                                Quick Party Room Booking
                            </a>
                    </div> -->
            </section>
            <section>
                <div class="image-wrapper">
                  <img :src="getBranchImage(branchID)" class="background-image1" />
                  <!-- 
                    <img src="https://images.squarespace-cdn.com/content/v1/637d8d8a7f609c521ddd5429/1672359448650-N89Q21OUSYRU8ROW18F1/Burwood+Plaza+Max3MB_72DPI_VCLAMedia+%2854+of+101%29.jpg" 
                    class="background-image1">
                  -->
                    </div>
            </section>
        </div>
        <Footer/>
    </div>
</template>
<script>
import axios from 'axios'; // axios를 import 추가
import NavBar from '@/components/NavBar.vue';
import Footer from '@/components/Footer.vue';

import { Swiper, SwiperSlide } from 'swiper/vue'; // 이 부분만 수정
import { Autoplay, Pagination, Navigation } from 'swiper/modules';
import 'swiper/swiper-bundle.css';  // Swiper 스타일 추가


export default {
    components:{
        NavBar,
        Footer,
        Swiper,
        SwiperSlide,
    },
    data() {
    return {
      branchID:this.$route.params.branchID,
      branchName: '',
      images: [],
    }
  },
  mounted(){

    const branchID = this.$route.params.branchID;
    console.log("branchID ****** = ", branchID)

    if(branchID === 'burwood'){
      this.images = [
      '/images/main/20250321_094707.jpg',
      '/images/main/20250321_094810.jpg',
      '/images/burwood/mmexport1746090489101.jpg',
      '/images/main/20250321_094831.jpg',
      '/images/main/20250321_094733.jpg',
      '/images/burwood/IMG-20250415-WA0085.jpg',
    ];
    }else if(branchID === 'hornsby'){
      this.images = [
        '/images/hornsby/IMG-20250430-WA0000.jpg',
        '/images/hornsby/IMG-20250430-WA0001.jpg',
        '/images/hornsby/IMG-20250430-WA0008.jpg',
        '/images/hornsby/IMG-20250430-WA0011.jpg',
        '/images/hornsby/IMG-20250415-WA0005.jpg',
        '/images/hornsby/IMG-20250415-WA0023.jpg',
        '/images/hornsby/IMG-20250415-WA0026.jpg',
        '/images/hornsby/IMG-20250415-WA0002.jpg',
        '/images/hornsby/IMG-20250430-WA0003.jpg',
        '/images/hornsby/IMG-20250430-WA0005.jpg',
      ];
    }

    this.$nextTick(() => {
      // Swiper 인스턴스를 가져와서 출력
      console.log("왜 사진 안넘어가냐고;; == ", this.$refs.mySwiper.swiper);
    });
  },
  async created() {
    const branchID = this.$route.params.branchID
    this.branchName = this.getBranchNameByID(branchID)
  },
    // Swiper 옵션은 템플릿 내에서 처리하므로 mounted에서 new Swiper()를 호출하지 않음
  setup() {
    return {
      modules: [Autoplay, Pagination, Navigation],
    };
  },
  methods: {
    getBranchNameByID(id) {
      const branchMap = {
        'burwood': 'Burwood',
        'hornsby': 'Hornsby',
        // 필요하면 계속 추가 가능
      }
      return branchMap[id] || 'Unknown Branch'
    },
    getBranchImage(branchID) {
      const images = {
        burwood: "https://images.squarespace-cdn.com/content/v1/637d8d8a7f609c521ddd5429/1672359448650-N89Q21OUSYRU8ROW18F1/Burwood+Plaza+Max3MB_72DPI_VCLAMedia+%2854+of+101%29.jpg",
        hornsby: "/images/hornsby.jpg"
      };
      return images[branchID]; // fallback
    }
  }
}
</script>

<style>
.swiper-wrapper-container {
  position: relative;
  width: 100%;
  height: 100vh;
  overflow: hidden;
}

.swiper-overlay-text {
  position: absolute;
  top: 85%; /* 세로 기준 중앙 */
  left: 25%; /* 가로 기준 중앙 */
  transform: translate(-50%, -50%);
  z-index: 10;
  color: white;
  text-shadow: 2px 2px 8px rgba(0,0,0,0.6); 
}

.swiper-overlay-text h1 {
  font-size: 3rem;
  font-weight: bold;
  margin: 0;
}

.swiper-overlay-text p {
  font-size: 1.5rem;
  text-align: left;
}

#app {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

.main-text h1,p{
    z-index: 1; /* h1이 버튼 위에 오도록 z-index 설정 */
    text-align: right;
    position: relative; /* z-index가 제대로 작동하도록 위치 지정 */
    color: black;
}

.slider {
  position: relative;
  z-index: 2; /* Swiper보다 위로 올라오지 않도록 2 이하로 설정 */
  padding: 60px 20px; /* 원하는 만큼 패딩 추가 */
  background-color: rgba(255, 255, 255, 0.95); /* 배경 설정 (필요시) */
}

.background-image1{
  width: 100%;
  height: auto;
  position: static; /* 기존 absolute 제거 */
  object-fit: cover;
  opacity: 0.7;
}

.toys{
    width: 100px;
    height: 100px;
}
ul {
  list-style-type: none; /* 동그라미 기호 제거 */
}

li {
  list-style-type: none; /* li에 대한 동그라미 기호 제거 (선택 사항) */
}


.swiper {
    width: 100vw; /* 브라우저 전체 너비 */
    height: 100vh; /* 전체 화면 높이 (또는 적절한 값) */
}

.swiper-slide {
  width: 100%;
  height: 800px; /* 원하는 높이 */
  text-align: center;
  font-size: 18px;
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


.btn-container {
    flex: 1; /* 버튼 영역도 동일한 비율로 배치 */
    display: flex;
    justify-content: left; 
    height: 50px; /* 버튼과 메뉴 항목 높이를 맞추기 위해 동일한 높이 설정 */
    margin-top: 30px;
    /*margin-left: 5px;  메뉴와 버튼 간의 간격 설정 */
}

/** button img */
.btn-book {
    display: flex; /* flexbox로 설정 */
    padding: 20px 20px;
    justify-content: center; /* 수평 중앙 정렬 */
    align-items: center; /* 수직 중앙 정렬 */
    margin-left: 10px; /* 버튼과 메뉴 간에 간격을 줌 */
    background-color: #4d0099; /* 버튼 배경 색상 */
    color: white; /* 텍스트 색상 */
    text-decoration: none; /* 링크 스타일 제거 */
    border-radius: 8px; /* 둥근 모서리 */
    font-size: 16px; /* 글자 크기 */
}

.btn-book:hover{
    background-color: #ff6600; /* 호버 시 색상 변경 */
    color: white; /* 텍스트 색상 */
}

.quickbtn-container {
    position: absolute;
    top: 50%; /* 원하는 위치 조정 */
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: 10; /* 이미지 위로 오도록 설정 */
}

.quickbtn-book {
    display: inline-block;
    padding: 12px 20px;
    background-color: #ffb3b3;
    color: white;
    text-decoration: none;
    font-size: 18px;
    font-weight: bold;
    border-radius: 5px;
    box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.2);
}

.quickbtn-book:hover {
    color: white;
    background-color: #ff8080;
}
</style>
