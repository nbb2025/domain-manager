<template>
  <div class="space-y-6">
    <!-- 页面标题 -->
    <div class="flex items-center justify-between">
      <div class="flex items-center space-x-4">
        <router-link :to="{ name: 'domains' }">
          <a-button>
            <template #icon>
              <icon-left />
            </template>
            返回
          </a-button>
        </router-link>
        <div>
          <h2 class="text-2xl font-bold text-gray-900">解析记录</h2>
          <p class="mt-1 text-sm text-gray-500">{{ domain?.name }}</p>
        </div>
      </div>
      <a-button type="primary" @click="showAddRecordModal = true">
        <template #icon>
          <icon-plus />
        </template>
        添加记录
      </a-button>
    </div>

    <!-- 记录列表 -->
    <a-card :bordered="false" class="general-card">
      <a-table :loading="loading" :data="records" :pagination="false">
        <template #columns>
          <a-table-column title="主机记录" data-index="Name">
            <template #cell="{ record }">
              <span class="text-primary">{{ record.Name }}</span>
            </template>
          </a-table-column>
          <a-table-column title="记录类型" data-index="Type" />
          <a-table-column title="线路类型" data-index="Line">
            <template #cell="{ record }">
              {{ record.Line || '默认' }}
            </template>
          </a-table-column>
          <a-table-column title="记录值" data-index="Value" />
          <a-table-column title="TTL" data-index="TTL">
            <template #cell="{ record }">
              {{ record.TTL }}秒
            </template>
          </a-table-column>
          <a-table-column title="备注" data-index="Remark">
            <template #cell="{ record }">
              {{ record.Remark || '-' }}
            </template>
          </a-table-column>
          <a-table-column title="状态">
            <template #cell="{ record }">
              <a-space>
                <a-tag :color="record.Enabled ? 'green' : 'gray'">
                  {{ record.Enabled ? '启用' : '禁用' }}
                </a-tag>
                <a-tag v-if="domain?.provider === 'cloudflare'" :color="record.Proxied ? 'orange' : 'gray'"
                  class="cursor-pointer" @click="toggleProxy(record)">
                  <template #icon>
                    <icon-shield />
                  </template>
                  {{ record.Proxied ? '已代理' : '未代理' }}
                </a-tag>
              </a-space>
            </template>
          </a-table-column>
          <a-table-column title="操作" align="right">
            <template #cell="{ record }">
              <a-space>
                <a-button @click="editRecord(record)">编辑</a-button>
                <a-button status="danger" @click="deleteRecord(record)">删除</a-button>
              </a-space>
            </template>
          </a-table-column>
        </template>
        <template #empty>
          <div class="text-center py-8">
            <icon-file class="text-gray-400 text-3xl mb-2" />
            <div class="text-gray-900 font-medium">暂无记录</div>
            <div class="text-gray-500 text-sm mt-1">开始添加您的第一条记录吧</div>
            <a-button type="primary" class="mt-4" @click="showAddRecordModal = true">
              <template #icon>
                <icon-plus />
              </template>
              添加记录
            </a-button>
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 添加/编辑记录对话框 -->
    <a-modal v-model:visible="showAddRecordModal" :title="currentRecord ? '编辑记录' : '添加记录'" @ok="handleSaveRecord"
      :ok-button-props="{ disabled: !isFormValid }" ok-text="保存" cancel-text="取消">
      <a-form :model="newRecord" layout="vertical">
        <a-form-item field="Type" label="记录类型">
          <a-select v-model="newRecord.Type">
            <a-option value="A">A</a-option>
            <a-option value="AAAA">AAAA</a-option>
            <a-option value="CNAME">CNAME</a-option>
            <a-option value="MX">MX</a-option>
            <a-option value="TXT">TXT</a-option>
            <a-option value="NS">NS</a-option>
            <a-option value="SRV">SRV</a-option>
            <a-option value="CAA">CAA</a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="Name" label="主机记录">
          <a-input v-model="newRecord.Name" placeholder="@" allow-clear />
        </a-form-item>
        <a-form-item field="Value" label="记录值">
          <a-input v-model="newRecord.Value" :placeholder="getRecordValuePlaceholder(newRecord.Type)" allow-clear />
        </a-form-item>
        <a-form-item field="TTL" label="TTL (秒)">
          <a-select v-model="newRecord.TTL">
            <a-option :value="600">10分钟</a-option>
            <a-option :value="1800">30分钟</a-option>
            <a-option :value="3600">1小时</a-option>
            <a-option :value="43200">12小时</a-option>
            <a-option :value="86400">1天</a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="Remark" label="备注">
          <a-input v-model="newRecord.Remark" placeholder="可选" allow-clear />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-checkbox v-model="newRecord.Enabled">启用</a-checkbox>
          </a-space>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { get, post, put, del } from '@/utils/api'
