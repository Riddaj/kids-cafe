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
                    <h1>Contact Information</h1>
                    <img src="/images/banquet.png" alt="food" class="food-image">
                    <div class="room-detail">
                        <!-- <p>Room ID: {{ bookingDetails.roomID }}</p> -->
                        <p>Room Name: {{ bookingDetails.roomName }}</p>
                        <p>Selected Date: {{ formattedDate }}</p>
                        <!-- <p>Day of Week: {{ formattedDate.dayofweek }}</p>
                        <p>Selected Time: {{ bookingDetails.selectedTime }}</p> -->
                        <p>Room price: {{ bookingDetails.selectedPrice }}</p>
                    </div>
                </div>
                <div class="client-info">
                    <div>
                        <label>Party Kid's Name: </label><input v-model="partyKidName" type="text" class="info-detail-input" />
                    </div>
                    <div>
                        <label>Party owner's Name: </label><input v-model="partyOwnerName" type="text" class="info-detail-input" />
                    </div>
                    <div>
                        <label>Party owner's Phone number: </label><input v-model="partyOwnersPhone" type="text" class="info-detail-input" />
                    </div>
                    <div class="button-container">
                        <router-link :to="`/book_a_party/booking-confirm`">
                            <button type="submit" class="submit-button" @click="submitBooking">Next</button>
                        </router-link>
                    </div>
                </div>
            </div>
        </div>
</template>

<script>
import axios from 'axios'; // axios를 import 추가
import CurrentTime from '../components/CurrentTime.vue';
import { ref, computed, watch } from "vue"; // ✅ watch 추가

export default{
    components:{
        CurrentTime
    },
    data(){
        return{
            bookingDetails:{
                roomID: "",
                roomName: "",
                selectedDate: "",
                selectedTime: "",
                selectedPrice: "",
                selectedFoodOptions: [],
                selectedroom: {}
            },
            partyKidName:"",
            partyOwnerName:"",
            partyOwnersPhone:""
        };
    },
    mounted(){
            // ✅ bookingDetails가 undefined가 아닌지 확인하고 초기화
        if (!this.bookingDetails) {
            this.bookingDetails = {};
        }

        console.log('####@@@@ client info  @@@@####:', this.$route);  // 라우터 정보를 출력하여 문제를 진단
        this.bookingDetails.roomID = this.$route.params.roomID || this.$route.query.roomID || "";
        this.bookingDetails.branchID = this.$route.params.branch_id || this.$route.query.branch_id || ""; 
        this.bookingDetails.roomName = decodeURIComponent(this.$route.query.roomName) || ""; // query string에서 가져오기
        this.bookingDetails.selectedDate = this.$route.params.selectedDate  || this.$route.query.selectedDate || "";
        this.bookingDetails.selectedTime = this.$route.params.selectedTime || this.$route.query.selectedTime || "";
        this.bookingDetails.selectedPrice = this.$route.params.selectedPrice || this.$route.query.selectedPrice || "";
        
        // ✅ selectedFoodOptions 가져오기
        if (this.$route.query.selectedFoodOptions) {
            try {
                this.bookingDetails.selectedFoodOptions = JSON.parse(this.$route.query.selectedFoodOptions);
            } catch (error) {
                console.error("Error parsing selectedFoodOptions:", error);
                this.bookingDetails.selectedFoodOptions = [];
            }
        } else {
            this.bookingDetails.selectedFoodOptions = [];
        }

        console.log('#### Room ID #### :', this.bookingDetails); // 이 값이 정상적으로 출력되는지 확인
        //const roomID = this.roomID;  // URL 파라미터에서 roomID 추출
        this.fetchSelectedroomData(this.bookingDetails.roomID);
        console.log("나와롸 ################# ", this.bookingDetails.roomID);
    },
    computed:{
        formattedDate() {
        if (!this.bookingDetails.selectedDate) return "";
        
        // URL에서 받은 날짜를 파싱
        const dateStr = decodeURIComponent(this.bookingDetails.selectedDate);
        const date = new Date(dateStr);

        if (isNaN(date.getTime())) {
            console.error("Invalid Date Format:", dateStr);
            return "Invalid Date";
        }

        // 날짜 포맷 변환 (예: Apr 24 2025(Thu))
        const options = { month: "short", day: "2-digit", year: "numeric", weekday: "short" };
        return date.toLocaleDateString("en-US", options).replace(",", ""); 
        }
    },
    methods:{
        async fetchSelectedroomData(roomId) {
        console.log("📌 Axios 요청 보냄 - room_Id:", this.bookingDetails.roomID);
        console.log("📌 Axios 요청 보냄 - branch_id:", this.branchID);
        console.log("📌 Axios 요청 보냄 - room_name:", this.bookingDetails.roomName);
        
            try {
                const response = await axios.get(`http://localhost:8081/api/selectedroom/${roomId}`
                , {params : {
                    room_id: this.roomID,
                    branch_id: this.branchID, 
                    room_name: this.bookingDetails.roomName}
                }); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

                this.selectedroom = response.data.partyroom;  // 받아온 데이터를 partyroom에 저장
                console.log("### selected data 언제 나오냐고 ### :", response.data);
            } catch (error) {
                console.error('Error fetching selected room data:', error);
            }
        },
        selectTime(option) {
        this.selectedTime = option;  // 선택된 옵션을 저장
        console.log("선택된 시간:", this.selectedTime); // 콘솔로 확인
        },
        async submitBooking() {
            // 제출할 데이터 객체 생성
            const bookingData = {
                partyroom_id: this.bookingDetails.roomID,
                partyroom_name: this.bookingDetails.roomName,
                partytime: this.formattedDate, // 수정된 부분
                total_price: this.bookingDetails.selectedPrice,
                partytime: this.bookingDetails.selectedTime,
                kid_name: this.partyKidName,
                owner_name: this.partyOwnerName,
                owner_phone: this.partyOwnersPhone,
            };

            console.log("📌📌📌📌 보내는 데이터 확인:", bookingData); // 🔴 확인용 로그

            try {
                // Go 백엔드로 POST 요청 보내기
                const response = await axios.post('http://localhost:8080/api/save-party', bookingData);
                alert("Booking completed🎉🎈");
                console.log('Booking successful', response.data);
                // 다음 페이지로 이동
                this.$router.push('/book_a_party/booking-confirm');
            } catch (error) {
                console.error('Error booking party:', error);
            }
        },
    }
}
</script>

