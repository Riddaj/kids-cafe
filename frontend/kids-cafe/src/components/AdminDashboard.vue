<template>
  <div id="app">
    <BookingBar/>
      <div class="party-table-wrapper">
        <h1>Admin Dashboard</h1>
        <div class="add-party">
          <div class="button-container" style="justify-content: flex-end; /* 오른쪽 정렬 */">
          <router-link :to="`/add_party`">
                    <button type="submit" class="submit-button">
                      Add Party
                    </button>
          </router-link>
        </div>
        </div>
        <h2>Party list</h2>
        <div class="table-container">
          <h2  style="text-align: left;">✅ Confirmed Bookings</h2>
          <div v-for="(parties, branch) in groupedConfirmedParties" :key="branch">
            <h3 :class="['branch-heading', branch]">{{ branch.toUpperCase() }}</h3>
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
                <th>Options</th>
                <!-- <th>Balloon Theme</th> -->
                <th>Food</th>
                <th>Additional Note</th>
                <th>Deposit Paid</th>
                <th>Confirm</th> <!-- ← 이게 버튼 열 제목 -->
                <th>Done</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="party in parties" :key="party.PartyID">
              <!-- <tr v-for="party in confirmedParties" :key="party.PartyID">  -->
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
                <td>{{ party.balloonDecorationsChecked ? '🎈' : '' }}</td>
                <!-- <td>{{ party.balloonDecorationsTheme }}</td> -->
                <td>{{ party.selected_food }}</td>
                <td>{{ party.addRequirement || '—' }}</td>
                <td>
                  <input type="checkbox" :checked="true" disabled />
                </td>
                <td>Confirmed</td>
                <td>
                  <button class="delete-button" @click="deleteParty(party)">🗑 Delete</button>
                </td>
              </tr>
            </tbody>
          </table>
          </div>
          <div style="margin-top: 80px;"></div>
          <h2 style="text-align: left;">📋 Unconfirmed Bookings</h2>
          <div v-for="(parties, branch) in groupedUnconfirmedParties" :key="branch">
            <h3 :class="['branch-heading', branch]">{{ branch.toUpperCase() }}</h3>
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
                <!-- <th>Payment</th> -->
                <th>Room Price</th>
                <th>Options</th>
                <!-- <th>Balloon Theme</th> -->
                <th>Food</th>
                <th>Additional Note</th>
                <th>Deposit Paid</th>
                <th>Deposit Image</th>
                <th>Confirm</th> <!-- ← 이게 버튼 열 제목 -->
              </tr>
            </thead>
            <tbody>
              <tr v-for="party in parties" :key="party.PartyID">
              <!-- <tr
              v-for="(party, index) in unconfirmedParties"
              :key="party.PartyID"
              :class="{ confirmed: confirmedRows.includes(index) }"
              >  -->
                <td>{{ party.kid_name }}</td>
                <td>{{ party.BranchID }}</td>
                <td>{{ party.partyroom_name }}</td>
                <td>{{ party.Partydate }}</td>
                <td>{{ party.Partytime }}</td>
                <td>{{ party.owner_name }}</td>
                <td>{{ party.owner_phone }}</td>
                <!-- <td>{{ party.payment_method }}</td> -->
                <td>{{ party.partyroom_price }}</td>
                <td>{{ party.balloonDecorationsChecked ? '🎈' : '' }}</td>
                <td>{{ party.selected_food }}</td>
                <td>{{ party.addRequirement || '—' }}</td>
                <td>
                  <input type="checkbox" v-model="paidStatus[party.PartyID]" />
                </td>
                <td>
                  <img v-if="party.imageUrl" :src="party.imageUrl" alt="party image" class="party-img" />
                  <!-- 
                    <img v-if="party.deposit_image_url" :src="party.deposit_image_url" alt="party image" class="party-img" />
                  -->
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
import { getStorage, ref, getDownloadURL } from "firebase/storage";
import { getAuth, onAuthStateChanged  } from "firebase/auth";

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
    //this.fetchParty();
    const auth = getAuth();
    onAuthStateChanged(auth, (user) => {
      if (user) {
        this.fetchParty();
      } else {
        console.warn("로그인 필요");
      }
    });
  },
  computed: {
    groupedConfirmedParties() {
      return this.groupByBranch(this.confirmedParties);
    },
    groupedUnconfirmedParties() {
      return this.groupByBranch(this.unconfirmedParties);
    }
  },
  methods: {  
      groupByBranch(parties) {
        const grouped = {};
        for (const party of parties) {
          if (!grouped[party.BranchID]) {
            grouped[party.BranchID] = [];
          }
          grouped[party.BranchID].push(party);
        }
        return grouped;
      },
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
          const api = process.env.VUE_APP_API_BASE;

          const response = await axios.get(`${api}/api/get-party`); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능
          //const response = await axios.get(`https://kids-cafe-rm9g.onrender.com/api/get-party`); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

          //this.parties = response.data.parties;

          // 오늘 날짜
          const today = new Date();
          today.setHours(0, 0, 0, 0); // 시간 부분 제거

          // 필터링: 오늘 이전 날짜는 제외
          const filteredParties = response.data.parties.filter(party => {
            const partyDate = this.parseDate(party.Partydate);
            return partyDate >= today;
          });

        const auth = getAuth();
        const user = auth.currentUser;
        //////이미지
        //if (user) {
        //const storage = getStorage();  
          for (const party of filteredParties) {
            console.log("🎯 deposit_filename 확인:", party.deposit_filename); // 👈 이거 추가
            if (party.deposit_filename) {
              try {
                const response = await axios.get(`${api}/api/signed-url`, {
                  params: { filename: party.deposit_filename },
                });
                party.imageUrl = response.data.signed_url;

                //const fileRef = ref(storage, party.deposit_filename); // 👉 필드명을 deposit_filename으로 변경
                //party.imageUrl = await getDownloadURL(fileRef);       // 👉 직접 URL 생성
                //const filePath = decodeURIComponent(
                //  party.deposit_image_url.split("/o/")[1].split("?")[0]
                //);
                //const fileRef = ref(storage, filePath);
                //party.imageUrl = await getDownloadURL(fileRef);
              } catch (error) {
                console.warn("❌ 이미지 URL 가져오기 실패:", error);
                party.imageUrl = null;
                //console.error("이미지 URL 가져오기 실패:", error);
                //party.imageUrl = null;
              }
            } else {
              party.imageUrl = null;
            }
          }
        //}else {
        //console.warn("🚫 로그인된 사용자 아님");
      //}

        // 🎯 확정 여부에 따라 나누기~
        this.confirmedParties = [];
        this.unconfirmedParties = [];

        filteredParties.forEach(party => {
          if (party.is_confirmed) {
            this.confirmedParties.push(party);
          } else {
            this.unconfirmedParties.push(party);
          }
        });

        // ✅ 정렬: 날짜 오름차순 + BranchID 알파벳 오름차순
        const sortByDateAndBranch = (a, b) => {
          const dateA = this.parseDate(a.Partydate);
          const dateB = this.parseDate(b.Partydate);
          if (dateA.getTime() !== dateB.getTime()) {
            return dateA - dateB;
          } else {
            return a.BranchID.localeCompare(b.BranchID);
          }
        };

        this.confirmedParties.sort(sortByDateAndBranch);
        this.unconfirmedParties.sort(sortByDateAndBranch);

        // ✅ 확인용 로그
        console.log("✅ 정렬된 확정 파티:", this.confirmedParties);
        console.log("✅ 정렬된 미확정 파티:", this.unconfirmedParties);

        // ✅ 체크 표시용 index 저장 (선택사항)
        this.confirmedRows = this.confirmedParties
          .map((party, index) => index);
        

        } catch (error) {
          console.error("#### Error fetching parties ##### :", error);
        }

        
      },
      async confirmParty(party, index) {
        this.confirmedRows.push(index);

        const api = process.env.VUE_APP_API_BASE;

        try {
          //const response = await axios.post("https://kids-cafe-rm9g.onrender.com/api/confirm-party", {
          const response = await axios.post(`${api}/api/confirm-party`, {
            party_id: party.PartyID
          });
          alert("✅ This party reservation is now confirmed!");
          console.log("✅ 확정됨:", response.data);
        } catch (error) {
          console.error("❌ 서버 업데이트 실패:", error);
        }
      },
      async deleteParty(party) {
        if (!confirm(`Delete ${party.kid_name}'s booking?`)) return;

        const api = process.env.VUE_APP_API_BASE;

        try {
          await axios.delete(`${api}/api/delete-party`, {
            params: {
              party_id: party.PartyID
            }
          });

          // 삭제 후 UI에서 제거
          this.confirmedParties = this.confirmedParties.filter(p => p.PartyID !== party.PartyID);
          alert("✅ Party deleted successfully.");
        } catch (error) {
          console.error("❌ Error deleting party:", error);
          alert("❌ Failed to delete the party.");
        }
      }

    }
}
</script>

