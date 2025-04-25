<template>
    <div id="app">
        <BookingBar/>
        <div>
            Congratulations!
            A birthday party always requires a lot of details!<br>
            We will contact you soon to discuss more details.<br>
            Please wait for a moment.<br>
            Thank you.
            <div>
                <h1>Booking Confirmation</h1>
                <table class="confirmation-table">
                    <thead>
                        <tr>
                            <th>Field</th>
                            <th>Details</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>Party Room Name</td>
                            <td>{{ bookingDetails.partyroom_name }}</td>
                        </tr>
                        <tr>
                            <td>Selected Date</td>
                            <td>{{ bookingDetails.partytime }}</td>
                        </tr>
                        <tr>
                            <td>Total Price</td>
                            <td>{{ bookingDetails.total_price }}</td>
                        </tr>
                        <tr>
                            <td>Kid's Name</td>
                            <td>{{ bookingDetails.kid_name }}</td>
                        </tr>
                        <tr>
                            <td>Owner's Name</td>
                            <td>{{ bookingDetails.owner_name }}</td>
                        </tr>
                        <tr>
                            <td>Phone</td>
                            <td>{{ bookingDetails.owner_phone }}</td>
                        </tr>
                        <tr>
                            <td>Email</td>
                            <td>{{ bookingDetails.email }}</td>
                        </tr>
                        <tr>
                            <td>Special Dietary Requirements</td>
                            <!-- <td>{{ bookingDetails.special_required.join(', ') }}</td> -->
                        </tr>
                    </tbody>
                </table>
            </div>
            <div class="button-container">
                <router-link :to="`/`">
                    <button type="submit" class="submit-button">go home</button>
                </router-link>
            </div>
        </div>  
</div>
</template>

<script>
import BookingBar from '../components/BookingBar.vue';

export default {
    components: {
        BookingBar
    },
    data(){
        return {
            bookingDetails: this.$route.query
        };
    },
    mounted() {
        const bookingData = this.$router.history.state;
        console.log("받은 정보:", bookingData);
    }

}
</script>

<style scoped>
.submit-button:hover{
    background-color: #6699ff; /* 버튼 배경 색상 */
}

.button-container {
    display: flex;
    justify-content: center;  /* 버튼을 오른쪽 정렬 */
    text-align: center; /* 버튼을 가로로 중앙 정렬 */
    margin-top: 20px; /* 버튼과 테이블 사이에 간격 추가 */
    margin-bottom: 30px;
}

.submit-button{
    background-color: #ffb3b3; /* 버튼 배경 색상 */
    color: white; /* 텍스트 색상 */
    text-decoration: none; /* 링크 스타일 제거 */
    padding: 10px 20px; /* 버튼 내부 여백 */
    border-radius: 5px; /* 둥근 모서리 */
    font-size: 16px; /* 글자 크기 */
    display: inline-block; /* 버튼 형태 유지 */
    text-align: center;
    margin-top: 30px; /* 버튼이 자동으로 글자 아래에 배치되도록 설정 */
}

</style>

