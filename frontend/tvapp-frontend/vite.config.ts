import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import path from 'path';
import { fileURLToPath } from 'url';

// Dynamically resolve __dirname
const __dirname = path.dirname(fileURLToPath(import.meta.url));

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'), // Alias to simplify import paths
    },
  },
  server: {
    proxy: {
      '/query': {
        target: 'http://tvapp_backend:8080', // Backend server URL
        changeOrigin: true, // Required for virtual hosted sites
        secure: false, // If using self-signed SSL certificates
        rewrite: (path) => path.replace(/^\/query/, '/query'), // Optional: rewrite paths
      },
    },
  },
});
