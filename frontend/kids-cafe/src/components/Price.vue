<template>
    <div id="">
        <NavBar/>
        <div class="content">
            <h1>Burwood Entry Ticket</h1>
            
            <h2>Kids Ticket</h2>


            <table style="border-collapse: collapse; width: 100%; font-family: 'Comic Sans MS', cursive;">
            <thead>
                <tr style="background-color: #ffe4e1;">
                <th style="padding: 12px;">🕒 Play Time per Child</th>
                <th style="padding: 12px;">📅 Weekday</th>
                <th style="padding: 12px;">🎉 Weekend & Public Holiday</th>
                </tr>
            </thead>
            <tbody>
                <tr style="background-color: #fffaf0;">
                <td style="padding: 12px 20px;">1</td>
                <td style="padding: 12px 20px;">$13</td>
                <td style="padding: 12px 20px;">$16</td>
                </tr>
                <tr style="background-color: #f0fff0;">
                <td style="padding: 12px 20px;">2</td>
                <td style="padding: 12px 20px;">$20</td>
                <td style="padding: 12px 20px;">$26</td>
                </tr>
                <tr style="background-color: #f0f8ff;"> <!-- Unlimited row -->
                <td style="padding: 12px 20px;">Unlimited</td>
                <td style="padding: 12px 20px;">$50</td>
                <td style="padding: 12px 20px;">$60</td>
                </tr>
            </tbody>
            </table>
            
                <p style="color: black;">
                    <strong>1 hour free</strong> 🎁 for Early bird (before 10 AM) & after 3 PM ⏰
                </p>
            <div id="announcement" style="text-align: left;">    
                
                <pre>
                    <strong>Important Notice:</strong>
                    🎈 Children under 12 months, free entry for 2 hours. (ID or Certificate may be required)
                    🎈 Kids over 12 years old will be charged as adults
                    🎈 15% surcharge applied on public holidays
                    🎈 Twinkle Kids Cafe prices vary by location
                </pre>
            </div>
            <!-- 안나옴. ㅎ -->
            <!-- <div class="table-container">
            <DataTable :value="products" tableStyle="min-width: 50rem">
                <Column v-for="col of columns" :key="col.field" :field="col.field" :header="col.header"></Column>
            </DataTable> 
            </div> -->
          
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
    
            menus:[],
            categorizedMenus: {},  // 카테고리별로 메뉴를 나눠 저장할 객체
            activeMenu: null, // 클릭된 메뉴를 추적하는 변수
           
        }
    },
    mounted() {
        this.fetchmenu();  // 컴포넌트가 마운트되면 fetchmenu 호출
    },
    methods:{
        setActiveMenu(headmenuName) {
            this.activeMenu = headmenuName; // 클릭된 메뉴를 추적
        },
        async fetchmenu(){
            try {
                const response = await axios.get("http://localhost:8081/api/cafe-menu");
                this.menus = response.data.menus;
                this.categorizeMenu(); 
                console.log("### menu data 나오라고 ### :", response.data.menus);
                
            } catch (error) {
                console.error("#### Error fetching menus ##### :", error);
            }
        },       
        categorizeMenu() {
            // menus 배열을 MenuCategory 기준으로 분류
            this.categorizedMenus = this.menus.reduce((categories, menu) => {
                const category = menu.MenuCategory;

                if (!categories[category]) {
                    categories[category] = [];  // 카테고리가 없으면 새 배열 생성
                }

                categories[category].push(menu);  // 카테고리에 해당하는 메뉴 추가
                return categories;
            }, {});
        }
    },
    created() {
        this.fetchmenu();  // 컴포넌트 생성 시 메뉴 데이터 가져오기
    }
}
</script>

<style scoped>


h1 {
    justify-content: center;

}

.table-container {
  display: flex;
  justify-content: center; /* 가로 중앙 정렬 */
  align-items: center; /* 세로 중앙 정렬 */
  
  width: 100%;
  }


.content {
    padding-top: 110px;

}




</style>