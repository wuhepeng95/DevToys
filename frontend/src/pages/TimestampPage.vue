<template>
  <AppToolLayout
    title="时间戳转换"
    description="时间戳与日期互转，支持多种时间单位和格式"
    :loading="store.loading"
    :error="store.error"
    :has-output="!!dateOutput || !!timestampOutput"
    @copy="copyResult"
    @clear="handleClear"
    @error-close="store.clearError"
  >
    <template #input>
      <div class="tool-section">
        <h3>时间戳 → 日期</h3>
        <AppTextarea
          v-model="timestampInput"
          placeholder="请输入时间戳..."
          :rows="2"
          @submit="handleUnixToDate"
        />
        <div class="tool-options">
          <el-radio-group v-model="timestampUnit">
            <el-radio value="s">秒</el-radio>
            <el-radio value="ms">毫秒</el-radio>
            <el-radio value="us">微秒</el-radio>
            <el-radio value="ns">纳秒</el-radio>
          </el-radio-group>
          <el-checkbox v-model="useUTC">UTC</el-checkbox>
        </div>
        <div class="tool-actions">
          <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleUnixToDate">
            转换为日期
          </el-button>
        </div>
        <el-input
          v-model="dateOutput"
          type="textarea"
          :rows="4"
          placeholder="日期结果..."
          readonly
        />
      </div>

      <el-divider />

      <div class="tool-section">
        <h3>日期 → 时间戳</h3>
        <AppTextarea
          v-model="dateInput"
          placeholder="请输入日期 (如 2024-01-01 00:00:00 或 RFC3339)..."
          :rows="2"
          @submit="handleDateToUnix"
        />
        <div class="tool-actions">
          <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleDateToUnix">
            转换为时间戳
          </el-button>
          <el-button @click="setCurrentTime">当前时间</el-button>
          <el-button @click="loadSample">示例数据</el-button>
        </div>
        <el-input
          v-model="timestampOutput"
          type="textarea"
          :rows="4"
          placeholder="时间戳结果..."
          readonly
        />
      </div>
    </template>
  </AppToolLayout>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import AppToolLayout from '../components/AppToolLayout.vue'
import AppTextarea from '../components/AppTextarea.vue'
import {useToolStore} from '../store/tools'
import {UnixToDate, DateToUnix} from '../../wailsjs/go/main/App'

interface UnixToDateData {
  local: string
  utc: string
  iso8601: string
  rfc3339: string
}

interface DateToUnixData {
  seconds: number
  milliseconds: number
  microseconds: number
  nanoseconds: number
}

const store = useToolStore()
const timestampInput = ref('')
const timestampUnit = ref('s')
const useUTC = ref(false)
const dateOutput = ref('')
const dateInput = ref('')
const timestampOutput = ref('')

async function handleUnixToDate() {
  store.setLoading(true)
  store.clearError()
  try {
    const result = await UnixToDate(timestampInput.value, timestampUnit.value, useUTC.value)
    if (result.code !== 0) {
      store.setError(result.message)
    } else {
      const data = result.data as UnixToDateData
      dateOutput.value = [
        `本地时间: ${data.local}`,
        `UTC 时间: ${data.utc}`,
        `ISO8601: ${data.iso8601}`,
        `RFC3339: ${data.rfc3339}`,
      ].join('\n')
    }
  } catch (e: any) {
    store.setError(e.message || '转换失败')
  } finally {
    store.setLoading(false)
  }
}

async function handleDateToUnix() {
  store.setLoading(true)
  store.clearError()
  try {
    const result = await DateToUnix(dateInput.value)
    if (result.code !== 0) {
      store.setError(result.message)
    } else {
      const data = result.data as DateToUnixData
      timestampOutput.value = [
        `秒: ${data.seconds}`,
        `毫秒: ${data.milliseconds}`,
        `微秒: ${data.microseconds}`,
        `纳秒: ${data.nanoseconds}`,
      ].join('\n')
    }
  } catch (e: any) {
    store.setError(e.message || '转换失败')
  } finally {
    store.setLoading(false)
  }
}

function setCurrentTime() {
  const now = new Date()
  dateInput.value = now.toISOString().replace('T', ' ').substring(0, 19)
  handleDateToUnix()
}

function loadSample() {
  timestampInput.value = '1700000000'
}

function handleClear() {
  timestampInput.value = ''
  dateOutput.value = ''
  dateInput.value = ''
  timestampOutput.value = ''
  store.reset()
}

async function copyResult() {
  const text = [dateOutput.value, timestampOutput.value].filter(Boolean).join('\n---\n')
  if (text) {
    await navigator.clipboard.writeText(text)
  }
}
</script>
