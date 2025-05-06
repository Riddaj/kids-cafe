<template>
  <div id="app">
    <h1>Party list</h1>
    <div class="table-container">

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
      async fetchParty() {
        console.log("Branch ID:", this.branchID);  // 값이 제대로 있는지 확인
        try {
          const response = await axios.get(`http://localhost:8081/api/get-party/${this.branchID}`); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

          this.parties = response.data.parties;
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

<style>

</style>