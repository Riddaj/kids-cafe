<template>
    <div id="app">
        <BookingBar/>
        <h1>Booking Confirmation</h1>
        <div class="table-container">
            <table class="confirmation-table">
                <thead>
                    <tr>
                        <th>Field</th>
                        <th>Details</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>Party Room Name</td>
                        <td>{{ bookingData.partyroom_name || 'N/A' }}</td>
                    </tr>
                    <tr>
                        <td>Selected Date</td>
                        <td>{{ bookingData.partytime || 'N/A' }}</td>
                    </tr>
                    <tr>
                        <td>Total Price(Partyroom + Food)</td>
                        <td>{{ (parseFloat(bookingData.partyroom_price) + parseFloat(bookingData.food_price)).toFixed(2) || 'N/A' }}</td>
                    </tr>
                    <tr>
                        <td>Kid's Name</td>
                        <td>{{ bookingData.kid_name || 'N/A' }}</td>
                    </tr>
                    <tr>
                        <td>Owner's Name</td>
                        <td>{{ bookingData.owner_name || 'N/A' }}</td>
                    </tr>
                    <tr>
                        <td>Phone</td>
                        <td>{{ bookingData.owner_phone || 'N/A' }}</td>
                    </tr>
                    <tr>
                        <td>Selected Foods</td>
                        <td>    {{ Array.isArray(bookingData.selected_food) 
                            ? bookingData.selected_food.join(', ') 
                            : (bookingData.selected_food || 'N/A') }}
                        </td>
                    </tr>
                    <!--<tr>
                        <td>Deposit Image</td>
                        <td>{{ bookingData.deposit_image_url }}</td>
                         <td>{{ bookingDetails.special_required.join(', ') }}</td> 
                    </tr>
                    -->
                </tbody>
            </table>
        </div>
        <div class="contact-msg">
            Congratulations!<br>
            A birthday party always requires a lot of details!<br>
            We will contact you soon to discuss more details.<br>
            Please wait for a moment.<br>
            Thank you.
        </div>
        <div class="button-container">
            <router-link :to="`/home/${bookingData.branch_id}`">
                <button type="submit" class="submit-button">go home</button>
            </router-link>
        </div>
        <!-- ✅ 아래 여백 주기 -->
        <div style="margin-top: 150px;"></div>
         <Footer/>
</div>
</template>

<script>
import BookingBar from '../components/BookingBar.vue';
import Footer from '../components/Footer.vue';

import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const bookingData = ref({});  // 기본값을 빈 객체로 설정

export default {
  components:{
    BookingBar,
    Footer
  },
  name: 'BookingConfirm',
  setup() {
    onMounted(() => {
      const storedData = sessionStorage.getItem('bookingData');
      console.log("storedData ===== ", storedData);
      if (storedData) {
        bookingData.value = JSON.parse(storedData);  // ref 객체의 값을 업데이트
        
      }
    });

    return {
      bookingData,  // ref 객체를 반환하여 템플릿에서 사용
    };
  },
}

</script>

<style scoped>
.table-container{
    display: flex;
    justify-content: center;
    text-align: center; /* 버튼을 가로로 중앙 정렬 */
    overflow-x: auto; /* ✅ 가로 스크롤 허용 (모바일에서 깨지지 않게) */
    /* box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); */
}

/* ✅ 왼쪽 열 너비 조정 */
.confirmation-table td:first-child {
    width: 30%;
}

/** 옵션 체크리스트 */
table {
    width: 80%;
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

.contact-msg{
    margin-top: 20px;
    margin-bottom: 20px;

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

@media (max-width: 768px) {
  .confirmation-table th,
  .confirmation-table td {
    font-size: 13px;
    padding: 6px;
  }

  h1 {
    font-size: 1.5rem;
  }

  .contact-msg {
    font-size: 14px;
    padding: 12px 15px;
    line-height: 1.5;
  }
}


</style>

