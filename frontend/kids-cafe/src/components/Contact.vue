<template>
    <div id="app" @click="shootConfetti"> 
        <BookingBar/>
        <component 
        :is="showBookingInfo ? 'SelectBranch' : 'div'"  
        :branches="branches" 
        v-if="branches.length > 0">
            <!-- Contact Info (예약 정보로 변경 가능) -->
            <template v-if="!showBookingInfo">
                <!-- <button class="button" @click="shootConfetti">
                  🎉 Like
                </button> -->
                
                <h1 >Contact us</h1>
                <div class="contact-info">

                <div class="contact-info">
                <div class="contact-info-item">
                  <div class="contact-info-icon">
                    <i class="fas fa-home"></i>
                  </div>
                  
                  <div class="contact-info-content">
                    <h4>Address</h4>
                    <p>{{ currentBranch.location }}</p>
                  </div>
                </div>
                
                <div class="contact-info-item">
                  <div class="contact-info-icon">
                    <i class="fas fa-phone"></i>
                  </div>
                  
                  <div class="contact-info-content">
                    <h4>Branch</h4>
                    <p>{{ currentBranch.branch_call }}</p>
                    <h4>Phone</h4>
                    <p>{{ currentBranch.phone }}</p>
                  </div>
                </div>
                
                <div class="contact-info-item">
                  <div class="contact-info-icon">
                    <i class="fas fa-envelope"></i>
                  </div>
                  
                  <div class="contact-info-content">
                    <h4>Email</h4>
                  <p>{{ currentBranch.email }}</p>
                  </div>
                </div>
              </div>
                    <!-- <li class="branch-item">
                      <div class="contact-info-icon">
                      <i class="fas fa-home"></i>
                    </div>
                    <strong>{{ currentBranch.branch_name }}</strong><br>
                    Email: {{ currentBranch.email }}<br>
                    Phone: {{ currentBranch.phone }}<br>
                    Store Phone Number: {{ currentBranch.branch_call }}
                  </li> -->
                    <!-- <ul v-if="branches && branches.length > 0">
                        <li v-for="branch in branches" :key="branch.id" class="branch-item">
                            <strong>{{ branch.branch_name }}</strong><br>
                            Email: {{ branch.email }}<br>
                            Phone: {{ branch.phone }}<br>
                            Store Phone Number: {{ branch.branch_call }}
                        
                        </li>
                        <br>
                    </ul> -->
                </div>
            </template>
        </component>
        <!-- ✅ 아래 여백 주기 -->
        <div style="margin-top: 150px;"></div>
        <!-- ✅ Footer 컴포넌트 추가 -->
        <Footer />
    </div>
</template>

<script setup>
import confetti from 'canvas-confetti'

function shootConfetti() {
  confetti({
    particleCount: 150,
    spread: 60
  })
}
</script>


<script>
import CurrentTime from '../components/CurrentTime.vue';
import axios from "axios";
import ConfirmDetail from '../components/ConfirmDetail.vue';
import Partyroom from '../components/Partyroom.vue';
import SelectBranch from '../components/SelectBranch.vue';
import BookingProcess from '../components/BookingProcess.vue';
import BookingBar from '../components/BookingBar.vue';
import Footer from './Footer.vue';
import confetti from 'canvas-confetti';


export default {
    data() {
        return {
            showBookingInfo : false,
            branches:[],
            //branchID: this.$route.params.branchID,
            currentBranch: null, // 현재 선택된 브랜치
        };
    },  
    mounted() {
      this.fetchBranches();
      
    },
 
    components: {
        CurrentTime,
        ConfirmDetail,
        Partyroom,
        SelectBranch,
        BookingProcess,
        BookingBar,
        Footer
    },  
    methods: {
        onClick() {
      // 전역 confetti 사용
      confetti({
        particleCount: 150,
        spread: 60
      });
    },
    async fetchBranches() {
      try {
        const response = await axios.get("https://kids-cafe-rm9g.onrender.com/api/branches"); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능
        //firebase관련 추가로 repsonse.data뒤에 branches 입력
        this.branches = response.data.branches;
        //console.log("### 전체 response 객체 ### :", response);
        console.log("### Branches data 나오라고 ### :", this.branches);
        //현재 브랜치에 해당하는 것만 필터.
        //console.log('BranchID ;:: ' +  ${branchID});
        const branchID = String(this.$route.params.branchID);
        console.log('BranchID ;:: ' + branchID);
        // branchID에 해당하는 데이터만 currentBranch에 저장
        this.currentBranch = this.branches.find(branch => branch.branch_id === branchID);
        console.log(this.currentBranch);
      } catch (error) {
        console.error("#### Error fetching branches ##### :", error);
      }
    }
  },
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap');


html, body, #app {
  width: 100vw;
  height: 100%;
  margin: 0;
  padding: 0;
  overflow-x: hidden;
}

ul {
  list-style-type: none; /* 동그라미 기호 제거 */
}

li {
  list-style-type: none; /* li에 대한 동그라미 기호 제거 (선택 사항) */
}

a{
    text-decoration: none;
}


.contact-info{
    display: flex;
    margin-top: 20px;
    justify-content: center;
    text-align: left;
    color: #ffffff;
    padding: 20px;
    /* box-shadow: 0px 4px 15px rgba(0, 0, 0, 0.08); */
    margin: 0 auto; /* 가로 중앙 정렬 */
    width: 80%; /* 또는 적절한 width 설정 */
}

.branch-item{
    margin-top: 10px;
    margin-bottom: 30px; /* 각 항목 간의 간격을 넓히기 */
    margin:40px;
}



button {
  cursor: pointer;
  font: inherit;
  margin: 0;
  padding: 0;
}

.button {
  background-color: #404663;
  color: #fff;
  border: 0;
  font-size: 2rem;
  font-weight: 400;
  padding: 0.5em 1.25em;
  border-radius: 0.5em;
  z-index: 999;
  position: relative;
  display: flex;
  gap: 0.5em;
  box-shadow:
    0px 1.7px 2.2px rgba(0, 0, 0, 0.02),
    0px 4px 5.3px rgba(0, 0, 0, 0.028),
    0px 7.5px 10px rgba(0, 0, 0, 0.035),
    0px 13.4px 17.9px rgba(0, 0, 0, 0.042),
    0px 25.1px 33.4px rgba(0, 0, 0, 0.05),
    0px 60px 80px rgba(0, 0, 0, 0.07);
}

.button:active {
  transform: scale(1.01);
}


.contact-info-item {
  display: flex;
  margin-bottom: 30px;
}

.contact-info-icon {
  height: 70px;
  width: 70px;
  background-color: #eb6d54;
  text-align: center;
  border-radius: 50%;
  flex-shrink: 0; /* 아이콘이 눌리는 걸 방지 */
}

.contact-info-icon i {
  font-size: 30px;
  line-height: 70px;
}

.contact-info-content {
  margin-left: 20px;
}

.contact-info-content h4 {
  color: #657C6A;
  font-size: 1.4em;
  margin-bottom: 5px;
}

.contact-info-content p {
  color: #657C6A;
  font-size: 1em;
}
</style>
