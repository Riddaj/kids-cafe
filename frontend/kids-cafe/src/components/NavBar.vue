<template>
    <header class="index-header">
        <!-- 로고 + 햄버거 (모바일용 상단 바) -->
        <div class="mobile-header" v-if="isMobile">
            <a :href="`/home/${$route.params.branchID}`">
                <img class="mobile-logo" src="./../assets/twinkle_logo.webp" alt="Logo" />
            </a>
            <div style="padding-right: 10px;"></div>
            <span v-if="branchName" class="branch-name">{{ branchName }}</span>
            <button class="hamburger-icon" @click="menuOpen = true">
                ☰
            </button>
        </div>
        
        <!-- 기존 데스크탑 메뉴 -->
        <!-- logo -->
        <div class="sc-layouts-logo-container">
            <a :href="`/home/${this.$route.params.branchID}`" class="sc_layouts_logo ">
            <img fetchpriority="high" class="logo_image" src="./../assets/twinkle_logo.webp" width="2390" height="924"  loading="eager" alt="Logo fallback" /></a>
            <!-- 브랜치 이름 표시 -->
            <span v-if="branchName" class="branch-name">{{ branchName }}</span>
        </div>

        <!-- menu + button -->
        <div class="sc_layouts_item">
            <nav class="sc_layouts_menu" id="trx_sc_layouts_menu">
                <ul id="sc_layouts_menu" class="sc_layouts_menu_nav" style="touch-action: pan-y;">
                    <li id="menu-item-3040" class="menu-item">
                        <a href="/"><span>Location</span></a></li>
                    <li id="menu-item-3044" class="menu-item">
                        <a :href="`/price/${this.$route.params.branchID}`" class="sf-with-ul">
                            <span>Price</span></a>
                    </li>
                    <!-- 
                        <li id="menu-item-3093" class="menu-item">
                            <a :href="`/about-us/${this.$route.params.branchID}`"><span>About Us</span></a>
                        </li>
                    -->
                    <li id="menu-item-3046" class="menu-item">
                        <a :href="`/parties-events/${this.$route.params.branchID}`"><span>Parties &amp; Events</span></a>
                    </li>
                    <!-- </li><li id="menu-item-3047" class="menu-item" data-width="110.012" :class="{ 'active': activeMenu === 'Cafe Menu' }" @click="setActiveMenu('Cafe Menu')"> -->
                    <!-- <li id="menu-item-3047" class="menu-item">
                        <a :href="`/menu/${this.$route.params.branchID}`" class="sf-with-ul"><span>Cafe Menu</span></a>
                    </li> -->
                    <li class="menu-item">
                        <a :href="`/entryrules/${this.$route.params.branchID}`"><span>Entry Rules</span></a>
                    </li>
                    <li id="menu-item-3048" class="menu-item">
                        <a :href="`/faq/${this.$route.params.branchID}`"><span>FAQ</span></a>
                    </li>
                    <li class="menu-item">
                        <a :href="`/contact/${this.$route.params.branchID}`"><span>Contact Us</span></a>
                    </li>
                    <li class="menu-item">
                        <a :href="`/login`"><span>Log In</span></a>
                    </li>
                </ul>
            </nav>
            <div class="btn-container">
                <a class="btn-book" :href="`/book_a_party/select_room/${this.$route.params.branchID}`" title="Book a party">
                    Book a party
                </a>
            </div>
        </div>
            <!-- 모바일 슬라이드 메뉴 -->
            <div class="mobile-slide-menu" v-if="menuOpen">
                <button class="close-btn" @click="menuOpen = false">×</button>
                <div style="padding-bottom: 30px;"></div>
                <ul>
                    <li><a href="/">Location</a></li>
                    <li><a :href="`/price/${$route.params.branchID}`">Price</a></li>
                    <li><a :href="`/parties-events/${$route.params.branchID}`">Parties & Events</a></li>
                    <li><a :href="`/entryrules/${$route.params.branchID}`">Entry Rules</a></li>
                    <li><a :href="`/faq/${$route.params.branchID}`">FAQ</a></li>
                    <li><a :href="`/contact/${$route.params.branchID}`">Contact Us</a></li>
                    <li><a href="/login">Log In</a></li>
                </ul>
                <div style="padding-bottom: 15px;"></div>
                <div>
                    <!-- 모바일 전용 고정 버튼 -->
                    <a class="mobile-book-button" :href="`/book_a_party/select_room/${$route.params.branchID}`">
                        Book a party
                    </a>
                </div>
            </div>
        </header>
