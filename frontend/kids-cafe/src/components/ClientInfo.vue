<template>
    <div id="app">
        <BookingBar/>
        <!-- booking process (2단계) -->
        <!-- <BookingProcess/> -->
        <!-- 현재 시간 -->
        <CurrentTime/>
        <div class="main-card">
            <div class="selected-room-card">
                <h1>Party Registration Form</h1>
                <img src="/images/banquet.png" alt="food" class="food-image">
                <div class="room-detail">
                    <!-- <p>Room ID: {{ bookingDetails.roomID }}</p> -->
                    <p>Room Name: {{ bookingDetails.roomName }}</p>
                    <p>Selected Date: {{ bookingDetails.selectedDate }}</p>
                    <!-- <p>Day of Week: {{ formattedDate.dayofweek }}</p> -->
                    <p>Selected Time: {{ bookingDetails.selectedTime }}</p>
                    <p>Room price: {{ bookingDetails.selectedPrice }}</p>
                </div>
            </div>
            <div class="client-info-card">
                <form @submit.prevent="submitBooking">
                    <div class="client-info">
                        <div class="contact-info">
                            <div class="section-title">Kid's Information</div>
                            <div class="form-row">
                                <label>Kid's Name: </label><input v-model="kid_name" type="text" class="info-detail-input" />
                                <label>Gender: </label><input v-model="kid_gender" type="text" class="info-detail-input" />
                                <label>Age Turning: </label><input v-model="kid_age" type="text" class="info-detail-input" />
                            </div>
                        </div>
                        <div class="contact-info">
                            <div class="section-title">Booking Contact Information</div>
                            <div class="form-row">
                                <div class="form-group">
                                    <label>Party owner's Name: </label><input v-model="owner_name" type="text" class="info-detail-input" />
                                    <label>Kid's relation: </label><input v-model="kid_relation" type="text" class="info-detail-input" />
                                </div>
                                <div class="form-group">
                                    <label>Phone Number: </label><input v-model="owner_phone" type="text" class="info-detail-input" />
                                    <label>E-mail: </label><input v-model="email" type="text" class="info-detail-input" />
                                </div>
                            </div>
                        </div>
                        <div class="contact-info">
                            <div class="section-title">Optional Service</div>
                            <div class="form-row">
                                <label><input type="checkbox" v-model="balloonDecorationsChecked" class="info-detail-check" />Balloon Decorations</label>
                            </div>
                            <div v-if="balloonDecorationsChecked" class="form-row">
                                <label>Balloon Decorations Theme: </label>
                                <input v-model="balloonDecorationsTheme" type="text" class="info-detail-input" />
                            </div>
                        </div>
                        <div class="contact-info">
                            <div class="section-title">Special Dietary Requirement(Allergy)</div>
                            <div class="form-row">
                                <label><input type="checkbox" value="Halal" v-model="selectedAllergies" class="info-detail-check"/>Halal</label>
                                <label><input type="checkbox" value="Non-Halal" v-model="selectedAllergies" class="info-detail-check"/>Non-Halal</label>
                                <label><input type="checkbox" value="Veggie" v-model="selectedAllergies" class="info-detail-check"/>Veggie</label>
                                <!-- 
                                    <p v-if="Array.isArray(selectedAllergies) && selectedAllergies.length > 0">
                                        {{ selectedAllergies }}
                                    </p>
                                -->
                            </div>
                        </div>
                        <div class="contact-info">
                            <div class="section-title">Additional Requirement</div>
                            <div class="form-row"><input type="text" v-model="addRequirement" class="addRequirement"/></div>
                        </div>
                        <div class="contact-info">
                            <div class="section-title">Deposit Screenshot | Selected Room Deposit : {{ this.bookingDetails.roomDeposit }}</div>
                            <div class="form-row"><input type="file" class="deposit"/></div>
                        </div>
                        <div class="agree-terms">
                            ✨ All deposit payments are non-refundable. If you need to reschedule, at least 4
                            weeks’ notice is required.<br/>
                            ✨ 50% of the deposit can be transferred to store credit if the cancellation is made
                            at least 4 weeks before the event.<br/>
                            ✨ Outside food and drinks are not allowed, except for a birthday cake and gift bags.<br/>
                            ✨ No outside balloon decorations are permitted.<br/>
                            ✨ All children and adults must wear non-slip socks in the venue.<br/>
                            ✨ Any painting, marking, or physical damage will incur a minimum damage fee of $200.<br/>
                            ✨ The hirer(s) and their guests will be held responsible for any costs incurred to
                            repair any damage caused to the venue.<br/>
                            <div class="agree">
                                <label>
                                    <input type="checkbox" v-model="agree_terms" />
                                </label>
                                By signing below, I acknowledge that I have read and understood the Playground
                                Rules and Policies above. I agree to comply with all Playground Rules and Policies and
                                pay for all the listed charges, if applicable.
                            </div>
                        </div>
                        <div class="button-container">
                                <button type="submit" class="submit-button">
                                    Next
                                </button>
                        </div>
                    </div>
                </form>
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
import BookingProcess from '../components/BookingProcess.vue';
import BookingBar from '../components/BookingBar.vue';

