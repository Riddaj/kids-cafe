<template>
 
    <NavBar/>
    <div class="rule-wrapper">
      <h1>{{branchName}} Entry Rules</h1>
      <div class="rule-content">
        <span class="title_text">We Provides a Safe and Entertaining Playground for Everyone</span>
        <div class="elementor-widget-container">
          <p style="text-align: center;">The following provides an outline of our basic conditions of entry. The purpose of these guidelines is to ensure all our children and adults feel comfortable and safe while enjoying our premises.</p>	
        </div>
      </div>

      <section class="play-area-section">
      <!-- Parallax elements -->
      <div class="parallax-item moon"></div>
      <div class="parallax-item star1"></div>
      <div class="parallax-item star2"></div>
      <div class="parallax-item star3"></div>

      <!-- Title and subtitle -->
      <div class="subtitle">If you encounter any problems within this playground, please contact the operator immediately</div>
      <h2>Play Area Rules</h2>

      <!-- Rules list -->
      <div class="play-area-rules">
        <ul>
          <li>Have Fun! That's what the playground is all about.</li>
          <li>Adult supervision required at all times.</li>
          <li>One child per equipment piece at any one time.</li>
          <li>One person at a time exiting all play events.</li>
          <li>Slides to be used feet first only, hands on knees.</li>
          <li>No shoes. Non-slip socks must be worn in the center.</li>
        </ul>
      </div>
    </section>


    </div>
    <Footer/>
</template>

<script>
import NavBar from '../components/NavBar.vue';
import Footer from '../components/Footer.vue';

export default {    
  data() {
        return {
          branchName: '',
          branchID: this.$route.params.branchID,
        };
    },  
    components:{
      NavBar,
      Footer,
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

<style scoped>

.rule-content {
  margin-top:10%;
}
.title_text {
 
  display: block;
  text-align: center;
  font-size: 2.5rem; /* 큼직한 글씨 */
  font-weight: 700; /* 굵게 */
  color: #eb6d54; /* 다홍색 */
  margin: 4rem 0; /* 위아래 여백 */
  
  
}

.elementor-widget-container p {
  text-align: center;
  font-size: 1.1rem; /* 주인공보다는 작지만 너무 작지 않게 */
  color: #555; /* 은은한 회색 톤 */
  margin-top: -1rem;
  margin-bottom: 2rem;
  line-height: 1.6;
  max-width: 700px;
  margin-left: auto;
  margin-right: auto;
  font-family: "Woodford_bournethin", sans-serif;
  font-weight: 300; /* 얇은 두께 */
  margin-bottom:9%;
}

/* General section styling */
.play-area-section {
  position: relative;
  overflow: hidden;
  background: #FCC953;
  padding: 60px 20px;
  text-align: center;
  font-family: 'Comic Sans MS', cursive, sans-serif;
}

/* Parallax elements (stars, moon) */
.parallax-item {
  position: absolute;
  background-size: auto;
  background-repeat: no-repeat;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: 1;
}

.star1 {
  top: 30px;
  left: 5%;
  background-image: url('https://twinklekidscafe.com.au/wp-content/uploads/2018/10/star_5.png');
}

.star2 {
  top: 75px;
  left: 8%;
  background-image: url('https://twinklekidscafe.com.au/wp-content/uploads/2018/10/star_5.png');
  margin-top:20%;
}

.star3 {
  top: 68px;
  left: 90%;
  background-image: url('https://twinklekidscafe.com.au/wp-content/uploads/2018/10/star_4.png');
  margin-top:30%;
}

.moon {
  top: 15px;
  left: 92%;
  background-image: url('https://twinklekidscafe.com.au/wp-content/uploads/2018/10/moon.png');
  
}

/* Title styling */
.play-area-section h2 {
  font-size: 36px;
  color: #333;
  margin: 50px 0;
}

.play-area-section .subtitle {
  font-family: "Playfair Display", serif !important;
  font-size: 1.15rem;
  line-height: 1.5em;
  font-style: italic !important;
  text-transform: none;
  letter-spacing: 0;
  display:block;
  color: #FD5C37;
  margin-bottom: 30px;
  position: relative;
  margin-top:9%;
}


/* Rules list */
.play-area-rules ul {
  column-count: 2; /* Divide the list into two columns */
  column-gap: 30px; /* Set the gap between columns */
  list-style: none;
  padding: 0;
  max-width: 70%;
  margin: 0 auto;
  text-align: left;
}

.play-area-rules li {
  position: relative;
  padding-left: 30px;
  margin-bottom: 15px;
  font-size: 16px;
  line-height: 1.5;
  color: #444;
  width:100%;
}

.play-area-rules li::before {
  content: "✔";
  color: #FD5C37;
  position: absolute;
  left: 0;
  font-weight: bold;
}

/**
divider를 흉내내고 싶었다 */
/* .play-area-rules li::after {
  content: ''; 
  position: absolute;
  width: 100%;
  height: 1px;
  background-color: white; 
  bottom: 0; 
  left: 0;
  margin-top:10px;
} */

</style>