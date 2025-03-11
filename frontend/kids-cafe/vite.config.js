import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path' // 추가

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
    resolve: {
      alias: {
        '@': path.resolve(__dirname, 'src') // 올바른 별칭 설정
      }
    },
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8081', // Go 서버의 주소
        changeOrigin: true, // 요청 헤더의 Origin을 변경
        rewrite: (path) => path.replace(/^\/api/, '') // '/api' 제거
      },
    },
  },
})

// vue.config.js
module.exports = {
  productionSourceMap: false,  // 소스맵을 비활성화하여 문제 해결
}
