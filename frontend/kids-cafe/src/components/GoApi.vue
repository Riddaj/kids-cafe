<template>
  <div>
    <h1>{{ message }}</h1>
    <button @click="fetchData">Fetch Message from Go API</button>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      message: ''
    };
  },
  methods: {
    async fetchData() {
      try {
        const response = await axios.get('/api/hello'); // '/api/hello'로 요청
        this.message = response.data.message; // Go API에서 받은 메시지
      } catch (error) {
        console.error('Error fetching data from Go API', error);
      }
    }
  }
};
</script>
