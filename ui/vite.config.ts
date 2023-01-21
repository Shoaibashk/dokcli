import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      "/spec-url": "http://localhost:1212"
    }
  },
  plugins: [react()],
})
