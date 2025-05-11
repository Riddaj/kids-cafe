<template>
    <div>
        <BookingBar/>
            <!-- 메인 사진 and booking button -->
            <div class="wrapper">
                <div><h1 class="company-logo-and-name">Twinkle Kids Cafe</h1></div>
                <div class="background-wrapper">
                    <img src="https://images.squarespace-cdn.com/content/v1/637d8d8a7f609c521ddd5429/1672359448650-N89Q21OUSYRU8ROW18F1/Burwood+Plaza+Max3MB_72DPI_VCLAMedia+%2854+of+101%29.jpg" 
                    alt="Background" class="background-image" />
                </div>
            </div>
            <!-- booking process (2단계) -->
            <BookingProcess/>
            <div>
                <!-- 현재 시간 -->
                <div><CurrentTime/></div>
            </div>
            <!-- ###########      메인 입력 내용        ########### -->
            <div class="main-card">
                <div class="selected-room-card">
                    <h1>Food Options</h1>
                    <img src="/images/banquet.png" alt="food" class="food-image">
                    <div v-if="computedbookingDetails.roomID && computedbookingDetails.roomName" class="room-detail">
                        <!-- <p>{{ bookingDetails }}</p>  -->
                        <!-- <p>Room ID: {{computedbookingDetails.roomID }}</p> -->
                        <p>Room Name: {{ computedbookingDetails.roomName }}</p>
                        <p>Room Deposit: {{ computedbookingDetails.roomDeposit }}</p>
                        <p>Selected Date: {{ formattedDate.date }}</p>
                        <p>Day of Week: {{ formattedDate.dayofweek }}</p>
                        <p>Selected Time: {{ computedbookingDetails.selectedTime }}</p>
                        <p>Room price: {{ computedbookingDetails.selectedPrice }}</p>
                    </div>
                    <!-- <div v-if="Object.keys(selectedroom).length > 0"  class="selectedroom-container">
                        <div class="room-name">{{ selectedroom.RoomName }}</div>
                        <div>{{ selectedroom.Description }}</div>
                    </div> -->
                </div>
                <div class="food-option">
                    <!-- ##### 푸드 옵션 관련 테이블 ###### -->
                    <div class="table-container">
                        <table class="info-detail-input-table">
                            <thead>
                                <tr>
                                <th>Select</th>
                                <th>No</th>
                                <th>Food Name</th>
                                <th>Quantity</th>
                                <th>Package</th>
                                <th>Price</th>
                                <th>Description</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="(foodoption, index) in foodoptions" :key="foodoption.FoodName">
                                <td>
                                    <input 
                                    type="checkbox" 
                                    v-model="selectedfoodoptions"
                                    :value="foodoption.FoodName"
                                    @change="logSelectedOption"
                                    />
                                </td>
                                <td>{{ index + 1 }}</td>
                                <td>{{ foodoption.FoodName }}</td>
                                <td>{{ foodoption.FoodQuantity }}</td>
                                <td>{{ foodoption.FoodPackage === 'true' ? 'Package' : 'Individual' }}</td>
                                <td>${{ foodoption.FoodPrice }}</td>
                                <td>{{ foodoption.FoodDescription }}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div> 
                    <div class="total-amount">
                        total: ${{ totalAmount }}
                    </div>
                    <div class="button-container">
                        <router-link :to="{path: `/book_a_party/client-info`, 
                        query: {
                            roomID: roomID,
                            roomName: roomName,
                            roomDeposit:roomDeposit,
                            selectedDate: formattedDate.date,
                            selectedTime: selectedTime,
                            selectedPrice: selectedPrice,
                            food_price: totalAmount,
                            selectedFoodOptions: JSON.stringify(selectedfoodoptions), // JSON 문자열로 변환
                        }
                        }">
                            <button type="submit" class="submit-button">Next</button>
                        </router-link>
                    </div>
                </div>
            </div>
        <!-- ✅ 아래 여백 주기 -->
        <div style="margin-top: 150px;"></div>
         <Footer/>
    </div>
</template>

<script>
import axios from 'axios'; // axios를 import 추가
import CurrentTime from '../components/CurrentTime.vue';
import Footer from '../components/Footer.vue';
import { useRoute, useRouter } from "vue-router"
//import { ref, computed, watch } from "vue"; // ✅ watch 추가
import { ref, computed, watch, onMounted } from "vue"; 
import BookingProcess from '../components/BookingProcess.vue';
import BookingBar from '../components/BookingBar.vue';

