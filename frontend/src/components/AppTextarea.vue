<template>
  <el-input
    ref="inputRef"
    :model-value="modelValue"
    type="textarea"
    :rows="rows"
    :placeholder="placeholder"
    :readonly="readonly"
    :disabled="disabled"
    class="app-textarea"
    @update:model-value="onInput"
    @dragenter.prevent
    @dragover.prevent
    @drop="onDrop"
    @keydown="onKeydown"
  />
</template>

<script setup lang="ts">
import type {ElInput} from 'element-plus'

const props = withDefaults(defineProps<{
  modelValue?: string
  placeholder?: string
  rows?: number
  readonly?: boolean
  disabled?: boolean
  autoFocus?: boolean
}>(), {
  modelValue: '',
  placeholder: '',
  rows: 8,
  readonly: false,
  disabled: false,
  autoFocus: false,
})

const emit = defineEmits<{
  'update:modelValue': [value: string]
  'submit': []
}>()

function onInput(val: string) {
  emit('update:modelValue', val)
}

function onDrop(e: DragEvent) {
  const text = e.dataTransfer?.getData('text')
  if (text) {
    emit('update:modelValue', text)
  }
}

function onKeydown(e: KeyboardEvent) {
  if ((e.metaKey || e.ctrlKey) && e.key === 'Enter') {
    emit('submit')
  }
}
</script>

<style scoped>
.app-textarea {
  width: 100%;
}
</style>
