<template>
    <div id="app">
        <BookingBar/>
        <div class="login-wrapper">
        <h2>Login</h2>
        <form @submit.prevent="handleLogin">
            <input v-model="email" type="email" placeholder="Email" required />
            <input v-model="password" type="password" placeholder="Password" required />
            <button type="submit">Login</button>
            <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
        </form>
        </div>
    </div>
  </template>
  
  <script>
  import { signInWithEmailAndPassword } from "firebase/auth";
  import { auth } from "@/firebase";
  import BookingBar from '../components/BookingBar.vue';
  
  export default {
    components:{
        BookingBar
    },
    data() {
      return {
        email: "",
        password: "",
        errorMessage: ""
      };
    },
    methods: {
      async handleLogin() {
        try {
          const userCredential = await signInWithEmailAndPassword(auth, this.email, this.password);
          const token = await userCredential.user.getIdToken();
          localStorage.setItem("authToken", token);
          this.$router.push("/admin"); // 로그인 후 이동할 경로
        } catch (error) {
          this.errorMessage = "Login failed: " + error.message;
        }
      }
    }
  };
  </script>
  
<style scoped>
.login-wrapper {
  max-width: 400px;
  margin: 80px auto;
  padding: 30px;
  background-color: #ffffff;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  text-align: center;
}

h2 {
  margin-bottom: 24px;
  font-size: 28px;
  color: #333;
}

input {
  display: block;
  width: 100%;
  padding: 14px 16px;
  margin: 12px 0;
  font-size: 16px;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-sizing: border-box;
  transition: border-color 0.2s ease;
}

input:focus {
  border-color: #4A90E2;
  outline: none;
}

button {
  width: 100%;
  padding: 14px;
  margin-top: 16px;
  background-color: #4A90E2;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

button:hover {
  background-color: #357ABD;
}

.error {
  margin-top: 10px;
  color: red;
  font-size: 14px;
}
</style>
  