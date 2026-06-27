<template>
  <AppToolLayout
    title="子网拆分"
    description="CIDR 子网拆分，点击按钮切换 split/merge"
    :loading="store.loading"
    :error="store.error"
    :has-output="flatRows.length > 0"
    @error-close="store.clearError"
    @clear="handleClear"
  >
    <template #input>
      <AppTextarea v-model="input" :rows="2" placeholder="例如: 192.168.1.0/24" @submit="handleCalc" />
    </template>
    <template #actions>
      <el-button type="primary" class="tool-primary-btn" :loading="store.loading" @click="handleCalc">
        拆分
      </el-button>
      <el-button @click="loadSample">示例数据</el-button>
    </template>
    <template #output>
      <div v-if="flatRows.length > 0">
        <el-table :data="flatRows" border size="small" max-height="500" :cell-style="cellStyle" row-key="rowKey">
          <el-table-column label="子网" min-width="200">
            <template #default="{row}">
              <span class="mono">{{ row.node.data.cidr }}</span>
            </template>
          </el-table-column>
          <el-table-column label="子网掩码" width="140">
            <template #default="{row}">
              <span class="mono">{{ row.node.data.subnetMask }}</span>
            </template>
          </el-table-column>
          <el-table-column label="广播地址" width="140">
            <template #default="{row}">
              <span class="mono">{{ row.node.data.broadcast }}</span>
            </template>
          </el-table-column>
          <el-table-column label="可用范围" width="290">
            <template #default="{row}">
              <span class="mono">{{ row.node.data.usableStart }} — {{ row.node.data.usableEnd }}</span>
            </template>
          </el-table-column>
          <el-table-column label="可用 IP" width="90">
            <template #default="{row}">
              {{ row.node.data.usableCount }}
            </template>
          </el-table-column>
          <el-table-column label="操作" width="90">
            <template #default="{row}">
              <el-button
                text
                size="small"
                :type="row.node.expanded ? 'warning' : 'primary'"
                @click="toggleRow(row.node)"
                :disabled="row.node.data.netSize >= 32"
              >
                {{ row.node.expanded ? '合并' : '/' + (row.node.data.netSize + 1) }}
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </template>
    <template #bottom-actions>
      <el-button @click="copyResult" :disabled="flatRows.length === 0">
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
import {ref, computed} from 'vue'
import AppToolLayout from '../components/AppToolLayout.vue'
import AppTextarea from '../components/AppTextarea.vue'
import {useToolStore} from '../store/tools'
import {CalculateSubnet, SplitSubnet} from '../../wailsjs/go/main/App'
import {CopyDocument, Delete} from '@element-plus/icons-vue'

interface SubnetData {
  cidr: string
  network: string
  netSize: number
  subnetMask: string
  wildcardMask: string
  broadcast: string
  usableStart: string
  usableEnd: string
  usableCount: number
  hostCount: number
  binaryNetwork: string
  binaryMask: string
}

interface SubnetNode {
  data: SubnetData
  children: SubnetNode[]
  expanded: boolean
}

interface FlatRow {
  node: SubnetNode
  depth: number
}

const depthColors = [
  'var(--el-color-primary-light-9)',
  'var(--el-color-success-light-9)',
  'var(--el-color-warning-light-9)',
  'var(--el-color-danger-light-9)',
  'var(--el-color-info-light-9)',
  'var(--el-color-primary-light-7)',
  'var(--el-color-success-light-7)',
  'var(--el-color-warning-light-7)',
]

const store = useToolStore()
const input = ref('')
const roots = ref<SubnetNode[]>([])

const flatRows = computed((): FlatRow[] => {
  const result: FlatRow[] = []
  function walk(nodes: SubnetNode[], depth: number) {
    for (const n of nodes) {
      result.push({node: n, depth})
      if (n.children.length > 0 && n.expanded) {
        walk(n.children, depth + 1)
      }
    }
  }
  walk(roots.value, 0)
  return result
})

function rowKey(row: FlatRow) {
  return row.node.data.cidr + '_' + row.depth
}

function cellStyle({row}: {row: FlatRow}) {
  if (row.depth > 0) {
    return {backgroundColor: depthColors[(row.depth - 1) % depthColors.length]}
  }
  return {}
}

async function handleCalc() {
  store.setLoading(true)
  store.clearError()
  try {
    const cidr = input.value.trim()
    if (!cidr.includes('/')) {
      store.setError('请输入 CIDR 格式，如 192.168.1.0/24')
      return
    }
    const res = await CalculateSubnet(cidr)
    if (res.code !== 0) {
      store.setError(res.message)
    } else {
      const data = res.data as SubnetData
      roots.value = [{
        data,
        children: [],
        expanded: false,
      }]
    }
  } catch (e: any) {
    store.setError(e.message || '拆分失败')
  } finally {
    store.setLoading(false)
  }
}

async function toggleRow(node: SubnetNode) {
  if (node.children.length > 0 && node.expanded) {
    node.expanded = false
    node.children = []
    return
  }

  if (node.children.length > 0) {
    node.expanded = true
    return
  }

  const nextSize = node.data.netSize + 1
  if (nextSize > 32) {
    store.setError('已达到最大子网掩码 /32')
    return
  }

  store.setLoading(true)
  try {
    const res = await SplitSubnet(node.data.cidr, nextSize)
    if (res.code === 0) {
      const data = res.data as {subnets: SubnetData[]}
      node.children = data.subnets.map(s => ({
        data: s,
        children: [],
        expanded: false,
      }))
      node.expanded = true
    } else {
      store.setError(res.message)
    }
  } catch (e: any) {
    store.setError(e.message || '拆分失败')
  } finally {
    store.setLoading(false)
  }
}

function loadSample() {
  input.value = '192.168.1.0/24'
}

function handleClear() {
  input.value = ''
  roots.value = []
  store.reset()
}

async function copyResult() {
  const lines: string[] = []
  for (const f of flatRows.value) {
    const n = f.node.data
    lines.push(`${n.cidr}  ${n.usableStart}—${n.usableEnd}  ${n.usableCount} IP`)
  }
  await navigator.clipboard.writeText(lines.join('\n'))
}
</script>

<style scoped>
</style>
