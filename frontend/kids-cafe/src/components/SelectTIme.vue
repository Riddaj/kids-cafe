<template>
  <div>
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
        <!-- 메인 사진 and booking button -->
        <div class="wrapper">
            <div><h1 class="company-logo-and-name">Twinkle Kids Cafe</h1></div>
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
        <!-- ###########      메인 입력 내용        ########### -->
         <div class="main-card">
             <div class="selected-room-card">
                 <img src="/images/happy-birthday.png" alt="party-hat" class="happy-birthday">
                <div v-if="Object.keys(selectedroom).length > 0"  class="selectedroom-container">
                    <div class="room-name">{{ selectedroom.room_name }}</div>
                    <div>{{ selectedroom.description }}</div>
                </div>
             </div>
             <div class="main-time-pick">
                <!-- <p>{{ partyroom.room_name }} Available times</p> -->
                <form @submit.prevent="submitForm">
                    <table>
                        <tr>
                            <td><label for="datepicker">Choose a date:</label></td>
                            <td>
                                <vue3-datepicker v-model="selectedDate" format="yyyy-MM-dd"></vue3-datepicker>
                            </td>
                        </tr>
                        <tr>
                            <td>Selected Date: <span class="highlight">{{ formattedDate }}</span></td>
                            <td style="display: none;">{{ dateInput }}</td>
                        </tr>
                        <tr>
                            <td>Selected Day of Week: <span class="highlight">{{ selectedDay  }}</span></td>
                        </tr>
                        <tr class="time-options-row">
                            <td
                                v-for="(timeOption, index) in timeOptions"
                                :key="index"
                                class="time-card"
                                :class="{ selected: selectedTime === timeOption }"
                                @click="selectTime(timeOption)"
                                >
                                {{ timeOption }}
                            </td>
                        </tr>
                        <tr>
                            <td class="button-td">
                                <button type="submit" class="submit-button">Next</button>
                            </td>
                        </tr>
                    </table>
                </form>
             </div>
         </div>
  </div>
</template>
<script>
import { useRoute } from 'vue-router';
import axios from 'axios'; // axios를 import 추가
import CurrentTime from '../components/CurrentTime.vue';
import { ref, computed, watch } from "vue"; // ✅ watch 추가
import Vue3Datepicker from 'vue3-datepicker';
// import 'vue3-datepicker/dist/main.css';
import Datepicker from 'vue3-datepicker';
const selectedDate = ref(null); // 날짜 상태 추가
// const route = useRoute();
// const selectedRoomName = route.query.room; // URL에서 'room' 값을 가져옴

export default {
    components: {
        CurrentTime,
        Vue3Datepicker
    }, 
     data() {
         return {
             selectedroom:{}, //배열 대신 객체로 변경
             branchID: "", // 기본값 설정
             roomID: "",
             roomName: "",
             selectedDate: null, // ✅ null로 초기화 (날짜 객체 저장)
             dayofweek: "",  // ✅ dayofweek 추가
             //dateTimeInput: "",   
             timeOptions: ['10:00~13:00', '16:00~18:00'],  // 선택 가능한 시간 옵션
             selectedTime: '',  // 선택된 시간
         };
    },
  mounted(){
    this.branchID = this.$route.params.branch_id || this.$route.query.branch_id || ""; 
    this.roomID = this.$route.params.roomID || this.$route.query.roomID || "";
    this.roomName = decodeURIComponent(this.$route.query.room_name) || ""; // query string에서 가져오기

    console.log('#### Room ID #### :', this.roomID, this.branchID, this.roomName); // 이 값이 정상적으로 출력되는지 확인
    const partyroomId = this.roomID;  // URL 파라미터에서 roomID 추출
    console.log("나와롸 ################# ", partyroomId);
    this.fetchSelectedroomData(partyroomId);
  },
    methods: {
        async fetchSelectedroomData(partyroomId) {
        console.log("📌 Axios 요청 보냄 - partyroomId:", partyroomId);
        console.log("📌 Axios 요청 보냄 - branch_id:", this.branchID);
        console.log("📌 Axios 요청 보냄 - room_name:", this.roomName);
        
            try {
                const response = await axios.get(`http://localhost:8081/api/selectedroom/${partyroomId}`
                , {params : {
                    branch_id: this.branchID
                    , room_name: this.roomName}
                }); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

                this.selectedroom = response.data;  // 받아온 데이터를 partyroom에 저장
                console.log("### selected data 언제 나오냐고 ### :", response.data);
            } catch (error) {
                console.error('Error fetching selected room data:', error);
            }
        },
        selectTime(option) {
        this.selectedTime = option;  // 선택된 옵션을 저장
        console.log("선택된 시간:", this.selectedTime); // 콘솔로 확인
        },
        async submitForm() {
            // console.log(this.selectedDate);
            console.log("이거 나오나", this.dateInput);
            console.log("이거 ", this.dateInput);
            console.log("time이 왜안나오는거야 ===== ", this.selectedTime );

            if (!this.selectedDate) {
                alert("Please enter the date.");
                return;
            }
            if (!this.selectedTime) {
            alert("Please enter the time.");
            return;
            }
            try {
                const response = await axios.post("http://localhost:8081/api/select-time", {
                    partyroom_id:  this.roomID,
                    selected_date: this.dateInput,
                    selected_time: this.selectedTime,
                },{
                    headers: {
                        'Content-Type': 'application/json'
                    }
                });
                console.log("$$$$ 응답 $$$$: ", response);  // 응답 상태 코드 및 데이터 확인
                alert("예약 성공적으로 완료되었습니다"); // 성공 메시지 출력
            } catch (error) {
                //console.log("===== date ======: ", date);
                console.error("Error:", error);
                alert("날짜 저장에 실패했습니다.");
            }
        },
    },
  computed: {
        formattedDate() {
            if (!this.selectedDate) return "";
            const date = new Date(this.selectedDate);
            const day = String(date.getDate()).padStart(2, '0');
            const month = String(date.getMonth() + 1).padStart(2, '0'); // 월은 0부터 시작하므로 +1
            const year = date.getFullYear();
            return `${day}/${month}/${year}`; // ✅ DD/MM/YYYY 형식으로 변환
        },
        selectedDay() {
            if (!this.selectedDate) return "";
            const date = new Date(this.selectedDate);
            return date.toLocaleDateString("en-US", { weekday: "short" }); // ✅ 요일 (Fri)
        },
        dateInput(){
            if (!this.selectedDate) return "";
            const date = new Date(this.selectedDate);
            const dateStr = date.toString(); // Date 객체를 문자열로 변환
            const dateandday = dateStr.split(" GMT")[0]; // 문자열에서 " GMT"를 기준으로 잘라내기
            return dateandday;

        }
    }

}
</script>
<style scoped>
tr {
  padding-top: 10px;    /* 위쪽 여백 */
  padding-bottom: 10px; /* 아래쪽 여백 */
}

