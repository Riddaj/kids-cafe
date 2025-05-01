<template>
    <div id="app"> 
        <BookingBar/>
        <!-- booking process (2단계) -->
        <BookingProcess/>
        <div class="btn-container">
            <!--<router-link class="btn-book" to="/book_a_party/simplybook/confirm_detail" title="Book Now">Book Now</router-link>-->
            <a class="btn-book" href="#" @click="showBookingInfo = !showBookingInfo" title="Book Now">Book Now</a>
        </div>
        <CurrentTime/>
        <component 
        :is="showBookingInfo ? 'SelectBranch' : 'div'"  
        :branches="branches" 
        v-if="branches.length > 0">
            <!-- Contact Info (예약 정보로 변경 가능) -->
            <template v-if="!showBookingInfo">
                <div class="contact-info">
                    <h1>Contact us</h1>
                    <ul v-if="branches && branches.length > 0">
                        <li v-for="branch in branches" :key="branch.id" class="branch-item">
                            <strong>{{ branch.branch_name }}</strong><br>
                            Email: {{ branch.email }}<br>
                            Phone: {{ branch.phone }}<br>
                            Store Phone Number: {{ branch.branch_call }}
                            <!-- <hr/> -->
                        </li>
                        <br>
                    </ul>
                </div>
            </template>
        </component>
        <!-- ✅ 아래 여백 주기 -->
        <div style="margin-top: 150px;"></div>
        <!-- ✅ Footer 컴포넌트 추가 -->
        <Footer />
    </div>
</template>
<script>
import CurrentTime from '../components/CurrentTime.vue';
import axios from "axios";
import ConfirmDetail from '../components/ConfirmDetail.vue';
import Partyroom from '../components/Partyroom.vue';
import SelectBranch from '../components/SelectBranch.vue';
import BookingProcess from '../components/BookingProcess.vue';
import BookingBar from '../components/BookingBar.vue';
import Footer from './Footer.vue';

export default {
    data() {
        return {
            showBookingInfo : false,
            branches:[]
        };
    },  
    mounted() {
    this.fetchBranches();
    },
    components: {
        CurrentTime,
        ConfirmDetail,
        Partyroom,
        SelectBranch,
        BookingProcess,
        BookingBar,
        Footer
    },  
    methods: {
    async fetchBranches() {
      try {
        const response = await axios.get("http://localhost:8081/api/branches"); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능
        //firebase관련 추가로 repsonse.data뒤에 branches 입력
        this.branches = response.data.branches;
        //console.log("### 전체 response 객체 ### :", response);
        console.log("### Branches data 나오라고 ### :", response.data);
      } catch (error) {
        console.error("#### Error fetching branches ##### :", error);
      }
    }
  },
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap');


html, body, #app {
  width: 100vw;
  height: 100%;
  margin: 0;
  padding: 0;
  overflow-x: hidden;
}

ul {
  list-style-type: none; /* 동그라미 기호 제거 */
}

li {
  list-style-type: none; /* li에 대한 동그라미 기호 제거 (선택 사항) */
}

a{
    text-decoration: none;
}

.booking-header{
    width: 100vw;
    display: flex;
    justify-content: flex-end; /* 요소들을 가로로 균등하게 정렬 */
    /*align-items: right;  세로 정렬 */
    /*padding: 20 20px;  좌우 여백 추가 (너무 화면 끝까지 붙지 않도록) */
    padding: 20px 40px 20px 20px; /* 위쪽 20px, 오른쪽 40px, 아래쪽 20px, 왼쪽 20px */
    background-color: #fff; /* 배경색 설정 (원하는 색상으로 변경 가능) */
    margin: 0; /* 헤더의 외부 여백 제거 */
    position: fixed; /* 화면 맨 위에 고정 */
    top: 0; /* 최상단 배치 */
    left: 0;
    /* right: 0; 오른쪽 끝까지 확장 */
    z-index: 1000; /* 다른 요소들 위에 배치 */
}

.wrapper{
    justify-content: center;
    width: 100%;
    height: 100%;
}
.company-logo-and-name{
  color: white;
  top: 50%; /* 수직 중앙 정렬 */
  left: 50%; /* 수평 중앙 정렬 */
  text-align: center;
  /* transform: translate(-50%, -50%); /* 정확한 중앙 정렬을 위해 transform 사용 */
  font-size: 2em; /* 텍스트 크기 */
  font-weight: bold; /* 텍스트 두껍게 */
  z-index: 2; /* h1이 버튼 위에 오도록 z-index 설정 */
  margin-bottom: 10px; /* h1과 버튼 간의 간격 */
}

