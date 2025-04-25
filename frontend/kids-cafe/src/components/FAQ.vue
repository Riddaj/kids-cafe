<template>
  <div id="app">
    <NavBar/>
        <h1>FREQUENTLY ASKED QUESTION</h1>
        <!-- 확장, 닫힘 버튼 -->
        <div class="flex flex-wrap gap-2 mb-6">
            <Button type="button" icon="pi pi-plus" label="Expand All" @click="expandAll" />
            <Button type="button" icon="pi pi-minus" label="Collapse All" @click="collapseAll" />
        </div>
        <Tree 
            :value="treeData" 
            :expanded-keys="expandedKeys" 
            class="w-full md:w-[30rem]"
            @update:expanded-keys="onToggle" 
            @node-click="onFAQClick" 
        />
        <div v-if="selectedFAQ">
        <h2>Answer</h2>
        <p>{{ selectedFAQ }}</p>
        </div>
    </div>
</template>

<script>
import axios from 'axios'; // axios를 import 추가
import NavBar from './NavBar.vue';
import PrimeVue from 'primevue/config';
import Tree from 'primevue/tree';

export default {
    components: {
        NavBar,
        Tree
    },
    data(){
        return{
            FAQs :[],
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
        async fetchFAQ() {
                //console.log("✨🎉✨ Branch ID ✨🎉✨:", this.branchID);  // 값이 제대로 있는지 확인
    
                try {
                    const response = await axios.get(`http://localhost:8081/api/faq`); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능
    
                    this.FAQs = response.data.FAQs;
                    //console.log("### 전체 response 객체 ### :", response);
                    console.log("### fetchFAQ data 나오라고 ### :", this.FAQs);

                    // Tree 컴포넌트에 적합한 형식으로 데이터를 변환
                    this.treeData = this.FAQs.map(faq => ({
                    label: faq.FaqQuestion,
                    data: faq.FaqAnswer,
                    }));
                } catch (error) {
                    console.error("#### Error fetching fetchFAQ ##### :", error);
                }
            },
        // Tree 항목 클릭 시 호출되는 메소드
        onFAQClick(event) {
            this.selectedFAQ = event.node.data; // 클릭된 항목의 답변을 선택
        },
        // Tree 펼침/축소 상태 업데이트 메소드
        onToggle(event) {
            this.expandedKeys = event.value;
        },
    },
}

</script>

<style scoped>

.faq-body{
    text-align: left;
    margin: 20px 20px 20px 50px;
}

.set{
    margin-bottom: 10px;
}

</style>