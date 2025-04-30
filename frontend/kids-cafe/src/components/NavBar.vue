<template>
    <header class="index-header">
        
        <!-- logo -->
        <div class="sc-layouts-logo-container">
            <a href="/" class="sc_layouts_logo ">
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
                    <li id="menu-item-3047" class="menu-item">
                        <a :href="`/menu/${this.$route.params.branchID}`" class="sf-with-ul"><span>Cafe Menu</span></a>
                    </li><li id="menu-item-3048" class="menu-item">
                        <a :href="`/faq/${this.$route.params.branchID}`"><span>FAQ</span></a>
                    </li>
                </ul>
            </nav>
            <div class="btn-container">
                <a class="btn-book" :href="`/book_a_party/${this.$route.params.branchID}`" title="Book a party">
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
    }
  },
  async created() {
    const branchID = this.$route.params.branchID
    this.branchName = this.getBranchNameByID(branchID)
  },
  methods: {
    getBranchNameByID(id) {
      const branchMap = {
        'burwood': 'Burwood',
        'hornsby': 'Hornsby',
        // 필요하면 계속 추가 가능
      }
      return branchMap[id] || 'Unknown Branch'
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

.menu-item.active {
    color: #7212a6; /* 선택된 메뉴 항목 색상 */
    font-weight: bold; /* 선택된 메뉴 항목 강조 */
}

.menu-item > a:hover{
    color: #ff6600; /* 호버 시 색상 변경 */
}
</style>
