import {defineStore} from 'pinia'
import {ref} from 'vue'

export interface ToolState {
  loading: boolean
  error: string
}

export const useToolStore = defineStore('tools', () => {
  const loading = ref(false)
  const error = ref('')

  function setLoading(val: boolean) {
    loading.value = val
  }

  function setError(msg: string) {
    error.value = msg
  }

  function clearError() {
    error.value = ''
  }

  function reset() {
    loading.value = false
    error.value = ''
  }

  return {loading, error, setLoading, setError, clearError, reset}
})
