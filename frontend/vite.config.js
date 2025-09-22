import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

const BACKEND_PORT = "8080";

export default defineConfig(({ mode }) => ({
  plugins: [svelte(), tailwindcss()],
  base: '/',
  preview: {
    allowedHosts: [
      '.ddns.net',
      'localhost',
    ],
    port: 5555,
    host: '0.0.0.0',
    proxy: {
      '/api': {
        target: `http://backend:${BACKEND_PORT}`,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  },
  server: {
    proxy: {
      '/api': {
        target: mode === 'development' ? `http://localhost:${BACKEND_PORT}` : `http://backend:${BACKEND_PORT}`,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    },
    allowedHosts: [
      '.ddns.net',
      'localhost'
    ],
    port: 5555,
    host: '0.0.0.0',
  },
  build: {
    assetsInlineLimit: 0
  }
}))

