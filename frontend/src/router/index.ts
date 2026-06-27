import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/json',
  },
  {
    path: '/json',
    name: 'JSON格式化',
    component: () => import('../pages/JsonPage.vue'),
  },
  {
    path: '/text-duplicate',
    name: '文本去重',
    component: () => import('../pages/TextDuplicatePage.vue'),
  },
  {
    path: '/text-sort',
    name: '文本排序',
    component: () => import('../pages/TextSortPage.vue'),
  },
  {
    path: '/base64',
    name: 'Base64',
    component: () => import('../pages/Base64Page.vue'),
  },
  {
    path: '/diff',
    name: '文本Diff',
    component: () => import('../pages/DiffPage.vue'),
  },
  {
    path: '/timestamp',
    name: '时间戳转换',
    component: () => import('../pages/TimestampPage.vue'),
  },
  {
    path: '/url',
    name: 'URL 编解码',
    component: () => import('../pages/UrlPage.vue'),
  },
  {
    path: '/password',
    name: '随机密码',
    component: () => import('../pages/PasswordPage.vue'),
  },
  {
    path: '/yaml',
    name: 'YAML 转换',
    component: () => import('../pages/YamlPage.vue'),
  },
  {
    path: '/sql',
    name: 'SQL 格式化',
    component: () => import('../pages/SqlPage.vue'),
  },
  {
    path: '/convert',
    name: '进制转换',
    component: () => import('../pages/ConvertPage.vue'),
  },
  {
    path: '/subnet',
    name: '子网计算',
    component: () => import('../pages/SubnetPage.vue'),
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router
