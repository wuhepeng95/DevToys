<template>
  <div class="app-container">
    <div class="app-body">
      <AppSidebar />
      <main class="app-main">
        <router-view v-slot="{Component}">
          <keep-alive>
            <component :is="Component" />
          </keep-alive>
        </router-view>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, onUnmounted} from 'vue'
import AppSidebar from './layouts/AppSidebar.vue'
import {useThemeStore} from './store/theme'

const theme = useThemeStore()

function onGlobalKeydown(e: KeyboardEvent) {
  if ((e.metaKey || e.ctrlKey) && e.key === 'Enter') {
    const active = document.activeElement
    if (active && active.closest('.tool-page')) {
      const submitBtn = active.closest('.tool-page')?.querySelector('.tool-primary-btn') as HTMLElement
      submitBtn?.click()
    }
  }
}

onMounted(() => {
  theme.init()
  window.addEventListener('keydown', onGlobalKeydown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', onGlobalKeydown)
})
</script>

<style scoped>
.app-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background-color: var(--el-bg-color);
}

.app-body {
  display: flex;
  flex: 1;
  overflow: hidden;
}

.app-main {
  flex: 1;
  overflow-y: auto;
  padding: 24px;
}
</style>
