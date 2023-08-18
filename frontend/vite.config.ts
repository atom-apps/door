import vue from '@vitejs/plugin-vue'
import { fileURLToPath } from 'url'
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      "^/v1/.*":  {
        target: 'http://127.0.0.1:9800',
        changeOrigin: true,
      }
    },
  },
  plugins: [
    vue({
      reactivityTransform: true,
    }),
  ],
  resolve: {
    alias: [
      { find: '@', replacement: fileURLToPath(new URL('./src', import.meta.url)) },
      { find: '@views', replacement: fileURLToPath(new URL('./src/views', import.meta.url)) },
      { find: '@components', replacement: fileURLToPath(new URL('./src/components', import.meta.url)) },
      { find: '@sdk', replacement: fileURLToPath(new URL('./src/sdk', import.meta.url)) },
    ],
  },
})