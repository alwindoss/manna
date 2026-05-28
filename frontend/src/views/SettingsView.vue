<template>
  <div class="view-placeholder">
    <div class="placeholder-icon"><Settings /></div>
    <h2>Settings</h2>
    <p>This section is under construction.</p>
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router'
import { onMounted } from 'vue';
import { useUiStore } from '@/stores/ui'
import SettingsRightPanel from '@/components/settings/SettingsRightPanel.vue';
import { Settings } from '@lucide/vue';

const route = useRoute()

const ui = useUiStore()

onMounted(() => {
  ui.setRightPanel(SettingsRightPanel, {
    title: 'Settings',
  })

})
</script>

<style scoped>
.view-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  min-height: 60vh;
  color: var(--ink-light);
  text-align: center;
  gap: 14px;
}
.placeholder-icon { font-size: 3rem; }
.view-placeholder h2 { font-size: 1.2rem; color: var(--ink-mid); }
.view-placeholder p  { font-size: 0.88rem; font-style: italic; }
</style>