<template>
    <div id="app">
      
        <NavBar/>
        <div class="form-container">
        <h2>🍽️ Register a New Menu</h2>
        <form @submit.prevent="submitForm">
        <!-- <label>
            Menu ID:
            <input v-model="menu.MenuID" type="text" required />
        </label> -->

        <label>
            Menu Name:
            <input v-model="menu.MenuName" type="text" required />
        </label>

        <label>
            Category:
            <!-- BURWOOD -->
            <select v-model="menu.MenuCategory" required >
                <option disabled value=""><label>-- Select Category --</label></option>
                <option>PIZZA</option>
                <option>TWINKLE STAR</option>
                <option>SNACK</option>
                <option>KIDS MENU</option>
                <option>ALL DAY BREAKFAST</option>
                <option>BURGERS & SPAGHETTI</option>
                <option>COLD DRINKS</option>
                <option>SALAD</option>
                <option>COFFEE & HOT TEA</option>
            </select>
        </label>

        <!-- Price || Options 필수-->
        <label>
            Price:
            <input v-model.number="menu.Price" type="number" step="0.1"/>
        </label>

        <label>
            Description:
            <textarea v-model="menu.Description" rows="3"></textarea>
        </label>

        <!-- <label>
            Image URL:
            <input v-model="menu.ImgUrl" type="url" />
        </label> -->


        <!-- Flavors. 맛 종류. 가격은 같음. -->
        <div class="flavors-section">
            <h3>Flavors</h3>
            <div v-for="(flavor, index) in menu.Flavors" :key="'flavor-' + index" class="flavor">
                <input v-model="menu.Flavors[index]" placeholder="Flavor Name" />
                <button type="button" @click="removeFlavor(index)">❌</button>
            </div>
            <button type="button" @click="addFlavor">➕ Add Flavor</button>
        </div>

        <!-- 옵션 부분. 옵션별 가격이 다름.-->
        <div class="options-section">
            <h3>Options</h3>
            <div v-for="(option, index) in menu.MenuOptions" :key="index" class="option">
            <input v-model="option.size" placeholder="Option" />
            <input v-model.number="option.price" placeholder="Price" type="number" step="0.1"/>
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
            Price: 0,
            branchID: "",
            Flavors: [],
            },
        };
    },
    mounted() {
        
        this.branchID = this.$route.params['branchID']
    },
    methods:{
      
        addOption() {
        this.menu.MenuOptions.push({ size: "", price: 0 });
        },
        removeOption(index) {
        this.menu.MenuOptions.splice(index, 1);
        },
        addFlavor() {
        this.menu.Flavors.push('');
        },
        removeFlavor(index) {
        this.menu.Flavors.splice(index, 1);
        },
        submitForm() {
            //console.log(this.$route.params['branchID'])
            console.log("Branch ID:", this.branchID);  // 값이 제대로 있는지 확인
            
            console.log("Submitting Menu:", this.menu);

        
        axios.post(`http://localhost:8081/api/addmenu/${this.branchID}`, this.menu)
        .then(res => {
            console.log("등록 성공:", res.data);

            this.$router.push({ name: 'menu', params: { branchID: this.branchID } });
        })
        .catch(err => {
            console.error("등록 실패:", err);
        });

     
        },
     
    },
    created() {
        
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
textarea,
select {
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