<style scoped>
.party-img {
  max-width: 150px;
  max-height: 200px;
  border-radius: 6px;
}


.branch-heading {
  font-size: 1.2rem;
  font-weight: bold;
  margin-top: 40px;
  padding: 10px;
  border-radius: 6px;
  color: #333;
}

.branch-heading.burwood {
  background-color: #fff3b0; /* 연노랑 */
  width: 30%;
  justify-content: center;  /* 버튼을 오른쪽 정렬 */
}

.branch-heading.hornsby {
  background-color: #cce5ff; /* 연파랑 */
  width: 30%;
  justify-content: center;  /* 버튼을 오른쪽 정렬 */
}

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

.delete-button{
  background-color: #ff6666;
  color: white;
  border: none;
  padding: 6px 10px;
  border-radius: 4px;
  cursor: pointer;
}

.delete-button:hover {
  background-color: #cc0000;
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

@media (max-width: 1024px) {
  .party-table-wrapper {
    padding: 10px;
  }

  .party-table th,
  .party-table td {
    font-size: 13px;
    padding: 6px;
  }

  .submit-button {
    width: 100%;
    font-size: 16px;
  }
}

@media (max-width: 768px) {
  h1 {
    font-size: 1.4rem;
    text-align: center;
  }

  h2 {
    font-size: 1.1rem;
    text-align: center;
  }

  .party-table-wrapper {
    padding: 8px;
    overflow-x: auto; /* 테이블 가로 스크롤 */
  }

  .party-table {
    min-width: 900px; /* 테이블이 너무 작아지지 않도록 */
  }

  .party-table th,
  .party-table td {
    font-size: 12px;
    padding: 5px 6px;
  }

  .button-container {
    flex-direction: column;
    align-items: center;
  }

  .submit-button {
    width: 100%;
    font-size: 15px;
  }
}

</style>