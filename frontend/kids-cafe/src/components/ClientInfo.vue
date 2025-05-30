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
                    <p>Room Deposit: {{ bookingDetails.roomDeposit }}</p>
                </div>
            </div>
            <div class="client-info-card">
                <form @submit.prevent="submitBooking">
                    <div class="client-info">
                        <div class="contact-info">
                            <div class="section-title">Kid's Information</div>
                            <div class="form-row">
                                <label>Kid's Name: </label><input v-model="kid_name" type="text" class="info-detail-input" />
                                <label>Gender: </label>
                                <select class="info-detail-input" v-model="kid_gender" >
                                    <option disabled value="">Please select</option>
                                    <option value="Girl">Girl</option>
                                    <option value="Boy">Boy</option>
                                    <option value="Other">Other</option>
                                </select>
                                <!-- <input v-model="kid_gender" type="text" class="info-detail-input" /> -->
                                <label>Age Turning: </label>
                                <select v-model="kid_age" class="info-detail-input">
                                    <option disabled value="">Please select</option>
                                    <option v-for="n in 12" :key="n" :value="n">{{ n }}</option>
                                </select>
                                <!-- <input v-model="kid_age" type="text" class="info-detail-input" /> -->
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
                            <div class="section-title">Optional Service 
                            </div>
                            <div class="form-row">
                                <label class="info-detail-check">
                                    <input type="checkbox" v-model="balloonDecorationsChecked"  />
                                    <span class="checkmark"></span>
                                        Balloon Decorations
                                    </label>
                            </div>
                            <div v-if="balloonDecorationsChecked" class="form-row">
                                <label>Balloon Decorations Theme: </label>
                                <input v-model="balloonDecorationsTheme" type="text" class="info-detail-input" />
                                <p style="font-size: 15px; color: #f0598b;">🎈For more details, we’ll be in contact with you soon.</p>
                            </div>
                        </div>
                        <div class="contact-info">
                            <div class="section-title">Special Dietary Requirement(Allergy)</div>
                            <div class="form-row">
                                <label class="info-detail-check">
                                    <input type="checkbox" value="Halal" v-model="selectedAllergies"/>
                                    <span class="checkmark"></span>
                                    Halal
                                </label>
                                <label class="info-detail-check">
                                    <input type="checkbox" value="Veggie" v-model="selectedAllergies" class="info-detail-check"/>
                                    <span class="checkmark"></span>
                                    Veggie(includes chicken)
                                </label>
                                    <span class="checkmark"></span>
                            </div>
                        </div>
                        <div class="contact-info">
                            <div class="section-title">Additional Requirement</div>
                            <div class="form-row">
                                <textarea class="addRequirement" v-model="addRequirement" ></textarea>
                            </div>
                        </div>
                        <div class="contact-info">
                            <div class="section-title">
                                <div class="deposit-row">Upload Deposit Screenshot
                                    <p style="color: #f0598b; font-size: 20px;">( Room Deposit: ${{ bookingDetails.roomDeposit }} ) </p>
                                </div>
                                <div class="contact-info" style="color: #f0598b; font-size: 15px;">
                                    bank_bsb : 062245<br>
                                    bank_account : 10556992<br>
                                    Account name : Bahareh Mirbagheri<br>
                                    To confirm your booking, the deposit must be verified.
                               <!-- Pay ID: --> 
                                </div>
                                <div style="padding: 10px;"></div>
                                <div class="form-row">
                                    <input type="file" @change="handleFileChange" accept="image/*" />
                                    <p v-if="depositFileName">📎 Selected: {{ depositFileName }}</p>
                                </div>
                            </div>
 
                        </div>

                        <div style="margin-bottom: 50px;"></div>
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
            depositImageFile: null,
            depositFileName: '',
            paymentMethod: '',         // ✅ 초기화되어 있어야 함!
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
            payment_method:''
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
        handleFileChange(event) {
            const file = event.target.files[0];
            if (file) {
                this.depositImageFile = file;
                this.depositFileName = file.name;
            }
        },
        //기존 정보 불러오기
        async fetchSelectedroomData(roomId) {
        console.log("📌 Axios 요청 보냄 - room_Id:", this.bookingDetails.roomID);
        
        console.log("📌 Axios 요청 보냄 - room_name:", this.bookingDetails.roomName);
        
        const api = process.env.VUE_APP_API_BASE;
            try {
                const response = await axios.get(`${api}/api/selectedroom/${roomId}`
                //const response = await axios.get(`https://kids-cafe-rm9g.onrender.com/api/selectedroom/${roomId}`
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

        //Next Button
        async submitBooking() {

            //validation check
            //kid's name
            if (!this.kid_name || this.kid_name.trim() === '') 
            {
                alert('Please enter the child\'s name.');
                return;
            }

            //kid's gender
            if (!this.kid_gender) 
            {
                alert('Please select the kid\'s gender.');
                return;
            }

            //age turning
            if (!this.kid_age) 
            {
                alert('Please select the kid\'s age.');
                return;
            }

            //Party owner's name
            if (!this.owner_name || this.owner_name.trim() === '') 
            {
                alert('Please enter the Owner\'s name.');
                return;
            }

            //Kid's relation
            if (!this.kid_relation || this.kid_relation.trim() === '') 
            {
                alert('Please enter the kid\'s relation.');
                return;
            }

             //phone number
             if (!this.owner_phone || this.owner_phone.trim() === '') 
            {
                alert('Please enter the phone number.');
                return;
            }

            // 호주 전화번호 정규식: 10자리, 04로 시작하는 모바일 번호 또는 0X로 시작하는 유선 번호
            const auPhonePattern = /^0[2-478]\d{8}$/;

            if (!auPhonePattern.test(this.owner_phone.replace(/\s+/g, ''))) {
                alert('Please enter a valid Australian phone number.');
                return;
            }

            //email
            if (!this.email || this.email.trim() === '') 
            {
                alert('Please enter the email.');
                return;
            }

            // 이메일 형식 검사
            const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

            if (!emailPattern.test(this.email.trim())) {
                alert('Please enter a valid email address.');
                return;
            }



            if (!this.agree_terms) {
                alert("Please agree to the terms and conditions before proceeding.");
                return;
            }

            try {
            
                console.log("📌 food_price 타입은:", typeof this.bookingDetails.food_price);
                const formData = new FormData();
                
                //const formData = new FormData();
                // form에 입력한 데이터 수집
                const bookingData = {
                //formData.append('bookingData', JSON.stringify({
                    partyroom_id: this.bookingDetails.roomID,
                    branch_id: this.selectedroom.BranchID,
                    partyroom_name: this.bookingDetails.roomName,
                    partydate: this.bookingDetails.selectedDate,
                    partytime: this.bookingDetails.selectedTime,
                    partyroom_price: Number(this.bookingDetails.selectedPrice),
                    food_price: Number(this.bookingDetails.food_price),
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
                    agree_terms: this.agree_terms,
                    payment_method : this.payment_method
                };
                //}));
                formData.append('bookingData', JSON.stringify(bookingData));

                // ✅ 이미지 파일이 선택되었으면 추가
                if (this.depositImageFile) {
                    formData.append('depositImage', this.depositImageFile);
                }

            console.log("📦 Booking Data 전송 준비:", bookingData);

            console.log("kid name 말해봐 = ", this.kid_name);
            console.log("📦📦📦 this.selectedroom.BranchID=== ", this.selectedroom.BranchID);
            // 백엔드로 POST 요청
            const api = process.env.VUE_APP_API_BASE;
            //const response = await axios.post(`${api}/api/save-party/${this.selectedroom.BranchID}`, 
            const response = await axios.post(`${api}/api/upload-deposit`,
            formData,
                {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
                }
            );

            // 📦 백엔드가 응답으로 보내준 image URL
            const imageURL = response.data.imagePath;

            // 📝 imageURL을 최종 예약정보에 저장하고 싶으면:
            bookingData.deposit_image_url = imageURL || "";

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

.deposit-row {
  display: flex;
  align-items: center;
  gap: 12px;
}

.deposit-row p {
  color: #f0598b;
  font-size: 20px;
  margin: 0;
}


.form-row{
    margin-bottom: 0.2rem;
}

.info-detail-check{
    margin-bottom: 20px;
}

/**
additional reqquirements
textarea
 */
.addRequirement {
    width: 600px;
    height: 150px;
    margin-right: 20px;
    margin-bottom: 20px;
    padding: 15px; /* 패딩을 더 주어서 퐁실하게 */
    border: 1px solid #ddd;  /* 부드러운 테두리 */
    border-radius: 15px;      /* 둥글게 */
    box-sizing: border-box;
    font-size: 16px;         /* 좀 더 큰 폰트 */
    line-height: 1.6;        /* 줄 간격을 여유 있게 */
    background-color: #f9f9f9;  /* 밝은 배경색 */
    color: #333;             /* 어두운 텍스트 색상 */
    font-family: 'Arial', sans-serif;
    transition: all 0.3s ease; /* 부드러운 전환 효과 */
}

.addRequirement:focus {
    outline: none;  /* 포커스 시 기본 테두리 없애기 */
    border-color: #88b7d5;  /* 포커스 시 테두리 색상 변경 */
    box-shadow: 0 0 10px rgba(136, 183, 213, 0.6); /* 푸른빛 그림자 */
    background-color: #fff;  /* 포커스 시 배경색 변경 */
}

.addRequirement::placeholder {
    color: #aaa;  /* 플레이스홀더 텍스트 색상 */
    font-style: italic; /* 플레이스홀더 이탤릭 */
}


.info-detail-input {
    width: 150px;
    max-width: 300px;
    padding: 10px 14px;
    margin-right: 20px;
    margin-bottom: 20px;
    
    border: 1px solid #ddd;
    border-radius: 8px;

    font-size: 16px;
    font-family: 'Segoe UI', 'Roboto', sans-serif;
    color: #333;

    background-color: #fafafa;

    box-shadow: 0 1px 3px rgba(0,0,0,0.05);
    transition: all 0.2s ease;

    box-sizing: border-box;
}

.info-detail-input:focus {
    outline: none;
    border-color: #4A90E2;
    box-shadow: 0 0 0 3px rgba(74, 144, 226, 0.2);
    background-color: #fff;
}

/**
체크박스 */
.info-detail-check {
    display: inline-flex;
    align-items: center;
    cursor: pointer;
    font-size: 16px;
    font-family: 'Segoe UI', 'Roboto', sans-serif;
    color: #333;
    margin-bottom: 20px;
    user-select: none;
}

/* 기본 체크박스 숨김 */
.info-detail-check input[type="checkbox"] {
    display: none;
}

/* 가짜 체크박스 */
.info-detail-check .checkmark {
    width: 20px;
    height: 20px;
    background-color: #fafafa;
    border: 2px solid #ddd;
    border-radius: 4px;
    margin-right: 10px;
    margin-left: 10px;
    position: relative;
    transition: all 0.2s ease;
}

/* 체크됐을 때 표시 */
.info-detail-check input[type="checkbox"]:checked + .checkmark::after {
    content: '';
    position: absolute;
    left: 5px;
    top: 1px;
    width: 6px;
    height: 12px;
    border: solid #4A90E2;
    border-width: 0 2px 2px 0;
    transform: rotate(45deg);
}

/* 체크되면 배경 강조 */
.info-detail-check input[type="checkbox"]:checked + .checkmark {
    background-color: #fff;
    border-color: #4A90E2;
    box-shadow: 0 0 0 3px rgba(74, 144, 226, 0.15);
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

/* 태블릿용 (max-width: 1024px 이하) */
@media (max-width: 1024px) {
  .main-card {
    flex-direction: column;
    align-items: center;
  }

  .selected-room-card {
    width: 80%;
    margin: 0 auto 30px;
  }

  .client-info-card {
    width: 90%;
    padding: 20px;
  }

  .info-detail-input {
    width: 100%;
    margin-right: 0;
  }

  .addRequirement {
    width: 100%;
  }

  .button-container {
    justify-content: center;
  }
}

/* 모바일용 (max-width: 768px 이하) */
@media (max-width: 768px) {
  .deposit-row {
    flex-direction: column;
    align-items: flex-start; /* 왼쪽 정렬 */
  }

  .deposit-row p {
    margin: 0; /* 여백 줄이기 */
  }

  .selected-room-card h1 {
    font-size: 1.3rem;
  }

  .food-image {
    width: 80px;
  }

  .info-detail-input {
    font-size: 15px;
    padding: 8px 12px;
  }

  .info-detail-check {
    font-size: 15px;
  }

  .addRequirement {
    height: 120px;
    font-size: 14px;
  }

  .submit-button {
    font-size: 16px;
    padding: 10px 16px;
    width: 100%;
  }

  .client-info-card {
    padding: 15px;
  }

  .section-title {
    font-size: 1.1rem;
  }

  .form-row {
    flex-direction: column;
  }

  .form-group {
    width: 100%;
  }
}

</style>

            <!-- 
                <label class="info-detail-check">
                    <input type="checkbox" value="Vegan" v-model="selectedAllergies" class="info-detail-check"/>
                    <span class="checkmark"></span>
                    Strict Vegan</label>
            -->
            
        <!-- 
            <p v-if="Array.isArray(selectedAllergies) && selectedAllergies.length > 0">
                {{ selectedAllergies }}
            </p>
        -->
<!-- <div class="contact-info" v-if="payment_method === 'deposit and text us'"> -->
    

    <!--</div> -->
<!-- Deposit 파일 업로드 영역 (선택된 경우에만 표시) 
<div class="contact-info" v-if="paymentMethod === 'deposit'">
    <div class="contact-info">
        <div class="section-title">Deposit Screenshot | Selected Room Deposit : {{ this.bookingDetails.roomDeposit }}</div>
        <div class="form-row">
            <input type="file" 
            style="display: none"
            ref="depositFileInput" 
            class="deposit" 
            @change="handleFileChange"/>
            <button class="submit-button" @click="triggerFileInput">Upload Screenshot</button>
            <p v-if="depositFileName">📎 Selected: {{ depositFileName }}</p>
        </div>
    </div>
</div>
-->


                           <!-- <div class="form-row">
                                 
                                    <label>Please let us know your Receipt No: </label><input v-model="receipt_no" type="text" class="info-detail-input" />
                                    
                            </div>-->
                            <!--<div class="form-row">
                                
                                <select v-model="payment_method" class="info-detail-input" style="width: 350px;">
                                 <option disabled value="">-- Please select --</option>   
                                <option value="deposit and text us">Bank Transfer and Send payment screenshot by text</option>-->
                                <!-- 
                                    <option value="pay at cafe and confirm with a staff">Pay at Café and confirm with a staff</option>
                                   
                                </select> 
                            </div>-->