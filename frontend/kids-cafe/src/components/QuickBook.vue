<template>
    <div>
        <header class="booking-header">
                <div id="sb_menu" class="header_menu_wrapper">
                    <ul class="header_navigation_nav" id="sb_menu_list_item_container">
                        <li class="header__nav-item">
                            <a class="go-home" href="/" target="_self">Twinkle Kids Cafe </a>
                        </li>
                        <li class="header__nav-item">
                            <a class="go-my-booking" href="#client/bookings/type/upcoming">My Bookings</a>
                        </li>
                        <li class="header__nav-item">
                            <a class="go-faq" href="#client/bookings/type/upcoming">FAQ</a>
                        </li>
                    </ul>
                </div>
            </header>
            <div class="wrapper">
                <div><h1 class="company-logo-and-name">Quick Party Room Booking</h1></div>
                <div class="background-wrapper">
                    <img src="/images/party-food.jpg" alt="Background" class="background-image" />
                </div>
            </div>
            <div class="button-and-time">
                <!-- 현재 시간 -->
                <div class="current-time"><CurrentTime/></div>
            </div>
        <form @submit.prevent="submitForm">
            <div class="quick-booking-info">
                <div class="info-detail">
                <label class="label-required">Branch: </label>
                <div class="select-container">
                    <select v-model="branchID" @change="fetchPartyrooms" class="styled-select">
                    <option value="">Select a branch</option>
                    <option v-for="branch in branches" :key="branch.id" :value="branch.id">
                        {{ branch.branch_name }}
                    </option>
                    </select>
                </div>
                </div>
                    <div class="info-detail">
                        <!-- <img src="/images/birthday-dino.png" /> -->
                        <label class="label-required">Contact Person Name: </label>
                        <!-- <input type="text" class="info-detail-input" /> -->
                        <input type="text" v-model="branchName" placeholder="Enter your name" class="styled-input" />

                    </div>
                    <div class="info-detail">
                        <!-- <img src="/images/birthday-dino.png" /> -->
                        <label>Email: </label><input type="text" class="info-detail-input" />
                    </div>
                    <div class="info-detail">
                        <!-- <img src="/images/birthday-dino.png" /> -->
                        <label class="label-required">Phone: </label><input type="text" class="info-detail-input" />
                    </div>
                    <div class="info-detail">
                        
                        <label class="label-required">Number of people</label>
                    
                            <div class="input-group">
                            <span>Adults:</span>
                            <input type="text" v-model="adultsValue" @input="validateInput('adults', adultsValue)" />
                            <span v-if="errorMessage.adults" class="error-message">{{ errorMessage.adults }}</span>
                          
                            <span >Children:</span>
                            <input type="text" v-model="childrenValue" @input="validateInput('children', childrenValue)" />
                            <span v-if="errorMessage.children" class="error-message">{{ errorMessage.children }}</span>
                            </div>
                      
                     
                        <!-- <img src="/images/birthday-dino.png" />
                        <label>Number of people</label>
                        <div class="info-detail">
                            <label class="label-required">Adults: </label>
                            <input type="text"
                            v-model="inputValue"
                            @input="validateInput('adults')" /><br>
                            <p v-if="errorMessage.adults" class="error-message">{{ errorMessage.adults }}</p>
                           
                            <label class="label-required">Children: </label>
                            <input type="text" class="info-detail-input-1" 
                            v-model="inputValue"
                            @input="validateInput('children')" /><br>
                            <p v-if="errorMessage.children" class="error-message">{{ errorMessage.children }}</p>
                        </div> -->
                    </div>
                    <div class="info-detail">
                        <!-- <img src="/images/birthday-dino.png" /> -->
                        <label class="label-required">Type of Party room: </label>
                        <select class="info-detail-input">
                            <option value="">Select a party room</option>
                            <option v-for="room in partyrooms" :key="room.id">
                                {{ room.RoomName }}
                            </option>
                        </select>
                    </div>
                    <!-- ###################  날짜 ############################## -->
                    <!-- ###################  날짜 ############################## -->
                    <div class="info-detail" style="gap: 0px; align-items: flex-start;">
                        <!-- <img src="/images/birthday-dino.png" /> -->
                        <label class="label-required " >Choose the date: </label>

                    <br/>
                  
                    <div class="calendar-container">
                          <!-- vue-cal 달력 컴포넌트 -->
                        <vue-cal
                        class="vuecal--blue-theme"
                        selected-date="2025-03-17"
                        xsmall
                        :time-from="10 * 60"
                        :disable-views="['day']"
                        events-count-on-year-view
                        active-view="month"
                        :events="events">
                
                    </vue-cal>
                        
                        
                    </div>
                        <pro-calendar
                            :events="events"
                        />
                    </div>
                    <!-- ##### 푸드 옵션 관련 테이블 ###### -->
                    <div class="info-detail">
                        <!-- <img src="/images/birthday-dino.png" /> -->
                        <label>Food Package Option</label>
                        <table class="info-detail-input-table">
                            <thead>
                                <tr>
                                <th>Select</th>
                                <th>No</th>
                                <th>Food Name</th>
                                <th>Quantity</th>
                                <th>Package</th>
                                <th>Price</th>
                                <th>For Kids</th>
                                <th>Description</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="(foodoption, index) in foodoptions" :key="foodoption.food_id">
                                <td>
                                    <input 
                                    type="checkbox" 
                                    v-model="selectedfoodoptions" 
                                    :value="foodoption.food_id" 
                                    @change="handleCheckboxChange(foodoption.food_id)"
                                    />
                                </td>
                                <td>{{ index + 1 }}</td>
                                <td>{{ foodoption.FoodName }}</td>
                                <td>{{ foodoption.FoodQuantity }}</td>
                                <td>{{ foodoption.FoodPackage === 'true' ? 'Package' : 'Individual' }}</td>
                                <td>${{ foodoption.FoodPrice }}</td>
                                <td>{{ foodoption.FoodForKids }}</td>
                                <td>{{ foodoption.FoodDescription }}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div> 
                    <div class="info-detail">
                        <!-- <img src="/images/birthday-dino.png" /> -->
                        <label>Party Decoration</label>
                        <input type="checkbox">Would you like to add party decorations?
                        (Please request detailed information regarding the theme upon completion of the reservation.)
                    </div> 
                    <div class="info-detail">
                        <!-- <img src="/images/birthday-dino.png" /> -->
                        <label>Please let us know if you need anything else for the party.</label>
                        <input class="info-detail-input-2"  type="text"/>
                    </div>    
                    <div>
          
                </div>
                    <div>
                        <button type="submit" class="quickbtn-book" href="#" title="Confirm Booking">Submit</button>
                    </div>
            </div>
       
        </form>
    </div>
