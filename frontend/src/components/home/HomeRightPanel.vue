<template>
  <aside class="right-panel">

    <div class="right-header">
      <h3 class="right-title">{{ ui.rightPanelContext.props.title ?? '' }}</h3>
      <button class="close-btn" @click="ui.closeRightPanel">✕</button>
    </div>

    <div class="right-content">
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
    </div>
  </aside>
</template>
<script setup>
import { useUiStore } from '@/stores/ui'

const ui = useUiStore()

const todaysPlan = [
  { ref: 'Genesis 3–4', done: true },
  { ref: 'Psalm 23', done: true },
  { ref: 'Matthew 5:1–12', done: false },
]


const quickLinks = ['Prayer Journal', 'Concordance', 'Bible Maps', 'Lexicon']


</script>
<style lang="css" scoped>
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
.rp-section {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

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
.plan-item {
  display: flex;
  gap: 10px;
  align-items: center;
  font-size: 0.88rem;
}

.plan-check {
  color: var(--parchment-dark);
  font-size: 0.9rem;
}

.plan-check.done {
  color: var(--gold);
}

.plan-ref {
  color: var(--ink-mid);
}

/* ── Streak ─────────────────────────────────── */
.streak-display {
  display: flex;
  align-items: baseline;
  gap: 8px;
}

.streak-num {
  font-size: 2.4rem;
  font-weight: 700;
  color: var(--gold);
  line-height: 1;
}

.streak-label {
  font-size: 0.85rem;
  color: var(--ink-light);
}

/* ── Quick links ────────────────────────────── */
.quick-link {
  font-size: 0.85rem;
  color: var(--ink-mid);
  padding: 7px 0;
  border-bottom: 1px solid var(--border-light);
  cursor: pointer;
  transition: color var(--transition);
}

.quick-link:hover {
  color: var(--gold);
}

.quick-link:last-child {
  border-bottom: none;
}
</style>