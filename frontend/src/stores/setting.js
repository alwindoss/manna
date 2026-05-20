import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useSettingStore = defineStore('setting', () => {
  const settings = ref({
    "leftMenuVisible": true,
  })

  function toggleSidePanel() {
    if(settings.value.leftMenuVisible) {
        settings.value.leftMenuVisible = false
    } else {
        settings.value.leftMenuVisible = true
    }
  }


  return { settings, toggleSidePanel }
})