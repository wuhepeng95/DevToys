<template>
  <AppToolLayout
    title="SQL 格式化"
    description="SQL 格式化与压缩"
    :loading="store.loading"
    :error="store.error"
    :has-output="!!output"
    @copy="copyResult"
    @clear="handleClear"
    @error-close="store.clearError"
  >
    <template #input>
      <CodeEditor v-model="input" language="sql" placeholder="请输入 SQL 语句..." @submit="handleFormat" />
    </template>
    <template #actions>
      <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleFormat">
        格式化
      </el-button>
      <el-button :loading="store.loading" @click="handleMinify">压缩</el-button>
      <el-button @click="loadSample">示例数据</el-button>
    </template>
    <template #output>
      <CodeEditor v-model="output" language="sql" placeholder="输出结果..." readonly />
    </template>
  </AppToolLayout>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import AppToolLayout from '../components/AppToolLayout.vue'
import CodeEditor from '../components/CodeEditor.vue'
import {useToolStore} from '../store/tools'
import {FormatSQL, MinifySQL} from '../../wailsjs/go/main/App'

const store = useToolStore()
const input = ref('')
const output = ref('')

async function handleFormat() {
  store.setLoading(true)
  store.clearError()
  try {
    const result = await FormatSQL(input.value)
    if (result.code !== 0) {
      store.setError(result.message)
    } else {
      output.value = result.data as string
    }
  } catch (e: any) {
    store.setError(e.message || '格式化失败')
  } finally {
    store.setLoading(false)
  }
}

async function handleMinify() {
  store.setLoading(true)
  store.clearError()
  try {
    const result = await MinifySQL(input.value)
    if (result.code !== 0) {
      store.setError(result.message)
    } else {
      output.value = result.data as string
    }
  } catch (e: any) {
    store.setError(e.message || '压缩失败')
  } finally {
    store.setLoading(false)
  }
}

function loadSample() {
  input.value = [
    'SELECT u.id, u.name, u.email, o.order_date, o.total',
    'FROM users u',
    'INNER JOIN orders o ON u.id = o.user_id',
    'WHERE u.active = 1',
    'AND o.total > 100',
    'ORDER BY o.order_date DESC',
    'LIMIT 20',
  ].join('\n')
}

function handleClear() {
  input.value = ''
  output.value = ''
  store.reset()
}

async function copyResult() {
  if (output.value) {
    await navigator.clipboard.writeText(output.value)
  }
}
</script>
