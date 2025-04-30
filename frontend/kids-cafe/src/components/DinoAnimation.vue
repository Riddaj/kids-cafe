<!-- DinoAnimation.vue -->
<template>
    <div class="dino-wrapper">
      <img src="@/assets/dino.png" class="dino" alt="Dino" />
    </div>
  </template>
  
  <style scoped>
  .dino-wrapper {
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 100px;
    pointer-events: none;
    overflow: hidden;
    z-index: 9999; /* 다른 요소 위에 올리기 */
  }
  
  .dino {
    position: fixed;
    width: 70px;
    height: auto;
    z-index: 1000;
    animation: moveAround 12s linear infinite,
               flipDirection 12s linear infinite;
  }

  @keyframes moveAround {
    0%   { top: 0; left: 0; }
    25%  { top: 0; left: calc(100vw - 70px); }
    50%  { top: calc(100vh - 70px); left: calc(100vw - 70px); }
    75%  { top: calc(100vh - 70px); left: 0; }
    100% { top: 0; left: 0; }
  }

  @keyframes flipDirection {
    0%   { transform: scaleX(1); }                        /* 오른쪽 */
    24.99% { transform: scaleX(1); }

    25%  { transform: rotate(90deg); }                    /* 아래쪽 */
    49.99% { transform: rotate(90deg); }

    50%  { transform: scaleX(-1); }                       /* 왼쪽 */
    74.99% { transform: scaleX(-1); }

    75%  { transform: rotate(-90deg); }                   /* 위쪽 */
    99.99% { transform: rotate(-90deg); }

    100% { transform: scaleX(1); }                        /* 다시 오른쪽 */
  }
  
  /* .dino {
    position: absolute;
    bottom: 0;
    height: 80px;
    animation: dinoMove 10s linear infinite alternate;
  }
  
  @keyframes dinoMove {
    0% {
      left: -100px;
    }
    100% {
      left: 100%;
    }
  } */
  </style>
  