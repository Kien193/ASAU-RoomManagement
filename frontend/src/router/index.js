import { createRouter, createWebHistory } from 'vue-router'

// import AppLayout from '@/layouts/AppLayout.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/:pathMatch(.*)*',
      redirect: '/admin/home',
    },
    {
      name: 'admin',
      path: '/admin',
      //component: AppLayout,
      redirect: { name: 'Home' },
      children: [
        {
          path: 'home',
          name: 'Home',
          component: () => import('@/views/HomeView.vue'),
        },
        {
          path: 'about',
          name: 'About',
          component: () => import('@/views/AboutView.vue'),
        },
      ]
    },
    {
      name: '404',
      path: '/404',
      component: () => import('@/views/404View.vue'),
    }
  ],
})

export default router