import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'

const backendPort = process.env.PORT || '8080'
const backendTarget = `http://localhost:${backendPort}`

export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: backendTarget,
        changeOrigin: true,
        secure: false,
      },
      '/lang': {
        target: backendTarget,
        changeOrigin: true,
        secure: false,
      }
    },
  },
  plugins: [
    Components({
      dirs: "resources/vue/",
      extensions: ['vue'],
      deep: true,
      dts: false,
    }),
    vue(),
  ],
})