</template>
<script>
import VueCal from 'vue-cal';   // vue-cal 임포트
import 'vue-cal/dist/vuecal.css'

import axios from 'axios'; // axios를 import 추가
import CurrentTime from '../components/CurrentTime.vue';
//import Vue3Datepicker from 'vue3-datepicker';
import { ref, computed, watch } from "vue"; // ✅ watch 추가

const config = {
  hideWeekends: true,
  time: true

}


export default {
    data() {
        return {
            selectedEvent: {},
            showDialog: false,
            branches:[],
            partyrooms:[],
            foodoptions: [], // 음식 옵션을 저장할 배열
            selectedfoodoptions: [], // 선택된 음식 옵션의 food_id를 저장할 배열
            branchID: "",  // 선택된 branch ID
            selectedDate: null, // ✅ null로 초기화 (날짜 객체 저장)
            adults: "",
            children: "",
            errorMessage: {
            adults: "",
            children: "",
            events: [  
                {
                    start: '2025-03-25 14:00',
                    end: '2025-03-25 18:00',
                    title: 'Need to go shopping',
                    icon: 'shopping_cart',
                    content: 'Click to see my shopping list',
                    contentFull: 'My shopping list is rather long:<br><ul><li>Avocados</li><li>Tomatoes</li><li>Potatoes</li><li>Mangoes</li></ul>',
                    class: 'leisure'
                }
            ]
            },

        };
    }, 
    components: {
        VueCal, 
        CurrentTime,
        
    }, 
    methods: {
        onEventClick (event, e) {
        this.selectedEvent = calendarEvents
        this.showDialog = true

        // Prevent navigating to narrower view (default vue-cal behavior).
        e.stopPropagation()
    },
        validateInput(field, inputValue) {
        // 입력값에서 숫자만 남기고 필터링
        inputValue = inputValue.replace(/\D/g, "");  // 숫자 이외의 모든 문자를 제거

        // 필드 값 갱신
        if (field === 'adults') {
            this.adultsValue = inputValue;
        } else if (field === 'children') {
            this.childrenValue = inputValue;
        }

        // 에러 메시지 처리
        if (inputValue === "") {
            this.errorMessage[field] = "Please enter number only";  // 숫자만 입력하라고 경고
        } else {
            this.errorMessage[field] = "";  // 유효한 값이면 에러 메시지 제거
        }
    },
        async fetchBranches() {
        try {
            const response = await axios.get(`${process.env.VUE_APP_API_BASE}/api/branches`); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능
            // const response = await axios.get("https://kids-cafe-rm9g.onrender.com/api/branches"); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

            this.branches = response.data.branches;
            //console.log("### 전체 response 객체 ### :", response);
            console.log("### branchs data 나오라고 ### :", response.data.branches);
        } catch (error) {
            console.error("#### Error fetching branchs ##### :", error);
        }
        },
        async fetchPartyrooms() {
            if (!this.branchID) {
                this.partyrooms = [];
                console.warn("branchID is empty, skipping API call");
                return;
            }
            try {
                console.log("Fetching party rooms for branchID:", this.branchID); // 디버깅용 로그
                //const response = await axios.get(`https://kids-cafe-rm9g.onrender.com/api/partyrooms/${this.branchID}`);
                const response = await axios.get(`${process.env.VUE_APP_API_BASE}/api/partyrooms/${this.branchID}`);

                this.partyrooms = response.data.partyrooms;
                console.log("$$$$ this.partyrooms.partyrooms $$$$$: ", this.partyrooms.partyrooms);
            } catch (error) {
                console.error("#### Error fetching party rooms:", error);
            }
        },
        async fetchfoodoptions() {
            try {
                //const response = await axios.get("https://kids-cafe-rm9g.onrender.com/api/foodoptions"); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능
                const response = await axios.get(`${process.env.VUE_APP_API_BASE}/api/foodoptions`); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능

                this.foodoptions = response.data.foodoptions;
                //console.log("### 전체 response 객체 ### :", response);
                console.log("### foodoptions data 나오라고 ### :", response.data.foodoptions);
            } catch (error) {
                console.error("#### Error fetching foodoptions ##### :", error);
            }
        }, 
        handleCheckboxChange(foodId) {
            if (this.selectedfoodoptions.length > 1) {
            // 한 번에 하나만 선택되도록 제한
            this.selectedfoodoptions = [foodId];
            }
        },
            // 날짜 클릭 시 실행되는 메서드
            handleDayClick(day) {
            const selectedDate = day.date.format('YYYY-MM-DD'); // 선택한 날짜
            console.log('선택한 날짜:', selectedDate);
            },
    },
    watch: {
            branchID(newVal) {
                console.log("@@@@ Selected branchID @@@@ :", newVal);
                this.fetchPartyrooms();
            }
    },
    mounted() {
        this.fetchBranches();
        this.fetchfoodoptions(); // 컴포넌트 마운트 후 데이터 로딩
        console.log("calendarEvents:", this.events);
        //console.log("Calendar Events:", this.calendarEvents);
        //this.fetchPartyrooms();
    },
}
</script>