</template>


<script>
export default {
  name: 'NavBar',
  data(){
    return{
        branchName: '', 
        menuOpen: false,
        isMobile: false
    }
  },
  async created() {
    const branchID = this.$route.params.branchID
    this.branchName = this.getBranchNameByID(branchID)
  },
  mounted() {
    this.checkIfMobile();
    window.addEventListener('resize', this.checkIfMobile);
  },
  beforeUnmount() {
    window.removeEventListener('resize', this.checkIfMobile);
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
    checkIfMobile() {
      this.isMobile = window.innerWidth <= 1024;
    }
  }
};
</script>

<style scoped>
#app {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

@media screen and (max-width: 1024px) {
  .sc_layouts_menu {
    display: none !important;
  }
  .mobile-header {
    display: flex;
  }
  .branch-name {
    display: none;
  }
  .logo_image{
    display: none;
  }
  .btn-container{
    display: none;
  }
  .mobile-book-button {
    bottom: 20px;
    right: 20px;
    background: #7212a6;
    color: white !important;
    padding: 12px 16px;
    border-radius: 30px;
    font-weight: bold;
    z-index: 999;
    /* box-shadow: 0 4px 10px rgba(0,0,0,0.2); */
    text-decoration: none;
  }

  .mobile-book-button:hover {
    background-color: #e65500;
    color: white;
 }
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
    max-height: 106px;
    height: auto;
    max-width: 70%;
}

.index-header {
    width: 100vw;
    display: flex;
    justify-content: space-between; /* 요소들을 가로로 균등하게 정렬 */
    align-items: center; /* 세로 정렬 */
    padding: 6px 6px; /* 여백 설정 */
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
    padding: 6px 11px; /* 메뉴 항목 내부 여백 */
    white-space: nowrap; /* 줄바꿈 방지 */
    min-width: 100px; /* 최소 너비 설정 */
    height: 100%; /* 메뉴 항목의 높이에 맞게 설정 */
    padding: 0 10px; /* 메뉴 항목의 좌우 여백 */
}

.sc_layouts_menu_nav .menu-item a {
    color: #595959;
    height: 100%; /* 메뉴 항목 높이에 맞게 100%로 설정 */
    display: block; /* 클릭 영역 확보 */
    padding: 10px 15px; /* 링크 내부 여백 */
    text-decoration: none;
    font-size: 16px; /* 글자 크기 조정 */
}

.menu-item.active {
    color: #7212a6; /* 선택된 메뉴 항목 색상 */
    font-weight: bold; /* 선택된 메뉴 항목 강조 */
}

.menu-item > a:hover{
    color: #ff6600; /* 호버 시 색상 변경 */
}


/* 모바일 헤더: 로고 + 햄버거 */
.mobile-header {
  width: 100vw;               /* ✅ 전체 뷰포트 너비를 꽉 채움 */
  max-width: 100%;            /* ✅ 오버플로우 방지 */
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  position: relative;
  z-index: 1001;
}

.mobile-logo {
  height: 48px;
  object-fit: contain;
}

.hamburger-icon {
  font-size: 28px;
  /* background: none; */
  border: none;
  cursor: pointer;
}

/* 슬라이드 메뉴 (오른쪽에서 나옴) */
.mobile-slide-menu {
  position: fixed;
  top: 0;
  right: 0;
  width: 260px;
  height: 100%;
  background: white;
  box-shadow: -2px 0 8px rgba(0,0,0,0.1);
  padding: 20px;
  z-index: 1002;
  animation: slideIn 0.3s ease forwards;
}

@keyframes slideIn {
  from {
    right: -260px;
  }
  to {
    right: 0;
  }
}

.mobile-slide-menu .close-btn {
  font-size: 26px;
  background: none;
  border: none;
  float: right;
  cursor: pointer;
}

.mobile-slide-menu ul {
  margin-top: 50px;
  list-style: none;
  padding: 0;
}

.mobile-slide-menu li {
  margin-bottom: 20px;
}

.mobile-slide-menu a {
  text-decoration: none;
  color: #333;
  font-size: 18px;
}


</style>
