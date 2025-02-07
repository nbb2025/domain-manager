<template>
  <div class="space-y-6">
    <!-- 页面标题 -->
    <div class="flex items-center justify-between">
      <h2 class="text-2xl font-bold text-gray-900">域名列表</h2>
      <a-button type="primary" @click="showAddDomainModal = true">
        <template #icon>
          <IconPlus />
        </template>
        添加域名
      </a-button>
    </div>

    <!-- 域名列表 -->
    <a-card :bordered="false" class="general-card">
      <a-table :loading="loading" :data="domains" :pagination="false">
        <template #columns>
          <a-table-column title="域名" data-index="name">
            <template #cell="{ record }">
              <div class="flex items-center">
                <IconPublic class="text-gray-400 mr-3" />
                <div>
                  <div class="text-primary">{{ record.name }}</div>
                  <div class="text-gray-500 text-sm mt-1">
                    密钥：{{ record.key_name }}
                  </div>
                  <div class="text-gray-500 text-sm">
                    服务商：{{ providerNames[record.provider] || record.provider }}
                  </div>
                </div>
              </div>
            </template>
          </a-table-column>
          <a-table-column title="操作" align="right">
            <template #cell="{ record }">
              <a-space>
                <a-button type="text" @click="$router.push({ name: 'domain-records', params: { id: record.id } })">
                  解析记录
                </a-button>
                <a-button type="text" status="danger" @click="handleDeleteDomain(record)">
                  删除
                </a-button>
              </a-space>
            </template>
          </a-table-column>
        </template>
        <template #empty>
          <div class="text-center py-8">
            <IconPublic class="text-gray-400 text-3xl mb-2" />
            <div class="text-gray-900 font-medium">暂无域名</div>
            <div class="text-gray-500 text-sm mt-1">开始添加您的第一个域名吧</div>
            <a-button type="primary" class="mt-4" @click="showAddDomainModal = true">
              <template #icon>
                <IconPlus />
              </template>
              添加域名
            </a-button>
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 添加域名对话框 -->
    <a-modal v-model:visible="showAddDomainModal" title="添加域名" @ok="handleAddDomain"
      :ok-button-props="{ disabled: !newDomain.name || !newDomain.provider }" ok-text="添加" cancel-text="取消">
      <a-form :model="newDomain" layout="vertical">
        <a-form-item field="key_id" label="密钥">
          <a-select v-model="newDomain.key_id" placeholder="请选择密钥" @change="handleKeyChange">
            <a-option v-for="key in providerKeys" :key="key.id" :value="key.id">
              {{ key.name }} ({{ providerNames[key.provider] || key.provider }})
            </a-option>
          </a-select>
        </a-form-item>
        <a-form-item v-if="newDomain.key_id" field="name" label="域名">
          <a-select v-model="newDomain.name" placeholder="请选择域名">
            <a-option v-for="domain in availableDomains" :key="domain.name" :value="domain.name"
              :disabled="domain.added">
              {{ domain.name }} {{ domain.added ? '(已添加)' : '' }}
            </a-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { get, post, del } from '@/utils/api'
import { IconPublic, IconPlus } from '@arco-design/web-vue/es/icon'

const domains = ref([])
const loading = ref(true)
const showAddDomainModal = ref(false)
const providerKeys = ref([])
const availableDomains = ref([])
const newDomain = ref({
  name: '',
  key_id: '',
  provider: ''
})

const providerNames = {
  aliyun: '阿里云',
  tencent: '腾讯云',
  cloudflare: 'Cloudflare'
}

// 获取域名列表
const fetchDomains = async () => {
  loading.value = true
  try {
    const response = await get('/api/v1/domains')
    domains.value = Array.isArray(response) ? response : []
  } catch (e) {
    // TODO: 处理错误
  } finally {
    loading.value = false
  }
}

// 获取服务商密钥列表
const fetchProviderKeys = async () => {
  try {
    providerKeys.value = await get('/api/v1/provider-keys')
  } catch (e) {
    // TODO: 处理错误
  }
}

// 获取可用域名列表
const handleKeyChange = async () => {
  if (!newDomain.value.key_id) return

  const selectedKey = providerKeys.value.find(key => key.id === newDomain.value.key_id)
  if (!selectedKey) return

  newDomain.value.provider = selectedKey.provider

  try {
    const availableDomainList = await get(`/api/v1/domains/available?key_id=${newDomain.value.key_id}`)
    // 标记已添加的域名
    debugger
    availableDomains.value = availableDomainList.map(domain => ({
      name: domain,
      added: domains.value.some(d => d.name == domain)
    }))
  } catch (e) {
    // TODO: 处理错误
  }
}

// 添加域名
const handleAddDomain = async () => {
  try {
    await post('/api/v1/domains', {
      name: newDomain.value.name,
      provider: newDomain.value.provider
    })
    showAddDomainModal.value = false
    newDomain.value = {
      name: '',
      key_id: '',
      provider: ''
    }
    await fetchDomains()
  } catch (e) {
    // TODO: 处理错误
  }
}

// 删除域名
const handleDeleteDomain = async (domain) => {
  try {
    await del(`/api/v1/domains/${domain.id}`)
    await fetchDomains()
  } catch (e) {
    // TODO: 处理错误
  }
}

onMounted(async () => {
  await Promise.all([fetchDomains(), fetchProviderKeys()])
})
</script>

<style scoped>
.general-card {
  border-radius: 4px;
}

.text-primary {
  color: rgb(var(--primary-6));
}
</style>