<style scoped>

.label-required::after {
    content: " *";           /* 별표 추가 */
    color: #990000;              /* 빨간색 */
    font-weight: bold;        /* 굵은 글씨 */
    margin-left: 5px;         /* 텍스트와 별표 사이 간격 */
    margin-right: 25px;
}

/* .label-required:hover::after {
    content: "필수 입력 항목";
    color: #990000;
    font-weight: bold;
    font-size: 12px;
} */

.company-logo-and-name{
  color: white;
  top: 50%; /* 수직 중앙 정렬 */
  left: 50%; /* 수평 중앙 정렬 */
  text-align: center;
  /* transform: translate(-50%, -50%); /* 정확한 중앙 정렬을 위해 transform 사용 */
  font-size: 2em; /* 텍스트 크기 */
  font-weight: bold; /* 텍스트 두껍게 */
  z-index: 2; /* h1이 버튼 위에 오도록 z-index 설정 */
  margin-bottom: 10px;  /* h1과 버튼 간의 간격 */
}

.background-image{
    width: 100vw; /* 뷰포트 전체 너비를 차지하도록 설정 */
    margin: 0 auto; /* 가로 중앙 정렬 */
    display: block; /* inline 요소가 가진 여백 제거 */
    z-index: 0; /* 이미지가 가장 뒤에 있도록 z-index 설정 */
    height: 100%;
    position: absolute; /* 이미지 배경으로 설정 */
    top: 0;
    left: 0;
    object-fit: cover; /* 이미지 비율에 맞게 채우기 */
    opacity: 0.8;  /* 이미지 불투명도 설정, 1이면 불투명, 0이면 완전 투명 */
    /* filter: brightness(50%);  이미지 어둡게 */
}

