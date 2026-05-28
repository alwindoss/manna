<template>
  <aside class="right-panel">

    <div class="right-header">
      <h3 class="right-title">{{ panelTitle }}</h3>
      <button class="close-btn" @click="ui.closeRightPanel">✕</button>
    </div>

    <div class="right-content">

      <!-- ── HOME panel ─────────────────────── -->
      <template v-if="routeName === 'home'">
        <div class="rp-section">
          <div class="rp-section-label">Today's Plan</div>
          <div class="plan-item" v-for="p in todaysPlan" :key="p.ref">
            <span class="plan-check" :class="{ done: p.done }">{{ p.done ? '✓' : '○' }}</span>
            <span class="plan-ref">{{ p.ref }}</span>
          </div>
        </div>

        <div class="rp-section">
          <div class="rp-section-label">Streak</div>
          <div class="streak-display">
            <span class="streak-num">7</span>
            <span class="streak-label">days in a row 🔥</span>
          </div>
        </div>

        <div class="rp-section">
          <div class="rp-section-label">Quick Links</div>
          <div class="quick-link" v-for="q in quickLinks" :key="q">{{ q }}</div>
        </div>
      </template>

      <!-- ── READ panel ─────────────────────── -->
      <template v-else-if="routeName === 'read'">
        <div class="rp-section">
          <div class="rp-section-label">Cross-References</div>
          <div class="cross-ref" v-for="ref in crossRefs" :key="ref.ref">
            <span class="cross-ref-tag">{{ ref.ref }}</span>
            <span class="cross-ref-text">{{ ref.snippet }}</span>
          </div>
        </div>

        <div class="rp-section">
          <div class="rp-section-label">Commentary</div>
          <p class="commentary-text">
            This passage is foundational to understanding the covenant narrative that runs
            through all of Scripture. Scholars note the parallel structure between the
            creation account and the tabernacle instructions.
          </p>
        </div>

        <div class="rp-section">
          <div class="rp-section-label">Original Language</div>
          <div class="orig-word" v-for="w in origWords" :key="w.word">
            <span class="script">{{ w.word }}</span>
            <span class="trans">{{ w.meaning }}</span>
          </div>
        </div>
      </template>

      <!-- ── NOTES panel ────────────────────── -->
      <template v-else-if="routeName === 'notes'">
        <div v-if="ui.selectedNote" class="note-editor">
          <div class="editor-ref">{{ ui.selectedNote.ref }}</div>
          <div class="editor-title">{{ ui.selectedNote.title }}</div>
          <div class="editor-body">{{ ui.selectedNote.body }}</div>
          <div class="editor-tags">
            <span class="tag" v-for="t in ui.selectedNote.tags" :key="t">{{ t }}</span>
          </div>
        </div>
        <div v-else class="no-selection">
          <p>Select a note to view its full content here.</p>
        </div>
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
const todaysPlan = [
  { ref: 'Genesis 3–4',    done: true  },
  { ref: 'Psalm 23',       done: true  },
  { ref: 'Matthew 5:1–12', done: false },
]
const quickLinks = ['Prayer Journal', 'Concordance', 'Bible Maps', 'Lexicon']

const crossRefs = [
  { ref: 'John 1:1',  snippet: '"In the beginning was the Word…"' },
  { ref: 'Heb 11:3',  snippet: 'By faith we understand the universe was created…' },
  { ref: 'Col 1:16',  snippet: 'For in him all things were created…' },
]
const origWords = [
  { word: 'בְּרֵאשִׁית', meaning: "b'reshit — In the beginning" },
  { word: 'בָּרָא',       meaning: 'bara — created (divine act)' },
  { word: 'אֱלֹהִים',    meaning: 'Elohim — God (plural majesty)' },
]
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

/* ── Section ────────────────────────────────── */
.rp-section { display: flex; flex-direction: column; gap: 10px; }
.rp-section-label {
  font-size: 0.68rem;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: var(--ink-light);
  font-family: sans-serif;
  font-weight: 700;
  border-bottom: 1px solid var(--border-light);
  padding-bottom: 6px;
}

/* ── Today's plan ───────────────────────────── */
.plan-item { display: flex; gap: 10px; align-items: center; font-size: 0.88rem; }
.plan-check { color: var(--parchment-dark); font-size: 0.9rem; }
.plan-check.done { color: var(--gold); }
.plan-ref { color: var(--ink-mid); }

/* ── Streak ─────────────────────────────────── */
.streak-display { display: flex; align-items: baseline; gap: 8px; }
.streak-num { font-size: 2.4rem; font-weight: 700; color: var(--gold); line-height: 1; }
.streak-label { font-size: 0.85rem; color: var(--ink-light); }

/* ── Quick links ────────────────────────────── */
.quick-link {
  font-size: 0.85rem;
  color: var(--ink-mid);
  padding: 7px 0;
  border-bottom: 1px solid var(--border-light);
  cursor: pointer;
  transition: color var(--transition);
}
.quick-link:hover { color: var(--gold); }
.quick-link:last-child { border-bottom: none; }

/* ── Cross-refs ─────────────────────────────── */
.cross-ref {
  display: flex;
  flex-direction: column;
  gap: 3px;
  padding: 8px 0;
  border-bottom: 1px solid var(--border-light);
}
.cross-ref:last-child { border-bottom: none; }
.cross-ref-tag { font-size: 0.75rem; font-weight: 700; color: var(--gold); font-family: sans-serif; }
.cross-ref-text { font-size: 0.82rem; color: var(--ink-light); font-style: italic; }

/* ── Commentary ─────────────────────────────── */
.commentary-text { font-size: 0.88rem; line-height: 1.7; color: var(--ink-mid); }

/* ── Original language ──────────────────────── */
.orig-word { display: flex; gap: 14px; align-items: baseline; padding: 6px 0; }
.script { font-size: 1rem; color: var(--ink); }
.trans  { font-size: 0.8rem; color: var(--ink-light); font-family: sans-serif; font-style: italic; }

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