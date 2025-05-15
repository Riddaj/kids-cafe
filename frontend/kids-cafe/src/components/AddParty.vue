<template>
    <div class="form-wrapper"> <!-- ✅ 노란 배경을 위한 감싸는 div -->
        <div class="booking-form">
        <h2>Party Booking Form</h2>
    
        <form @submit.prevent="submitBooking">
            <!-- 지점 ID -->
            <div>
             <label>Branch:</label> 
             <!-- <input type="hidden" v-model="branch_id"  /> -->
                <select v-model="branch_id">
                    <option disabled value="">Select</option>
                    <option>burwood</option>
                    <option>hornsby</option>
                </select>
            </div>

            <!-- 아이 이름 -->
            <div>
            <label>Child's Name:</label>
            <input type="text" v-model="kid_name" required />
            </div>
    
            <!-- 아이 성별 -->
            <div>
            <label>Child's Gender:</label>
            <select v-model="kid_gender">
                <option disabled value="">Select</option>
                <option>Boy</option>
                <option>Girl</option>
            </select>
            </div>
    
            <!-- 아이 나이 -->
            <div>
            <label>Child's Age:</label>
            <input type="number" v-model.number="kid_age" min="1" />
            </div>
    
            <!-- 보호자 이름 -->
            <div>
            <label>Parent/Guardian Name:</label>
            <input type="text" v-model="owner_name" />
            </div>
    
            <!-- 보호자 연락처 -->
            <div>
            <label>Phone Number:</label>
            <input type="tel" v-model="owner_phone" />
            </div>
    
            <!-- 이메일 -->
            <div>
            <label>Email:</label>
            <input type="email" v-model="email" />
            </div>
    
            <!-- 보호자와 아이 관계 -->
            <div>
            <label>Relationship to Child:</label>
            <input type="text" v-model="kid_relation" />
            </div>
    
            <!-- 파티 날짜 -->
            <div>
            <label>Party Date:</label>
            <input type="date" v-model="partydate" required />
            <p>Selected Date: {{ formattedPartyDate }}</p>
            </div>

            <!-- 파티룸 이름 -->
            <div>
            <label>Party Room Name:</label>
            <!-- <input type="text" v-model="partyroom_name" required /> -->
            <select v-model="partyroom_name" required>
                <option disabled value="">Please select</option>
                <option value="Party room hiring">Party room hiring</option>
                <option value="Private party">Private party</option>
            </select>
            </div>

            <!-- 파티 시간 -->
            <div>
            <label>Party Time:</label>
            <!-- <input type="time" v-model="partytime" required /> -->
            <select v-model="partytime" required>
                <option disabled value="">Select time</option>
                <option
                v-for="time in availableTimeOptions"
                :key="time"
                :value="time"
                >
                {{ time }}
                </option>
            </select>
            </div>
    

    
            <!-- 파티룸 ID -->
            <div>
            <!-- <label>Party Room ID:</label> -->
            <input type="hidden" v-model="partyroom_id"  />
            </div>
    

    
            <!-- 파티룸 가격 -->
            <div>
            <label>Room Price ($):</label>
            <input type="number" v-model.number="partyroom_price" />
            </div>
    
            <!-- 음식 가격 -->
            <div>
            <label>Food Price ($):</label>
            <input type="number" v-model.number="food_price" />
            </div>
    
            <!-- 음식 옵션 -->
            <div>
            <label>Selected Food (comma-separated):</label>
            <input type="text" v-model="selected_food" placeholder="e.g. nuggets, pizza" />
            </div>
    
            <!-- 알러지 정보 -->
            <div>
            <label>Allergy or Special Requirement:</label>
            <textarea v-model="special_required"></textarea>
            </div>
    
            <!-- 풍선 데코레이션 여부 -->
            <div>
            <label>
                <input type="checkbox" v-model="balloonDecorationsChecked" />
                Include Balloon Decorations
            </label>
            </div>
    
            <!-- 풍선 테마 -->
            <div v-if="balloonDecorationsChecked">
            <label>Balloon Theme:</label>
            <input type="text" v-model="balloonDecorationsTheme" />
            </div>
    
            <!-- 추가 요구사항 -->
            <div>
            <label>Additional Requirements:</label>
            <textarea v-model="addRequirement"></textarea>
            </div>
    
            <!-- 결제 방식 -->
            <div>
            <label>Payment Method:</label>
            <select v-model="payment_method">
                <option disabled value="">Select</option>
                <option value="bank">Bank Transfer</option>
                <option value="card">Credit Card</option>
            </select>
            </div>
    
            <!-- 약관 동의 -->
            <div>
            <label>
                <input type="hidden" v-model="agree_terms" />
            </label>
            <!--I agree to the terms and conditions -->
            </div>
    
            <div class="button-container">
            <button type="submit" class="submit-button">Submit Booking</button>
            </div>
        </form>
        </div>
    </div>
  </template>
  

  <script>
  import axios from 'axios'; // axios를 import 추가

  export default {
    data() {
      return {
        kid_name: '',
        kid_gender: '',
        kid_age: null,
        owner_name: '',
        owner_phone: '',
        email: '',
        kid_relation: '',
        partydate: '',
        partytime: '',
        partyroom_name: '',
        partyroom_id: 'B01',
        branch_id: '',
        partyroom_price: null,
        food_price: 0,
        selected_food: '',
        special_required: '',
        balloonDecorationsChecked: false,
        balloonDecorationsTheme: '',
        addRequirement: '',
        agree_terms: true,
        payment_method: ''
      };
    },
    watch: {
        partydate(newVal) {
            console.log("📅 선택된 날짜:", newVal);
        }
    },
    computed: {
    formattedPartyDate() {
      if (!this.partydate) return '';
      const [y, m, d] = this.partydate.split('-');
      return `${d}-${m}-${y}`;
    },
    availableTimeOptions() {
        if (this.partyroom_name === 'Party room hiring') {
        return ['10:00~13:00', '13:30~16:30'];
        } else if (this.partyroom_name === 'Private party') {
        return ['16:00~19:00', '17:00~20:00'];
        } else {
        return [];
        }
    }
  },
    methods: {
      async submitBooking() {
        if (!this.kid_name) {
          alert("Please fill in required fields.");
          return;
        }
  
        const bookingData = {
          kid_name: this.kid_name,
          kid_gender: this.kid_gender,
          kid_age: this.kid_age,
          owner_name: this.owner_name,
          owner_phone: this.owner_phone,
          email: this.email,
          kid_relation: this.kid_relation,
          partydate: this.formattedPartyDate,
          partytime: this.partytime,
          partyroom_name: this.partyroom_name,
          partyroom_id: this.partyroom_id,
          branch_id: this.branch_id,
          partyroom_price: this.partyroom_price,
          food_price: this.food_price,
          selected_food: this.selected_food.split(',').map(f => f.trim()),
          special_required: this.selected_food.split(',').map(f => f.trim()),
          balloonDecorationsChecked: this.balloonDecorationsChecked,
          balloonDecorationsTheme: this.balloonDecorationsTheme,
          addRequirement: this.addRequirement,
          agree_terms: this.agree_terms,
          payment_method: this.payment_method
        };
  
        try {
          const api = process.env.VUE_APP_API_BASE;
          const res = await axios.post(`${api}/api/save-party/${this.branch_id}`, bookingData);
          sessionStorage.setItem('bookingData', JSON.stringify(bookingData));
          
            console.log("🎯 전체 응답 객체:", res);
            console.log("🎯 res.data:", res.data);
            console.log("🎯 res.data.success:", res.data.success);


          if (res.data && res.data.success) {
            console.log("✅ 예약 성공:", res.data);
            this.$router.push({ name: 'admin' });
            } else {
            alert("예약은 되었지만 응답이 예상과 달라 페이지 이동에 실패했습니다.");
            }

        } catch (err) {
          console.error("Error saving booking:", err);
          alert("Failed to submit booking. Please try again.");
        }
      }
    }
  };
  </script>
  

