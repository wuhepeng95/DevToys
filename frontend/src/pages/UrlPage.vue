<template>
  <AppToolLayout
    title="URL 编解码"
    description="URL 编码 / 解码 / 解析"
    :loading="store.loading"
    :error="store.error"
    :has-output="!!output || !!parseResult"
    @copy="copyResult"
    @clear="handleClear"
    @error-close="store.clearError"
  >
    <template #options>
      <div class="tool-options">
        <el-radio-group v-model="mode">
          <el-radio value="encode">编码</el-radio>
          <el-radio value="decode">解码</el-radio>
          <el-radio value="parse">解析</el-radio>
        </el-radio-group>
        <el-radio-group v-model="encodingType" v-if="mode !== 'parse'">
          <el-radio value="url">URL 编码</el-radio>
          <el-radio value="query">Query 参数</el-radio>
        </el-radio-group>
      </div>
    </template>
    <template #input>
      <AppTextarea
        v-model="input"
        :placeholder="inputPlaceholder"
        :rows="mode === 'parse' ? 3 : 8"
        @submit="handleConvert"
      />
    </template>
    <template #actions>
      <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleConvert">
        {{ mode === 'encode' ? '编码' : mode === 'decode' ? '解码' : '解析' }}
      </el-button>
      <el-button @click="loadSample">示例数据</el-button>
    </template>
    <template #output>
      <div v-if="mode === 'parse' && parseResult" class="parse-table">
        <el-table :data="parseRows" border stripe size="small">
          <el-table-column prop="key" label="字段" width="120" />
          <el-table-column prop="value" label="值" />
        </el-table>
      </div>
      <AppTextarea
        v-else
        v-model="output"
        placeholder="输出结果..."
        readonly
      />
    </template>
  </AppToolLayout>
</template>

<script setup lang="ts">
import {ref, computed} from 'vue'
import AppToolLayout from '../components/AppToolLayout.vue'
import AppTextarea from '../components/AppTextarea.vue'
import {useToolStore} from '../store/tools'
import {URLEncode, URLDecode, QueryEncode, QueryDecode, ParseURL} from '../../wailsjs/go/main/App'

interface URLParts {
  scheme: string
  host: string
  port: string
  path: string
  query: string
  fragment: string
  raw: string
}

const store = useToolStore()
const input = ref('')
const output = ref('')
const mode = ref('encode')
const encodingType = ref('url')
const parseResult = ref<URLParts | null>(null)

const inputPlaceholder = computed(() => {
  if (mode.value === 'encode') return '请输入要编码的文本...'
  if (mode.value === 'decode') return '请输入要解码的 URL 编码字符串...'
  return '请输入 URL (如 https://example.com/path?key=val)...'
})

const parseRows = computed(() => {
  const p = parseResult.value
  if (!p) return []
  return [
    {key: '完整 URL', value: p.raw},
    {key: '协议 (Scheme)', value: p.scheme},
    {key: '主机 (Host)', value: p.host},
    {key: '端口 (Port)', value: p.port || '(默认)'},
    {key: '路径 (Path)', value: p.path || '/'},
    {key: '查询参数 (Query)', value: p.query || '(无)'},
    {key: '片段 (Fragment)', value: p.fragment || '(无)'},
  ]
})

async function handleConvert() {
  store.setLoading(true)
  store.clearError()
  parseResult.value = null
  try {
    if (mode.value === 'parse') {
      const result = await ParseURL(input.value)
      if (result.code !== 0) {
        store.setError(result.message)
      } else {
        parseResult.value = result.data as URLParts
        output.value = ''
      }
    } else {
      const fn = mode.value === 'encode'
        ? (encodingType.value === 'url' ? URLEncode : QueryEncode)
        : (encodingType.value === 'url' ? URLDecode : QueryDecode)
      const result = await fn(input.value)
      if (result.code !== 0) {
        store.setError(result.message)
      } else {
        output.value = result.data as string
      }
    }
  } catch (e: any) {
    store.setError(e.message || '处理失败')
  } finally {
    store.setLoading(false)
  }
}

function loadSample() {
  if (mode.value === 'encode') {
    input.value = 'hello world & foo=bar'
  } else if (mode.value === 'decode') {
    input.value = 'hello+world+%26+foo%3Dbar'
  } else {
    input.value = 'https://example.com:443/path/to/page?name=hello&key=val#section'
  }
}

function handleClear() {
  input.value = ''
  output.value = ''
  parseResult.value = null
  store.reset()
}

async function copyResult() {
  if (parseResult.value) {
    const text = parseRows.value.map(r => `${r.key}: ${r.value}`).join('\n')
    await navigator.clipboard.writeText(text)
  } else if (output.value) {
    await navigator.clipboard.writeText(output.value)
  }
}
</script>

<style scoped>
.parse-table {
  margin-top: 12px;
}
</style>
