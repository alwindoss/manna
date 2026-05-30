<template>
  <!-- SettingRow: a single labelled row with a control slot -->
  <div v-if="type === 'row'" class="setting-row" :class="{ danger: danger, disabled: disabled }">
    <div class="row-label-group">
      <label class="row-label">{{ label }}</label>
      <p v-if="hint" class="row-hint">{{ hint }}</p>
    </div>
    <div class="row-control">
      <slot />
    </div>
  </div>

  <!-- SettingGroup: a titled card wrapping multiple rows -->
  <div v-else-if="type === 'group'" class="setting-group">
    <div v-if="label" class="group-header">
      <h3 class="group-title">{{ label }}</h3>
      <p v-if="hint" class="group-subtitle">{{ hint }}</p>
    </div>
    <div class="group-body">
      <slot />
    </div>
  </div>

  <!-- SettingDivider -->
  <div v-else-if="type === 'divider'" class="setting-divider" />

  <!-- SettingAction: a full-width button row -->
  <div v-else-if="type === 'action'" class="setting-action" :class="{ danger: danger }">
    <div class="row-label-group">
      <span class="row-label">{{ label }}</span>
      <p v-if="hint" class="row-hint">{{ hint }}</p>
    </div>
    <button class="action-btn" :class="danger ? 'btn-danger' : 'btn-default'" @click="$emit('click')">
      {{ buttonLabel }}
    </button>
  </div>
</template>

<script setup>
defineProps({
  type:        { type: String, default: 'row' },  // 'row' | 'group' | 'divider' | 'action'
  label:       { type: String, default: '' },
  hint:        { type: String, default: '' },
  buttonLabel: { type: String, default: 'Action' },
  danger:      { type: Boolean, default: false },
  disabled:    { type: Boolean, default: false },
})
defineEmits(['click'])
</script>

<style scoped>
/* ── Group ────────────────────────────────── */
.setting-group {
  background: var(--white);
  border: 1px solid var(--border-light);
  border-radius: var(--radius);
  overflow: hidden;
  margin-bottom: 20px;
}
.group-header {
  padding: 14px 18px 10px;
  border-bottom: 1px solid var(--border-light);
}
.group-title {
  font-size: 0.82rem;
  font-weight: 700;
  letter-spacing: 0.06em;
  text-transform: uppercase;
  color: var(--ink-light);
  font-family: sans-serif;
}
.group-subtitle { font-size: 0.78rem; color: var(--ink-light); margin-top: 2px; font-style: italic; }
.group-body > *:not(:last-child) { border-bottom: 1px solid var(--border-light); }

/* ── Row ──────────────────────────────────── */
.setting-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 13px 18px;
  transition: background var(--transition);
}
.setting-row:hover { background: var(--parchment); }
.setting-row.disabled { opacity: 0.45; pointer-events: none; }
.setting-row.danger .row-label { color: #8b2e2e; }

.row-label-group { flex: 1; min-width: 0; }
.row-label { font-size: 0.88rem; color: var(--ink); display: block; }
.row-hint  { font-size: 0.75rem; color: var(--ink-light); margin-top: 2px; font-style: italic; }
.row-control { flex-shrink: 0; }

/* ── Action row ───────────────────────────── */
.setting-action {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 13px 18px;
}
.action-btn {
  border: none;
  border-radius: 7px;
  padding: 7px 16px;
  font-family: inherit;
  font-size: 0.82rem;
  cursor: pointer;
  transition: background var(--transition), opacity var(--transition);
  white-space: nowrap;
}
.btn-default {
  background: var(--parchment-mid);
  color: var(--ink);
  border: 1px solid var(--border);
}
.btn-default:hover { background: var(--parchment-dark); }
.btn-danger  { background: #8b2e2e; color: white; }
.btn-danger:hover { background: #6e2222; }

/* ── Divider ──────────────────────────────── */
.setting-divider {
  height: 1px;
  background: var(--border-light);
  margin: 8px 0;
}

/* ── Shared controls ──────────────────────── */
:global(.s-toggle) {
  position: relative;
  display: inline-flex;
  width: 40px;
  height: 22px;
  cursor: pointer;
}
:global(.s-toggle input) { opacity: 0; width: 0; height: 0; }
:global(.s-toggle-track) {
  position: absolute;
  inset: 0;
  background: var(--parchment-dark);
  border-radius: 11px;
  transition: background var(--transition);
}
:global(.s-toggle input:checked ~ .s-toggle-track) { background: var(--gold); }
:global(.s-toggle-thumb) {
  position: absolute;
  top: 3px; left: 3px;
  width: 16px; height: 16px;
  background: white;
  border-radius: 50%;
  transition: transform var(--transition);
  box-shadow: 0 1px 3px rgba(0,0,0,0.2);
}
:global(.s-toggle input:checked ~ .s-toggle-track .s-toggle-thumb) {
  transform: translateX(18px);
}

:global(.s-select) {
  appearance: none;
  background: var(--parchment-mid);
  border: 1px solid var(--border);
  border-radius: 7px;
  padding: 6px 28px 6px 10px;
  font-family: inherit;
  font-size: 0.82rem;
  color: var(--ink);
  cursor: pointer;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='10' height='6'%3E%3Cpath d='M0 0l5 6 5-6z' fill='%238b7355'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 10px center;
}
:global(.s-select:focus) { outline: 2px solid var(--gold); outline-offset: 1px; }

:global(.s-input) {
  background: var(--parchment-mid);
  border: 1px solid var(--border);
  border-radius: 7px;
  padding: 6px 10px;
  font-family: inherit;
  font-size: 0.82rem;
  color: var(--ink);
  width: 100%;
}
:global(.s-input:focus) { outline: 2px solid var(--gold); outline-offset: 1px; }
</style>