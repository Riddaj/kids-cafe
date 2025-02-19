<template>
    <!-- 예약 정보 받기 -->
    <div class="booking_info">
        <div class="confirm-booking">
            <h2>Please, confirm details</h2>
                <div class="booking-info-detail">
                    <div class="info-detail">
                        <label>Branch: </label>
                        <select class="info-detail-input">
                            <!-- v-for로 branch.name을 반복 -->
                            <option v-for="branch in branches" :key="branch.id">{{ branch.branch_name }}</option>
                        </select>
                    </div>
                    <div class="info-detail">
                        <label>Name: </label><input type="text" class="info-detail-input" />
                    </div>
                    <div class="info-detail">
                        <label>Email: </label><input type="text" class="info-detail-input" />
                    </div>
                    <div class="info-detail">
                        <label>Phone: </label><input type="text" class="info-detail-input" />
                    </div>
                    <div class="info-detail">
                        <label>Number of people: </label>
                        <label>Adults: </label><input type="text" class="info-detail-input-1"       
                        v-model="inputValue"
                        @input="validateInput" /><br>
                        <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
                        <label>Children: </label><input type="text" class="info-detail-input-1" 
                        v-model="inputValue"
                        @input="validateInput" /><br>
                        <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
                    </div>
                <div>
                    <button class="btn-book2" href="#" title="Confirm Booking">Submit</button>
                </div>
            </div>
        </div>
        </div>  
</template>

<script>
export default {
    data(){
        return{
            inputValue: "",  // 빈 문자열로 초기화하여 undefined 방지
            errorMessage: "" // 에러 메시지 상태
        };
    },
    props: {
        branches: Array,  // 부모에서 전달받은 branches 데이터를 props로 받음
    },
    mounted() {
        console.log("%%%% Received branches in ConfirmDetail %%%% :", this.branches);  // props가 잘 전달되는지 확인
    },
    methods: {
        validateInput() {
            // inputValue가 undefined인 경우 빈 문자열로 처리
            this.inputValue = this.inputValue || "";
            // 숫자만 남기고 나머지 문자 제거
            const numericValue = this.inputValue.replace(/[^0-9]/g, "");
            
            if (this.inputValue !== numericValue) {
                this.errorMessage = "Number only, please."; // 에러 메시지 표시
            } else {
                this.errorMessage = ""; // 올바른 입력이면 메시지 숨김
            }

            this.inputValue = numericValue; // 숫자만 입력되도록 반영
        }
    }    
}
</script>

<style scoped>
.error-message {
  color: red;
  font-size: 12px;
  margin-top: 5px;
}

/** booking info css*/
.booking-info-detail {
    display: flex;               /* flexbox로 배치 */
    flex-direction: column;      /* 세로 방향으로 정렬 */
    gap: 15px;                   /* 항목 간의 간격 */
    padding: 20px;               /* 내부 여백 */
}

.info-detail {
    display: flex;               /* 내부 요소들 역시 flexbox */
    align-items: center;         /* 레이블과 입력 필드 수직 중앙 정렬 */
    gap: 10px;                   /* 레이블과 입력 필드 간의 간격 */
}

.info-detail label {
    width: 120px;                /* 레이블의 고정 너비 설정 */
    text-align: right;           /* 레이블 텍스트를 오른쪽 정렬 */
}

.info-detail-input {
    /*flex-grow: 1;                 입력 필드는 가용 공간을 모두 차지 */
    width: 250px;
    padding: 5px;                /* 입력 필드에 패딩 추가 */
    border: 1px solid #ccc;      /* 입력 필드 테두리 */
    border-radius: 5px;          /* 입력 필드 테두리 둥글게 */
    box-sizing: border-box;  /* box-sizing 설정 */
}

.info-detail-input-1{
    /* flex-grow: 1;                /* 입력 필드는 가용 공간을 모두 차지 */
    width: 175px; /* 입력 필드의 너비를 200px로 설정 */
    padding: 5px;                /* 입력 필드에 패딩 추가 */
    border: 1px solid #ccc;      /* 입력 필드 테두리 */
    border-radius: 5px;          /* 입력 필드 테두리 둥글게 */
    box-sizing: border-box;  /* box-sizing 설정 */
}

/** button img */
.btn-book2 {
    display: inline-block;
    padding: 12px 20px;
    margin-left: 30px; /* 버튼과 메뉴 간에 간격을 줌 */
    background-color: #4d0099; /* 버튼 배경 색상 */
    color: white; /* 텍스트 색상 */
    text-align: center;
    text-decoration: none; /* 링크 스타일 제거 */
    border-radius: 8px; /* 둥근 모서리 */
    font-size: 16px; /* 글자 크기 */
}
</style>