<template>
    <div id="app">
        <header>
            <div id="sb_menu" class="header_menu_wrapper">
                <ul class="header_navigation_nav" id="sb_menu_list_item_container">
                    <li class="header__navigation-item menu-item ">
                        <a class="popup-hide header__navigation-link icon-pos--none icon-Twinkle Kids Cafe" href="/" target="_self">Twinkle Kids Cafe </a>
                    </li>
                    <li class="header__navigation-item menu-item ">
                        <a class="popup-hide header__navigation-link icon-pos--none icon-menu_bookings" href="#client/bookings/type/upcoming">My Bookings</a>
                    </li>
                </ul>
            </div>
        </header>
        <div class="wrapper">
            <h1 class="company-logo-and-name">Twinkle Kids Cafe</h1>
            <img src="https://images.squarespace-cdn.com/content/v1/637d8d8a7f609c521ddd5429/1672359448650-N89Q21OUSYRU8ROW18F1/Burwood+Plaza+Max3MB_72DPI_VCLAMedia+%2854+of+101%29.jpg" 
            alt="Background" class="background-image" />
        </div>
        <!-- 
            <div id="sb_booking_info" class="sb_booking_info">
                <ul class="sb_booking_item">
                    <li class="step_info_item"><a href="/book/count/1">Service</a></li>
                    <li class="step_info_item"><a href="/book/count/2">Service Add-ons</a></li>
                    <li class="step_info_item"><a href="/book/count/3">Time</a></li>
                    <li class="step_info_item"><a href="/book/count/4">Client</a></li>
                </ul>
            </div>
        -->
        <div id="booking_info" class="booking_info">

        </div>
    </div>
    
</template>

<script>
export default {

}
</script>

<style>
#app{
  margin: 0;
  padding: 0;
  height: 100%; /* 전체 높이 설정 */
  width: 100%; /* 전체 너비 설정 */
}

/* 메뉴 전체 스타일 */
.header_menu_wrapper {
  background-color:  white; /* 링크 텍스트 색상 */
  margin: 0 15px; /* 좌우 여백 */
}

/* 메뉴 항목들을 일렬로 정렬 */
.header_navigation_nav {
  list-style: none; /* 기본 목록 스타일 제거 */
  padding: 0;
  margin: 0;
  display: flex; /* 메뉴 항목들을 가로로 배치 */
  justify-content: flex-end; /* 왼쪽 정렬 */
  align-items: center; /* 수직 가운데 정렬 */
}

/* 메뉴 항목 간 간격 설정 */
.header__navigation-item {
  margin: 0 15px; /* 메뉴 항목 간의 간격을 15px로 설정 */
}

/* 메뉴 항목 스타일 */
.header__navigation-link {
  color: #333333;
  text-decoration: none; /* 링크 밑줄 제거 */
  font-size: 14px;
  font-weight: bold;
}

/* 링크 hover 시 스타일 */
.header__navigation-link:hover {
  color: #f39c12; /* hover 시 색상 변경 */
}


.wrapper{
  position: relative; /* 컨테이너에 상대적인 위치를 설정 */
  width: 100%;
  height: 300px; /* 이미지 크기 설정 */
  text-align: center;
}

.background-image{
  width: 100vw; /* 뷰포트 전체 너비를 차지하도록 설정 */
  height: 100%;
  object-fit: cover; /* 이미지 비율에 맞게 채우기 */
  /* opacity: 0.8; 이미지 불투명도 설정, 1이면 불투명, 0이면 완전 투명 */
}

.company-logo-and-name{
  color: white;
  position: absolute; /* 텍스트를 절대 위치로 설정 */
  top: 50%; /* 수직 중앙 정렬 */
  left: 50%; /* 수평 중앙 정렬 */
  transform: translate(-50%, -50%); /* 정확한 중앙 정렬을 위해 transform 사용 */
  font-size: 2em; /* 텍스트 크기 */
  font-weight: bold; /* 텍스트 두껍게 */
}

ul {
  list-style-type: none; /* 동그라미 기호 제거 */
}

li {
  list-style-type: none; /* li에 대한 동그라미 기호 제거 (선택 사항) */
}

.sb_booking_item {
    width: 100vw; /* 뷰포트 전체 너비를 차지하도록 설정 */
    display: flex;
    height: 70px;
    list-style: none;
    padding: 0;
    margin: 0;
    background-color: #ffe6ff;
}

.step_info_item{
    width: 25%; /* 각 항목의 너비를 25%로 설정 */
    padding: 10px; /* 항목 간 간격을 위한 padding (옵션) */
    box-sizing: border-box; /* padding이 width에 포함되도록 설정 */
    align-items: center; /* 세로(수직) 가운데 정렬 */
    justify-content: center; /* 가로(수평) 가운데 정렬 */
    text-align: center;
    cursor: pointer;
    border: 1px solid #e6e6e6; /* ul 테두리 추가 */
}

.sb_booking_item > li > a {
    text-decoration: none;
    font-size: 16px;
    font-weight: bold;
    color: #333;
    padding: 10px 15px;
    transition: color 0.3s ease;
}

</style>