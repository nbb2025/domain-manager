<template>
  <div class="register-container flex items-center justify-center min-h-screen bg-gray-50">
    <div class="register-card relative w-full max-w-md p-8 rounded-lg shadow-lg">
      <h1 class="register-title-text text-2xl font-bold text-center mb-8">注册新账号</h1>

      <a-form :model="formData" @submit="handleRegister">
        <a-form-item field="username" label="用户名">
          <a-input v-model="formData.username" placeholder="请输入用户名" allow-clear />
        </a-form-item>

        <a-form-item field="password" label="密码">
          <a-input-password v-model="formData.password" placeholder="请输入密码" allow-clear />
        </a-form-item>

        <a-button type="primary" html-type="submit" class="register-button w-full mb-4" :loading="loading">
          {{ loading ? '注册中...' : '注册' }}
        </a-button>
      </a-form>

      <div class="text-center">
        <a class="login-link text-primary" @click="router.push('/login')">返回登录</a>
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
const formData = reactive({
  username: '',
  password: ''
})

const handleRegister = async () => {
  // 表单验证
  if (!formData.username || !formData.password) {
    Message.warning('请填写完整的注册信息')
    return
  }

  // 密码强度验证
  if (formData.password.length < 6) {
    Message.warning('密码长度至少6位')
    return
  }

  try {
    loading.value = true

    // 注册请求
    await api.post('/api/v1/users/register', {
      username: formData.username.trim(),
      password: formData.password
    })

    // 注册成功处理
    Message.success({
      content: '注册成功，即将跳转到登录页面',
      duration: 2000
    })

    // 延迟跳转，让用户看到成功提示
    setTimeout(() => {
      router.push('/login')
    }, 1500)

  } catch (error) {
    // 错误处理优化
    const errorMsg = error.response?.data?.message || error.message || '注册失败，请稍后重试'
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
.register-container {
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
.register-card {
  border: 1px solid rgba(var(--primary-6), 0.12);
  background: linear-gradient(145deg, rgba(255, 255, 255, 0.96) 0%, rgba(255, 255, 255, 0.98) 100%);
  overflow: visible;
  z-index: 1;
}

.register-card::before {
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
.register-title-text {
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
.register-button {
  position: relative;
  overflow: hidden;
  border: none;
  background: linear-gradient(135deg, rgb(var(--primary-6)) 0%, rgb(var(--primary-5)) 100%);
  transition: all 0.3s ease;
  z-index: 2;
}

.register-button::after {
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

/* 优化登录链接动效 */
.login-link {
  position: relative;
  display: inline-block;
  transition: color 0.3s ease;
  z-index: 2;
}

.login-link::after {
  content: '';
  position: absolute;
  bottom: -2px;
  left: 0;
  width: 0;
  height: 1px;
  background: rgb(var(--primary-6));
  transition: width 0.3s ease;
}

.login-link:hover::after {
  width: 100%;
}

/* 优化移动端体验 */
@media (max-width: 640px) {
  .register-card {
    border-radius: 12px;
  }

  :deep(.a-input) {
    padding: 10px 14px;
  }

  .register-button {
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