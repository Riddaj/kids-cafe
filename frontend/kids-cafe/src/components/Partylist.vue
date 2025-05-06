<template>
  <div class="party-table-wrapper">
    <h1>Party list</h1>
    <div class="table-container">
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
          </tr>
        </thead>
        <tbody>
          <tr v-for="(party, index) in parties" :key="index">
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
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import axios from 'axios'; // axios를 import 추가

export default {
  data() {
        return {
            parties:[],
            branchID: this.$route.params.branchID,
              // URL 파라미터에서 branchID 가져오기
        };
    },
  mounted(){
    this.fetchParty();
  },
  methods: {
        // 날짜 문자열을 Date 객체로 변환
      parseDate(str) {
        const [day, month, year] = str.split('-');
        return new Date(`${year}-${month}-${day}`);
      },
      async fetchParty() {
        console.log("Branch ID:", this.branchID);  // 값이 제대로 있는지 확인
        try {
          const response = await axios.get(`http://localhost:8081/api/get-party/${this.branchID}`); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

        //this.parties = response.data.parties;

          // 정렬: 날짜 오름차순
        this.parties = response.data.parties.sort((a, b) => {
          return this.parseDate(a.Partydate) - this.parseDate(b.Partydate);
        });

          // 전체 응답 객체 찍어보기
          console.log("### 전체 response 객체 ### :", response);
          console.log("### parties data ### :", response.data.parties);

          //console.log("### 전체 response 객체 ### :", response);
          console.log("### parties data 나오라고 ### :", response.data.parties);
        } catch (error) {
          console.error("#### Error fetching parties ##### :", error);
        }
      }
    }
}
</script>

<style scoped>
.party-table-wrapper {
  padding: 40px;
  max-width: 100%;
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
  padding: 12px 16px;
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