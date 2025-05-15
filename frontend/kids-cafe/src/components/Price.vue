<template>
    <div id="app">
        <NavBar/>
        <div class="price-wrapper">
            <!-- <div class="image-container">
                <img v-if="branchID" :src="getBranchImage(branchID)" 
                            class="branch-image">
            </div> -->
            <div class="table-wrapper">
                <div class="content">
                    <h1>{{branchName}} Entry Ticket</h1>
                    <h2>Kids Ticket</h2> 
                    <table class="price-table">
                        <thead>
                            <tr style="background-color: #ffe4e1;">
                            <th style="padding: 12px;">🕒 Play Time per Child</th>
                            <th style="padding: 12px;">📅 Weekday</th>
                            <th style="padding: 12px;">🎉 Weekend & Public Holiday</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(item, index) in sortedPrices" :key="index" :style="{ backgroundColor: getRowColor(index) }">
                                <td style="padding: 12px 20px;">{{ item.Duration.replace('_', ' ') }}</td>
                                <td style="padding: 12px 20px;">${{ item.WeekdayPrice }}</td>
                                <td style="padding: 12px 20px;">${{ item.WeekendPrice }}</td>
                            </tr>
                        </tbody>
                    </table>
                    <!-- ✅ 점선 구분선 추가 -->
                    <hr style="border: none; border-top: 1px dashed #eb6d54; margin: 40px 0;" />
                    <h2>Adult Ticket</h2> 
                    <p>$5 per session. Comes with a $5 drink voucher.</p>
                    <hr style="border: none; border-top: 1px dashed #eb6d54; margin: 40px 0;" />
                        <p style="text-align: right; clear:both; margin-top: 8px; ">
                            <strong>1 hour free</strong> 🎁 for Early bird (before 10 AM)⏰
                        </p>
                    <!-- 안나옴. ㅎ -->
                    <!-- <div class="table-container">
                    <DataTable :value="products" tableStyle="min-width: 50rem">
                        <Column v-for="col of columns" :key="col.field" :field="col.field" :header="col.header"></Column>
                    </DataTable> 
                    </div> -->
                    <div id="announcement">    
                        <!-- <pre>
                        <strong>Important Notice:</strong>
                        🎈 Children under 12 months, free entry for 2 hours. (ID or Certificate may be required)
                        🎈 Kids over 12 years old will be charged as adults
                        🎈 15% surcharge applied on public holidays
                        🎈 Twinkle Kids Cafe prices vary by location
                        </pre> -->
                        <p><strong>Important Notice:</strong></p>
                        <p>🎈 Children under 12 months, free entry for 1 hour. (ID or Certificate should be required)</p>
                        <p>🎈 Kids over 12 years old will be charged as adults</p>
                        <!-- <p>🎈 15% surcharge applied on public holidays</p> -->
                        <p>🎈 Twinkle Kids Cafe prices vary by location</p>
                    </div>
                </div>
            </div>
        </div>
        <div class="event-wrapper">
            <h1>Our Special</h1>
            <div class="event-cards-container"> 
                <!-- 
                <div class="event-card">
                    <div class="event-title">Buy 1 Hour, Get 1 Hour FREE</div>
                    <hr style="border: none; border-top: 1px dashed #014739; margin: 25px 0;" />
                    <div class="event-content">
                        <li>Valid for entrance before 10am</li>
                        <li  style="color: #f0598b;">Monday to Friday Only, <br>Except Public holidays and School holidays</li><br>
                    </div>
                </div>
                -->
                <div class="event-card">
                    <div class="event-title">
                        2nd Child Half Price<br>(Unlimited Ticket Purchase Only)
                    </div>
                    <hr style="border: none; border-top: 1px dashed #014739; margin: 25px 0;" />
                    <div class="event-content">
                        <li>The Second Child in the family will be charged half price.</li>
                        <li>Full price applies to every odd-numbered child. Every second child in a pair receives 50% off.</li>
                    </div>
                </div>
                <!-- 
                    <div class="event-card">
                        <div class="event-title">
                            Over $40, 1 hour FREE Extended
                        </div>
                        <hr style="border: none; border-top: 1px dashed #014739; margin: 25px 0;" />
                        <div class="event-content">
                            <li>When you order FOOD from our cafe menu over $40, you will get another free hour for one child.</li>
                        </div>
                        
                            <div class="event-title">Under 12 months, Free entry for 2 hours</div>
                            <hr style="border: none; border-top: 1px dashed #014739; margin: 25px 0;" />
                            <div class="event-content">
                                <li>We required ID or Certificate</li>
                                <li>A regular kids' entry fee will be charged after 2 hours</li>
                            </div>
                    
                    </div>
                -->
            </div>
        </div>
        <Footer/>
    </div>
