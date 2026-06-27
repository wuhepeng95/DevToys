<template>
  <AppToolLayout
    title="文本去重"
    description="去除重复行，支持保留顺序/排序、大小写控制"
    :loading="store.loading"
    :error="store.error"
    :has-output="!!output"
    @copy="copyResult"
    @clear="handleClear"
    @error-close="store.clearError"
  >
    <template #input>
      <AppTextarea v-model="input" placeholder="请输入文本..." @submit="handleRemoveDuplicate" />
    </template>
    <template #options>
      <div class="tool-options">
        <el-checkbox v-model="keepOrder">保留顺序</el-checkbox>
        <el-checkbox v-model="sortAfter">排序后去重</el-checkbox>
        <el-checkbox v-model="ignoreCase">忽略大小写</el-checkbox>
        <el-checkbox v-model="removeEmpty">删除空行</el-checkbox>
      </div>
    </template>
    <template #actions>
      <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleRemoveDuplicate">
        去重
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
import {RemoveDuplicate} from '../../wailsjs/go/main/App'

const store = useToolStore()
const input = ref('')
const output = ref('')
const keepOrder = ref(true)
const sortAfter = ref(false)
const ignoreCase = ref(false)
const removeEmpty = ref(true)

async function handleRemoveDuplicate() {
  store.setLoading(true)
  store.clearError()
  try {
    const result = await RemoveDuplicate(input.value, {
      keepOrder: keepOrder.value,
      ignoreCase: ignoreCase.value,
      removeEmpty: removeEmpty.value,
    })
    if (result.code !== 0) {
      store.setError(result.message)
    } else {
      output.value = result.data as string
    }
  } catch (e: any) {
    store.setError(e.message || '去重失败')
  } finally {
    store.setLoading(false)
  }
}

function loadSample() {
  input.value = ['a', 'b', 'a', 'c', 'b', 'd'].join('\n')
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
