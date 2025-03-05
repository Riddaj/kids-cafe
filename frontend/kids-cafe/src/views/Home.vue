<template>
    <div id="app">
        <header class="index-header">
            <div class="sc-layouts-logo-container">
                <a href="/" id="trx_sc_layouts_logo_598080678" 
                class="sc_layouts_logo sc_layouts_logo_default trx_addons_inline_2112412811">
                <!-- /.sc_layouts_logo -->		
                <img fetchpriority="high" class="logo_image" src="//twinklekidscafe.com.au/wp-content/uploads/2018/08/Twinkle-Kids-Cafe_logo.png" 
                alt="Twinkle Kids Cafe" width="2390" height="924"></a>
            </div>
            <div class="sc_layouts_item">
                    <nav class="sc_layouts_menu" id="trx_sc_layouts_menu">
                        <ul id="sc_layouts_menu" class="sc_layouts_menu_nav" style="touch-action: pan-y;">
                            <li id="menu-item-3040" class="menu-item">
                                <a href="https://twinklekidscafe.com.au/" class="sf-with-ul"><span>Location</span></a></li>
                            <li id="menu-item-3044" class="menu-item">
                                <a href="https://twinklekidscafe.com.au/macarthur-visiting-us/" class="sf-with-ul">
                                    <span>Visiting Us</span></a>
                            </li><li id="menu-item-3093" class="menu-item">
                                <a href="https://twinklekidscafe.com.au/macarthur-about-us/" class="sf-with-ul"><span>About Us</span></a>
                            </li><li id="menu-item-3046" class="menu-item">
                                <a href="https://twinklekidscafe.com.au/macarthur-parties-events/" class="sf-with-ul"><span>Parties &amp; Events</span></a>
                            </li><li id="menu-item-3047" class="menu-item" data-width="110.012">
                                <a href="https://twinklekidscafe.com.au/macarthur-cafe-menu/" class="sf-with-ul"><span>Cafe Menu</span></a>
                            </li><li id="menu-item-3048" class="menu-item">
                                <a href="https://twinklekidscafe.com.au/macarthur-contact-us/"><span>Contact Us</span></a>
                            </li>
                        </ul>
                    </nav>
                    <div class="btn-container">
                        <a class="btn-book" href="/book_a_party" title="Book a party">
                            Book a party
                        </a>
                    </div>
            </div>

        </header>
        <section class="slider">
				<div class="main-text">
			        <h1 class="elementor-heading-title elementor-size-default">Welcome to <br>Twinkle <br>Kids Cafe</h1>
				</div>
				<div class="main-text">
                    <p>Welcome to the center where your children are able <br> 
                        to play, learn, excercise and have the best Birthday 
                        <br>parties!&nbsp;</p>
                </div>
                <div class="quickbtn-container">
                        <a class="quickbtn-book" href="/book_a_party/quickbook" title="Quick Party Room Booking">
                            Quick Party Room Booking
                        </a>
                </div>
                <div class="image-wrapper">
                    <img src="https://images.squarespace-cdn.com/content/v1/637d8d8a7f609c521ddd5429/1672359448650-N89Q21OUSYRU8ROW18F1/Burwood+Plaza+Max3MB_72DPI_VCLAMedia+%2854+of+101%29.jpg" 
                    class="background-image1">
                </div>
		</section>


        <!--
            <section class="slider">
                <div class="info-01">
                    <div class="price-menu-info">
                        <h1 class="playland">Playland</h1>
                    </div>
                    <div class="price-menu-info">
                        <p>We are  the best destination<br>for exciting and dynamic family<br>activities! Let’s have fun!</p>
                    </div>
                </div>
                
            </section>
        -->
    </div>
</template>
<script>
export default {

}
</script>

<style>
#app {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

.main-text h1,p{
    z-index: 1; /* h1이 버튼 위에 오도록 z-index 설정 */
    text-align: right;
    position: relative; /* z-index가 제대로 작동하도록 위치 지정 */
    color: #262626;
}

