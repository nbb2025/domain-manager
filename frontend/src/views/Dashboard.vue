<template>
  <a-layout class="h-screen">
    <a-layout-header
      class="fixed w-full z-10 px-6 flex justify-between items-center bg-white border-b border-gray-100 shadow-sm">
      <div class="flex items-center">
        <div class="text-2xl font-bold text-gray-900">域名管理系统</div>
      </div>
      <div class="flex items-center">
        <a-button type="text" size="large" @click="handleLogout">
          <template #icon>
            <IconExport />
          </template>
          退出登录
        </a-button>
      </div>
    </a-layout-header>

    <a-layout style="margin-top: 72px">
      <a-layout-sider :width="256" class="h-[calc(100vh-72px)] bg-white border-r border-gray-100">
        <a-menu :selected-keys="[$route.name]" :style="{ height: '100%', borderRight: 0 }"
          @menuItemClick="handleMenuClick">
          <a-menu-item key="domains">
            <template #icon>
              <IconPublic />
            </template>
            域名管理
          </a-menu-item>
          <a-menu-item key="provider-keys">
            <template #icon>
              <IconLock />
            </template>
            服务商密钥
          </a-menu-item>
        </a-menu>
      </a-layout-sider>

      <a-layout-content class="p-6 bg-[var(--color-fill-2)]">
        <router-view></router-view>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { IconPublic, IconLock, IconExport } from '@arco-design/web-vue/es/icon'

const router = useRouter()

const handleLogout = () => {
  localStorage.removeItem('token')
  router.push('/login')
}

const handleMenuClick = (key) => {
  router.push(`/dashboard/${key}`)
}
</script>

<style scoped>
.arco-layout-header {
  height: 72px;
  line-height: 72px;
}

.arco-menu-item .router-link-active {
  color: rgb(var(--primary-6));
}
</style>