.btn-container {
    display: flex;
    justify-content: center;  /* 가로 중앙 */
    align-items: center;      /* 세로 중앙 */
    margin-top: 0;             /* margin-top 없애기 */
    position: relative;        /* 위치 기준 잡기 */
    z-index: 10;               /* 버튼이 다른 요소들 위로 오게 */
}

.btn-book {
    background-color: #ffb3b3;
    color: white;
    text-decoration: none;
    padding: 20px 20px;
    border-radius: 5px;
    font-size: 16px;
    z-index: 1;
    position: relative;
    /* margin-top: 220px; <-- 이거 삭제!!! */
}

.background-image{
    width: 100vw; /* 뷰포트 전체 너비를 차지하도록 설정 */
    margin: 0 auto; /* 가로 중앙 정렬 */
    display: block; /* inline 요소가 가진 여백 제거 */
    z-index: 0; /* 이미지가 가장 뒤에 있도록 z-index 설정 */
    height: 100%;
    position: absolute; /* 이미지 배경으로 설정 */
    top: 0;
    left: 0;
    object-fit: cover; /* 이미지 비율에 맞게 채우기 */
    opacity: 0.8; /* 이미지 불투명도 설정, 1이면 불투명, 0이면 완전 투명 */
    filter: brightness(50%);  /*이미지 어둡게 */
}

.booking-process-item {
    width: 100vw; /* 뷰포트 전체 너비를 차지하도록 설정 */
    display: flex;
    height: 70px;
    list-style: none;
    padding: 0;
    margin: 0;
    background-color: #ffe6e6;
    margin-bottom: 30px;
}

.step-info-item{
    width: 50%; /* 각 항목의 너비를 25%로 설정 */
    padding: 10px; /* 항목 간 간격을 위한 padding (옵션) */
    box-sizing: border-box; /* padding이 width에 포함되도록 설정 */
    align-items: center; /* 세로(수직) 가운데 정렬 */
    justify-content: center; /* 가로(수평) 가운데 정렬 */
    text-align: center;
    cursor: pointer;
    border: 1px solid #e6e6e6; /* ul 테두리 추가 */
}

.header_navigation_nav{
    display: flex;
    list-style: none;
    text-align: center;
    /* padding-right: 50px; /* 네비게이션 바 오른쪽 여백 추가 ############# 이게 오른쪽 여백 적용되는 부분################*/  
    margin: 0;
    height: 30px; /* 메뉴 높이 설정 (예: 50px) */
}

.header__nav-item .header__nav-item{
    padding: 10px 15px; /* 메뉴 항목 내부 여백 */
    white-space: nowrap; /* 줄바꿈 방지 */
    min-width: 120px; /* 최소 너비 설정 */
    height: 100%; /* 메뉴 항목의 높이에 맞게 설정 */
    padding: 0 15px; /* 메뉴 항목의 좌우 여백 */

}
.header__nav-item a{
    color: #595959;
    height: 100%; /* 메뉴 항목 높이에 맞게 100%로 설정 */
    display: block; /* 클릭 영역 확보 */
    padding: 10px 15px; /* 링크 내부 여백 */
    text-decoration: none;
    font-size: 16px; /* 글자 크기 조정 */
}

.header__nav-item a:hover{
    border-bottom: 3px solid #6699ff; /* hover 시 선 표시 */
}

.contact-info{
    display: flex;
    margin-top: 20px;
    justify-content: center;
    text-align: left;
    color: #003366;
    padding: 20px;
    box-shadow: 0px 4px 15px rgba(0, 0, 0, 0.08);
    margin: 0 auto; /* 가로 중앙 정렬 */
    width: 80%; /* 또는 적절한 width 설정 */
}

.branch-item{
    margin-top: 10px;
    margin-bottom: 30px; /* 각 항목 간의 간격을 넓히기 */
}


.btn-bar {
    flex: 1; /* 버튼 영역도 동일한 너비로 설정 */
    display: flex;
    justify-content: center; /* 버튼을 중앙 정렬 */
}

.btn-book:hover{
    background-color: #6699ff; /* 호버 시 색상 변경 */
}
</style>
