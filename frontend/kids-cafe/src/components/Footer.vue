<template>
    <footer>
        <div class="footer-content">
              <img src="@/assets/Logo-with-white-border-02-370x146.png" class="logo_image" />
              <!-- 
              <strong style="font-size: 25px; text-align: center;">Twinkle Kids Cafe</strong>
              --> 
              <div style="padding: 20px;"></div>
              <p>{{ branch.location }}</p>
              <p>⏰ OPEN 7DAYS</p>
              <p>{{ weekdayHours }}</p>
              <p>{{ weekendHours }}</p>
              <p>Check Weekly Updates from our Google Map & Instagram</p>
        </div>
        <div class="insta">
                Follow us!
                <div>
                  <a :href="getInstaLink"
                  target="_blank" class="instagram-link">
                  <img src="@/assets/insta.png" alt="Instagram" class="insta-icon" />
                  </a>
                </div>          
        </div>
    </footer>
</template>
<script>
import axios from 'axios'; // axios를 import 추가

export default {
    data() {
    return {
      // branchID를 기본 값으로 비워두거나 빈 값으로 초기화
      branchID: this.$route.params.branchID || this.$route.query.branch_id || '',
      branches:[]
    };
  },
  computed: {
    branch() {
    return this.branches.find(branch => branch.branch_id === this.branchID) || {};
    },
    weekdayHours() {
      return this.branchID === 'hornsby'
        ? 'Mon~Fri: 9am - 5:30pm'
        : 'Mon~Fri: 9am - 6pm';
    },
    weekendHours() {
      return this.branchID === 'hornsby'
        ? 'Sat~Sun: 9am - 6pm'
        : 'Sat~Sun: 9am - 7pm';
    },
    // branchID에 맞는 인스타그램 링크를 동적으로 반환
    getInstaLink() {
      const instagramBaseUrl = 'https://www.instagram.com/'; // 기본 URL

      console.log("#$#$#$#$ footer branchID #$#$@#$$ = ", this.branchID);

      // branchID에 따라 다르게 링크 설정
      if (this.branchID === 'burwood') {
        return instagramBaseUrl + 'twinklekidscafe_burwood'; // Burwood branch 인스타그램 링크
      } else if (this.branchID === 'hornsby') {
        return instagramBaseUrl + 'twinklekidscafe_hornsby'; // Hornsby branch 인스타그램 링크
      } else {
        return instagramBaseUrl + 'twinklekidscafe_burwood'; // 기본 링크 (기본값)
      }
    },
  },
  mounted() {
    this.fetchBranches();
  },
  methods: {
    async fetchBranches() {
      
      try {
        //const response = await axios.get("https://kids-cafe-rm9g.onrender.com/api/branches"); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능
        const response = await axios.get(`${process.env.VUE_APP_API_BASE}/api/branches`); // Proxy를 설정했으므로 백엔드 주소 없이 호출 가능
        
        //firebase관련 추가로 repsonse.data뒤에 branches 입력
        this.branches = response.data.branches;
        //console.log("### 전체 response 객체 ### :", response);
        console.log("### Branches data 나오라고 ### :", response.data);
      } catch (error) {
        console.error("#### Error fetching branches ##### :", error);
      }
    }
  },

}
</script>

<style scoped>

@font-face {
    font-family: 'Ownglyph_ParkDaHyun';
    src: url('https://fastly.jsdelivr.net/gh/projectnoonnu/2411-3@1.0/Ownglyph_ParkDaHyun.woff2') format('woff2');
    font-weight: normal;
    font-style: normal;
}

footer {
  padding: 20px 0;
  text-align: left;
  background-color: #faeca7;
  width: 100%;
}


.logo-container {
    flex: 1;  /* 필요에 따라 비율을 조절하세요 */
    align-items: flex-start;
    margin: 50px;
}


.logo_image{
  max-height: 60px;
  height: auto;
  max-width: 100%;
}

.insta {
    flex: 1;  /* 필요에 따라 비율을 조절하세요 */
    margin-right: 50px;
    text-align: center;
    color: black;
}

.footer-content {
  font-size: 18px;
  text-align: center;
  font-weight: bold;
  letter-spacing: 1px;
  text-transform: uppercase;
  /* padding: 5%; */
  max-width: 600px;   /* ✅ 최대 너비 제한 */
  margin: 0 auto;     /* ✅ 수평 가운데 정렬 */
  padding: 40px 20px;
  color: black;
}

.footer-content p {
  font-size: 12px;
  text-align: left;
  white-space: nowrap;
  justify-content: center; 
}

.instagram-link {
  display: inline-block;
  padding: 10px;
}

.insta-icon {
  width: 40px;
  height: 40px;
  object-fit: contain;
  transition: transform 0.3s ease;
}

.insta-icon:hover {
  transform: scale(1.1);
}

@media screen and (max-width: 768px) {
  .footer-content {
    font-size: 14px;
    padding: 20px 15px;
    text-align: center;
    word-break: break-word;       /* ✅ 긴 단어 줄바꿈 */
    overflow-wrap: break-word;    /* ✅ 줄바꿈 보조 */
    white-space: normal;          /* ✅ 자동 줄바꿈 */
    box-sizing: border-box;       /* ✅ 패딩 포함 */
  }

  .footer-content p {
    font-size: 13px;
    text-align: center;
    white-space: normal;          /* ✅ 자동 줄바꿈 */
    margin: 5px 0;
  }

  .footer-content strong {
    font-size: 18px;
    display: block;
    margin-bottom: 10px;
  }

  .insta {
    text-align: center;
    margin-top: 20px;
  }

  footer {
    width: 100%;
    overflow-x: hidden;   /* ✅ 가로 스크롤 제거 */
  }
}

</style>