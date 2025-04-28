<template>
    <!-- booking process (2단계) -->
    <div class="booking-process">
        <ul class="booking-process-item">
            <li class="step-info-item"><a href="#">Category</a></li>
            <li class="step-info-item"><a href="#">Time</a></li>
            <li class="step-info-item"><a href="#">Client</a></li>
        </ul>
    </div>  
</template>

<script>
export default {

}
</script>

<style>
html, body, #app {
  width: 100vw;
  height: 100%;
  margin: 0;
  padding: 0;
  overflow-x: hidden;
}

.booking-process-item {
    width: 100%; /* 뷰포트 전체 너비를 차지하도록 설정 */
    display: flex;
    height: 70px;
    list-style: none;
    padding: 0;
    margin: 0;
    background-color: #ffe6e6;
    margin-bottom: 30px;
}

.step-info-item{
    width: 33.33%; /* 1/3로 수정! */
    padding: 10px; /* 항목 간 간격을 위한 padding (옵션) */
    box-sizing: border-box; /* padding이 width에 포함되도록 설정 */
    align-items: center; /* 세로(수직) 가운데 정렬 */
    display: flex; /* 추가 */
    justify-content: center; /* 가로(수평) 가운데 정렬 */
    text-align: center;
    cursor: pointer;
    border: 1px solid #e6e6e6; /* ul 테두리 추가 */
}
</style>