.booking-header{
    width: 100vw;
    display: flex;
    justify-content: flex-end; /* 요소들을 가로로 균등하게 정렬 */
    /*align-items: right;  세로 정렬 */
    /*padding: 20 20px;  좌우 여백 추가 (너무 화면 끝까지 붙지 않도록) */
    padding: 20px 40px 20px 20px; /* 위쪽 20px, 오른쪽 40px, 아래쪽 20px, 왼쪽 20px */
    background-color: #fff; /* 배경색 설정 (원하는 색상으로 변경 가능) */
    margin: 0; /* 헤더의 외부 여백 제거 */
    position: fixed; /* 화면 맨 위에 고정 */
    top: 0; /* 최상단 배치 */
    left: 0;
    /* right: 0; 오른쪽 끝까지 확장 */
    z-index: 1000; /* 다른 요소들 위에 배치 */
}

.current-time{
    display: flex;           /* flexbox 활성화 */
    width: 100vw;
    justify-content: flex-end;
    padding-right: 50px;;
}

.header_navigation_nav{
    display: flex;
    list-style: none;
    text-align: center;
    padding-right: 50px; /* 네비게이션 바 오른쪽 여백 추가 ############# 이게 오른쪽 여백 적용되는 부분################*/ 
    margin: 0;
    height: 30px; /* 메뉴 높이 설정 (예: 50px) */
}

.header__nav-item .header__nav-item{
    padding: 10px 15px; /* 메뉴 항목 내부 여백 */
    white-space: nowrap; /* 줄바꿈 방지 */
    min-width: 120px; /* 최소 너비 설정 */
    height: 100%; /* 메뉴 항목의 높이에 맞게 설정 */
    padding: 0 15px; /* 메뉴 항목의 좌우 여백 */

}
.header__nav-item a{
    color: #595959;
    height: 100%; /* 메뉴 항목 높이에 맞게 100%로 설정 */
    display: block; /* 클릭 영역 확보 */
    padding: 10px 15px; /* 링크 내부 여백 */
    text-decoration: none;
    font-size: 16px; /* 글자 크기 조정 */
}

.header__nav-item a:hover{
    border-bottom: 3px solid #6699ff; /* hover 시 선 표시 */
}

/** booking info css*/
.quick-booking-info {
    display: flex;               /* flexbox로 배치 */
    flex-direction: column;      /* 세로 방향으로 정렬 */
    gap: 15px;                   /* 항목 간의 간격 */
    padding: 20px;               /* 내부 여백 */
}

    /** 푸드 옵션 체크리스트 */
