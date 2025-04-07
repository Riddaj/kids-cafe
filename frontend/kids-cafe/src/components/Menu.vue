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
                                <a href="/menu" class="sf-with-ul"><span>Cafe Menu</span></a>
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
        <div class="main">
            <div class="menu-wrapper">
                <div v-for="(categoryMenus, category) in categorizedMenus" :key="category">
                <h2 class="category-title">{{ category }}</h2>
                    <div v-for="menu in categoryMenus" :key="menu.MenuID" class="menu-item">
                        <div class="menu-name">{{ menu.MenuName }}</div>
                            <!-- MenuOptions 배열 반복 -->
                        <ul class="menu-options">
                            <li v-for="(option, index) in menu.MenuOptions" :key="index">
                                <span v-if="option.Size">{{ option.Size }} - </span>${{ option.Price }}
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'; // axios를 import 추가

export default {
    data(){
        return{
            menus:[],
            categorizedMenus: {},  // 카테고리별로 메뉴를 나눠 저장할 객체
        }
    },
    mounted() {
        this.fetchmenu();  // 컴포넌트가 마운트되면 fetchmenu 호출
    },
    methods:{
        async fetchmenu(){
            try {
                const response = await axios.get("http://localhost:8081/api/cafe-menu");
                this.menus = response.data.menus;
                this.categorizeMenu(); 
                console.log("### menu data 나오라고 ### :", response.data.menus);
                
            } catch (error) {
                console.error("#### Error fetching menus ##### :", error);
            }
        },       
        categorizeMenu() {
            // menus 배열을 MenuCategory 기준으로 분류
            this.categorizedMenus = this.menus.reduce((categories, menu) => {
                const category = menu.MenuCategory;

                if (!categories[category]) {
                    categories[category] = [];  // 카테고리가 없으면 새 배열 생성
                }

                categories[category].push(menu);  // 카테고리에 해당하는 메뉴 추가
                return categories;
            }, {});
        }
    },
    created() {
        this.fetchmenu();  // 컴포넌트 생성 시 메뉴 데이터 가져오기
    }
}
</script>

<style scoped>
#app {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

.main {
  display: flex;
  justify-content: left;
  padding: 120px 0 40px; /* top padding으로 헤더 피하기 */
}

.menu-wrapper {
  width: 90%;
  max-width: 800px; /* 최대 너비 설정해서 가운데 정렬 */
  text-align: left;
  padding: 60px;
}

.menu-item {
  display: flex;
  align-items: center;
  gap: 12px;
  flex-wrap: wrap; /* 혹시라도 넘치면 자동 줄바꿈 */
  /*margin-bottom: 24px;*/
}

.menu-name {
  font-weight: bold;
  font-size: 1.2rem;
  margin-bottom: 6px;
  white-space: nowrap;
}

.menu-options {
  display: flex;
  /* flex-wrap: wrap; 한 줄 넘치면 줄바꿈 허용 */
  gap: 12px;        /* 옵션 간 간격 */
  padding: 0;
  margin: 0;
  list-style: none;
}

.menu-options li {
  /*display: inline;
  display: flex;
  align-items: center;
  margin-right: 12px;  옵션 간 간격 */
  white-space: nowrap;
}

.category-title{
    color: #ff6600; 
    
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
</style>
