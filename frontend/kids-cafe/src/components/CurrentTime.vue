<template>
    <div id="app" class="current-time">
      <p>Our time: {{ currentTime }} Australia/Sydney</p>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        currentTime: this.getFormattedTime()
      };
    },
    methods: {
      getFormattedTime() {
        const now = new Date();
        return now.toLocaleTimeString('en-US', {hour: '2-digit', minute: '2-digit', hour12: true}); // 현재 시간을 HH:MM:SS 형식으로 변환
      }
    },
    mounted() {
      this.timer = setInterval(() => {
        this.currentTime = this.getFormattedTime();
      }, 1000); // 1초마다 시간 갱신
    },
    beforeUnmount() {
      clearInterval(this.timer); // 컴포넌트가 제거될 때 타이머 정리
    }
  };
  </script>
  
  <style>
  html, body {
    width: 100%;
    overflow-x: hidden;  /* 가로 스크롤 방지 */
  }

  #app {
    position: relative;
    width: 100%;         /* 가로를 화면에 맞게 설정 */
    justify-content: flex-end; /* 콘텐츠 가로 중앙 정렬 */
    padding-right: 50px; /* 오른쪽 여백 추가 */
    font-size: 16px;
    font-weight: bold;
  }

  .current-time p{
    color: black !important;
    padding-right: 50px; /* 오른쪽 여백 추가 */
  }

  </style>
  