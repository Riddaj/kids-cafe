<template>
  <div id="app">
    <BookingBar/>
      <div class="party-table-wrapper">
        <h1>Admin Dashboard</h1>
        <h2>Party list</h2>
        <div class="table-container">
          <h2  style="text-align: left;">✅ Confirmed Bookings</h2>
            <table class="party-table">
            <thead>
              <tr>
                <th>Kid's Name</th>
                <th>Branch</th>
                <th>Room type</th>
                <th>Party Date(dd/mm/yyyy)</th>
                <th>Time</th>
                <th>Owner</th>
                <th>Phone</th>
                <th>Payment</th>
                <th>Room Price</th>
                <th>Food Price</th>
                <!-- <th>Balloon Theme</th> -->
                <th>Food</th>
                <th>Additional Note</th>
                <th>Deposit Paid</th>
                <th>Confirm</th> <!-- ← 이게 버튼 열 제목 -->
              </tr>
            </thead>
            <tbody>
              <tr v-for="party in confirmedParties" :key="party.PartyID">
              <!-- 
                <tr
                v-for="(party, index) in parties"
                :key="index"
                :class="{ confirmed: confirmedRows.includes(index) }"
                >
              -->
                <td>{{ party.kid_name }}</td>
                <td>{{ party.BranchID }}</td>
                <td>{{ party.partyroom_name }}</td>
                <td>{{ party.Partydate }}</td>
                <td>{{ party.Partytime }}</td>
                <!-- <td>${{ party.food_price }}</td> -->
                <td>{{ party.owner_name }}</td>
                <td>{{ party.owner_phone }}</td>
                <td>{{ party.payment_method }}</td>
                <td>{{ party.partyroom_price }}</td>
                <td>{{ party.food_price }}</td>
                <!-- <td>{{ party.balloonDecorationsTheme }}</td> -->
                <td>{{ party.selected_food }}</td>
                <td>{{ party.addRequirement || '—' }}</td>
                <td>
                  <input type="checkbox" :checked="true" disabled />
                </td>
                <td>Confirmed</td>
              </tr>
            </tbody>
          </table>
          <div style="margin-top: 80px;"></div>
          <h2 style="text-align: left;">📋 Unconfirmed Bookings</h2>
          <table class="party-table">
            <thead>
              <tr>
                <th>Kid's Name</th>
                <th>Branch</th>
                <th>Room type</th>
                <th>Party Date(dd/mm/yyyy)</th>
                <th>Time</th>
                <th>Owner</th>
                <th>Phone</th>
                <th>Payment</th>
                <th>Room Price</th>
                <th>Food Price</th>
                <!-- <th>Balloon Theme</th> -->
                <th>Food</th>
                <th>Additional Note</th>
                <th>Deposit Paid</th>
                <th>Confirm</th> <!-- ← 이게 버튼 열 제목 -->
              </tr>
            </thead>
            <tbody>
              <tr
              v-for="(party, index) in unconfirmedParties"
              :key="party.PartyID"
              :class="{ confirmed: confirmedRows.includes(index) }"
              >
                <td>{{ party.kid_name }}</td>
                <td>{{ party.BranchID }}</td>
                <td>{{ party.partyroom_name }}</td>
                <td>{{ party.Partydate }}</td>
                <td>{{ party.Partytime }}</td>
                <td>{{ party.owner_name }}</td>
                <td>{{ party.owner_phone }}</td>
                <td>{{ party.payment_method }}</td>
                <td>{{ party.partyroom_price }}</td>
                <td>{{ party.food_price }}</td>
                <td>{{ party.selected_food }}</td>
                <td>{{ party.addRequirement || '—' }}</td>
                <td>
                  <input type="checkbox" v-model="paidStatus[party.PartyID]" />
                </td>
                <td>
                  <button
                    :disabled="!paidStatus[party.PartyID]"
                    @click="confirmParty(party)"
                  >
                    Confirm
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="button-container">
          <router-link :to="`/`">
                    <button type="submit" class="submit-button">go home</button>
          </router-link>
        </div>
      </div>
  </div>
</template>

<script>
import axios from 'axios'; // axios를 import 추가
import BookingBar from '../components/BookingBar.vue';

