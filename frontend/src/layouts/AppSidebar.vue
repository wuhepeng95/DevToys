<template>
  <aside class="sidebar">
    <div class="sidebar-header">
      <div class="sidebar-title-row">
        <h2 class="sidebar-title">DevToys</h2>
        <el-button
          :icon="theme.isDark ? Sunny : Moon"
          text
          @click="theme.toggle"
          class="theme-toggle"
        />
      </div>
      <el-input
        ref="searchRef"
        v-model="search"
        placeholder="⌘⇧F"
        :prefix-icon="SearchIcon"
        clearable
        size="small"
        class="sidebar-search"
        @keydown.escape="search = ''"
      />
    </div>
    <el-menu
      :default-active="currentRoute"
      router
      class="sidebar-menu"
    >
      <template v-for="group in filteredGroups" :key="group.title">
        <el-menu-item-group :title="group.title">
          <el-menu-item
            v-for="item in group.items"
            :key="item.path"
            :index="item.path"
          >
            <el-icon><component :is="item.icon" /></el-icon>
            <span>{{ item.name }}</span>
          </el-menu-item>
        </el-menu-item-group>
      </template>
    </el-menu>
  </aside>
</template>

<script setup lang="ts">
import {computed, ref, onMounted, onUnmounted, markRaw, type Component} from 'vue'
import {useRoute} from 'vue-router'
import {
  Remove,
  Sort,
  CopyDocument,
  Key,
  Link,
  Switch,
  Edit,
  Document,
  DataAnalysis,
  Clock,
  Connection,
  Lock,
  Search as SearchIcon,
  Sunny,
  Moon,
} from '@element-plus/icons-vue'
import {useThemeStore} from '../store/theme'

interface MenuItem {
  path: string
  name: string
  keywords: string[]
  icon: Component
}

interface MenuGroup {
  title: string
  items: MenuItem[]
}

const theme = useThemeStore()
const route = useRoute()
const currentRoute = computed(() => route.path)
const search = ref('')
const searchRef = ref()

const menuGroups: MenuGroup[] = [
  {
    title: '文本工具',
    items: [
      {path: '/text-duplicate', name: '文本去重', keywords: ['删除重复', 'unique', '去重', '重复行'], icon: markRaw(Remove)},
      {path: '/text-sort', name: '文本排序', keywords: ['sort', 'order', '排序', '升序', '降序', '整理'], icon: markRaw(Sort)},
      {path: '/diff', name: '文本 Diff', keywords: ['对比', '差异', '比较', 'difference', 'compare'], icon: markRaw(CopyDocument)},
    ],
  },
  {
    title: '编码工具',
    items: [
      {path: '/base64', name: 'Base64', keywords: ['base64', '编码', '解码', 'encode', 'decode'], icon: markRaw(Key)},
      {path: '/url', name: 'URL 编解码', keywords: ['url', 'encode', 'decode', '编码', '解码', 'query', '参数'], icon: markRaw(Link)},
      {path: '/convert', name: '进制转换', keywords: ['进制', 'base', 'binary', 'binary', 'hex', 'hexadecimal', 'octal', 'decimal', '2进制', '8进制', '10进制', '16进制'], icon: markRaw(Switch)},
    ],
  },
  {
    title: '格式化工具',
    items: [
      {path: '/json', name: 'JSON 格式化', keywords: ['json', '格式化', '压缩', 'format', 'pretty', 'validate', '校验'], icon: markRaw(Edit)},
      {path: '/yaml', name: 'YAML 转换', keywords: ['yaml', 'yml', 'json', '转换', 'convert'], icon: markRaw(Document)},
      {path: '/sql', name: 'SQL 格式化', keywords: ['sql', '格式化', '压缩', 'format', 'minify', '数据库'], icon: markRaw(DataAnalysis)},
    ],
  },
  {
    title: '时间工具',
    items: [
      {path: '/timestamp', name: '时间戳转换', keywords: ['timestamp', '时间戳', 'unix', '日期', 'time', '转换'], icon: markRaw(Clock)},
    ],
  },
  {
    title: '网络工具',
    items: [
      {path: '/subnet', name: '子网计算', keywords: ['子网', 'subnet', 'cidr', 'ip', '网络', '掩码', 'mask', '拆分'], icon: markRaw(Connection)},
    ],
  },
  {
    title: '安全工具',
    items: [
      {path: '/password', name: '随机密码', keywords: ['密码', 'password', '随机', 'generate', '生成', '安全'], icon: markRaw(Lock)},
    ],
  },
]

const filteredGroups = computed(() => {
  const q = search.value.trim().toLowerCase()
  if (!q) return menuGroups
  return menuGroups
    .map(group => ({
      ...group,
      items: group.items.filter(item =>
        item.name.toLowerCase().includes(q) ||
        item.keywords.some(k => k.toLowerCase().includes(q)),
      ),
    }))
    .filter(g => g.items.length > 0)
})

function onGlobalKeydown(e: KeyboardEvent) {
  if ((e.metaKey || e.ctrlKey) && e.shiftKey && e.code === 'KeyF') {
    e.preventDefault()
    searchRef.value?.focus()
  }
  if (e.key === '/' && searchRef.value) {
    const tag = (e.target as HTMLElement)?.tagName
    if (tag !== 'INPUT' && tag !== 'TEXTAREA') {
      e.preventDefault()
      searchRef.value?.focus()
    }
  }
}

onMounted(() => {
  window.addEventListener('keydown', onGlobalKeydown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', onGlobalKeydown)
})
</script>

<style scoped>
.sidebar {
  width: 220px;
  min-width: 220px;
  background-color: var(--el-bg-color-page);
  border-right: 1px solid var(--el-border-color-light);
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.sidebar-header {
  padding: 16px 12px 8px;
  flex-shrink: 0;
}

.sidebar-title-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 8px;
}

.sidebar-title {
  font-size: 18px;
  font-weight: 700;
  margin: 0;
  color: var(--el-text-color-primary);
}

.theme-toggle {
  font-size: 18px;
}

.sidebar-search {
  width: 100%;
}

.sidebar-menu {
  border-right: none;
  flex: 1;
  overflow-y: auto;
}
</style>
