<template>
    <div id="app">
        <NavBar/>
        Please note that prices may vary depending on the location.
        <div class="main">
            
            <div class="menu-wrapper">
                
                <div v-for="(categoryMenus, category) in categorizedMenus" :key="category">
                <div class="category-wrapper" :style="{ backgroundColor: coloredCategories[category] }">
                    <h2 class="category-title">
                        <!-- <i :class="['fas', categoryIcons[category]]"></i> -->
                        <span class="category-emoji">{{ categoryEmojis[category] }}</span> <!-- 이모지 추가 -->
                        {{ category }}
                    </h2>
                        <!-- 메뉴 항목들. -->
                        <div v-for="menu in categoryMenus" :key="menu.MenuID" class="cafemenu-item">
                            <div class="menu-name">{{ menu.MenuName }}</div>
                            
                            <ul class="menu-options">
                                <!-- <li v-for="(option, index) in menu.MenuOptions" :key="index">
                                    <span v-if="option.Size">{{ option.Size }} - </span>${{ option.Price }}
                                </li> -->

                                <!-- MenuOptions가 있고 길이가 1 이상일 경우 -->
                                <li v-if="menu.MenuOptions && menu.MenuOptions.length">
                                    <li v-for="(option, index) in menu.MenuOptions" :key="index">
                                        <span v-if="option.Size">{{ option.Size }} - </span>${{ option.Price }}
                                    </li>
                                </li>

                                <!-- MenuOptions가 없을 경우 단일 가격 보여주기 -->
                                <li v-else>
                                ${{ menu.Price }}
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
                <router-link :to="`/admin/menu/${branchID}`">
                    <button>메뉴 등록 버튼</button>
                </router-link>

            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'; // axios를 import 추가
import NavBar from './NavBar.vue';

export default {
    components: {
        NavBar,
    },
    data(){
        return{
            menus:[],
            categorizedMenus: {},  // 카테고리별로 메뉴를 나눠 저장할 객체
            activeMenu: null, // 클릭된 메뉴를 추적하는 변수
            branchID: this.$route.params.branchID,
            categoryColors: {
            'COFFEE & HOT TEA': '#FFEBEE',
            'KIDS MENU': '#E8F5E9',
            'SNACK': '#E3F2FD',
            'COLD DRINKS': '#FFF3E0',
            // 기본 색상
            'default': '#F5F5F5'
            }
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

            console.log("Branch ID:", this.branchID);  // 값이 제대로 있는지 확인
            try {
                const response = await axios.get(`http://localhost:8081/api/menu/${this.branchID}`);
                this.menus = response.data.menus;
                this.categorizeMenu(); 
                console.log("### menu data 나오라고 ### :", response.data.menus);
                
            } catch (error) {
                console.error("#### Error fetching menus ##### :", error);
            }
        },       
        categorizeMenu() {
            //카테고리 순서를 정해둬야함..
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
    },
    computed: {
        coloredCategories() {
            const colors = ['#FFEBEE', '#E8F5E9', '#E3F2FD', '#FFF3E0', '#F3E5F5']
            const result = {}
            const categories = Object.keys(this.categorizedMenus)
            categories.forEach((cat, index) => {
            result[cat] = colors[index % colors.length]
            })
            return result
        }
    }
}
</script>

<script setup>
const categoryIcons = {
  'PIZZA': 'fa-pizza-slice',
  'TWINKLE STAR': 'fa-star',
  'SNACK': 'fa-cookie-bite',
  'KIDS MENU': 'fa-child',
  'ALL DAY BREAKFAST': 'fa-egg',
  'BURGERS & SPAGHETTI': 'fa-hamburger',
  'COLD DRINKS': 'fa-glass-whiskey',
  'SALAD': 'fa-leaf',
  'COFFEE & HOT TEA': 'fa-mug-hot',
}

const categoryEmojis = {
  'PIZZA': '🍕',
  'TWINKLE STAR': '🌟',
  'SNACK': '🍪',
  'KIDS MENU': '👶',
  'ALL DAY BREAKFAST': '🍳',
  'BURGERS & SPAGHETTI': '🍔🍝',
  'COLD DRINKS': '🥤',
  'SALAD': '🥗',
  'COFFEE & HOT TEA': '☕',
}
</script>

<style scoped>
#app {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

.main {
  display: flex;
  justify-content: left;
  padding: 120px 0 40px; /* top padding으로 헤더 피하기 */
}

.menu-wrapper {
  width: 100%;
  max-width: none; /* 최대 너비 설정해서 가운데 정렬 */
  text-align: left;
  padding: 60px 1rem;

  display: grid;
  grid-template-columns: repeat(2, 1fr); /* 2열 */
  gap: 2rem;
}

@media (max-width: 768px) {
  .menu-wrapper {
    grid-template-columns: 1fr; /* 모바일에선 1열 */
  }
}

.cafemenu-item {
  display: flex;
  align-items: center;
  gap: 12px;
  flex-wrap: wrap; /* 혹시라도 넘치면 자동 줄바꿈 */
  /*margin-bottom: 24px;*/

  word-wrap: break-word; /* 긴 단어를 적절히 잘라서 다음 줄로 넘기기 */
  white-space: normal;   /* 텍스트가 너무 길면 자동으로 줄 바꿈 */
}

.menu-name {
  font-weight: bold;
  font-size: 1.2rem;
  margin-bottom: 6px;
  white-space: nowrap;

  word-wrap: break-word; /* 긴 단어를 적절히 잘라서 다음 줄로 넘기기 */
  white-space: normal;   /* 텍스트가 너무 길면 자동으로 줄 바꿈 */
}

.menu-options {
  display: flex;
  /* flex-wrap: wrap; 한 줄 넘치면 줄바꿈 허용 */
  gap: 12px;        /* 옵션 간 간격 */
  padding: 0;
  margin: 0;
  list-style: none;
}

.menu-options li {
  /*display: inline;
  display: flex;
  align-items: center;
  margin-right: 12px;  옵션 간 간격 */
  white-space: nowrap;
}

.category-title{
    color: #ff6600; 
    
}


ul {
  list-style-type: none; /* 동그라미 기호 제거 */
}

li {
  list-style-type: none; /* li에 대한 동그라미 기호 제거 (선택 사항) */
}

.logo_image{
    width: 100%;  /* 부모 요소 크기에 맞춤 */
    height: auto;  /* 가로 비율에 맞춰 세로 크기 자동 조정 */
    max-width: 200px; /* 최대 크기 제한 */
}

.category-wrapper {
    
    padding: 1rem;          /* 안쪽 여백 */
    margin: 1rem 0;         /* 위아래 여백 */

    border-radius: 12px;
 
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s ease;
    
    /** 2열로 보여주고 싶을 때 */
    /* display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1rem; 
  margin-bottom: 2rem; */
}

</style>
