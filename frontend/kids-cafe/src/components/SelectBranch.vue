<template>
    <div class="branch_info">
        <h1>Branch</h1>
        <div v-if="branches && branches.length > 0"   class="branch-container">
            <div v-for="branch in branches" :key="branch.branch_id">
                <router-link :to="`/book_a_party/select_room/${branch.branch_id}`" class="branch-card">
                    <div class="branch-card">
                        <div>
                            {{ branch.branch_name }}<br>
                        </div>
                        <div class="img-div">
                            <img :src="getBranchImage(branch.branch_id)" 
                            :alt="`Branch ${branch.branch_name}`" 
                            class="branch-image">
                        </div>
                    </div>
                </router-link>
            </div>
        </div>
        <div v-else> <!-- 데이터가 없으면 이 메시지 표시 -->
            No branches available.
        </div>
    </div>
</template>

<script>
import axios from 'axios'; // axios를 import 추가

export default {
    data() {
        return {
            branches:[]
        };
    },  
    mounted() {
    this.fetchBranches();
    console.log("🔥🔥 Mounted! branches 🔥🔥:", this.branches);
    },
    methods: {
    async fetchBranches() {
      try {
        const response = await axios.get("https://kids-cafe-rm9g.onrender.com/api/branches"); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

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
                'hornsby': "https://twinklekidscafe.com.au/wp-content/uploads/2024/06/gallery-27-1.jpg"
            };
            return images[branch_id]; // 기본 이미지
        }
  }
}
</script>

<style scoped>

a{
    text-decoration: none;
}

.branch-image{
    width: 500px;
    height: 350px;
    opacity: 0.8;
    border-radius: 15px;
    justify-content: center; /* 중앙 정렬 */
}

.branch-container {
    margin: 20px 20px 20px 20px;
    display: flex;
    flex-wrap: wrap;
    gap: 300px; /* 카드 사이 간격 */
    justify-content: center; /* 중앙 정렬 */
}

.branch-card {
    display: flex; /* 내부 요소를 가로 배치 */
    flex-direction: column; /* 세로 방향 배치 */
    justify-content: space-between; /* 내용과 버튼 사이 여백 자동 조정 */
    align-items: center; /* 중앙 정렬 */
    background: transparent;
    border-radius: 10px;
    padding: 15px;
    width: 250px;
    /* box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); */
    text-align: left;
    transition: transform 0.2s, box-shadow 0.2s;
}

.branch-card:hover {
    transform: translateY(-5px);
    /* box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15); */
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

.select-room {
    background-color: #ffb3b3; /* 버튼 배경 색상 */
    color: white; /* 텍스트 색상 */
    text-decoration: none; /* 링크 스타일 제거 */
    padding: 10px 20px; /* 버튼 내부 여백 */
    border-radius: 5px; /* 둥근 모서리 */
    font-size: 16px; /* 글자 크기 */
    display: inline-block; /* 버튼 형태 유지 */
    margin-top: auto; /* 버튼이 자동으로 글자 아래에 배치되도록 설정 */
}

.select-room:hover{
    background-color: #6699ff; /* 버튼 배경 색상 */
}
</style>

