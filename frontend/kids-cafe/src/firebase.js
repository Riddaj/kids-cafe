import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";


// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
    apiKey: "AIzaSyCwAIELJOTVPKfc6UlfJtBp4JrqwBoWCDA",
    authDomain: "kids-cafe-booking-project.firebaseapp.com",
    projectId: "kids-cafe-booking-project",
    storageBucket: "kids-cafe-booking-project.appspot.com",
    messagingSenderId: "117428498096",
    appId: "1:117428498096:web:6f176062066731e7829575",
    measurementId: "G-T7JW6ZECVJ"
  };

const app = initializeApp(firebaseConfig);
const auth = getAuth(app);

export { auth };
