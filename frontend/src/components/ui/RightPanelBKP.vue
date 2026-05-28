<template>
  <aside class="right-panel">

    <div class="right-header">
      <h3 class="right-title">{{ panelTitle }}</h3>
      <button class="close-btn" @click="ui.closeRightPanel">✕</button>
    </div>

    <div class="right-content">

      <!-- ── HOME panel ─────────────────────── -->
      <template v-if="routeName === 'home'">
        <HomeRightPanel></HomeRightPanel>
      </template>

      <!-- ── READ panel ─────────────────────── -->
      <template v-else-if="routeName === 'read'">
        <ReadBibleRightPanel></ReadBibleRightPanel>
      </template>

      <!-- ── NOTES panel ────────────────────── -->
      <template v-else-if="routeName === 'notes'">
        <NotesRightPanel></NotesRightPanel>
      </template>

      <!-- ── FALLBACK ───────────────────────── -->
      <template v-else>
        <p class="empty-panel">No contextual panel for this section.</p>
      </template>

    </div>
  </aside>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useUiStore } from '@/stores/ui'
import ReadBibleRightPanel from '../bible/ReadBibleRightPanel.vue'
import NotesRightPanel from '../notes/NotesRightPanel.vue'
import HomeRightPanel from '../home/HomeRightPanel.vue'

const ui    = useUiStore()
const route = useRoute()

const routeName = computed(() => route.name)

const panelTitles = {
  home:      'Today',
  read:      'Study Tools',
  notes:     'Note Detail',
  bookmarks: 'Details',
  plans:     'Progress',
  search:    'Filters',
  settings:  'Options',
  profile:   'Account',
}
const panelTitle = computed(() => panelTitles[routeName.value] ?? 'Info')

/* ── Static data (replace with your store/API) ── */

</script>

<style scoped>
.right-panel {
  grid-column: 3;
  background: var(--white);
  border-left: 1px solid var(--border);
  display: flex;
  flex-direction: column;
  width: var(--right-panel-w);
  overflow: hidden;
}

.right-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  height: var(--header-h);
  border-bottom: 1px solid var(--border-light);
  flex-shrink: 0;
}
.right-title {
  font-size: 1rem;
  color: var(--ink);
  letter-spacing: 0.04em;
}
.close-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: var(--ink-light);
  font-size: 0.85rem;
  padding: 4px 8px;
  border-radius: 4px;
  transition: background var(--transition);
}
.close-btn:hover { background: var(--parchment-mid); color: var(--ink); }

.right-content {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 24px;
}






/* ── Tablet: fixed overlay ──────────────────── */
@media (max-width: 1024px) {
  .right-panel {
    position: fixed;
    right: 0; top: 0; bottom: 0;
    z-index: 100;
    box-shadow: -4px 0 20px rgba(44,31,14,0.15);
  }
}

/* ── Mobile: full-width overlay ─────────────── */
@media (max-width: 640px) {
  .right-panel {
    width: 100% !important;
    max-width: 100%;
  }
}
</style>