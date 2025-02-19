<template>
    <div id="app"> 
        <header>
            <div id="sb_menu" class="header_menu_wrapper">
                <ul class="header_navigation_nav" id="sb_menu_list_item_container">
                    <li class="header__navigation-item menu-item ">
                        <a class="popup-hide header__navigation-link icon-pos--none icon-Twinkle Kids Cafe" href="/" target="_self">Twinkle Kids Cafe </a>
                    </li>
                    <li class="header__navigation-item menu-item ">
                        <a class="popup-hide header__navigation-link icon-pos--none icon-menu_bookings" href="#client/bookings/type/upcoming">My Bookings</a>
                    </li>
                </ul>
            </div>
        </header>
        <div class="wrapper">
            <h1 class="company-logo-and-name">
                Twinkle Kids Cafe
            </h1>

        </div>
        <div class="book-online-body">
            <div class="book-online">
                <h1 class="heading-title">Book online</h1>
                <p>2% surcharge will be apply for online booking. </p>
            </div>
            <div class="btn-bar1">
                    <a class="btn-book1" href="/book_a_party/simplybook" title="Book Now">
                        Book Now
                    </a>
            </div>
        </div>
        <div class="book-email-body">
            <div class="book-email">
                <h1 class="heading-title">Book via Email</h1>
                <p>There is no surcharge for email booking. </p>
            </div>
            <div class="bank-info">
                <p>Bank details are provided below.<br>Account Name: Twinkle Australia Pty Ltd<br>BSB: 062-128
                    <br>Account NO.: 11151130<br>* Please use your child’s name and party date as the description for the bank transfer.</p>
            </div>
        </div>
    </div>
</template>

<script>
export default {

}
</script>

<style scoped>
.book-online-body{
    display: flex;
    width: 100vw;
    text-align: left;
    padding: 20px;
}

/* 내부 div들이 각각 50% 비율을 차지하도록 설정 */
.book-online-body > div {
    flex: 1; /* 동일한 비율로 크기 조정 */
    max-width: 50%; /* 최대 너비 50%로 제한 */
}

.book-email-body{
    background-color: #ffe6ff;
    display: flex;
    width: 100vw;
    text-align: left;
    padding: 20px;
    box-sizing: border-box; /* padding 포함하여 width 계산 */
}

/* 내부 div들이 각각 50% 비율을 차지하도록 설정 */
.book-email-body > div {
    flex: 1; /* 동일한 비율로 크기 조정 */
    max-width: 50%; /* 최대 너비 50%로 제한 */
}

.btn-book1{
    text-align: center;
    text-decoration: none; /* 링크 스타일 제거 */
    border-radius: 8px; /* 둥근 모서리 */
    font-size: 16px; /* 글자 크기 */
    height: 30px;
    padding: 12px 20px;
    color: white;
    background-color: #4d0099;    
    text-decoration: none; /* 링크 스타일 제거 */ 
}

.btn-book1 a:hover{
    background-color: #ffcc00;
    color: white;
}
</style>