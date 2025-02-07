<template>
  <div class="space-y-6">
    <!-- 页面标题 -->
    <div class="flex items-center justify-between">
      <h2 class="text-2xl font-bold text-gray-900">服务商密钥</h2>
      <a-button type="primary" @click="showAddKeyModal = true">
        <template #icon>
          <IconPlus />
        </template>
        添加密钥
      </a-button>
    </div>

    <!-- 密钥列表 -->
    <a-card :bordered="false" class="general-card">
      <a-table :loading="loading" :data="keys" :pagination="false">
        <template #columns>
          <a-table-column title="服务商" data-index="provider">
            <template #cell="{ record }">
              <div class="flex items-center">
                <IconTool class="text-gray-400 mr-3" />
                <div>
                  <div class="text-primary text-lg font-bold">{{ record.name }}</div>
                  <div class="text-gray-500 text-sm mt-1">
                    服务商：{{ providerNames[record.provider] || record.provider }}
                  </div>
                  <div class="text-gray-500 text-sm">
                    {{ getKeyFieldName(record.provider, 'access_key') }}: {{ maskKey(record.access_key) }}
                  </div>
                  <div v-if="showSecretKeyForProvider(record.provider)" class="text-gray-500 text-sm">
                    {{ getKeyFieldName(record.provider, 'secret_key') }}: {{ maskKey(record.secret_key) }}
                  </div>
                  <div class="text-gray-500 text-sm">
                    创建时间: {{ new Date(record.created_at).toLocaleString() }}
                  </div>
                </div>
              </div>
            </template>
          </a-table-column>
          <a-table-column title="操作" align="right">
            <template #cell="{ record }">
              <a-space>
                <a-button type="text" @click="handleEditKeyName(record)">
                  编辑名称
                </a-button>
                <a-button type="text" status="danger" @click="handleDeleteKey(record)">
                  删除
                </a-button>
              </a-space>
            </template>
          </a-table-column>
        </template>
        <template #empty>
          <div class="text-center py-8">
            <IconTool class="text-gray-400 text-3xl mb-2" />
            <div class="text-gray-900 font-medium">暂无密钥</div>
            <div class="text-gray-500 text-sm mt-1">添加服务商密钥以管理域名</div>
            <a-button type="primary" class="mt-4" @click="showAddKeyModal = true">
              <template #icon>
                <IconPlus />
              </template>
              添加密钥
            </a-button>
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 添加密钥对话框 -->
    <a-modal v-model:visible="showAddKeyModal" title="添加服务商密钥" @ok="handleAddKey" @cancel="resetForm"
      :ok-button-props="{ disabled: !newKey.name || !newKey.provider || !newKey.access_key || !newKey.secret_key }"
      ok-text="添加" cancel-text="取消">
      <a-form :model="newKey" layout="vertical">
        <a-form-item field="name" label="密钥名称" required>
          <a-input v-model="newKey.name" placeholder="请输入密钥名称" allow-clear />
        </a-form-item>
        <a-form-item field="provider" label="服务商">
          <a-select v-model="newKey.provider" placeholder="请选择服务商">
            <a-option value="aliyun">阿里云</a-option>
            <a-option value="tencent">腾讯云</a-option>
            <a-option value="cloudflare">Cloudflare</a-option>
          </a-select>
        </a-form-item>
        <template v-if="newKey.provider">
          <a-form-item field="access_key" :label="getKeyFieldName(newKey.provider, 'access_key')">
            <a-input v-model="newKey.access_key" :placeholder="`请输入 ${getKeyFieldName(newKey.provider, 'access_key')}`"
              allow-clear />
          </a-form-item>
          <a-form-item v-if="showSecretKey" field="secret_key" :label="getKeyFieldName(newKey.provider, 'secret_key')">
            <a-input-password v-model="newKey.secret_key"
              :placeholder="`请输入 ${getKeyFieldName(newKey.provider, 'secret_key')}`" allow-clear />
          </a-form-item>
        </template>
      </a-form>
    </a-modal>

    <!-- 编辑密钥名称对话框 -->
    <a-modal v-model:visible="showEditNameModal" title="编辑密钥名称" @ok="handleUpdateKeyName" @cancel="resetEditForm"
      :ok-button-props="{ disabled: !editKey.newName }" ok-text="保存" cancel-text="取消">
      <a-form :model="editKey" layout="vertical">
        <a-form-item field="newName" label="密钥名称" required>
          <a-input v-model="editKey.newName" placeholder="请输入新的密钥名称" allow-clear />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { IconPlus, IconTool } from '@arco-design/web-vue/es/icon'
