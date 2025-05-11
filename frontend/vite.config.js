import { defineConfig } from 'vite'
import pugPlugin from '@vituum/vite-plugin-pug'

import vue from '@vitejs/plugin-vue'

const options = { pretty: false }
const locals = { name: "SpaghettiCannon" }

export default defineConfig({
  plugins: [
    pugPlugin(options, locals),
    vue(),
  ],
  build: {
    rollupOptions: {
      input: ["index.pug"],
    }
  }
})
