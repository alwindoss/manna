<template>
  <div class="view-placeholder">
    <div class="placeholder-icon">{{ route.meta?.icon ?? '📄' }}</div>
    <h2>{{ route.meta?.title ?? 'Coming Soon' }}</h2>
    <p>This section is under construction.</p>
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router'
const route = useRoute()
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