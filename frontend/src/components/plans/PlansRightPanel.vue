<template v-else>
    <aside class="right-panel">

        <div class="right-header">
            <h3 class="right-title">{{ ui.rightPanelContext.props.title ?? '' }}</h3>
            <button class="close-btn" @click="ui.closeRightPanel">✕</button>
        </div>

        <div class="right-content">
            <p class="empty-panel">No contextual panel for this section.</p>
        </div>
    </aside>
</template>
<script setup>
import { useUiStore } from '@/stores/ui'

const ui = useUiStore()

</script>


<style scoped>
.right-panel {
  grid-column: 3;
  background: var(--white);
  border-left: 1px solid var(--border);
  display: flex;
  flex-direction: column;
  width: var(--right-panel-w);
  min-height: 0;
  height: 100%;
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

.close-btn:hover {
  background: var(--parchment-mid);
  color: var(--ink);
}

.right-content {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 24px;
}

/* ── Notes: note editor ─────────────────────── */
.note-editor { display: flex; flex-direction: column; gap: 12px; }
.editor-ref {
  font-size: 0.75rem;
  color: var(--gold);
  font-weight: 700;
  font-family: sans-serif;
  letter-spacing: 0.08em;
}
.editor-title { font-size: 1.05rem; font-weight: 700; color: var(--ink); }
.editor-body  { font-size: 0.88rem; line-height: 1.75; color: var(--ink-mid); white-space: pre-wrap; }
.editor-tags  { display: flex; gap: 6px; flex-wrap: wrap; }
.tag {
  background: var(--parchment-mid);
  border: 1px solid var(--border);
  border-radius: 20px;
  padding: 3px 10px;
  font-size: 0.72rem;
  color: var(--ink-mid);
  font-family: sans-serif;
}

.no-selection,
.empty-panel { color: var(--ink-light); font-size: 0.88rem; font-style: italic; }






/* ── Tablet: fixed overlay ──────────────────── */
@media (max-width: 1024px) {
  .right-panel {
    position: fixed;
    right: 0;
    top: 0;
    bottom: 0;
    z-index: 100;
    box-shadow: -4px 0 20px rgba(44, 31, 14, 0.15);
  }
}

/* ── Mobile: full-width overlay ─────────────── */
@media (max-width: 640px) {
  .right-panel {
    width: 100% !important;
    max-width: 100%;
  }
}
/* ── Notes: note editor ─────────────────────── */
.note-editor {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.editor-ref {
  font-size: 0.75rem;
  color: var(--gold);
  font-weight: 700;
  font-family: sans-serif;
  letter-spacing: 0.08em;
}

.editor-title {
  font-size: 1.05rem;
  font-weight: 700;
  color: var(--ink);
}

.editor-body {
  font-size: 0.88rem;
  line-height: 1.75;
  color: var(--ink-mid);
  white-space: pre-wrap;
}

.editor-tags {
  display: flex;
  gap: 6px;
  flex-wrap: wrap;
}

.tag {
  background: var(--parchment-mid);
  border: 1px solid var(--border);
  border-radius: 20px;
  padding: 3px 10px;
  font-size: 0.72rem;
  color: var(--ink-mid);
  font-family: sans-serif;
}

.no-selection,
.empty-panel {
  color: var(--ink-light);
  font-size: 0.88rem;
  font-style: italic;
}
</style>