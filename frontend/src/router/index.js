import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '@/views/Home.vue'
import Farming from "../views/Farming.vue"
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: "/module/:id",
    name: "ModuleInfo",
    component: () => import('@/views/Plant')
  },
  {
    path: '/admin',
    name: 'AdminSettings',
    component: ()  => import('@/views/AdminSettings')

  },
  {
    path: '/pin',
    name: 'AdminMenu',
    component: ()  => import('@/views/PinAndAdminMenu')

  },
  {
    path: '/adminmodule',
    name: 'TypeChanger',
    component: ()  => import('@/views/ModuleTypeSettings')

  },

  {
    path: '/batchfarming',
    name: 'BatchFarming',
    component: () => import('@/views/BatchFarming')
  },
  
  {
    path: '/plant',
    name: 'Plant',
    component: () => import('@/views/Plant')
  },
  {
    path: '/harvest',
    name: 'Harvest',
    component: () => import('@/views/Harvest')
  },
  {
    path: '/farming',
    name: 'Farming',
    component: Farming

  },

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
