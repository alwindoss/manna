<template>
  <div class="section">

    <S type="group" label="Storage Overview">
      <div class="storage-overview">
        <div class="storage-bar">
          <div
            v-for="seg in storageSegments"
            :key="seg.label"
            class="storage-seg"
            :style="{ width: seg.pct + '%', background: seg.color }"
            :title="seg.label + ': ' + seg.size"
          />
        </div>
        <div class="storage-legend">
          <div v-for="seg in storageSegments" :key="seg.label" class="legend-item">
            <span class="legend-dot" :style="{ background: seg.color }"></span>
            <span class="legend-label">{{ seg.label }}</span>
            <span class="legend-size">{{ seg.size }}</span>
          </div>
        </div>
        <div class="storage-total">
          Total used: <strong>2.4 GB</strong> of available local storage
        </div>
      </div>
    </S>

    <S type="group" label="Bible Data">
      <S type="row" label="Cache chapter text on read"
         hint="Stores recently read chapters for instant offline access.">
        <label class="s-toggle">
          <input type="checkbox" v-model="cacheOnRead" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="Pre-cache entire translation"
         hint="Download all chapters of installed translations upfront.">
        <label class="s-toggle">
          <input type="checkbox" v-model="preCache" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="action" label="Clear Bible cache"
         hint="Re-downloads chapters on next read. Installed translations are kept."
         button-label="Clear cache"
         @click="clearBibleCache" />
    </S>

    <S type="group" label="Commentary & Cross-References">
      <S type="row" label="Store commentary offline">
        <label class="s-toggle">
          <input type="checkbox" v-model="offlineCommentary" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="Store cross-references offline">
        <label class="s-toggle">
          <input type="checkbox" v-model="offlineCrossRefs" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
    </S>

    <S type="group" label="Danger Zone">
      <S type="action" label="Clear all cached data"
         hint="Removes all cached content. Your notes, highlights and bookmarks are not affected."
         button-label="Clear all cache"
         :danger="true"
         @click="clearAll" />
      <S type="action" label="Reset app data"
         hint="Permanently deletes all local data including notes, highlights and bookmarks."
         button-label="Reset everything"
         :danger="true"
         @click="resetAll" />
    </S>

  </div>
</template>

<script setup>
import { ref } from 'vue'
import S from './SettingPrimitives.vue'

const cacheOnRead        = ref(true)
const preCache           = ref(false)
const offlineCommentary  = ref(true)
const offlineCrossRefs   = ref(true)

const storageSegments = [
  { label: 'Translations', size: '1.2 GB', pct: 50, color: '#c9973a' },
  { label: 'Commentary',   size: '0.6 GB', pct: 25, color: '#8b6f3e' },
  { label: 'Notes',        size: '0.3 GB', pct: 12, color: '#5a4230' },
  { label: 'Cache',        size: '0.3 GB', pct: 13, color: '#d9cfb8' },
]

function clearBibleCache() { console.log('Clear bible cache') }
function clearAll()        { console.log('Clear all cache') }
function resetAll()        { console.log('Reset app data') }
</script>

<style scoped>
.section { display: flex; flex-direction: column; }

.storage-overview { padding: 16px 18px; display: flex; flex-direction: column; gap: 14px; }

.storage-bar {
  height: 10px;
  border-radius: 5px;
  overflow: hidden;
  display: flex;
  background: var(--parchment-dark);
}
.storage-seg { height: 100%; transition: width 0.4s ease; }

.storage-legend { display: flex; flex-wrap: wrap; gap: 12px 20px; }
.legend-item { display: flex; align-items: center; gap: 6px; }
.legend-dot { width: 10px; height: 10px; border-radius: 50%; flex-shrink: 0; }
.legend-label { font-size: 0.78rem; color: var(--ink-mid); }
.legend-size { font-size: 0.72rem; color: var(--ink-light); font-family: sans-serif; }

.storage-total { font-size: 0.78rem; color: var(--ink-light); }
.storage-total strong { color: var(--ink); }
</style>