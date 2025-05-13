<template>
    <div>
        <header class="booking-header">
            <div id="sb_menu" class="header_menu_wrapper">
                <ul class="header_navigation_nav" id="sb_menu_list_item_container">
                    <li class="header__nav-item">
                        <a class="go-home" :href="`/home/${branchID}`" target="_self">Twinkle Kids Cafe </a>
                    </li>
                    <!-- 
                        <li class="header__nav-item">
                            <a class="go-my-booking" href="#client/bookings/type/upcoming">My Bookings</a>
                        </li>
                    -->
                    <li class="header__nav-item">
                        <a class="go-faq" :href="`/faq/${branchID}`">FAQ</a>
                    </li>
                </ul>
            </div>
        </header>
        <!-- 메인 사진 and booking button -->
        <!-- <div class="wrapper">
            <div><h1 class="company-logo-and-name">Twinkle Kids Cafe</h1></div>
            <div class="background-wrapper">
                <img src="https://images.squarespace-cdn.com/content/v1/637d8d8a7f609c521ddd5429/1672359448650-N89Q21OUSYRU8ROW18F1/Burwood+Plaza+Max3MB_72DPI_VCLAMedia+%2854+of+101%29.jpg" 
                alt="Background" class="background-image" />
            </div>
        </div> -->
    </div>
</template>

<script>
export default {
    data(){
    return{
    }
  },
    async created() {
        // URL에 branch_id가 있을 경우 localStorage에 저장
        const branchFromQuery = this.$route.query.branch_id || this.$route.query.branchID;
        if (branchFromQuery) {
        localStorage.setItem('branchID', branchFromQuery);
        }
        //const branchID = this.$route.params.branchID
        //const branchID = this.$route.query.branch_id;
  },
  computed: {
        branchID() {
        // localStorage에서 branchID 가져오기
        return localStorage.getItem('branchID') || 'burwood'; // 기본값 설정
        }
    }
}
</script>

<style>
#app {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
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
    /* position: fixed; 
    top: 0; 
    left: 0; */
    /* right: 0; 오른쪽 끝까지 확장 */
    z-index: 1000; /* 다른 요소들 위에 배치 */
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

.header_navigation_nav{
    display: flex;
    list-style: none;
    text-align: center;
    padding-right: 50px; /* 네비게이션 바 오른쪽 여백 추가 ############# 이게 오른쪽 여백 적용되는 부분################*/ 
    margin: 0;
    height: 30px; /* 메뉴 높이 설정 (예: 50px) */
}

/* 태블릿용 (max-width: 1024px) */
@media (max-width: 1024px) {
  .header_navigation_nav {
    gap: 10px;
  }

  .header__nav-item a {
    font-size: 15px;
  }
}

@media (max-width: 768px) {
  .booking-header {
    padding: 10px 15px;
    justify-content: flex-start;
    overflow-x: auto;
  }

  .header_menu_wrapper {
    width: 100%;
    overflow-x: auto;
  }

  .header_navigation_nav {
    flex-wrap: nowrap;
    white-space: nowrap;
    overflow-x: auto;
    -webkit-overflow-scrolling: touch;
    padding: 0 10px;
    gap: 8px;

    /* ✅ 스크롤은 되지만 스크롤바는 숨김 */
    scrollbar-width: none; /* Firefox */
  }

  .header_navigation_nav::-webkit-scrollbar {
    display: none; /* Chrome, Safari */
  }

  .header__nav-item {
    flex: 0 0 auto;
  }

  .header__nav-item a {
    font-size: 14px;
    padding: 8px 12px;
  }
}


</style>