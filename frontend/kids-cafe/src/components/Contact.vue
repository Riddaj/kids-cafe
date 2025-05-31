<template>
    <div id="app" @click="shootConfetti"> 
        <NavBar/>
        <!-- ✅ 아래 여백 주기 -->
        <div style="margin-top: 200px;"></div>
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
                    </div>
                  </div>

                  <div class="contact-info-item">
                    <div class="contact-info-icon">
                      <i class="fas fa-mobile-alt"></i> <!-- 전화(Phone) 아이콘 -->
                    </div>
                    <div class="contact-info-content">
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
import NavBar from '../components/NavBar.vue';
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
        NavBar,
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

        const api = process.env.VUE_APP_API_BASE;
        const response = await axios.get(`${api}/api/branches`);
        //const response = await axios.get("https://kids-cafe-rm9g.onrender.com/api/branches"); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능
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
    flex-wrap: wrap; /* 줄 바꿈 허용 */
    justify-content: space-around;
    align-items: flex-start;
    margin-top: 20px;
    text-align: left;
    color: #ffffff;
    padding: 40px 20px;
    /* box-shadow: 0px 4px 15px rgba(0, 0, 0, 0.08); */
    margin: 0 auto; /* 가로 중앙 정렬 */
    max-width: 1000px;
    width: 80%; /* 또는 적절한 width 설정 */
    gap: 40px; /* 가로 여백 설정 */
    
}

.branch-item{
    margin-top: 10px;
    margin-bottom: 40px; /* 각 항목 간의 간격을 넓히기 */
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

/*
.contact-info-item {
  display: flex;
  margin-bottom: 30px;
  min-width: 200px;  아이템 최소 너비 (선택사항) 
}
*/
.contact-info-item {
  display: flex;
  align-items: center;
  /*flex: 1 1 300px;
  max-width: 400px;*/
  min-width: 250px;
  max-width: 400px;
}

.contact-info-icon {
  height: 60px;
  width: 60px;
  background-color: #eb6d54;
  border-radius: 50%;
  text-align: center;
  display: flex;               /* ✅ flex로 설정 */
  justify-content: center;     /* ✅ 수평 가운데 */
  align-items: center;         /* ✅ 수직 가운데 */
  font-size: 24px;
  color: white;
  flex-shrink: 0; /* 아이콘이 눌리는 걸 방지 */
}

.contact-info-icon i {
  font-size: 30px;
}

.contact-info-content {
  margin-left: 15px;
}

.contact-info-content h4 {
  color: #657C6A;
  font-size: 1.4em;
  margin-bottom: 5px;
  margin: 0;
}

.contact-info-content p {
  margin: 5px 0 0;
  color: #657C6A;
  font-size: 1em;
}

@media screen and (max-width: 1024px) {
  .contact-info {
    flex-direction: column; /* 세로로 쌓이게 */
    margin-top: 50px; /* ✅ 메뉴바에 안 가리게 충분한 여백 확보 */
    align-items: center;
    padding: 20px 10px;
    width: 90%;
  }

  .contact-info-item {
    flex-direction: row;
    align-items: flex-start;
    width: 100%;
    margin-bottom: 20px;
  }

  .contact-info-icon {
    width: 60px;
    height: 60px;
  }

  .contact-info-icon i {
    font-size: 24px;
    line-height: 60px;
  }

  .contact-info-content {
    margin-left: 15px;
  }

  .contact-info-content h4 {
    font-size: 1.2em;
  }

  .contact-info-content p {
    font-size: 0.95em;
  }
}

@media screen and (max-width: 768px) {
  .contact-info {
    padding: 20px 10px;
    width: 95%;
  }

  .contact-info-icon {
    width: 50px;
    height: 50px;
  }

  .contact-info-icon i {
    font-size: 20px;
    line-height: 50px;
  }

  .contact-info-content h4 {
    font-size: 1.1em;
  }

  .contact-info-content p {
    font-size: 0.9em;
  }

  h1 {
    font-size: 1.4rem;
    text-align: center;
  }
}

</style>