<style scoped>
.main-card{
    display: flex;
    justify-content: center; /* 전체 가운데 정렬 */
    align-items: center;
    margin-top: 20px; /* 위쪽 여백 추가 */
    padding: 20px; /* 좌우 여백 추가 */
}

.room-detail p{
    color: black;
    text-align: left;
}

.client-info{
    width: 60%;
    padding: 20px;  /* 박스 내부 여백 */
    display: block;
    justify-content: center;
    text-align: center; /* 버튼을 가로로 중앙 정렬 */
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.selected-room-card{
    border-radius: 10px;
    padding: 15px;
    margin-right: 50px;
    width: 350px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    text-align: center;
    transition: transform 0.2s, box-shadow 0.2s;
}

.food-image{
    width: 120px;
    height: auto; /* 높이는 비율에 맞게 자동 조정 */
    object-fit: contain; /* 이미지 비율을 유지하면서 가능한 공간에 맞추기 */
    display: block;
    margin: 0 auto; /* 이미지 중앙 정렬 %이거 하니까 적용됨% */
}

.info-detail-input {
        /*flex-grow: 1;                 입력 필드는 가용 공간을 모두 차지 */
        width: 60%;
        margin-bottom: 20px;
        padding: 5px;                /* 입력 필드에 패딩 추가 */
        border: 1px solid #ccc;      /* 입력 필드 테두리 */
        border-radius: 2px;          /* 입력 필드 테두리 둥글게 */
        box-sizing: border-box;  /* box-sizing 설정 */
    }

.info-detail label {
        margin-left: 15%;
        width: 30%;                 /* 레이블의 고정 너비 설정 */
        text-align: left;           /* 레이블 텍스트를 오른쪽 정렬 */
    }




.client-info {
    display: block;               /* 내부 요소들 역시 flexbox */
    align-items: left;         /* 레이블과 입력 필드 수직 중앙 정렬 */
    gap: 10px;                   /* 레이블과 입력 필드 간의 간격 */
}

/* 각 입력 항목을 감싸는 div 스타일 */
.client-info div {
    display: flex;
    flex-direction: row;  /* label과 input을 가로로 정렬 */
    align-items: center;  /* label과 input 세로 정렬 */
    gap: 10px;  /* label과 input 사이에 간격 추가 */
}

/* 레이블 스타일 */
.client-info label {
  font-weight: bold;
  margin-bottom: 5px;  /* 레이블과 입력란 간격 */
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

.current-time{
    display: flex;           /* flexbox 활성화 */
    width: 100vw;
    justify-content: flex-end;
    padding-right: 50px;;
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

.submit-button{
    background-color: #ffb3b3; /* 버튼 배경 색상 */
    color: white; /* 텍스트 색상 */
    text-decoration: none; /* 링크 스타일 제거 */
    padding: 10px 20px; /* 버튼 내부 여백 */
    border-radius: 5px; /* 둥근 모서리 */
    font-size: 16px; /* 글자 크기 */
    display: inline-block; /* 버튼 형태 유지 */
    text-align: center;
    margin-top: 30px; /* 버튼이 자동으로 글자 아래에 배치되도록 설정 */
}

.submit-button:hover{
    background-color: #6699ff; /* 버튼 배경 색상 */
}

.button-container {
    display: flex;
    justify-content: center;  /* 버튼을 오른쪽 정렬 */
    text-align: center; /* 버튼을 가로로 중앙 정렬 */
    margin-top: 20px; /* 버튼과 테이블 사이에 간격 추가 */
    margin-bottom: 30px;
}
</style>