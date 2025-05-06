<template>
  <div id="app">
    <BookingBar/>
    <div class="rule-wrapper">
      <h1>{{branchName}} Entry Rules</h1>
      <div class="rule-content">

      </div>
    </div>
  </div>
</template>

<script>
import BookingBar from '../components/BookingBar.vue';

export default {    
  data() {
        return {
          branchName: '',
          branchID: this.$route.params.branchID,
        };
    },  
    components:{
        BookingBar
    },
    created() {
        //const branchID = this.$route.params.branchID;
        const branchID = this.$route.params.branchID
        this.branchID = branchID; // 👈 이게 빠져있었음!
        this.branchName = this.getBranchNameByID(branchID)
    },
    methods:{
      getBranchNameByID(id) {
        const branchMap = {
            'burwood': 'Burwood',
            'hornsby': 'Hornsby',
            // 필요하면 계속 추가 가능
        }
        return branchMap[id] || 'Unknown Branch'
        },
    }
}
</script>

<style>

</style>