export default {
    components:{
        CurrentTime,
        BookingProcess,
        BookingBar,
        Footer
    },
    data(){
        return{
            balloonDecorationsChecked: false, // 체크박스 상태
            balloonDecorationsTheme: '',      // 테마 입력
            bookingDetails:{
                roomID: "",
                roomName: "",
                roomDeposit:"",
                selectedDate: "",
                selectedTime: "",
                selectedPrice: "",
                selectedFoodOptions: [],
                food_price:"",
                selectedroom: {},
            },
            kid_name: '',
            kid_gender: '',
            kid_age: '',
            owner_name: '',
            kid_relation: '',
            owner_phone: '',
            email: '',
            option_service: '',
            selectedAllergies: [], // checkbox는 배열로
        }
    },
    computed:{
    },
    mounted(){
        if (!this.bookingDetails) {
            this.bookingDetails = {};
        }

        console.log('####@@@@ client info  @@@@####:', this.$route);  // 라우터 정보를 출력하여 문제를 진단
        this.bookingDetails.roomID = this.$route.params.roomID || this.$route.query.roomID || "";
        this.bookingDetails.branchID = this.$route.params.branch_id || this.$route.query.branch_id || ""; 
        this.bookingDetails.roomName = decodeURIComponent(this.$route.query.roomName) || ""; // query string에서 가져오기
        this.bookingDetails.roomDeposit =  this.$route.params.roomDeposit || this.$route.query.roomDeposit || ""; 
        this.bookingDetails.selectedDate = this.$route.params.selectedDate  || this.$route.query.selectedDate || "";
        this.bookingDetails.selectedTime = this.$route.params.selectedTime || this.$route.query.selectedTime || "";
        this.bookingDetails.selectedPrice = this.$route.params.selectedPrice || this.$route.query.selectedPrice || "";
        this.bookingDetails.food_price = this.$route.params.food_price || this.$route.query.food_price || "";

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
    },
    methods:{
        //기존 정보 불러오기
        async fetchSelectedroomData(roomId) {
        console.log("📌 Axios 요청 보냄 - room_Id:", this.bookingDetails.roomID);
        
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
                console.log("📌 Axios 요청 보냄 - branch_id:", this.selectedroom.BranchID);
            } catch (error) {
                console.error('Error fetching selected room data:', error);
            }
        },
        selectTime(option) {
            this.selectedTime = option;  // 선택된 옵션을 저장
            console.log("선택된 시간:", this.selectedTime); // 콘솔로 확인
        },
        async submitBooking() {
            try {
            // form에 입력한 데이터 수집
            const bookingData = {
                partyroom_id: this.bookingDetails.roomID,
                branch_id: this.selectedroom.BranchID,
                partyroom_name: this.bookingDetails.roomName,
                partydate: this.bookingDetails.selectedDate,
                partytime: this.bookingDetails.selectedTime,
                partyroom_price: this.bookingDetails.selectedPrice,
                food_price: this.bookingDetails.food_price,
                selected_food: this.bookingDetails.selectedFoodOptions,
                kid_name: this.kid_name,
                kid_gender: this.kid_gender,
                kid_age: Number(this.kid_age),
                owner_name: this.owner_name,
                kid_relation: this.kid_relation,
                owner_phone: this.owner_phone,
                email: this.email,
                option_service: this.option_service,
                special_required: this.selectedAllergies,
                balloonDecorationsChecked: this.balloonDecorationsChecked,
                balloonDecorationsTheme: this.balloonDecorationsTheme,
                addRequirement : this.addRequirement,
                agree_terms: this.agree_terms
            };

            console.log("📦 Booking Data 전송 준비:", bookingData);

            console.log("kid name 말해봐 = ", this.kid_name);
            console.log("📦📦📦 this.selectedroom.BranchID=== ", this.selectedroom.BranchID);
            // 백엔드로 POST 요청
            const response = await axios.post(`http://localhost:8081/api/save-party/${this.selectedroom.BranchID}`, bookingData);

            console.log("✅ 예약 성공:", response.data);

            // 성공 후 다음 페이지로 이동
            //this.$router.push('/book_a_party/booking-confirm');
            //state: bookingData
            sessionStorage.setItem('bookingData', JSON.stringify(bookingData));

            this.$router.push({ name: 'booking-confirm'});
        } catch (error) {
            console.error('❌ 예약 중 오류 발생:', error);
            alert('예약 중 오류가 발생했습니다. 다시 시도해 주세요.');
        }
        }
    }
}
</script>

