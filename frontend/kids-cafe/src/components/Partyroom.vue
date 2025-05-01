<template>
    <div>
        <BookingBar/>
        <!-- booking process (2단계) -->
        <BookingProcess/>
            <!-- 현재 시간 -->
        <CurrentTime/>
        <div class="partyroom_info">
            <h1>Party room</h1>
            <div v-if="partyrooms && partyrooms.length > 0"  class="partyroom-container">
                <div v-for="partyroom in partyrooms" :key="partyroom.RoomID">
                    <div class="partyroom-card">
                        <div class="img-div">
                            <img src="/images/party-hat.png" alt="party-hat" class="party-hat">
                        </div>
                        <br>
                        <div>
                            <!-- {{ partyroom.RoomID }} -->
                            {{ partyroom.RoomName }}<br>
                            <hr/>
                            Deposit: {{ partyroom.RoomDeposit }}<br>
                            <div>{{ partyroom.Description }}</div>
                        </div>
                        <div class="description">
                            <router-link :to="`/book_a_party/select-time/${partyroom.RoomID}?branch_id=${partyroom.BranchID}&room_name=${partyroom.RoomName}`" class="select-room">
                                Select</router-link>
                        </div>
                    </div>
                </div>
            </div>
            <div v-else> <!-- 데이터가 없으면 이 메시지 표시 -->
                No party rooms available.
            </div>
        </div>
        <!-- ✅ 아래 여백 주기 -->
        <div style="margin-top: 150px;"></div>
        <!-- ✅ Footer 컴포넌트 추가 -->
        <Footer />
    </div>
</template>

<script>
import axios from 'axios'; // axios를 import 추가
import CurrentTime from '../components/CurrentTime.vue';
import BookingProcess from '../components/BookingProcess.vue';
import BookingBar from '../components/BookingBar.vue';
import Footer from './Footer.vue';

export default {
    components: {
        CurrentTime,
        BookingProcess,
        BookingBar,
        Footer
    }, 
    data() {
        return {
            partyrooms:[],
            branchID: this.$route.params.branchID,
              // URL 파라미터에서 branchID 가져오기
        };
    },
    mounted() {
    this.fetchPartyrooms();
    }, 
    methods: {
    async fetchPartyrooms() {
      console.log("Branch ID:", this.branchID);  // 값이 제대로 있는지 확인
      try {
        const response = await axios.get(`http://localhost:8081/api/partyrooms/${this.branchID}`); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

        this.partyrooms = response.data.partyrooms;
        // 전체 응답 객체 찍어보기
        console.log("### 전체 response 객체 ### :", response);
        console.log("### partyrooms data ### :", response.data.partyrooms);

        // partyrooms 배열에서 RoomID가 정확히 무엇으로 반환되는지 확인
        response.data.partyrooms.forEach(partyroom => {
        console.log("###@@@ RoomID @@@### :", partyroom.RoomID);
        });
        //console.log("### 전체 response 객체 ### :", response);
        console.log("### partyrooms data 나오라고 ### :", response.data.partyrooms);
      } catch (error) {
        console.error("#### Error fetching partyrooms ##### :", error);
      }
    }
  }
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




a{
    text-decoration: none;
}

.party-hat{
    width: 120px;
    height: auto; /* 높이는 비율에 맞게 자동 조정 */
    object-fit: contain; /* 이미지 비율을 유지하면서 가능한 공간에 맞추기 */
}
.partyroom-container {
    display: flex;
    flex-wrap: wrap;
    gap: 25px; /* 카드 사이 간격 */
    justify-content: center; /* 중앙 정렬 */
}

.partyroom-card {
    display: flex; /* 내부 요소를 가로 배치 */
    flex-direction: column; /* 세로 방향 배치 */
    justify-content: space-between; /* 내용과 버튼 사이 여백 자동 조정 */
    align-items: center; /* 중앙 정렬 */
    background: white;
    border-radius: 10px;
    padding: 15px;
    width: 400px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    text-align: left;
    transition: transform 0.2s, box-shadow 0.2s;
}

.partyroom-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
}

.partyroom-card h3 {
    font-size: 18px;
    font-weight: bold;
    color: #333;
    margin-bottom: 10px;
}

.partyroom-card p {
    font-size: 14px;
    color: #666;
}

.select-room {
    background-color: #ffb3b3; /* 버튼 배경 색상 */
    color: white; /* 텍스트 색상 */
    text-decoration: none; /* 링크 스타일 제거 */
    padding: 10px 20px; /* 버튼 내부 여백 */
    border-radius: 5px; /* 둥근 모서리 */
    font-size: 16px; /* 글자 크기 */
    display: inline-block; /* 버튼 형태 유지 */
    margin-top: auto; /* 버튼이 자동으로 글자 아래에 배치되도록 설정 */
}

.select-room:hover{
    background-color: #6699ff; /* 버튼 배경 색상 */
}
</style>

