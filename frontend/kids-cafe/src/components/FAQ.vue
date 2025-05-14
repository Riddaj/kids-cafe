<template>
  <div id="app">
    <NavBar/>
    <div class="main">
        <h1>FREQUENTLY ASKED QUESTION</h1>
        <!-- 확장, 닫힘 버튼 
        <div class="flex flex-wrap gap-2 mb-6">
            <Button class="btn" type="button" icon="pi pi-plus" label="Expand All" @click="expandAll">Expand All</Button>
            <Button class="btn" type="button" icon="pi pi-minus" label="Collapse All" @click="collapseAll">Collapse All</Button>
        </div>-->
        <div  v-if="FAQs && FAQs.length > 0"  class="faq-card">
            <div v-for="(FAQ, index) in FAQs" :key="FAQ.FaqID" 
            :style="{ backgroundColor: getRandomColor(index) }" class="faq-item">
                <div class="faq-question">
                    {{ FAQ.FaqQuestion }}<br>
                </div>
                <div class="faq-answer">
                    {{ FAQ.FaqAnswer }}<br>
                </div>
            </div>
        </div>
        <!-- <Tree 
            :value="treeData" 
            :expanded-keys="expandedKeys" 
            expandMode="multiple"  
            class="w-full md:w-[30rem]"
            @update:expanded-keys="onToggle" 
            @select="onFAQClick" 
        /> -->
        <!-- <div v-if="selectedFAQ">
            <p>{{ selectedFAQ }}</p>
        </div> -->
    </div>
    <Footer/>
    </div>
</template>

<script>
import axios from 'axios'; // axios를 import 추가
import NavBar from './NavBar.vue';
import PrimeVue from 'primevue/config';
import Tree from 'primevue/tree';
import Button from 'primevue/button'; // <-- Button 추가!
import Footer from './Footer.vue';

export default {
    components: {
        NavBar,
        Tree,
        Button,
        PrimeVue,
        Footer
    },
    data(){
        return{
            FAQs :[],
            colors: ['#FFEBEE', '#E8F5E9', '#E3F2FD', '#FFF3E0', '#F3E5F5'], // 색상 배열
            //branchID: this.$route.params.branchID,
            // URL 파라미터에서 branchID 가져오기
            treeData: [],
            expandedKeys: {},
            selectedFAQ: null, // 클릭된 FAQ의 답변을 저장할 변수
        }
    },
    mounted() {
        //console.log("✨🎉✨ Route params:", this.$route.params);  // 여기에서 branchID가 있는지 확인
        this.fetchFAQ();
    },
    methods:{
        // 색상을 랜덤으로 선택하는 메소드
        getRandomColor(index) {
            return this.colors[index % this.colors.length]; // 색상을 인덱스를 기준으로 순차적으로 선택
        },
        async fetchFAQ() {
                //console.log("✨🎉✨ Branch ID ✨🎉✨:", this.branchID);  // 값이 제대로 있는지 확인
    
                const api = process.env.VUE_APP_API_BASE;
                try {
                    //const response = await axios.get(`https://kids-cafe-rm9g.onrender.com/api/faq`); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능
                    const response = await axios.get(`${api}/api/faq`); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능
    
                    this.FAQs = response.data.FAQs;
                    //console.log("### 전체 response 객체 ### :", response);
                    console.log("### fetchFAQ data 나오라고 ### :", this.FAQs);

                    // Tree 컴포넌트에 적합한 형식으로 데이터를 변환
                    this.treeData = this.FAQs.map(faq => ({
                        key: faq.FaqID, // (optional) 에러 방지 위해 key 추가
                        label: faq.FaqQuestion,
                        children: [
                            {
                            key: faq.FaqID + '_answer',
                            label: faq.FaqAnswer,  // 답변을 child로 넣기
                            }
                        ]
                        }));

                        console.log("✨ 변환된 Tree 데이터:", this.treeData);
                } catch (error) {
                    console.error("#### Error fetching fetchFAQ ##### :", error);
                }
            },
        // Tree 항목 클릭 시 호출되는 메소드
        onFAQClick(event) {
            console.log("왜안나와 답변 == ", this.selectedFAQ);
            this.selectedFAQ = {
                question: event.node.label,
                answer: event.node.children[0].label,  // 첫 번째 자식 항목에 있는 답변을 가져옴
            };
            //this.selectedFAQ = event.node.children[0].label;  // 첫 번째 자식 항목에 있는 답변을 가져옴
        },
        // Tree 펼침/축소 상태 업데이트 메소드
        onToggle(event) {
            this.expandedKeys = event.value || {}; // 여기를 수정!
        },
        expandAll() {
            const expanded = {};
            this.treeData.forEach((node, index) => {
            expanded[index] = true;
            });
            this.expandedKeys = expanded;
        },
        collapseAll() {
            this.expandedKeys = {};
        }
    },
}

</script>

<style scoped>
.main {
  display: flex;
  flex-direction: column; /* 👉 세로로 나열 */
  justify-content: left;
  padding: 120px 0 40px; /* top padding으로 헤더 피하기 */
  padding-bottom: 150px;
}

h1 {
  font-size: 2rem;
  margin-bottom: 20px;
  text-align: center;
}

.faq-card {
  display: flex;
  flex-direction: column;
  gap: 20px;
  margin-top: 20px;
  margin-left: 80px;
  margin-right: 80px;
}

.faq-item {
  text-align: left;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.faq-item:hover {
  transform: translateY(-5px);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.1);
}

.faq-question {
  font-size: 1.1rem;
  font-weight: bold;
  color: #333;
}

.faq-answer {
  font-size: 1rem;
  color: #555;
  margin-top: 10px;
}

@media (max-width: 768px) {
  .faq-item {
    padding: 15px;
  }

}

@media (max-width: 768px) {
  .faq-card {
    margin-left: 16px;
    margin-right: 16px;
  }

  .faq-item {
    padding: 16px;
  }

  .faq-question {
    font-size: 1rem;
  }

  .faq-answer {
    font-size: 0.95rem;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  .faq-card {
    margin-left: 40px;
    margin-right: 40px;
  }
}

</style>