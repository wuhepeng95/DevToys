<template>
  <AppToolLayout
    title="文本排序"
    description="对文本行进行排序"
    :loading="store.loading"
    :error="store.error"
    :has-output="!!output"
    @copy="copyResult"
    @clear="handleClear"
    @error-close="store.clearError"
  >
    <template #input>
      <AppTextarea v-model="input" placeholder="请输入文本..." @submit="handleSort" />
    </template>
    <template #options>
      <div class="tool-options">
        <el-radio-group v-model="sortOrder">
          <el-radio value="asc">升序</el-radio>
          <el-radio value="desc">降序</el-radio>
        </el-radio-group>
        <el-checkbox v-model="numericSort">数字排序</el-checkbox>
        <el-checkbox v-model="ignoreCase">忽略大小写</el-checkbox>
      </div>
    </template>
    <template #actions>
      <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleSort">
        排序
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
import {SortLines} from '../../wailsjs/go/main/App'

const store = useToolStore()
const input = ref('')
const output = ref('')
const sortOrder = ref('asc')
const numericSort = ref(false)
const ignoreCase = ref(false)

async function handleSort() {
  store.setLoading(true)
  store.clearError()
  try {
    const result = await SortLines(input.value, {
      order: sortOrder.value,
      numeric: numericSort.value,
      ignoreCase: ignoreCase.value,
    })
    if (result.code !== 0) {
      store.setError(result.message)
    } else {
      output.value = result.data as string
    }
  } catch (e: any) {
    store.setError(e.message || '排序失败')
  } finally {
    store.setLoading(false)
  }
}

function loadSample() {
  input.value = ['banana', 'apple', 'Cherry', 'date', '3', '10', '1'].join('\n')
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
