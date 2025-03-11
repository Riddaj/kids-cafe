<template>
    <div id="app"> 
        <header class="booking-header">
            <div id="sb_menu" class="header_menu_wrapper">
                <ul class="header_navigation_nav" id="sb_menu_list_item_container">
                    <li class="header__nav-item">
                        <a class="go-home" href="/" target="_self">Twinkle Kids Cafe </a>
                    </li>
                    <li class="header__nav-item">
                        <a class="go-my-booking" href="#client/bookings/type/upcoming">My Bookings</a>
                    </li>
                    <li class="header__nav-item">
                        <a class="go-faq" href="/faq">FAQ</a>
                    </li>
                </ul>
            </div>
        </header>
        <!-- 메인 사진 and booking button -->
        <div class="wrapper">
            <div><h1 class="company-logo-and-name">Twinkle Kids Cafe</h1></div>
            <div class="btn-container">
                <!--<router-link class="btn-book" to="/book_a_party/simplybook/confirm_detail" title="Book Now">Book Now</router-link>-->
                <a class="btn-book" href="#" @click="showBookingInfo = !showBookingInfo" title="Book Now">Book Now</a>
            </div>
            <div class="background-wrapper">
                <img src="https://images.squarespace-cdn.com/content/v1/637d8d8a7f609c521ddd5429/1672359448650-N89Q21OUSYRU8ROW18F1/Burwood+Plaza+Max3MB_72DPI_VCLAMedia+%2854+of+101%29.jpg" 
                alt="Background" class="background-image" />
            </div>
        </div>
        <!-- booking process (2단계) -->
        <div class="booking-process">
            <ul class="booking-process-item">
                <li class="step-info-item"><a href="#">Category</a></li>
                <li class="step-info-item"><a href="#">Time</a></li>
                <li class="step-info-item"><a href="#">Client</a></li>
            </ul>
        </div>
        <div class="button-and-time">
            <!-- 현재 시간 -->
            <div class="current-time"><CurrentTime/></div>
        </div>
        <!-- 변경될 부분 -->
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
                            Phone: {{ branch.phone }}
                            <!-- <hr/> -->
                        </li>
                        <br>
                    </ul>
                </div>
            </template>
        </component>
    </div>
</template>
<script>
import CurrentTime from '../components/CurrentTime.vue';
import axios from "axios";
import ConfirmDetail from '../components/ConfirmDetail.vue';
import Partyroom from '../components/Partyroom.vue';
import SelectBranch from '../components/SelectBranch.vue';

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
        SelectBranch
    },  
    methods: {
    async fetchBranches() {
      try {
        const response = await axios.get("http://localhost:8081/api/branches"); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

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

.btn-container{
    display: flex;
    justify-content: center; 
}

.btn-book {
    margin-top: 220px; /* h1과 버튼 간의 간격 */
    background-color: #ffb3b3; /* 버튼 배경 색상 */
    color: white; /* 텍스트 색상 */
    text-decoration: none; /* 링크 스타일 제거 */
    padding: 20px 20px; /* 버튼의 내부 여백을 조정하여 길이를 늘림 */
    border-radius: 5px; /* 둥근 모서리 */
    font-size: 16px; /* 글자 크기 */
    z-index: 1; /* h1이 버튼 위에 오도록 z-index 설정 */
    position: relative; /* z-index가 제대로 작동하도록 위치 지정 */
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

.current-time{
    display: flex;           /* flexbox 활성화 */
    width: 100vw;
    justify-content: flex-end;
    padding-right: 50px;;
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
    padding-right: 50px; /* 네비게이션 바 오른쪽 여백 추가 ############# 이게 오른쪽 여백 적용되는 부분################*/ 
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
