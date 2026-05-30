<template>
  <div class="section">

    <S type="group" label="Theme">
      <div class="theme-picker">
        <div
          v-for="t in themes"
          :key="t.id"
          class="theme-card"
          :class="{ selected: selectedTheme === t.id }"
          @click="selectedTheme = t.id"
        >
          <div class="theme-preview" :style="t.previewStyle">
            <div class="theme-preview-sidebar" :style="{ background: t.sidebarColor }"></div>
            <div class="theme-preview-lines">
              <div class="preview-line" :style="{ background: t.lineColor, width: '70%' }"></div>
              <div class="preview-line" :style="{ background: t.lineColor, width: '50%' }"></div>
              <div class="preview-line" :style="{ background: t.lineColor, width: '60%' }"></div>
            </div>
          </div>
          <span class="theme-name">{{ t.name }}</span>
          <span v-if="selectedTheme === t.id" class="theme-check">✓</span>
        </div>
      </div>
    </S>

    <S type="group" label="Accent Colour"
       hint="Used for active states, highlights and gold accents.">
      <S type="row" label="Accent colour">
        <div class="accent-picker">
          <div
            v-for="a in accents"
            :key="a.value"
            class="accent-swatch"
            :class="{ selected: selectedAccent === a.value }"
            :style="{ background: a.value }"
            :title="a.name"
            @click="selectedAccent = a.value"
          />
        </div>
      </S>
    </S>

    <S type="group" label="Interface Font">
      <S type="row" label="UI font" hint="Used for menus, labels and controls.">
        <select class="s-select" v-model="uiFont">
          <option value="Georgia">Georgia</option>
          <option value="system-ui">System default</option>
          <option value="'Helvetica Neue'">Helvetica Neue</option>
          <option value="'Times New Roman'">Times New Roman</option>
        </select>
      </S>
    </S>

    <S type="group" label="Window">
      <S type="row" label="Compact mode" hint="Reduces padding and spacing throughout the app.">
        <label class="s-toggle">
          <input type="checkbox" v-model="compactMode" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="Show reading progress bar"
         hint="Displays a thin progress bar at the top of the reading view.">
        <label class="s-toggle">
          <input type="checkbox" v-model="progressBar" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
    </S>

  </div>
</template>

<script setup>
import { ref } from 'vue'
import S from './SettingPrimitives.vue'

const selectedTheme  = ref('parchment')
const selectedAccent = ref('#c9973a')
const uiFont         = ref('Georgia')
const compactMode    = ref(false)
const progressBar    = ref(true)

const themes = [
  {
    id: 'parchment',
    name: 'Parchment',
    previewStyle: { background: '#f7f2e8' },
    sidebarColor: '#1c1610',
    lineColor: '#d9cfb8',
  },
  {
    id: 'dark',
    name: 'Dark',
    previewStyle: { background: '#1a1a1a' },
    sidebarColor: '#0d0d0d',
    lineColor: '#333',
  },
  {
    id: 'light',
    name: 'Light',
    previewStyle: { background: '#ffffff' },
    sidebarColor: '#f0f0f0',
    lineColor: '#e0e0e0',
  },
  {
    id: 'sepia',
    name: 'Sepia',
    previewStyle: { background: '#f5ebe0' },
    sidebarColor: '#2c1a0e',
    lineColor: '#d4b896',
  },
]

const accents = [
  { name: 'Gold',       value: '#c9973a' },
  { name: 'Amber',      value: '#d4850a' },
  { name: 'Crimson',    value: '#8b2e2e' },
  { name: 'Forest',     value: '#2d6a4f' },
  { name: 'Slate',      value: '#4a5568' },
  { name: 'Indigo',     value: '#4c51bf' },
]
</script>

<style scoped>
.section { display: flex; flex-direction: column; }

.theme-picker {
  display: flex;
  gap: 14px;
  padding: 16px;
  flex-wrap: wrap;
}
.theme-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  position: relative;
}
.theme-preview {
  width: 88px;
  height: 60px;
  border-radius: 8px;
  border: 2px solid transparent;
  overflow: hidden;
  display: flex;
  transition: border-color var(--transition), transform var(--transition);
}
.theme-card:hover .theme-preview { transform: translateY(-2px); }
.theme-card.selected .theme-preview { border-color: var(--gold); }

.theme-preview-sidebar { width: 22px; flex-shrink: 0; }
.theme-preview-lines {
  flex: 1;
  padding: 8px 6px;
  display: flex;
  flex-direction: column;
  gap: 5px;
  justify-content: center;
}
.preview-line { height: 4px; border-radius: 2px; }
.theme-name { font-size: 0.75rem; color: var(--ink-mid); }
.theme-check {
  position: absolute;
  top: 4px; right: 4px;
  font-size: 0.7rem;
  background: var(--gold);
  color: white;
  width: 16px; height: 16px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.accent-picker { display: flex; gap: 10px; align-items: center; }
.accent-swatch {
  width: 26px; height: 26px;
  border-radius: 50%;
  cursor: pointer;
  border: 2px solid transparent;
  transition: transform var(--transition), border-color var(--transition);
}
.accent-swatch:hover { transform: scale(1.15); }
.accent-swatch.selected { border-color: var(--ink); transform: scale(1.15); }
</style>