<style scoped>
.booking-form {
  max-width: 700px;            /* 너무 넓지 않게 */
  margin: 40px auto;           /* 화면 중앙으로 */
  padding: 24px;
  background-color: #ffffff;   /* 흰 배경 */
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.booking-form h2 {
  margin-bottom: 20px;
  font-size: 24px;
  font-weight: bold;
  color: #333;
  text-align: left;
}

.booking-form form > div {
  margin-bottom: 16px;
  text-align: left;
}

.booking-form label {
  display: block;
  margin-bottom: 6px;
  font-weight: 500;
  color: #444;
}

.booking-form input[type="text"],
.booking-form input[type="number"],
.booking-form input[type="email"],
.booking-form input[type="tel"],
.booking-form input[type="date"],
.booking-form input[type="time"],
.booking-form select,
.booking-form textarea {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #ccc;
  border-radius: 6px;
  font-size: 14px;
  box-sizing: border-box;
}

.booking-form textarea {
  resize: vertical;
  min-height: 80px;
}

.booking-form input[type="checkbox"] {
  margin-right: 8px;
}

.button-container {
  display: flex;
  justify-content: flex-end; /* 오른쪽 정렬 */
  margin-top: 20px;
}

.submit-button {
  background-color: #4CAF50;
  color: white;
  border: none;
  padding: 10px 24px;
  font-size: 14px;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.submit-button:hover {
  background-color: #45a049;
}

/* ✅ 태블릿: 양쪽 16px 고정 여백 */
@media (max-width: 768px) {
  .booking-form {
    padding: 54px 50px; /* 동일한 좌우 여백 유지 */
    margin: 30px auto;
    max-width: 100%;
    box-sizing: border-box;
  }

  body {
    padding: 0 36px; /* 📌 뷰포트 전체에 좌우 여백 추가 */
  }

  .button-container {
    justify-content: center;
  }
  .form-wrapper {
    padding: 24px 16px;
  }

}

/* ✅ 모바일: 동일하게 좌우 여백 보장 */
@media (max-width: 600px) {
  body {
    padding: 0 36px; /* 📌 모바일도 균등한 좌우 여백 */
  }

  .booking-form {
    padding: 54px 50px; /* 동일한 좌우 여백 유지 */
    margin: 20px auto;
    max-width: 100%;
  }
  .form-wrapper {
    padding: 24px 16px;
  }

  .submit-button {
    width: 100%;
    font-size: 16px;
    padding: 12px;
  }

  .button-container {
    justify-content: center;
  }
}
</style>