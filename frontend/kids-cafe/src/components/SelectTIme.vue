<template>
  <div>
        <BookingBar/>
        <!-- <BookingProcess/> -->
        <!-- 현재 시간 -->
        <CurrentTime/>
        <!-- ###########      메인 입력 내용        ########### -->
         <div class="main-card">
             <div class="selected-room-card">
                 <img src="/images/happy-birthday.png" alt="birthday" class="happy-birthday">
                <div v-if="Object.keys(selectedroom).length > 0"  class="selectedroom-container">
                    <div class="room-header">
                        <div class="room-name">{{ selectedroom.RoomName }}</div>
                        <p class="room-session">(3 hours/Session)</p>
                    </div>
                    <div>Room deposit: {{ selectedroom.RoomDeposit }}</div>
                    <div>{{ displayWeekdayPrice  }}</div>
                    <div>{{ displayWeekendPrice }}</div>
                    <div><p class="multiline">{{ selectedroom.Description }}</p></div>
                </div>
             </div>
             <div><h1>Select Time</h1></div>
             <div class="main-time-pick">
                <!-- <p>{{ partyroom.room_name }} Available times</p> -->
                    <table>
                        <tr>
                            <td><label for="datepicker">Choose a date:</label></td>
                        </tr>
                        <tr>
                            <td>
                                <vue3-datepicker v-model="selectedDate" format="yyyy-MM-dd" 
                                inline
                                :auto-apply="true" 
                                :disabled-dates="disablePastDates"
                                @update:modelValue="onDateChange"></vue3-datepicker>
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
                                v-for="(timeOption, index) in filteredTimeOptions"
                                :key="index"
                                class="time-card"
                                :class="{ selected: selectedTime === timeOption }"
                                @click="selectTime(timeOption)"
                                >
                                {{ timeOption }}
                            </td>
                        </tr>
                        <tr>
                            <td>Price of Selected Day: <span class="highlight">{{ selectedPrice  }}</span></td>
                        </tr>
                        <tr style="color:yellowgreen;">
                            <td>Deposit required to confirm your reservation: <span class="highlight">{{ selectedroom.RoomDeposit  }}</span></td>
                        </tr>
                        <tr>
                            <td class="button-td">
                                <!-- <button type="submit" class="submit-button">Next</button>  -->
                                <router-link 
                                v-if="selectedroom.RoomName === 'Private party'"
                                :to="{path: `/book_a_party/foodopitions`,     
                                query: {
                                roomID: roomID,
                                roomName: roomName,
                                roomDeposit : selectedroom.RoomDeposit,
                                selectedDate: selectedDate,
                                selectedTime: selectedTime,
                                selectedPrice: selectedPrice
                                }
                                }">
                                    <button type="submit" @click="goToNext" class="submit-button"
                                    :disabled="!formattedDate || !selectedTime"
                                    >Next</button>
                                </router-link>

                                <!-- 🧁 Party Room Hiring일 경우 -->
                                <router-link
                                v-else-if="selectedroom.RoomName === 'Party room hiring'"
                                :to="{
                                    path: '/book_a_party/client-info',
                                    query: {
                                    roomID: roomID,
                                    roomName: roomName,
                                    roomDeposit: selectedroom.RoomDeposit,
                                    selectedDate: formattedDate,
                                    selectedTime: selectedTime,
                                    selectedPrice: selectedPrice
                                    }
                                }"
                                >
                                <button type="submit" @click="goToNext" class="submit-button"
                                :disabled="!formattedDate || !selectedTime"
                                >Next</button>
                            </router-link>
                            </td>
                        </tr>
                    </table>
             </div>
         </div>
        <!-- ✅ 아래 여백 주기 -->
        <div style="margin-top: 150px;"></div>
         <Footer/>
  </div>
</template>
<script>
import { useRoute } from 'vue-router';
import axios from 'axios'; // axios를 import 추가
import CurrentTime from '../components/CurrentTime.vue';
import BookingBar from '../components/BookingBar.vue';
import Footer from '../components/Footer.vue';
import BookingProcess from '../components/BookingProcess.vue';
import { ref, computed, watch } from "vue"; // ✅ watch 추가
import Vue3Datepicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css';
const selectedDate = ref(null); // 날짜 상태 추가
// import 'vue3-datepicker/dist/main.css';

// const route = useRoute();
// const selectedRoomName = route.query.room; // URL에서 'room' 값을 가져옴