<style scoped>
html,body {
  color: black;
}

.form-row{
    margin-bottom: 0.2rem;
}

.info-detail-check{
    margin-bottom: 20px;
}

.addRequirement{
    width: 600px;
    height: 150px;
    margin-right: 20px;
    margin-bottom: 20px;
    padding: 5px;                /* 입력 필드에 패딩 추가 */
    border: 1px solid #ccc;      /* 입력 필드 테두리 */
    border-radius: 2px;          /* 입력 필드 테두리 둥글게 */
    box-sizing: border-box;  /* box-sizing 설정 */
}

.info-detail-input {
    /*flex-grow: 1;                 입력 필드는 가용 공간을 모두 차지 */
    width: 150px;
    margin-right: 20px;
    margin-bottom: 20px;
    padding: 5px;                /* 입력 필드에 패딩 추가 */
    border: 1px solid #ccc;      /* 입력 필드 테두리 */
    border-radius: 2px;          /* 입력 필드 테두리 둥글게 */
    box-sizing: border-box;  /* box-sizing 설정 */
    }

.main-card{
    display: flex;
    justify-content: center;
}

.client-info-card{
    border-radius: 10px;
    padding: 15px;
    width: 1000px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    text-align: left;
    transition: transform 0.2s, box-shadow 0.2s;
}

.section-title {
  font-size: 1.25rem;
  font-weight: bold;
  color: #374151;
  margin-bottom: 1rem;
  padding-bottom: 0.5rem;
}

.selected-room-card{
    color: black;
    border-radius: 10px;
    padding: 15px;
    margin-right: 50px;
    width: 350px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    text-align: center;
    transition: transform 0.2s, box-shadow 0.2s;
}

.room-detail p{
    color: black;
    text-align: left;
}

.food-image{
    width: 120px;
    height: auto; /* 높이는 비율에 맞게 자동 조정 */
    object-fit: contain; /* 이미지 비율을 유지하면서 가능한 공간에 맞추기 */
    display: block;
    margin: 0 auto; /* 이미지 중앙 정렬 %이거 하니까 적용됨% */
}

.agreeTerms{
    text-align: left;
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

.agree{
    margin-top: 20px;
}
</style>