export default {
    components:{
        CurrentTime,
        BookingProcess,
        Footer,
        BookingBar
    },
    data(){
        return{
            foodoptions: [],
            //selectedroom:{}, //배열 대신 객체로 변경
            selectedfoodoptions: [], // 선택된 음식 옵션의 food_id를 저장할 배열
        };
    },
    computed: {
        // computed는 항상 최신 query 값을 반영하여 자동으로 업데이트됨
        computedbookingDetails() {
            return {
                roomID: this.$route.query.roomID || "",
                roomName: this.$route.query.roomName || "",
                roomDeposit: this.$route.query.roomDeposit || "",
                selectedDate: this.$route.query.selectedDate || "",
                selectedTime: this.$route.query.selectedTime || "",
                selectedPrice: this.$route.query.selectedPrice || ""
            };
        },
        formattedDate() {
        if (!this.$route.query.selectedDate) return "";

        const dateObj = new Date(this.$route.query.selectedDate);

        // 날짜 형식 변환 (28-03-2025)
        const formatted = dateObj.toLocaleDateString("en-GB").replace(/\//g, "-");

        // 요일 추출 (Fri)
        const dayOfWeek = dateObj.toLocaleDateString("en-US", { weekday: "short" });

        return {
                date: formatted,
                dayofweek: dayOfWeek
            };
        },
        // 선택된 음식 옵션들의 가격 합산
        totalAmount() {
            let total = 0;
            this.selectedfoodoptions.forEach(foodName => {
                const food = this.foodoptions.find(f => f.FoodName === foodName);
                if (food) {
                    total += food.FoodPrice;
                }
            });
            return total.toFixed(2); // 두 자릿수로 포맷팅
        }
  },
  watch: {
    '$route.query': {
        handler(newQuery) {
            this.computedbookingDetails = { ...newQuery };
        },
        immediate: true, // 처음 로드될 때도 실행
        deep: true // 중첩 객체 감지
    }
  },
    mounted(){
        this.fetchfoodoptions();
        console.log("📌 라우터 정보:", this.$route);
        console.log("📌 전달받은", this.computedbookingDetails);
        console.log("Formatted Date: ", this.formattedDate);
        
        this.roomID = this.$route.query.roomID || "";
        this.roomName = this.$route.query.roomName || "";
        this.roomDeposit = this.$route.query.roomDeposit || "";
        this.selectedTime = this.$route.query.selectedTime || "";
        this.selectedDate = this.$route.query.selectedDate || "";
        this.selectedPrice = this.$route.query.selectedPrice || "";

        // formatting Date 값도 처리
        //this.formattedDate = this.formattedDate || "";
            // 음식 옵션을 가져오는 fetch
    },
    methods:{
        logSelectedOption() {
            console.log('선택된 항목 ID:', this.selectedfoodoptions);
        },
        async fetchfoodoptions() {
            try {
                const response = await axios.get("https://kids-cafe-rm9g.onrender.com/api/foodoptions"); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

                this.foodoptions = response.data.foodoptions;
                //console.log("### 전체 response 객체 ### :", response);
                console.log("### foodoptions data 나오라고 ### :", response.data.foodoptions);
            } catch (error) {
                console.error("#### Error fetching foodoptions ##### :", error);
            }
        },
        fetchUpdatedData(query) {
            console.log("새로운 데이터를 불러오는 중...", query);
            this.computedbookingDetails = {
                roomID: query.roomID || "",
                roomName: query.roomName || "",
                roomDeposit: query.roomDeposit || "",
                selectedDate: query.selectedDate || "",
                selectedTime: query.selectedTime || "",
                selectedPrice: query.selectedPrice || ""
            };
            // API 호출 또는 추가적인 데이터 처리 가능
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
    padding: 50px; /* 좌우 여백 추가 */
}

.room-detail{
    text-align: center;
}
.room-detail p {
    color: black; /* room-detail 내의 p 태그 텍스트 색상도 블랙으로 설정 */
    text-align: left;
}

.selected-room-card{
    color: black; /* 모든 텍스트를 블랙으로 설정 */
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

.table-container {
    width: 100%; /* 테이블 너비를 화면의 80%로 설정 (원하는 비율로 조정 가능) */
    margin: 0 auto; /* 좌우 여백을 자동으로 설정하여 중앙 정렬 */
    padding: 20px; /* 테이블 주변에 20px의 내부 여백 추가 */
    box-sizing: border-box; /* 패딩을 너비에 포함시켜 전체 크기를 조정 */
    
}

.food-option{
    display: block;
    justify-content: center;
    text-align: center; /* 버튼을 가로로 중앙 정렬 */
    /* box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);*/
}
    
/** 푸드 옵션 체크리스트 */
table {
    width: 100%;
    border-collapse: collapse;
    }

th, td {
    border: 1px solid #ccc;
    padding: 8px;
    text-align: left;
    }

th {
    background-color: #c6ecc6;
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
    text-align: center; /* 버튼을 가로로 중앙 정렬 */
    margin-top: 20px; /* 버튼과 테이블 사이에 간격 추가 */
    margin-bottom: 30px;
}
</style>