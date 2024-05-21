import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/404',
      component: () => import('@/views/404View.vue'),
      hidden: true
    },
    {
      path: '/',
      name: 'Home',
      component: HomeView
    },
    {
      path: '/about',
      component: () => import('@/views/AboutView.vue')
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: '/404',
      hidden: true
    }
  ],
})

export default router