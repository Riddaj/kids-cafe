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
                        <a class="go-faq" href="#client/bookings/type/upcoming">FAQ</a>
                    </li>
                </ul>
            </div>
    </header>
    <h1>FREQUENTLY ASKED QUESTION</h1>

    </div>
</template>

<script>
import axios from 'axios'; // axios를 import 추가
import CurrentTime from '../components/CurrentTime.vue';
import Tree from 'primevue/tree';

export default {

}
</script>

<style scoped>
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

.faq-body{
    text-align: left;
    margin: 20px 20px 20px 50px;
}

.set{
    margin-bottom: 10px;
}

</style>