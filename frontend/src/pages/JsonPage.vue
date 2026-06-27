<template>
  <AppToolLayout
    title="JSON 格式化"
    description="格式化、压缩和校验 JSON 数据"
    :loading="store.loading"
    :error="store.error"
    :has-output="!!output"
    @copy="copyResult"
    @clear="handleClear"
    @error-close="store.clearError"
  >
    <template #input>
      <CodeEditor v-model="input" language="json" placeholder="请输入 JSON 字符串..." @submit="handleFormat" />
    </template>
    <template #actions>
      <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleFormat">
        格式化
      </el-button>
      <el-button :loading="store.loading" @click="handleCompress">压缩</el-button>
      <el-button @click="loadSample">示例数据</el-button>
    </template>
    <template #output>
      <CodeEditor v-model="output" language="json" placeholder="输出结果..." readonly />
    </template>
  </AppToolLayout>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import AppToolLayout from '../components/AppToolLayout.vue'
import CodeEditor from '../components/CodeEditor.vue'
import {useToolStore} from '../store/tools'
import {FormatJSON, CompressJSON} from '../../wailsjs/go/main/App'

const store = useToolStore()
const input = ref('')
const output = ref('')

async function handleFormat() {
  store.setLoading(true)
  store.clearError()
  try {
    const result = await FormatJSON(input.value, '  ')
    if (result.code !== 0) {
      store.setError(result.message)
    } else {
      output.value = result.data as string
    }
  } catch (e: any) {
    store.setError(e.message || 'JSON 格式化失败')
  } finally {
    store.setLoading(false)
  }
}

async function handleCompress() {
  store.setLoading(true)
  store.clearError()
  try {
    const result = await CompressJSON(input.value)
    if (result.code !== 0) {
      store.setError(result.message)
    } else {
      output.value = result.data as string
    }
  } catch (e: any) {
    store.setError(e.message || 'JSON 压缩失败')
  } finally {
    store.setLoading(false)
  }
}

function loadSample() {
  input.value = JSON.stringify({name: 'DevToys', version: 1, tools: ['json', 'base64', 'text']}, null, 2)
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