.background-image1{
    z-index: 0; /* 이미지가 가장 뒤에 있도록 z-index 설정 */
    width: 100vw; /* 뷰포트 전체 너비를 차지하도록 설정 */
    height: 100%;
    position: absolute; /* 이미지 배경으로 설정 */
    top: 0;
    left: 0;
    object-fit: cover; /* 이미지 비율에 맞게 채우기 */
    opacity: 0.7; /* 이미지 불투명도 설정, 1이면 불투명, 0이면 완전 투명 */
}

ul {
  list-style-type: none; /* 동그라미 기호 제거 */
}

li {
  list-style-type: none; /* li에 대한 동그라미 기호 제거 (선택 사항) */
}
.logo_image{
    width: 100%;  /* 부모 요소 크기에 맞춤 */
    height: auto;  /* 가로 비율에 맞춰 세로 크기 자동 조정 */
    max-width: 200px; /* 최대 크기 제한 */
}

.index-header {
    width: 100vw;
    display: flex;
    justify-content: space-between; /* 요소들을 가로로 균등하게 정렬 */
    align-items: center; /* 세로 정렬 */
    padding: 10px 10px; /* 여백 설정 */
    background-color: #fff; /* 배경색 설정 (원하는 색상으로 변경 가능) */
    margin: 0; /* 헤더의 외부 여백 제거 */
    position: fixed; /* 화면 맨 위에 고정 */
    top: 0; /* 최상단 배치 */
    left: 0;
    z-index: 1000; /* 다른 요소들 위에 배치 */
}

.sc-layouts-logo-container {
    flex: 1; /* 동일한 너비로 분배 */
    display: flex;
    justify-content: center; /* 로고를 가운데 정렬 */
}

/** 로고를 왼쪽으로 정렬 */
.sc_layouts_logo {  
    display: inline-block; /* 로고를 inline-block으로 만들어 다른 요소와 함께 배치할 수 있도록 설정 */
    margin-left: 0; /* 왼쪽 여백을 0으로 설정 */
}

.elementor-widget-wrap {
    flex: 2; /* 메뉴가 로고보다 더 넓게 배치되도록 설정 */
    display: flex;
    justify-content: center; /* 메뉴를 가운데 정렬 */
}

.sc_layouts_item{
    flex: 3; /* 메뉴와 버튼 영역 */
    display: flex;
    align-items: center; /* 수직 정렬 */
    justify-content: space-between; /* 메뉴와 버튼 정렬 */
}

.sc_layouts_menu {
    flex: 2; /* 메뉴 너비 조정 */
    display: flex;
    justify-content: center;
    gap: 20px; /* 메뉴 간 간격 */
    margin-top: 40px; /* 원하는 값으로 조절 (10px, 15px 등) */
}

.sc_layouts_menu_nav {
    display: flex;
    list-style: none;
    text-align: center;
    padding: 0;
    margin: 0;
    height: 50px; /* 메뉴 높이 설정 (예: 50px) */
}

.sc_layouts_menu_nav .menu-item {
    padding: 10px 15px; /* 메뉴 항목 내부 여백 */
    white-space: nowrap; /* 줄바꿈 방지 */
    min-width: 120px; /* 최소 너비 설정 */
    height: 100%; /* 메뉴 항목의 높이에 맞게 설정 */
    padding: 0 15px; /* 메뉴 항목의 좌우 여백 */
}

.sc_layouts_menu_nav .menu-item a {
    color: #595959;
    height: 100%; /* 메뉴 항목 높이에 맞게 100%로 설정 */
    display: block; /* 클릭 영역 확보 */
    padding: 10px 15px; /* 링크 내부 여백 */
    text-decoration: none;
    font-size: 16px; /* 글자 크기 조정 */
}

.menu-item > a:hover{
    color: #ff6600; /* 호버 시 색상 변경 */
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
