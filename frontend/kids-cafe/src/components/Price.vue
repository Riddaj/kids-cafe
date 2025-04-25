<template>
  
        <NavBar/>
        <div class="price-wrapper">
            <div class="image-container">
                <img src="../assets/kidsparty_grass.png" >
            </div>
            <div class="table-wrapper">
                <div class="content">
                    <h1>Burwood Entry Ticket</h1>
                    <!-- <h2>Kids Ticket</h2> -->
                    <table style="border-collapse: collapse; width: 100%; font-family: 'Comic Sans MS', cursive;">
                    <thead>
                        <tr style="background-color: #ffe4e1;">
                        <th style="padding: 12px;">🕒 Play Time per Child</th>
                        <th style="padding: 12px;">📅 Weekday</th>
                        <th style="padding: 12px;">🎉 Weekend & Public Holiday</th>
                        </tr>
                    </thead>
                    <tbody>
                        <!-- <tr style="background-color: #fffaf0;">
                        <td style="padding: 12px 20px;">1</td>
                        <td style="padding: 12px 20px;">$13</td>
                        <td style="padding: 12px 20px;">$16</td>
                        </tr>
                        <tr style="background-color: #f0fff0;">
                        <td style="padding: 12px 20px;">2</td>
                        <td style="padding: 12px 20px;">$20</td>
                        <td style="padding: 12px 20px;">$26</td>
                        </tr>
                        <tr style="background-color: #f0f8ff;">
                        <td style="padding: 12px 20px;">Unlimited</td>
                        <td style="padding: 12px 20px;">$50</td>
                        <td style="padding: 12px 20px;">$60</td>
                        </tr> -->

                        <tr v-for="(item, index) in sortedPrices" :key="index" :style="{ backgroundColor: getRowColor(index) }">
                            <td style="padding: 12px 20px;">{{ item.Duration.replace('_', ' ') }}</td>
                            <td style="padding: 12px 20px;">${{ item.WeekdayPrice }}</td>
                            <td style="padding: 12px 20px;">${{ item.WeekendPrice }}</td>
                        </tr>
                    </tbody>
                    </table>
                    
                        <p style="clear:both; margin-top: 8px; color: black;">
                            <strong>1 hour free</strong> 🎁 for Early bird (before 10 AM) & after 3 PM ⏰
                        </p>
                    
                    <!-- 안나옴. ㅎ -->
                    <!-- <div class="table-container">
                    <DataTable :value="products" tableStyle="min-width: 50rem">
                        <Column v-for="col of columns" :key="col.field" :field="col.field" :header="col.header"></Column>
                    </DataTable> 
                    </div> -->
                    <div id="announcement" style="text-align: left;">    
                        <pre>
<strong>Important Notice:</strong>
🎈 Children under 12 months, free entry for 2 hours. (ID or Certificate may be required)
🎈 Kids over 12 years old will be charged as adults
🎈 15% surcharge applied on public holidays
🎈 Twinkle Kids Cafe prices vary by location
                        </pre>
                    </div>
                </div>
                
            </div>
        </div>

</template>

<script>
import axios from 'axios'; // axios를 import 추가
import NavBar from './NavBar.vue';
// import DataTable from 'primevue/datatable';
// import Column from 'primevue/column';

export default {
    components: {
        NavBar,
    
    },
    data(){
        return{
            // products: [
            //     { code: 'P001', name: 'Product 1', category: 'Category 1', quantity: 10 },
            //     { code: 'P002', name: 'Product 2', category: 'Category 2', quantity: 20 },
            // ],
            branchID: this.$route.params.branchID,
            prices:[],
            sortedPrices:[],
            categorizedMenus: {},  // 카테고리별로 메뉴를 나눠 저장할 객체
            activeMenu: null, // 클릭된 메뉴를 추적하는 변수
           
        }
    },
    mounted() {
        this.fetchprice();  // 컴포넌트가 마운트되면 fetchmenu 호출
    },
    methods:{
        async fetchprice(){
            console.log("Branch ID:", this.branchID);  // 값이 제대로 있는지 확인
            try {
                const response = await axios.get(`http://localhost:8081/api/price/${this.branchID}`);
                this.prices = response.data.prices;
                //this.categorizeMenu(); 
                console.log("### price data 나오라고 ### :", response.data.prices);

                // duration 순서대로 정렬
                let sortedItems = this.prices.sort((a, b) => {
                    const order = ['1_hour', '2_hour', 'Unlimited']; // 정렬 순서 정의
                    return order.indexOf(a.Duration) - order.indexOf(b.Duration);
                });
                //정렬 후.
                
                this.sortedPrices = [...sortedItems];
                console.log("###정렬 후## :", this.sortedPrices);
                
            } catch (error) {
                console.error("#### Error fetching prices ##### :", error);
            }
        },       
        categorizeMenu() {
            // prices 배열을 MenuCategory 기준으로 분류
            this.categorizedMenus = this.prices.reduce((categories, menu) => {
                const category = menu.MenuCategory;

                if (!categories[category]) {
                    categories[category] = [];  // 카테고리가 없으면 새 배열 생성
                }

                categories[category].push(menu);  // 카테고리에 해당하는 메뉴 추가
                return categories;
            }, {});
        },
        getRowColor(index) {
            // 색상 규칙을 위한 조건을 설정
            if (index === 0) return '#fffaf0'; // 첫 번째 행
            else if (index === 1) return '#f0fff0'; // 두 번째 행
            else return '#f0f8ff'; // Unlimited 행
        }
    },
    created() {
        this.fetchprice();  // 컴포넌트 생성 시 메뉴 데이터 가져오기
    }
}
</script>

<style scoped>
.table-wrapper {
  display: flex;
  flex-direction: column; /* ← 요거만 세로로 만드는 핵심 */
  gap: 16px;
  
  padding: 16px;
}


h1 {
    justify-content: center;

}

.table-container {
  display: flex;
  justify-content: center; /* 가로 중앙 정렬 */
  align-items: center; /* 세로 중앙 정렬 */
  flex: 1;
  width: 100%;
  }
.price-wrapper {
  display: flex;         /* 가로 배치의 핵심 */
  gap: 16px;             /* 두 div 사이 간격 */
  
  padding: 16px;
}

.content {
    padding-top: 64px;
    display: flex;
    flex-wrap: wrap;
    flex-direction: column;
 
}

.image-container img {
  flex: 1;
  width: 100%;
  height: 93%;
  object-fit: cover;
  border-radius: 8px;

}

.info-container {
  flex: 1;
}

pre {
  text-align: left;
  margin: 0;         /* 여백 제거 */
  padding: 0 16px;   /* 좌우 약간의 여백 */
}

</style>