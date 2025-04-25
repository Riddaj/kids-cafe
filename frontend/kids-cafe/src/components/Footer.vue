<template>
    <footer>
        <div class="footer-content">
            <div>
                Follow us!
            </div>
            <a :href="getInstaLink"
            target="_blank" class="instagram-link">
            <img src="@/assets/insta.png" alt="Instagram" class="insta-icon" />
            </a>
        </div>
    </footer>
</template>
<script>
export default {
    data() {
    return {
      // branchID를 기본 값으로 비워두거나 빈 값으로 초기화
      branchID: this.$route.params.branchID || '',
    };
  },
  computed: {
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

}
</script>

<style>
footer {
  padding: 20px 0;
  text-align: left;
  background-color: #faeca7;
  width: 100%;
}

.footer-content {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 20px; /* Add space between text and icon */
}

.footer-content div {
  font-size: 18px;
  font-weight: bold;
  letter-spacing: 1px;
  text-transform: uppercase;
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
</style>