export default {
  components:{
    BookingBar
  },
  data() {
        return {
            parties:[],
            confirmedParties: [],
            unconfirmedParties: [],
            branchID: this.$route.params.branchID,
              // URL 파라미터에서 branchID 가져오기
            paidStatus: {}, // {index: true/false} 형태로 저장
            confirmedRows: []       // 확정된 행 index 저장
        };
    },
  mounted(){
    this.fetchParty();
    
  },
  methods: {
      // confirmParty(party, index) {
      //   this.confirmedRows.push(index);
      //   console.log("Confirmed party:", party);
      //   // 서버로 보내도 OK
      // },

        // 날짜 문자열을 Date 객체로 변환
      parseDate(str) {
        const [day, month, year] = str.split('-');
        return new Date(`${year}-${month}-${day}`);
      },
      async fetchParty() {
        console.log("Branch ID:", this.branchID);  // 값이 제대로 있는지 확인
        try {
          const response = await axios.get(`http://localhost:8081/api/get-party`); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

          //this.parties = response.data.parties;

          // 오늘 날짜
          const today = new Date();
          today.setHours(0, 0, 0, 0); // 시간 부분 제거

          // 필터링: 오늘 이전 날짜는 제외
          const filteredParties = response.data.parties.filter(party => {
            const partyDate = this.parseDate(party.Partydate);
            return partyDate >= today;
          });

          // 정렬: 날짜 오름차순
        this.parties = filteredParties.sort((a, b) => {
          return this.parseDate(a.Partydate) - this.parseDate(b.Partydate);
        });

        // ✅ 여기에 추가!
        this.confirmedRows = this.parties
          .map((party, index) => (party.is_confirmed ? index : null))
          .filter(index => index !== null);

        console.log("✅ 필터링된 파티 리스트:", this.parties);
          // 전체 응답 객체 찍어보기
          console.log("### 전체 response 객체 ### :", response);
          console.log("### parties data ### :", response.data.parties);

          //console.log("### 전체 response 객체 ### :", response);
          console.log("### parties data 나오라고 ### :", response.data.parties);

          // 데이터를 불러온 후
          this.confirmedParties = [];
          this.unconfirmedParties = [];

          filteredParties.forEach((party, index) => {
            if (party.is_confirmed) {
              this.confirmedParties.push(party);
            } else {
              this.unconfirmedParties.push(party);
            }
          });

          // ✅ 여기서 정렬
          this.confirmedParties.sort((a, b) => b.BranchID.localeCompare(a.BranchID));
          this.unconfirmedParties.sort((a, b) => b.BranchID.localeCompare(a.BranchID));
        

        } catch (error) {
          console.error("#### Error fetching parties ##### :", error);
        }

        
      },
      async confirmParty(party, index) {
        this.confirmedRows.push(index);

        try {
          const response = await axios.post("http://localhost:8081/api/confirm-party", {
            party_id: party.PartyID
          });
          alert("✅ This party reservation is now confirmed!");
          console.log("✅ 확정됨:", response.data);
        } catch (error) {
          console.error("❌ 서버 업데이트 실패:", error);
        }
      },
    }
}
</script>

<style scoped>
input[type="checkbox"][disabled] {
  accent-color: green; /* 브라우저가 지원하면 초록색 체크 */
  cursor: not-allowed;
}

.confirmed {
  background-color: #d6f5d6 !important; /* 연한 초록색 배경 */
}

button[disabled] {
  background-color: #ccc;
  cursor: not-allowed;
}

.button-container {
    display: flex;
    justify-content: center;  /* 버튼을 오른쪽 정렬 */
    text-align: center; /* 버튼을 가로로 중앙 정렬 */
    margin-top: 20px; /* 버튼과 테이블 사이에 간격 추가 */
    margin-bottom: 30px;
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

.party-table-wrapper {
  padding: 20px;
  max-width: 100%;
  margin: 0 auto;     /* 가운데 정렬 */
  overflow-x: auto;
}

.party-table {
  width: 100%;
  border-collapse: collapse;
  font-family: Arial, sans-serif;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.party-table thead {
  background-color: #f97e7e;
  color: white;
}

.party-table th,
.party-table td {
  border: 1px solid #ddd;
  padding: 8px 10px;
  text-align: left;
  font-size: 14px;
}

.party-table tbody tr:nth-child(even) {
  background-color: #f9f9f9;
}

.party-table tbody tr:hover {
  background-color: #ffe6e6;
}
</style>