</template>

<script>
import axios from 'axios'; // axios를 import 추가
import NavBar from './NavBar.vue';
import Footer from '@/components/Footer.vue';
// import DataTable from 'primevue/datatable';
// import Column from 'primevue/column';

export default {
    components: {
        NavBar,
        Footer
    },
    data(){
        return{
            // products: [
            //     { code: 'P001', name: 'Product 1', category: 'Category 1', quantity: 10 },
            //     { code: 'P002', name: 'Product 2', category: 'Category 2', quantity: 20 },
            // ],
            branchName: '',
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
        getBranchImage(branch_id) {
            const images = {
                'burwood': "https://images.squarespace-cdn.com/content/v1/637d8d8a7f609c521ddd5429/1672359522132-RU2ZPENTVALEBF0Z47PG/285887484_694866768237604_5851615251096205906_n.jpg",
                'hornsby': "/images/hornsby_a.jpg"
            };
            return images[branch_id]; // 기본 이미지
        },
        getBranchNameByID(id) {
        const branchMap = {
            'burwood': 'Burwood',
            'hornsby': 'Hornsby',
            // 필요하면 계속 추가 가능
        }
        return branchMap[id] || 'Unknown Branch'
        },
        async fetchprice(){
            console.log("Branch ID:", this.branchID);  // 값이 제대로 있는지 확인
            try {
                const response = await axios.get(`${process.env.VUE_APP_API_BASE}/api/price/${this.branchID}`);
                //const response = await axios.get(`https://kids-cafe-rm9g.onrender.com/api/price/${this.branchID}`);

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
            if (index === 0) return '#ffffff'; // 첫 번째 행
            else if (index === 1) return '#f0fff0'; // 두 번째 행
            else return '#f0f8ff'; // Unlimited 행
        }
    },
    created() {
        //const branchID = this.$route.params.branchID;
        this.fetchprice();  // 컴포넌트 생성 시 메뉴 데이터 가져오기
        const branchID = this.$route.params.branchID
        this.branchID = branchID; // 👈 이게 빠져있었음!
        this.branchName = this.getBranchNameByID(branchID)
    }
}
</script>

<style scoped>

h2{
    text-align: left;
    color: #eb6d54;
}

.event-wrapper{
    background-color: white;
    padding: 40px 20px 100px 20px;
    text-align: center;
}

.event-content li{
    text-align: left;
    margin-left: 30px; /* 또는 padding-left: 20px; */
}

.event-title{
    color: #014739;
    font-size: 20px;
    font-weight: bold;
    margin-top: 20px;
    margin-bottom: 25px;
}
.event-cards-container {
  display: flex;
  justify-content: center;
  gap: 20px; /* 카드 사이 간격 */
  flex-wrap: wrap; /* 화면 작으면 줄바꿈 */
  margin-top: 20px;
}

.event-card {
  background-color: #d9fff7;
  border: 1px solid #d9fff7;
  border-radius: 20px; /* 둥근 테두리 */
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
  padding: 20px;
  width: 400px;
  min-height: 150px;
  display: columns;
  align-items: center;
  justify-content: center;
  text-align: center;
  transition: transform 0.3s;
}

.event-card:hover {
  transform: translateY(-5px);
}

p{
    text-align: left;
   
}

.table-wrapper {
  display: flex;
  flex-direction: column; /* ← 요거만 세로로 만드는 핵심 */
  gap: 16px;
  margin: 0 auto; /* 가운데 정렬 */
  padding: 16px;
}

.price-table{
    width: 100%;
    color: black;
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