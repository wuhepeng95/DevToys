<template>
  <AppToolLayout
    title="随机密码生成"
    description="生成高强度随机密码，支持多种字符组合"
    :loading="store.loading"
    :error="store.error"
    :has-output="passwords.length > 0"
    @error-close="store.clearError"
    @clear="handleClear"
  >
    <template #options>
      <div class="password-opts">
        <div class="opt-row">
          <label>密码长度：</label>
          <el-slider v-model="length" :min="4" :max="64" :step="1" style="width: 200px" />
          <span class="opt-value">{{ length }}</span>
        </div>
        <div class="opt-row">
          <label>生成数量：</label>
          <el-slider v-model="count" :min="1" :max="50" :step="1" style="width: 200px" />
          <span class="opt-value">{{ count }}</span>
        </div>
        <div class="opt-row">
          <el-checkbox v-model="includeUpper">大写字母 (A-Z)</el-checkbox>
          <el-checkbox v-model="includeLower">小写字母 (a-z)</el-checkbox>
          <el-checkbox v-model="includeDigits">数字 (0-9)</el-checkbox>
          <el-checkbox v-model="includeSpecial">特殊字符</el-checkbox>
          <el-checkbox v-model="excludeAmbiguous">排除易混淆字符 (0O1lI)</el-checkbox>
        </div>
      </div>
    </template>
    <template #actions>
      <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleGenerate">
        生成密码
      </el-button>
    </template>
    <template #output>
      <div v-if="passwords.length > 0" class="password-list">
        <div
          v-for="(pwd, idx) in passwords"
          :key="idx"
          class="password-item"
        >
          <el-input v-model="passwords[idx]" readonly>
            <template #append>
              <el-button @click="copySingle(pwd)">
                <el-icon><CopyDocument /></el-icon>
              </el-button>
            </template>
          </el-input>
        </div>
      </div>
    </template>
    <template #bottom-actions>
      <el-button @click="copyAll" :disabled="passwords.length === 0">
        <el-icon><CopyDocument /></el-icon>
        复制全部
      </el-button>
      <el-button @click="handleClear">
        <el-icon><Delete /></el-icon>
        清空
      </el-button>
    </template>
  </AppToolLayout>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import AppToolLayout from '../components/AppToolLayout.vue'
import {useToolStore} from '../store/tools'
import {GeneratePasswords} from '../../wailsjs/go/main/App'
import {CopyDocument, Delete} from '@element-plus/icons-vue'

const store = useToolStore()
const length = ref(12)
const count = ref(5)
const includeUpper = ref(true)
const includeLower = ref(true)
const includeDigits = ref(true)
const includeSpecial = ref(false)
const excludeAmbiguous = ref(true)
const passwords = ref<string[]>([])

async function handleGenerate() {
  store.setLoading(true)
  store.clearError()
  try {
    const result = await GeneratePasswords({
      length: length.value,
      count: count.value,
      includeUpper: includeUpper.value,
      includeLower: includeLower.value,
      includeDigits: includeDigits.value,
      includeSpecial: includeSpecial.value,
      excludeAmbiguous: excludeAmbiguous.value,
    })
    if (result.code !== 0) {
      store.setError(result.message)
    } else {
      passwords.value = (result.data as {passwords: string[]}).passwords
    }
  } catch (e: any) {
    store.setError(e.message || '生成失败')
  } finally {
    store.setLoading(false)
  }
}

function handleClear() {
  passwords.value = []
  store.reset()
}

async function copySingle(pwd: string) {
  await navigator.clipboard.writeText(pwd)
}

async function copyAll() {
  await navigator.clipboard.writeText(passwords.value.join('\n'))
}
</script>

<style scoped>
.password-opts {
  margin: 12px 0;
}

.opt-row {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 12px;
  flex-wrap: wrap;
}

.opt-row label {
  font-size: 13px;
  color: var(--el-text-color-secondary);
  min-width: 80px;
}

.opt-value {
  font-size: 14px;
  font-weight: 600;
  color: var(--el-color-primary);
  min-width: 30px;
}

.password-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-top: 12px;
}

.password-item {
  font-family: 'SF Mono', 'Fira Code', monospace;
}
</style>
