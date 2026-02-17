// vue
import { createRouter, createWebHistory } from 'vue-router'

// local
import ParseView from '../views/ParseView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'parse',
      component: ParseView,
      meta: { title: '解析' },
    },
    {
      path: '/download',
      name: 'download',
      component: () => import('../views/DownloadView.vue'),
      meta: { title: '下载' },
    },
  ],
})

export default router
