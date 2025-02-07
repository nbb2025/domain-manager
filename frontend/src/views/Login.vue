<template>
  <div class="login-container flex items-center justify-center min-h-screen bg-gray-50">
    <div class="login-card relative w-full max-w-md p-8 rounded-lg shadow-lg">
      <h1 class="login-title-text text-2xl font-bold text-center mb-8">域名管理系统</h1>

      <a-form :model="form" @submit="handleSubmit">
        <a-form-item field="username" label="用户名">
          <a-input v-model="form.username" placeholder="请输入用户名" allow-clear />
        </a-form-item>

        <a-form-item field="password" label="密码">
          <a-input-password v-model="form.password" placeholder="请输入密码" allow-clear />
        </a-form-item>

        <a-button type="primary" html-type="submit" class="login-button w-full mb-4" :loading="loading">
          {{ loading ? '登录中...' : '登录' }}
        </a-button>
      </a-form>

      <div class="text-center">
        <a class="register-link text-primary" @click="router.push('/register')">注册新账号</a>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { Message } from '@arco-design/web-vue'
import { useRouter } from 'vue-router'
import api from '@/utils/api'

const router = useRouter()
const loading = ref(false)
const form = reactive({
  username: '',
  password: ''
})

const handleSubmit = async () => {
  // 表单验证
  if (!form.username || !form.password) {
    Message.warning('请填写完整的登录信息')
    return
  }

  try {
    loading.value = true
    const data = await api.post('/api/v1/users/login', {
      username: form.username.trim(),
      password: form.password
    })

    // 登录成功处理
    Message.success({
      content: '登录成功，正在跳转...',
      duration: 2000
    })

    localStorage.setItem('token', data.token)

    // 延迟跳转，让用户看到成功提示
    setTimeout(() => {
      router.push('/dashboard')
    }, 1500)
  } catch (error) {
    // 错误处理优化
    const errorMsg = error.response?.data?.message || error.message || '登录失败，请稍后重试'
    Message.error({
      content: errorMsg,
      duration: 3000
    })
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
/* 新增整体入场动画 */
.login-container {
  animation: fadeIn 0.6s ease-out;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.9), rgba(240, 240, 255, 0.95));
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 优化卡片设计 */
.login-card {
  border: 1px solid rgba(var(--primary-6), 0.12);
  background: linear-gradient(145deg, rgba(255, 255, 255, 0.96) 0%, rgba(255, 255, 255, 0.98) 100%);
  overflow: visible;
  z-index: 1;
}

.login-card::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: conic-gradient(from 180deg at 50% 50%,
      rgba(var(--primary-6), 0.1) 0deg,
      rgba(var(--primary-6), 0.03) 180deg,
      rgba(var(--primary-6), 0.1) 360deg);
  animation: rotate 12s linear infinite;
  z-index: -1;
}

@keyframes rotate {
  to {
    transform: rotate(360deg);
  }
}

/* 优化标题效果 */
.login-title-text {
  position: relative;
  z-index: 2;
  padding: 8px 0;
  letter-spacing: -0.5px;
  text-shadow: 0 2px 4px rgba(var(--primary-6), 0.1);
}

/* 优化输入框交互 */
:deep(.a-input) {
  transition: all 0.3s ease;
  border-radius: 8px;
  background: rgba(var(--gray-1), 0.8);
  z-index: 3;
}

:deep(.a-input:focus-within) {
  background: var(--color-bg-1);
  box-shadow: 0 0 0 2px rgba(var(--primary-6), 0.2);
}

/* 增强按钮质感 */
.login-button {
  position: relative;
  overflow: hidden;
  border: none;
  background: linear-gradient(135deg, rgb(var(--primary-6)) 0%, rgb(var(--primary-5)) 100%);
  transition: all 0.3s ease;
  z-index: 2;
}

.login-button::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: linear-gradient(45deg,
      transparent 25%,
      rgba(255, 255, 255, 0.1) 50%,
      transparent 75%);
  animation: shine 3s infinite;
}

@keyframes shine {
  to {
    transform: translateX(100%);
  }
}

/* 优化注册链接动效 */
.register-link {
  position: relative;
  display: inline-block;
  transition: color 0.3s ease;
  z-index: 2;
}

.register-link::after {
  content: '';
  position: absolute;
  bottom: -2px;
  left: 0;
  width: 0;
  height: 1px;
  background: rgb(var(--primary-6));
  transition: width 0.3s ease;
}

.register-link:hover::after {
  width: 100%;
}

/* 优化移动端体验 */
@media (max-width: 640px) {
  .login-card {
    border-radius: 12px;
  }

  :deep(.a-input) {
    padding: 10px 14px;
  }

  .login-button {
    height: 48px;
    font-size: 15px;
  }
}

/* 新增输入框标签动画 */
:deep(.a-form-item-label) {
  display: block;
  transform-origin: left center;
  transition: all 0.3s ease;
  z-index: 2;
}

:deep(.a-input:focus-within + .a-form-item-label) {
  transform: translateY(-2px) scale(0.95);
  color: rgb(var(--primary-6));
}
</style>