<template>
  <AppToolLayout
    title="文本 Diff"
    description="对比两段文本差异"
    :loading="store.loading"
    :error="store.error"
    :has-output="!!output"
    @copy="copyResult"
    @clear="handleClear"
    @error-close="store.clearError"
  >
    <template #input>
      <div class="diff-inputs">
        <AppTextarea v-model="inputLeft" placeholder="输入左侧文本..." @submit="handleDiff" />
        <AppTextarea v-model="inputRight" placeholder="输入右侧文本..." @submit="handleDiff" />
      </div>
    </template>
    <template #actions>
      <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleDiff">
        对比
      </el-button>
      <el-button @click="loadSample">示例数据</el-button>
    </template>
    <template #output>
      <div class="diff-output">
        <div
          v-for="(line, idx) in diffLines"
          :key="idx"
          :class="['diff-line', 'diff-' + line.type]"
        ><pre>{{ line.text }}</pre></div>
      </div>
    </template>
  </AppToolLayout>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import AppToolLayout from '../components/AppToolLayout.vue'
import AppTextarea from '../components/AppTextarea.vue'
import {useToolStore} from '../store/tools'
import {CompareText} from '../../wailsjs/go/main/App'

interface DiffLine {
  type: 'same' | 'add' | 'remove'
  text: string
}

const store = useToolStore()
const inputLeft = ref('')
const inputRight = ref('')
const output = ref('')
const diffLines = ref<DiffLine[]>([])

async function handleDiff() {
  store.setLoading(true)
  store.clearError()
  try {
    const result = await CompareText(inputLeft.value, inputRight.value)
    if (result.code !== 0) {
      store.setError(result.message)
    } else {
      const data = result.data as {lines: DiffLine[]; text: string}
      diffLines.value = data.lines
      output.value = data.text
    }
  } catch (e: any) {
    store.setError(e.message || '对比失败')
  } finally {
    store.setLoading(false)
  }
}

function loadSample() {
  inputLeft.value = ['Hello World', 'Line 2', 'Line 3'].join('\n')
  inputRight.value = ['Hello World', 'Line 2 Modified', 'Line 4'].join('\n')
}

function handleClear() {
  inputLeft.value = ''
  inputRight.value = ''
  output.value = ''
  diffLines.value = []
  store.reset()
}

async function copyResult() {
  if (output.value) {
    await navigator.clipboard.writeText(output.value)
  }
}
</script>

<style scoped>
.diff-inputs {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.diff-output {
  margin-top: 12px;
  border: 1px solid var(--el-border-color);
  border-radius: 4px;
  padding: 8px 12px;
  background-color: var(--el-fill-color-light);
  min-height: 200px;
  max-height: 400px;
  overflow: auto;
  font-family: 'SF Mono', 'Fira Code', monospace;
  font-size: 13px;
  line-height: 1.5;
}

.diff-line pre {
  margin: 0;
  white-space: pre-wrap;
}

.diff-add {
  background-color: rgba(0, 200, 83, 0.15);
}

.diff-add pre {
  color: var(--el-color-success);
}

.diff-remove {
  background-color: rgba(255, 71, 87, 0.15);
}

.diff-remove pre {
  color: var(--el-color-danger);
}
</style>
