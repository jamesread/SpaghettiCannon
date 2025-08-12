import { createRouter, createWebHistory } from 'vue-router'

import Home from './resources/vue/views/Home.vue'

const routes = [
  {
    name: "home",
    path: '/',
    component: Home,
  },
  {
    name: "weight",
    path: '/weight',
    component: () => import('./resources/vue/views/Weight.vue'),
  },
  {
    name: "heart",
    path: '/heart',
    component: () => import('./resources/vue/views/Heart.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router;
