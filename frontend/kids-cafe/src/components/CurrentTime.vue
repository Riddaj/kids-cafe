<template>
    <div class="current-time">
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
  
  <style scoped>
  .current-time {
    font-size: 16px;
    font-weight: bold;
    color: black;
  }
  </style>
  