table {
    width: 100%;
    border-collapse: collapse;
    }

th, td {
    border: 1px solid #ccc;
    padding: 8px;
    text-align: left;
    }

th {
    background-color: #c6ecc6;
    }

    .info-detail-input {
        /*flex-grow: 1;                 입력 필드는 가용 공간을 모두 차지 */
        width: 60%;
        padding: 5px;                /* 입력 필드에 패딩 추가 */
        border: 1px solid #ccc;      /* 입력 필드 테두리 */
        border-radius: 2px;          /* 입력 필드 테두리 둥글게 */
        box-sizing: border-box;  /* box-sizing 설정 */
    }

    .info-detail label {
        margin-left: 30%;
        width: 30%;                 /* 레이블의 고정 너비 설정 */
        text-align: left;           /* 레이블 텍스트를 오른쪽 정렬 */
    }

    .info-detail-input-1{
        /* flex-grow: 1;                /* 입력 필드는 가용 공간을 모두 차지 */
        width: 60%; /* 입력 필드의 너비를 200px로 설정 */
        padding: 5px;                /* 입력 필드에 패딩 추가 */
        border: 1px solid #ccc;      /* 입력 필드 테두리 */
        border-radius: 2px;          /* 입력 필드 테두리 둥글게 */
        box-sizing: border-box;  /* box-sizing 설정 */
    }

    .info-detail {
        display: flex;               /* 내부 요소들 역시 flexbox */
        align-items: center;         
        gap: 10px;                   /* 레이블과 입력 필드 간의 간격 */
    }

/** for tab screen */
@media (max-width: 768px){

    .error-message {
        color: red;
        font-size: 12px;
        margin-top: 5px;
    }

    .info-detail {
        display: flex;               /* 내부 요소들 역시 flexbox */
        align-items: center;         /* 레이블과 입력 필드 수직 중앙 정렬 */
        gap: 10px;                   /* 레이블과 입력 필드 간의 간격 */
    }

    .info-detail-input {
        /*flex-grow: 1;                 입력 필드는 가용 공간을 모두 차지 */
        width: 100%;
        padding: 5px;                /* 입력 필드에 패딩 추가 */
        border: 1px solid #ccc;      /* 입력 필드 테두리 */
        border-radius: 5px;          /* 입력 필드 테두리 둥글게 */
        box-sizing: border-box;  /* box-sizing 설정 */
    }

    .info-detail label {
        margin-left: 5%;
        width: 70%;                 /* 레이블의 고정 너비 설정 */
        text-align: left;           /* 레이블 텍스트를 오른쪽 정렬 */
    }
    
    
    
    .info-detail-input-1{
        /* flex-grow: 1;                /* 입력 필드는 가용 공간을 모두 차지 */
        width: 50px; /* 입력 필드의 너비를 200px로 설정 */
        padding: 5px;                /* 입력 필드에 패딩 추가 */
        border: 1px solid #ccc;      /* 입력 필드 테두리 */
        border-radius: 5px;          /* 입력 필드 테두리 둥글게 */
        box-sizing: border-box;  /* box-sizing 설정 */
    }
    
    .info-detail-input-table{
        width: 80%; /* 입력 필드의 너비를 200px로 설정 */
        padding: 5px;                /* 입력 필드에 패딩 추가 */
    }
}


.info-detail-input-2{
    /* flex-grow: 1;                /* 입력 필드는 가용 공간을 모두 차지 */
    width: 50%; /* 입력 필드의 너비를 200px로 설정 */
    height: 150px;
    padding: 5px;                /* 입력 필드에 패딩 추가 */
    border: 1px solid #ccc;      /* 입력 필드 테두리 */
    border-radius: 5px;          /* 입력 필드 테두리 둥글게 */
    box-sizing: border-box;  /* box-sizing 설정 */
}