import { Message, Modal } from '@arco-design/web-vue'
import { IconLeft, IconPlus, IconFile } from '@arco-design/web-vue/es/icon'

const route = useRoute()
const router = useRouter()

const domain = ref(null)
const records = ref([])
const loading = ref(true)
const showAddRecordModal = ref(false)
const currentRecord = ref(null)

const newRecord = ref({
  Type: 'A',
  Name: '',
  Value: '',
  TTL: 600,
  Remark: '',
  Enabled: true
})

const isFormValid = computed(() => {
  return newRecord.value.Type && newRecord.value.Name && newRecord.value.Value
})

// 获取域名信息
const fetchDomain = async () => {
  try {
    domain.value = await get(`/api/v1/domains/${route.params.id}`)
  } catch (e) {
    Message.error('获取域名信息失败')
  }
}

// 获取记录列表
const fetchRecords = async () => {
  loading.value = true
  try {
    records.value = await get(`/api/v1/domains/${route.params.id}/records`)
  } catch (e) {
    Message.error('获取记录列表失败')
  } finally {
    loading.value = false
  }
}

// 获取记录值的占位符
const getRecordValuePlaceholder = (type) => {
  const placeholders = {
    A: '例如：192.168.1.1',
    AAAA: '例如：2001:db8::1',
    CNAME: '例如：example.com',
    MX: '例如：mail.example.com',
    TXT: '例如：v=spf1 include:spf.example.com ~all',
    NS: '例如：ns1.example.com',
    SRV: '例如：0 5 5060 sip.example.com',
    CAA: '例如：0 issue "letsencrypt.org"'
  }
  return placeholders[type] || ''
}

// 编辑记录
const editRecord = (record) => {
  currentRecord.value = record
  newRecord.value = { ...record }
  showAddRecordModal.value = true
}

// 保存记录
const handleSaveRecord = async () => {
  try {
    if (currentRecord.value) {
      await put(`/api/v1/domains/${route.params.id}/records/${currentRecord.value.Id}`, newRecord.value)
      Message.success('更新记录成功')
    } else {
      await post(`/api/v1/domains/${route.params.id}/records`, newRecord.value)
      Message.success('添加记录成功')
    }
    showAddRecordModal.value = false
    currentRecord.value = null
    newRecord.value = {
      Type: 'A',
      Name: '',
      Value: '',
      TTL: 600,
      Remark: '',
      Enabled: true
    }
    await fetchRecords()
  } catch (e) {
    Message.error(currentRecord.value ? '更新记录失败' : '添加记录失败')
  }
}

// 删除记录
const deleteRecord = (record) => {
  Modal.warning({
    title: '确认删除',
    content: '确定要删除这条记录吗？',
    okText: '确定',
    cancelText: '取消',
    async onOk() {
      try {
        await del(`/api/v1/domains/${route.params.id}/records/${record.Id}`)
        Message.success('删除记录成功')
        await fetchRecords()
      } catch (e) {
        Message.error('删除记录失败')
      }
    }
  })
}

// 切换代理状态（仅Cloudflare）
const toggleProxy = async (record) => {
  try {
    await put(`/api/v1/domains/${route.params.id}/records/${record.Id}`, {
      ...record,
      Proxied: !record.Proxied
    })
    await fetchRecords()
  } catch (e) {
    Message.error('更新代理状态失败')
  }
}

onMounted(async () => {
  await fetchDomain()
  await fetchRecords()
})
</script>

<style scoped>
.general-card {
  border-radius: 8px;
}
</style>