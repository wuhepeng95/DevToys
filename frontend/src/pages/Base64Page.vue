<template>
  <AppToolLayout
    title="Base64 编解码"
    description="Base64 编码与解码，支持 UTF-8 和 URL Base64"
    :loading="store.loading"
    :error="store.error"
    :has-output="!!output"
    @copy="copyResult"
    @clear="handleClear"
    @error-close="store.clearError"
  >
    <template #input>
      <AppTextarea v-model="input" placeholder="请输入文本..." @submit="handleConvert" />
    </template>
    <template #options>
      <div class="tool-options">
        <el-radio-group v-model="mode">
          <el-radio value="encode">编码</el-radio>
          <el-radio value="decode">解码</el-radio>
        </el-radio-group>
        <el-checkbox v-model="urlSafe">URL Base64</el-checkbox>
      </div>
    </template>
    <template #actions>
      <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleConvert">
        执行
      </el-button>
      <el-button @click="loadSample">示例数据</el-button>
    </template>
    <template #output>
      <AppTextarea v-model="output" placeholder="输出结果..." readonly />
    </template>
  </AppToolLayout>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import AppToolLayout from '../components/AppToolLayout.vue'
import AppTextarea from '../components/AppTextarea.vue'
import {useToolStore} from '../store/tools'
import {EncodeBase64, DecodeBase64} from '../../wailsjs/go/main/App'

const store = useToolStore()
const input = ref('')
const output = ref('')
const mode = ref('encode')
const urlSafe = ref(false)

async function handleConvert() {
  store.setLoading(true)
  store.clearError()
  try {
    const fn = mode.value === 'encode' ? EncodeBase64 : DecodeBase64
    const result = await fn(input.value, urlSafe.value)
    if (result.code !== 0) {
      store.setError(result.message)
    } else {
      output.value = result.data as string
    }
  } catch (e: any) {
    store.setError(e.message || 'Base64 处理失败')
  } finally {
    store.setLoading(false)
  }
}

function loadSample() {
  if (mode.value === 'encode') {
    input.value = 'Hello, DevToys!'
  } else {
    input.value = 'SGVsbG8sIERldlRveXMh'
  }
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
