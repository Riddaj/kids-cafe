<template>
  <div id="app" class="pick-branch-background">
    <div class="logo-image">
        <a href="/">
            <!-- /.sc_layouts_logo -->		
            <img fetchpriority="high" class="logo_image" src="//twinklekidscafe.com.au/wp-content/uploads/2018/08/Twinkle-Kids-Cafe_logo.png">
        </a>
    </div>
    <div class="select-branch">
        <div v-if="branches && branches.length > 0"   class="branch-container">
            <div v-for="branch in branches" :key="branch.branch_id">
                <router-link :to="`/home/${branch.branch_id}`">
                    <div class="branch-card">
                        <div class="branch-name">
                            {{ branch.branch_name }}<br>
                        </div>
                        <div class="img-div">
                            <img :src="getBranchImage(branch.branch_id)" 
                            :alt="`Branch ${branch.branch_name}`" 
                            class="branch-image">
                            <div class="image-overlay"></div> <!-- 그라데이션용 오버레이 -->
                        </div>
                    </div>
                </router-link>
            </div>
        </div>
        <div v-else> <!-- 데이터가 없으면 이 메시지 표시 -->
            No branches available.
        </div>
    </div>
    <DinoAnimation/>
  </div>
</template>

<script>
import SelectBranch from '@/components/SelectBranch.vue'  
import axios from 'axios'; // axios를 import 추가
import DinoAnimation from '@/components/DinoAnimation.vue';

export default {
    data() {
        return {
            branches:[]
        };
    },  
    components:{
        DinoAnimation
    },
    mounted() {
    this.fetchBranches();
    console.log("🔥🔥 Mounted! branches 🔥🔥:", this.branches);
    },
    methods: {
    async fetchBranches() {
      try {
        const response = await axios.get("http://localhost:8081/api/branches"); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

        this.branches = response.data.branches;
        //console.log("### 전체 response 객체 ### :", response);
        console.log("### branchs data 나오라고 ### :", response.data.branches);
      } catch (error) {
        console.error("#### Error fetching branchs ##### :", error);
      }
    }, 
    getBranchImage(branch_id) {
            const images = {
                'burwood': "https://images.squarespace-cdn.com/content/v1/637d8d8a7f609c521ddd5429/1672359522132-RU2ZPENTVALEBF0Z47PG/285887484_694866768237604_5851615251096205906_n.jpg",
                'hornsby': "/images/hornsby.jpg"
            };
            return images[branch_id]; // 기본 이미지
        }
  }
}
</script>

<style scoped>

/* 전체 앱 배경 */
a {
    text-decoration: none;
}


/* html과 body 태그에 높이 설정 */
body {
    height: 100%; /* 화면 전체를 차지하도록 설정 */
    margin: 0; /* 기본 margin 제거 */
    overflow-x: hidden; /* ✨ 가로 스크롤 제거 핵심 */
    overflow-y: auto; /* ✨ 가로 스크롤 제거 핵심 */
    background-color: #eefabb; /* 🌈 원하는 배경색 */
}

#app {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 40px;
  box-sizing: border-box;
  min-height: 100%;
}

.img-div {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

/* 로고 이미지 */
.logo_image {
  margin-top: 50px;
  max-width: 400px;
  width: 100%;
  height: auto;
}

.branch-container {
    margin: 20px 20px 20px 20px;
    display: flex;
    flex-wrap: wrap;
    gap: 100px; /* 카드 사이 간격 */
    justify-content: center; /* 중앙 정렬 */
}

.branch-name{
  margin-bottom: 5px;
  font-size: 20px;
  font-weight: bold;
  color: #2c3e50; /* 원하는 색상으로 조정 가능 */
  text-align: center;
}

/* 카드 스타일 */
.branch-card {
    align-items: center; /* 가운데 정렬 */
    position: relative;  /* 자식 요소에 absolute 적용을 위해 */
    display: flex;
    flex-direction: column;
    justify-content: flex-end;  /* 이미지와 텍스트가 겹치지 않도록 아래로 배치 */
    align-items: center;
    border-radius: 15px;
    padding: 0px;
    width: 500px;           /* 가로 크기 */
    /*height: 350px;           세로 크기 (정사각형) */
    text-align: center;
    transition: transform 0.2s;
    overflow: hidden; /* ✅ 이미지 넘어가면 잘라내기 */
    /* box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); */
    background-color: transparent; /* ✅ 배경색 제거 or 투명하게 */
}

/* Hover 효과 */
.branch-card:hover {
    transform: translateY(-5px);
}

/* 이미지 스타일 */
.branch-image{
    width: 100%;
    height: 350px;
    opacity: 0.8;
    border-radius: 15px;
    display: block; /* ✅ inline 요소 여백 제거 */
    justify-content: center; /* 중앙 정렬 */
    object-fit: cover; /* 이미지 비율 유지 */
    opacity: 0.6;       /* 이미지를 어둡게 */
}

/* 그라데이션 오버레이 */
.image-overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 100%; /* 전체 덮거나, 원하면 height: 40% 등으로 조정 */
  background: linear-gradient(to bottom, rgba(0, 0, 0, 0) 60%, #ffeb91 100%);
  border-radius: 15px;
  pointer-events: none; /* 클릭 이벤트 방지 */
}

.branch-card h3 {
    font-size: 18px;
    font-weight: bold;
    color: #333;
    margin-bottom: 10px;
}

.branch-card p {
    font-size: 14px;
    color: #666;
}

</style>
