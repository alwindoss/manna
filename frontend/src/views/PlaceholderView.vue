<template>
  <div class="view-placeholder">
    <div class="placeholder-icon">🔖</div>
    <h2>Bookmarks</h2>
    <p>This section is under construction.</p>
  </div>
</template>

<script setup>
import PlaceholderRightPanel from '@/components/ui/PlaceholderRightPanel.vue';
import { useRoute } from 'vue-router'
import { onMounted } from 'vue';

const route = useRoute()
import { useUiStore } from '@/stores/ui'

const ui = useUiStore()

onMounted(() => {
  ui.setRightPanel(PlaceholderRightPanel, {
    title: 'Placeholder Panel',
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