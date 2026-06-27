<template>
  <AppToolLayout
    title="进制转换"
    description="二进制 / 八进制 / 十进制 / 十六进制互转"
    :loading="store.loading"
    :error="store.error"
    :has-output="!!result.binary"
    @error-close="store.clearError"
  >
    <template #options>
      <div class="tool-options">
        <el-radio-group v-model="fromBase">
          <el-radio :value="2">二进制</el-radio>
          <el-radio :value="8">八进制</el-radio>
          <el-radio :value="10">十进制</el-radio>
          <el-radio :value="16">十六进制</el-radio>
        </el-radio-group>
      </div>
    </template>
    <template #input>
      <AppTextarea v-model="input" :rows="2" placeholder="请输入数字..." @submit="handleConvert" />
    </template>
    <template #actions>
      <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleConvert">
        转换
      </el-button>
      <el-button @click="loadSample">示例数据</el-button>
    </template>
    <template #output>
      <el-table v-if="result.binary" :data="rows" border stripe size="small">
        <el-table-column prop="label" label="进制" width="120" />
        <el-table-column prop="value" label="值">
          <template #default="{row}">
            <span class="mono">{{ row.value }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="100">
          <template #default="{row}">
            <el-button size="small" @click="copySingle(row.value)">
              复制
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </template>
    <template #bottom-actions>
      <el-button @click="handleClear">
        <el-icon><Delete /></el-icon>
        清空
      </el-button>
    </template>
  </AppToolLayout>
</template>

<script setup lang="ts">
import {ref, computed} from 'vue'
import AppToolLayout from '../components/AppToolLayout.vue'
import AppTextarea from '../components/AppTextarea.vue'
import {useToolStore} from '../store/tools'
import {ConvertBase} from '../../wailsjs/go/main/App'
import {Delete} from '@element-plus/icons-vue'

interface ConvertResult {
  binary: string
  octal: string
  decimal: string
  hexadecimal: string
}

const store = useToolStore()
const input = ref('')
const fromBase = ref(10)
const result = ref<ConvertResult>({binary: '', octal: '', decimal: '', hexadecimal: ''})

const rows = computed(() => [
  {label: '二进制', value: result.value.binary},
  {label: '八进制', value: result.value.octal},
  {label: '十进制', value: result.value.decimal},
  {label: '十六进制', value: result.value.hexadecimal},
])

async function handleConvert() {
  store.setLoading(true)
  store.clearError()
  try {
    const res = await ConvertBase(input.value, fromBase.value)
    if (res.code !== 0) {
      store.setError(res.message)
    } else {
      result.value = res.data as ConvertResult
    }
  } catch (e: any) {
    store.setError(e.message || '转换失败')
  } finally {
    store.setLoading(false)
  }
}

function loadSample() {
  input.value = '42'
}

function handleClear() {
  input.value = ''
  result.value = {binary: '', octal: '', decimal: '', hexadecimal: ''}
  store.reset()
}

async function copySingle(val: string) {
  if (val) await navigator.clipboard.writeText(val)
}
</script>