import { get, post, del, put } from '@/utils/api'
import { Message } from '@arco-design/web-vue'

const keys = ref([])
const loading = ref(true)
const showAddKeyModal = ref(false)
const showEditNameModal = ref(false)
const newKey = ref({
  name: '',
  provider: '',
  access_key: '',
  secret_key: ''
})
const editKey = ref({
  provider: '',
  newName: ''
})

const providerNames = {
  aliyun: '阿里云',
  tencent: '腾讯云',
  cloudflare: 'Cloudflare'
}

const providerKeyNames = {
  aliyun: {
    access_key: 'AccessKey ID',
    secret_key: 'AccessKey Secret'
  },
  tencent: {
    access_key: 'SecretId',
    secret_key: 'SecretKey'
  },
  cloudflare: {
    access_key: 'API Token',
    secret_key: ''
  }
}

// 获取密钥列表
const fetchKeys = async () => {
  loading.value = true
  try {
    keys.value = await get('/api/v1/provider-keys')
  } catch (e) {
    Message.error('获取密钥列表失败')
  } finally {
    loading.value = false
  }
}

// 添加密钥
const resetForm = () => {
  newKey.value = {
    name: '',
    provider: '',
    access_key: '',
    secret_key: ''
  }
}

const handleAddKey = async () => {
  try {
    await post('/api/v1/provider-keys', newKey.value)
    Message.success('添加密钥成功')
    showAddKeyModal.value = false
    resetForm()
    await fetchKeys()
  } catch (e) {
    Message.error('添加密钥失败：' + (e.message || '未知错误'))
  }
}

// 删除密钥
const handleDeleteKey = async (key) => {
  try {
    await del(`/api/v1/provider-keys/${key.provider}`)
    Message.success('删除密钥成功')
    await fetchKeys()
  } catch (e) {
    Message.error('删除密钥失败：' + (e.message || '未知错误'))
  }
}

// 编辑密钥名称
const handleEditKeyName = (key) => {
  editKey.value = {
    id: key.id,
    newName: key.name
  }
  showEditNameModal.value = true
}

const resetEditForm = () => {
  editKey.value = {
    provider: '',
    newName: ''
  }
}

const handleUpdateKeyName = async () => {
  try {
    await put(`/api/v1/provider-keys/${editKey.value.id}/name`, {
      new_name: editKey.value.newName
    })
    Message.success('更新密钥名称成功')
    showEditNameModal.value = false
    resetEditForm()
    await fetchKeys()
  } catch (e) {
    Message.error('更新密钥名称失败：' + (e.message || '未知错误'))
  }
}

// 掩码显示密钥
const maskKey = (key) => {
  if (!key) return ''
  return key.slice(0, 4) + '*'.repeat(key.length - 8) + key.slice(-4)
}

// 获取当前服务商的字段名称
const getKeyFieldName = (provider, field) => {
  if (!provider) return field === 'access_key' ? 'Access Key' : 'Secret Key'
  return providerKeyNames[provider]?.[field] || ''
}

// 是否显示 Secret Key 字段
const showSecretKey = computed(() => {
  return !newKey.value.provider || providerKeyNames[newKey.value.provider]?.secret_key !== ''
})

// 判断指定服务商是否需要显示 Secret Key
const showSecretKeyForProvider = (provider) => {
  return providerKeyNames[provider]?.secret_key !== ''
}

onMounted(fetchKeys)
</script>

<style scoped>
.general-card {
  border-radius: 4px;
}

.text-primary {
  color: rgb(var(--primary-6));
}
</style>
