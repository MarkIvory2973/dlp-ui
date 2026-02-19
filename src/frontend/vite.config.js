// vite
import { defineConfig } from 'vite'
import { fileURLToPath, URL } from 'node:url'
// vue
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'
// varlet
import components from 'unplugin-vue-components/vite'
import autoImport from 'unplugin-auto-import/vite'
import { VarletImportResolver } from '@varlet/import-resolver'

// https://vite.dev/config/
export default defineConfig({
  base: '/ui',
  plugins: [
    vue(),
    vueJsx(),
    vueDevTools(),
    components({
      resolvers: [VarletImportResolver()],
    }),
    autoImport({
      resolvers: [VarletImportResolver({ autoImport: true })],
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
})