export default {
    components: {
        CurrentTime,
        Vue3Datepicker,
        BookingProcess,
        BookingBar,
        Footer
    }, 
     data() {
         return {
             selectedroom:{}, //배열 대신 객체로 변경
             branchID: "", // 기본값 설정
             roomID: "",
             roomName: "",
             roomDeposit: "",
             selectedDate: null, // ✅ null로 초기화 (날짜 객체 저장)
             dayofweek: "",  // ✅ dayofweek 추가
             //dateTimeInput: "",   
             // 선택 가능한 시간 옵션
             timeOptions: {
                '01': ['10:00~13:00', '13:30~16:30'],
                '02': ['17:00~20:00', '16:00~19:00'],
             },
             selectedTime: '',  // 선택된 시간
             selectedPrice: ''
         };
    },
  mounted(){
    console.log('####@@@@ URL Parameters @@@@####:', this.$route);  // 라우터 정보를 출력하여 문제를 진단
    this.roomID = this.$route.params.roomID || this.$route.query.roomID || "";
    this.branchID = this.$route.params.branch_id || this.$route.query.branch_id || ""; 
    this.roomName = decodeURIComponent(this.$route.query.room_name) || ""; // query string에서 가져오기

    console.log('#### Room ID #### :', this.roomID, this.branchID, this.roomName); // 이 값이 정상적으로 출력되는지 확인
    const roomID = this.roomID;  // URL 파라미터에서 roomID 추출
    console.log("나와롸 ################# ", roomID);
    this.fetchSelectedroomData(roomID);
  },
    methods: {
        disablePastDates(date) {
            const today = new Date()
            today.setHours(0, 0, 0, 0) // 시간 제거 (자정 기준)

            return date <= today // 오늘보다 이전이면 비활성화
        },
        // 선택된 날짜가 바뀔 때 시간, 가격,,, 초기화
        onDateChange(date) {
            console.log(`선택된 날짜는 ${date}입니다.`);
            this.selectedTime = null;

        },
        async fetchSelectedroomData(roomId) {
        console.log("📌 Axios 요청 보냄 - room_Id:", this.roomID);
        console.log("📌 Axios 요청 보냄 - branch_id:", this.branchID);
        console.log("📌 Axios 요청 보냄 - room_name:", this.roomName);
        
            try {
                // const response = await axios.get(`https://kids-cafe-rm9g.onrender.com/api/selectedroom/${roomId}`

                const response = await axios.get(`${process.env.VUE_APP_API_BASE}/api/selectedroom/${roomId}`
                , {params : {
                    room_id: this.roomID,
                    branch_id: this.branchID, 
                    room_name: this.roomName}
                }); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

                this.selectedroom = response.data.partyroom;  // 받아온 데이터를 partyroom에 저장
                console.log("### selected data 언제 나오냐고 ### :", response.data);
            } catch (error) {
                console.error('Error fetching selected room data:', error);
            }
        },
        /**
         * 시간 선택
         */
        async selectTime(option) {
            if (!this.selectedDate) {
                alert("Please select a date first.🎈");
                return;
            }


            this.selectedTime = option;  // 선택된 옵션을 저장
            console.log("선택된 시간:", this.selectedTime); // 콘솔로 확인

            const selectedDate = this.formattedDate;
            console.log("선택된 날짜:", selectedDate);
           
            // 오늘 날짜 가져오기 (YYYY-MM-DD 형식)
            const today = new Date();
            const todayString = today.toISOString().split('T')[0]; // '2025-04-24' 형식으로 변환
            console.log("%^%^%^%^ today ^%^%^%^ = ", todayString);
            

            // DD-MM-YYYY 형식을 YYYY-MM-DD 형식으로 변환
            const [day, month, year] = selectedDate.split('-');
            const formattedSelectedDate = `${year}-${month}-${day}`;

            // 오늘 날짜 이전의 날짜가 선택되었는지 확인
            if (formattedSelectedDate < todayString) {
                alert("The selected date has already passed. Kindly choose another date.");
                return;
            }

            try {
                // const response = await axios.get("https://kids-cafe-rm9g.onrender.com/api/bookings/check", {
                const response = await axios.get(`${process.env.VUE_APP_API_BASE}/api/bookings/check`, {
                    params: {
                        date: selectedDate,  // DD-MM-YYYY 형식
                        time: this.selectedTime
                    }
                });

                if (response.data.available) {
                    console.log("예약 가능한 시간입니다.");
                } else {
                    alert("This time is already booked! Please select a different time, or contact us!");
                    this.selectedTime = ""; // 초기화
                }
                } catch (error) {
                console.error("예약 확인 중 오류:", error);
                alert("An error occurred while confirming the reservation.");
                }
            }
    },
  computed: {
        displayWeekdayPrice() {
            console.log("this.selectedroom.RoomPriceWeekday:", this.selectedroom.RoomPriceWeekday);
            console.log("this.roomId.slice(-2) @#$#$@#$# :", this.roomID);
            const suffix = this.roomID.slice(-2);
            if (suffix === '01') {
                return `weekday price: $${this.selectedroom.RoomPriceWeekday}`;
            } else if (suffix === '02') {
                return `(17:00~20:00 price): $${this.selectedroom.RoomPriceWeekday}`;
            }
        },
        displayWeekendPrice() {
            const suffix = this.roomID.slice(-2);
            if (suffix === '01') {
                return `weekend price: $${this.selectedroom.RoomPriceWeekend}`;
            } else if (suffix === '02') {
                return `(16:00~19:00 price): $${this.selectedroom.RoomPriceWeekend}`;
            } 
        },
        filteredTimeOptions() {
            // roomId의 끝자리 추출
            console.log("!@#!$@#$#$^%^%$ room id @#@$#!%!#: =", this.roomID);
            const roomSuffix = this.roomID.slice(-2);
            // 해당하는 시간 옵션 반환, 없을 경우 빈 배열 반환
            return this.timeOptions[roomSuffix] || [];
        },
        formattedDate() {
            if (!this.selectedDate) return "";
            const date = new Date(this.selectedDate);
            const day = String(date.getDate()).padStart(2, '0');
            const month = String(date.getMonth() + 1).padStart(2, '0'); // 월은 0부터 시작하므로 +1
            const year = date.getFullYear();
            return `${day}-${month}-${year}`; // ✅ DD/MM/YYYY 형식으로 변환
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

        }, 
        // 날짜 선택
        selectedPrice(){
            // if (!this.selectedTime || !this.selectedDate) 
            //     return alert("Please select date and time for the party.");
            console.log("previously selected Date ::::" + this.selectedDate);

            const date = new Date(this.selectedDate);
            const day = date.getDay(); // 0: 일요일, 6: 토요일
            const isWeekend = (day === 0 || day === 6);

            //지금 선택한 날짜
            console.log("date --= ===== ", day);
            console.log("isweekend --= ===== ", isWeekend);
            const suffix = this.roomID.slice(-2);
            //const day = this.selectedDay; // Accessing the selectedDay computed property
            const time = this.timeOptions;

            if (suffix === '01') {
                if(isWeekend){
                    return this.selectedroom.RoomPriceWeekend;
                }else{
                    return this.selectedroom.RoomPriceWeekday;
                }
            } else if (suffix === '02') {
                if(this.selectedTime === '16:00~19:00'){
                    return this.selectedroom.RoomPriceWeekend;
                }else if(this.selectedTime === '17:00~20:00'){
                    return this.selectedroom.RoomPriceWeekday;
                }
            }
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

.room-header {
  display: flex;
  align-items: baseline; /* 글자 수직 정렬 자연스럽게 */
  gap: 10px;              /* 글자 간 여백 */
}

.room-session {
  color: #41ba80;
  font-size: 1rem;
  margin: 0;
}


.multiline {
  white-space: pre-line;
  text-align: left;
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

.submit-button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
  opacity: 0.6;
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
    flex-direction: column; /* ← 세로로 배치 */
    align-items: center;     /* 가운데 정렬 */
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
    width: 80%; /* 또는 70% */
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
    /*width: 80%;  */
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

@media (max-width: 1024px) {
  .main-card {
    flex-direction: column; /* 세로로 배치 */
    align-items: center;
  }

  .selected-room-card,
  .main-time-pick {
    width: 90%; /* 양쪽 여백 확보 */
    margin: 10px 0; /* 위아래 여백 */
  }

  .submit-button {
    width: 100%; /* 전체 너비 사용 */
    font-size: 18px;
  }

  .time-card {
    width: 100px;
    font-size: 14px;
  }
}

@media (max-width: 768px) {
  .selected-room-card,
  .main-time-pick {
    width: 95%;
    padding: 12px;
  }

  .room-name {
    font-size: 20px;
  }

  .happy-birthday {
    width: 80px;
  }

  .time-card {
    width: 90px;
    height: auto;
    padding: 8px;
    font-size: 13px;
  }

}

</style>