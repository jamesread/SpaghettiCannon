import { createRouter, createWebHistory } from 'vue-router'
import { HomeIcon, WeightScaleIcon, HeartAddIcon } from '@hugeicons/core-free-icons'

import Home from './resources/vue/views/Home.vue'

const routes = [
  {
    name: "home",
    path: '/',
    component: Home,
    meta: {
      title: 'Home',
      icon: HomeIcon,
    }
  },
  {
    name: "weight",
    path: '/weight',
    component: () => import('./resources/vue/views/Weight.vue'),
    meta: {
      title: 'Weight',
      icon: WeightScaleIcon,
    }
  },
  {
    name: "heart",
    path: '/heart',
    component: () => import('./resources/vue/views/Heart.vue'),
    meta: {
      title: 'Heart',
      icon: HeartAddIcon,
    }
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router;
