<template>
  <div class="code-editor" :class="{'is-readonly': readonly}">
    <Codemirror
      :model-value="modelValue"
      :extensions="extensions"
      :disabled="readonly"
      :placeholder="placeholder"
      :style="{height: editorHeight}"
      :autofocus="autoFocus"
      @update:model-value="onChange"
    />
  </div>
</template>

<script setup lang="ts">
import {computed} from 'vue'
import {Codemirror} from 'vue-codemirror'
import {basicSetup} from 'codemirror'
import {EditorView} from '@codemirror/view'
import {oneDark} from '@codemirror/theme-one-dark'
import {json} from '@codemirror/lang-json'
import {sql} from '@codemirror/lang-sql'

import {useThemeStore} from '../store/theme'
import type {Extension} from '@codemirror/state'

const props = withDefaults(defineProps<{
  modelValue?: string
  placeholder?: string
  readonly?: boolean
  disabled?: boolean
  autoFocus?: boolean
  language?: 'json' | 'yaml' | 'sql' | 'text'
  rows?: number
}>(), {
  modelValue: '',
  placeholder: '',
  readonly: false,
  disabled: false,
  autoFocus: false,
  language: 'text',
  rows: 8,
})

const emit = defineEmits<{
  'update:modelValue': [value: string]
  'submit': []
}>()

const themeStore = useThemeStore()

const editorHeight = computed(() => `${props.rows * 24}px`)

const langExt = computed((): Extension => {
  switch (props.language) {
    case 'json': return json()
    case 'sql': return sql()
    case 'yaml': return []
    default: return []
  }
})

const extensions = computed((): Extension[] => {
  const ext: Extension[] = [
    basicSetup,
    langExt.value,
    EditorView.theme({
      '&': {fontSize: '13px'},
      '.cm-scroller': {fontFamily: "'SF Mono', 'Fira Code', 'JetBrains Mono', monospace"},
      '.cm-placeholder': {color: 'var(--el-text-color-placeholder)'},
      '&.cm-focused': {outline: 'none'},
      '.cm-gutters': {borderRight: '1px solid var(--el-border-color-light)'},
    }),
    EditorView.lineWrapping,
    EditorView.domEventHandlers({
      keydown: (e) => {
        if ((e.metaKey || e.ctrlKey) && e.key === 'Enter') {
          emit('submit')
        }
      },
    }),
  ]

  if (themeStore.isDark) {
    ext.push(oneDark)
  }

  return ext
})

function onChange(val: string) {
  emit('update:modelValue', val)
}


</script>

<style scoped>
.code-editor {
  border: 1px solid var(--el-border-color);
  border-radius: 4px;
  overflow: hidden;
  transition: border-color 0.2s;
}

.code-editor:focus-within {
  border-color: var(--el-color-primary);
}

.code-editor.is-readonly {
  background-color: var(--el-fill-color-light);
}
</style>
