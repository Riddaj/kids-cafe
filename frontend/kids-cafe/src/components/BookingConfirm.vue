<template>
   <div>
      Congratulations!
      A birthday party always requires a lot of details!<br>
      We will contact you soon to discuss more details.<br>
      Please wait for a moment.<br>
      Thank you.
      <div class="button-container">
         <router-link :to="`/`">
               <button type="submit" class="submit-button" @click="submitBooking">go home</button>
         </router-link>
      </div>
   </div>  

</template>

<script>
export default {

}
</script>

<style scoped>
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

</style>