td {
  padding: 10px;        /* 셀 안의 내용에 여백 */
}

.button-td{
    display: flex;            /* Flexbox 활성화 */
    justify-content: center;  /* 수평 가운데 정렬 */
    align-items: center;      /* 수직 가운데 정렬 */
    height: 100px;            /* 버튼을 세로로 충분히 공간을 두려면 높이를 지정 */

}

.submit-button{
    background-color: #ffb3b3; /* 버튼 배경 색상 */
    color: white; /* 텍스트 색상 */
    text-decoration: none; /* 링크 스타일 제거 */
    padding: 10px 20px; /* 버튼 내부 여백 */
    border-radius: 5px; /* 둥근 모서리 */
    font-size: 16px; /* 글자 크기 */
    display: inline-block; /* 버튼 형태 유지 */
    margin-top: 30px; /* 버튼이 자동으로 글자 아래에 배치되도록 설정 */
}

.submit-button:hover{
    background-color: #6699ff; /* 버튼 배경 색상 */
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

.main-card{
    display: flex;
    justify-content: center;
}

.highlight{
    color: #6699ff;
    font-weight: bold;
}

.room-name{
    font-size: 25px;
    font-weight: bold;
}

.selected-room-card{
    border-radius: 10px;
    padding: 15px;
    margin-right: 50px;
    width: 350px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    text-align: left;
    transition: transform 0.2s, box-shadow 0.2s;
    align-items: center; /* 가로 가운데 정렬 */
    justify-content: center; /* 수평 가운데 정렬 */
}

.happy-birthday{
    width: 120px;
    height: auto; /* 높이는 비율에 맞게 자동 조정 */
    object-fit: contain; /* 이미지 비율을 유지하면서 가능한 공간에 맞추기 */
    display: block;
    margin: 0 auto; /* 이미지 중앙 정렬 %이거 하니까 적용됨% */
}

.selectedroom-container {
    width: 100%; /* 부모 요소 크기 전체 사용 */
    display: flex;
    flex-direction: column;
    align-items: flex-start; /* 왼쪽 정렬 */
    padding: 10px; /* 텍스트와 카드 간 여백 추가 */
}

.main-time-pick{
    border-radius: 10px;
    padding: 15px;
    width: 600px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    text-align: left;
    transition: transform 0.2s, box-shadow 0.2s;
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

.time-options-row {
  display: flex; /* 가로로 배열 */
  justify-content: center; /* 가운데 정렬 */
  align-items: center; /* 수직 가운데 정렬 */
  gap: 10px; /* 각 카드 사이에 간격 추가 */
}

.time-card {
  width: 120px;
  height: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
  border: 2px solid #ccc;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s, transform 0.2s;
  text-align: center;
  padding: 10px;
}

.time-card:hover {
  background-color: #f0f0f0;
  transform: scale(1.05);
}

.time-card.selected {
  background-color: #4CAF50;
  color: white;
  border-color: #4CAF50;
}
</style>