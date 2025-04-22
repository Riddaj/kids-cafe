<template>
    <div id="app">
      
        <NavBar/>
        <div class="form-container">
        <h2>🍽️ Register a New Menu</h2>
        <form @submit.prevent="submitForm">
        <label>
            Menu ID:
            <input v-model="menu.MenuID" type="text" required />
        </label>

        <label>
            Menu Name:
            <input v-model="menu.MenuName" type="text" required />
        </label>

        <label>
            Category:
            <input v-model="menu.MenuCategory" type="text" required />
        </label>

        <label>
            Description:
            <textarea v-model="menu.Description" rows="3"></textarea>
        </label>

        <label>
            Image URL:
            <input v-model="menu.ImgUrl" type="url" />
        </label>

        <div class="options-section">
            <h3>Options</h3>
            <div v-for="(option, index) in menu.MenuOptions" :key="index" class="option">
            <input v-model="option.name" placeholder="Option Name" />
            <input v-model.number="option.price" placeholder="Price" type="number" />
            <button type="button" @click="removeOption(index)">❌</button>
            </div>
            <button type="button" @click="addOption">➕ Add Option</button>
        </div>

        <button type="submit" class="submit-btn">✅ Register</button>
        </form>
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
            menu: {
            MenuID: "",
            MenuName: "",
            MenuCategory: "",
            Description: "",
            ImgUrl: "",
            MenuOptions: [],
            branchID: this.$route.params.branchID,
            },
        };
    },
    mounted() {
        this.fetchmenu();  // 컴포넌트가 마운트되면 fetchmenu 호출
    },
    methods:{
        addOption() {
        this.menu.MenuOptions.push({ name: "", price: 0 });
        },
        removeOption(index) {
        this.menu.MenuOptions.splice(index, 1);
        },
        submitForm() {
        console.log("Submitting Menu:", this.menu);
        // 실제 등록 로직 여기에 추가 (예: axios POST)
        },
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
#app {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}
.form-container {
  max-width: 600px;
  margin: auto;
  padding: 2rem;
  padding-top: 110px;
  background-color: #fdfdfd;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

h2 {
  text-align: center;
  margin-bottom: 1.5rem;
}

form label {
  display: block;
  margin-bottom: 1rem;
  font-weight: bold;
}

input,
textarea {
  width: 100%;
  padding: 8px;
  margin-top: 0.3rem;
  border-radius: 6px;
  border: 1px solid #ccc;
}

.options-section {
  margin-top: 1.5rem;
}

.option {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}

.submit-btn {
  margin-top: 1.5rem;
  padding: 10px 20px;
  background-color: #4caf50;
  color: white;
  font-weight: bold;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  display: block;
  width: 100%;
}

.submit-btn:hover {
  background-color: #45a049;
}


</style>