/** button img */
.quickbtn-book {
    display: inline-block;
    padding: 12px 20px;
    margin-left: 30px; /* 버튼과 메뉴 간에 간격을 줌 */
    background-color: #00b33c; /* 버튼 배경 색상 */
    color: white; /* 텍스트 색상 */
    text-align: center;
    text-decoration: none; /* 링크 스타일 제거 */
    border-radius: 8px; /* 둥근 모서리 */
    font-size: 16px; /* 글자 크기 */
}

.quickbtn-book:hover{
    background-color:#4d0099;
}

h1 {
  text-align: center;
  font-family: 'Arial', sans-serif;
  margin-bottom: 20px;
}

.calendar-container {
  display: flex;
  justify-content: center; /* 가로 중앙 정렬 */
  align-items: center; /* 세로 중앙 정렬 */
  /* gap: 20px; 캘린더 사이 여백 */
  width: 100%;
  /* padding: 20px; */
  /* //margin-left: 30%; */
}

.move-label {
    display: flex; /* 인풋을 블록 요소로 변경 */
    margin-left: 15%;
  }




.input-group {
  display: flex;
  align-items: center; /* 라벨과 입력 필드를 가로로 정렬 */
  gap: 10px;
}

.label-required {
  font-weight: bold;
  margin-bottom: 0;
}

.input-details input {
  padding: 8px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 4px;
  width: 100px;
}

.error-message {
  color: red;
  font-size: 12px;
  margin-left: 10px; /* 오류 메시지를 인풋 옆에 배치 */
}

/* Select 박스 컨테이너 */
.select-container {
  position: relative;
  width: 100%; /* 부모 너비에 맞춰서 늘어날 수 있도록 설정 */
  max-width: 300px; /* 최대 너비를 설정해줘서 너무 커지지 않게 조정 */
}

/* 세련된 Select 박스 */
.styled-select {
  width: 100%;
  padding: 12px 20px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 8px;
  background-color: #f9f9f9;
  color: #333;
  font-family: 'Arial', sans-serif;
  transition: all 0.3s ease;
  box-sizing: border-box;
  -webkit-appearance: none; /* 기본 스타일 제거 */
  -moz-appearance: none;
  appearance: none;
}

/* 셀렉트 박스의 화살표 스타일링 */
.styled-select::-ms-expand {
  display: none; /* IE에서 기본 화살표 제거 */
}

/* 포커스 시 변화 */
.styled-select:focus {
  border-color: #4CAF50; /* 포커스 시 초록색 테두리 */
  background-color: #fff; /* 포커스 시 배경 흰색 */
  outline: none;
  box-shadow: 0 0 5px rgba(0, 255, 0, 0.3); /* 부드러운 그림자 */
}

/* 옵션을 위한 스타일 */
.styled-select option {
  padding: 10px;
  font-size: 14px;
}

/* Select 박스 뒤에 커스텀 화살표 아이콘 추가 */
.select-container::after {
  content: '▼'; /* 화살표 기호 */
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  color: #888;
  font-size: 18px;
  pointer-events: none; /* 클릭 방지 */
}

/* 세련된 인풋 박스 */
.styled-input {
  width: 30%; /* 부모 컨테이너에 맞춰 100% 확장 */
  padding: 12px 20px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 8px;
  background-color: #f9f9f9;
  color: #333;
  font-family: 'Arial', sans-serif;
  transition: all 0.3s ease;
  box-sizing: border-box;
  outline: none;
}

/* 포커스 시 변화 */
.styled-input:focus {
  border-color: #4CAF50;
  background-color: #fff;
  box-shadow: 0 0 5px rgba(0, 255, 0, 0.3);
}

/* 인풋 텍스트를 입력할 때 placeholder 스타일 */
.styled-input::placeholder {
  color: #bbb;
}

/* 인풋 뒤에 화살표 아이콘을 원할 경우 추가할 수 있음 */
.styled-input-wrapper {
  position: relative;
}

.styled-input-wrapper::after {
  content: '▼'; /* 화살표 기호 */
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  color: #888;
  font-size: 18px;
  pointer-events: none;
}
</style>