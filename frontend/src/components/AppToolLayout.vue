<template>
  <div class="tool-page">
    <h1 class="tool-title">{{ title }}</h1>
    <p class="tool-desc">{{ description }}</p>

    <div v-if="loading" class="tool-loading">
      <el-icon class="loading-icon" :size="24"><Loading /></el-icon>
      <span>处理中...</span>
    </div>

    <div v-if="error" class="tool-error">
      <el-alert :title="error" type="error" show-icon :closable="true" @close="onErrorClose" />
    </div>

    <slot name="input" />

    <slot name="options" />

    <div class="tool-actions">
      <slot name="actions" />
    </div>

    <slot name="output" />

    <div class="tool-bottom-actions">
      <slot name="bottom-actions">
        <el-button @click="onCopy" :disabled="!hasOutput">
          <el-icon><CopyDocument /></el-icon>
          复制结果
        </el-button>
        <el-button @click="onClear">
          <el-icon><Delete /></el-icon>
          清空
        </el-button>
      </slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import {Loading, CopyDocument, Delete} from '@element-plus/icons-vue'

defineProps<{
  title: string
  description: string
  loading?: boolean
  error?: string
  hasOutput?: boolean
}>()

const emit = defineEmits<{
  copy: []
  clear: []
  'error-close': []
}>()

function onCopy() {
  emit('copy')
}

function onClear() {
  emit('clear')
}

function onErrorClose() {
  emit('error-close')
}
</script>

<style scoped>
.tool-page {
  max-width: 960px;
}

.tool-title {
  font-size: 22px;
  font-weight: 700;
  margin-bottom: 4px;
  color: var(--el-text-color-primary);
}

.tool-desc {
  font-size: 13px;
  color: var(--el-text-color-secondary);
  margin-bottom: 20px;
}

.tool-loading {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 0;
  color: var(--el-text-color-secondary);
}

.loading-icon {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.tool-error {
  margin-bottom: 12px;
}

.tool-actions {
  display: flex;
  gap: 8px;
  margin: 12px 0;
  flex-wrap: wrap;
}

.tool-bottom-actions {
  display: flex;
  gap: 8px;
  margin-top: 12px;
}
</style>
