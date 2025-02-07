import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/dashboard'
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login.vue')
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/Register.vue')
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('../views/Dashboard.vue'),
      meta: { requiresAuth: true },
      redirect: '/dashboard/domains',
      children: [
        {
          path: 'domains',
          name: 'domains',
          component: () => import('../views/Domains.vue')
        },
        {
          path: 'domains/:id/records',
          name: 'domain-records',
          component: () => import('../views/DomainRecords.vue')
        },
        {
          path: 'provider-keys',
          name: 'provider-keys',
          component: () => import('../views/ProviderKeys.vue')
        }
      ]
    }
  ]
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!token) {
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    } else {
      next()
    }
  } else {
    if (token && to.path === '/login') {
      next({ path: '/dashboard' })
    } else {
      next()
    }
  }
})

export default router