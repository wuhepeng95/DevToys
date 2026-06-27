<template>
  <AppToolLayout
    title="YAML 转换"
    description="JSON ↔ YAML 互转"
    :loading="store.loading"
    :error="store.error"
    :has-output="!!output"
    @copy="copyResult"
    @clear="handleClear"
    @error-close="store.clearError"
  >
    <template #options>
      <div class="tool-options">
        <el-radio-group v-model="direction">
          <el-radio value="json2yaml">JSON → YAML</el-radio>
          <el-radio value="yaml2json">YAML → JSON</el-radio>
        </el-radio-group>
      </div>
    </template>
    <template #input>
      <CodeEditor
        v-model="input"
        :language="inputLang"
        :placeholder="inputPlaceholder"
        @submit="handleConvert"
      />
    </template>
    <template #actions>
      <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleConvert">
        转换
      </el-button>
      <el-button v-if="direction === 'yaml2json'" @click="compressOutput" :loading="store.loading">
        压缩 JSON
      </el-button>
      <el-button @click="loadSample">示例数据</el-button>
    </template>
    <template #output>
      <CodeEditor
        v-model="output"
        :language="outputLang"
        placeholder="输出结果..."
        readonly
      />
    </template>
  </AppToolLayout>
</template>

<script setup lang="ts">
import {ref, computed} from 'vue'
import AppToolLayout from '../components/AppToolLayout.vue'
import CodeEditor from '../components/CodeEditor.vue'
import {useToolStore} from '../store/tools'
import {JSONToYAML, YAMLToJSON, YAMLToJSONCompact} from '../../wailsjs/go/main/App'

const store = useToolStore()
const input = ref('')
const output = ref('')
const direction = ref('json2yaml')

const inputLang = computed(() => direction.value === 'json2yaml' ? 'json' : 'yaml')
const outputLang = computed(() => direction.value === 'json2yaml' ? 'yaml' : 'json')

const inputPlaceholder = computed(() => {
  return direction.value === 'json2yaml'
    ? '请输入 JSON 字符串...'
    : '请输入 YAML 字符串...'
})

async function handleConvert() {
  store.setLoading(true)
  store.clearError()
  try {
    const fn = direction.value === 'json2yaml' ? JSONToYAML : YAMLToJSON
    const result = await fn(input.value)
    if (result.code !== 0) {
      store.setError(result.message)
    } else {
      output.value = result.data as string
    }
  } catch (e: any) {
    store.setError(e.message || '转换失败')
  } finally {
    store.setLoading(false)
  }
}

async function compressOutput() {
  store.setLoading(true)
  store.clearError()
  try {
    const result = await YAMLToJSONCompact(input.value)
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
  if (direction.value === 'json2yaml') {
    input.value = JSON.stringify({
      name: 'DevToys',
      version: 1,
      features: ['json', 'yaml', 'base64'],
      config: {enabled: true, timeout: 30},
    }, null, 2)
  } else {
    input.value = 'name: DevToys\nversion: 1\nfeatures:\n  - json\n  - yaml\n  - base64\nconfig:\n  enabled: true\n  timeout: 30\n'
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
