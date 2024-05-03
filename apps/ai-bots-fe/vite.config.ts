import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api': 'http://localhost:8080'
    }
  },
  resolve: {
    alias: {
      'api': '/src/api',
      'types': '/src/types',
      'hooks': '/src/hooks',
      'pages': '/src/pages',
    }
  }
})
