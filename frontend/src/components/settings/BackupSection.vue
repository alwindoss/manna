<template>
  <div class="section">

    <S type="group" label="Export">
      <S type="row" label="Export format">
        <select class="s-select" v-model="exportFormat">
          <option value="json">JSON (full backup)</option>
          <option value="csv">CSV (spreadsheet)</option>
          <option value="markdown">Markdown (notes only)</option>
        </select>
      </S>
      <S type="row" label="Include in export">
        <div class="checkbox-list">
          <label v-for="item in exportItems" :key="item.id" class="check-label">
            <input type="checkbox" v-model="item.selected" />
            {{ item.label }}
          </label>
        </div>
      </S>
      <S type="action" label="Export now"
         hint="A file will be saved to your Downloads folder."
         button-label="Export"
         @click="exportNow" />
    </S>

    <S type="group" label="Restore">
      <S type="action" label="Restore from file"
         hint="Import a previously exported Manna backup file (.json)."
         button-label="Choose file…"
         @click="restoreFromFile" />
      <div class="restore-warning">
        ⚠ Restoring will merge imported data with your current data.
        Conflicts will keep the most recent version.
      </div>
    </S>

    <S type="group" label="Automatic Backups">
      <S type="row" label="Auto-backup to Hub"
         hint="Requires a Manna Hub account. Encrypted at rest.">
        <label class="s-toggle">
          <input type="checkbox" v-model="autoBackup" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="Backup frequency" :disabled="!autoBackup">
        <select class="s-select" v-model="backupFreq" :disabled="!autoBackup">
          <option value="daily">Daily</option>
          <option value="weekly">Weekly</option>
          <option value="monthly">Monthly</option>
        </select>
      </S>
      <S type="row" label="Keep last N backups" :disabled="!autoBackup">
        <select class="s-select" v-model="keepBackups" :disabled="!autoBackup">
          <option value="3">3</option>
          <option value="5">5</option>
          <option value="10">10</option>
        </select>
      </S>
    </S>

  </div>
</template>
<script setup>
import { ref } from 'vue'
import S from './SettingPrimitives.vue'
const exportFormat = ref('json')
const autoBackup   = ref(false)
const backupFreq   = ref('weekly')
const keepBackups  = ref('5')
const exportItems  = ref([
  { id: 'notes',      label: 'Notes',           selected: true  },
  { id: 'highlights', label: 'Highlights',      selected: true  },
  { id: 'bookmarks',  label: 'Bookmarks',       selected: true  },
  { id: 'plans',      label: 'Reading plans',   selected: true  },
  { id: 'history',    label: 'Reading history', selected: false },
  { id: 'settings',   label: 'App settings',    selected: false },
])
function exportNow()       { console.log('Export:', exportFormat.value) }
function restoreFromFile() { console.log('Restore from file') }
</script>
<style scoped>
.section { display: flex; flex-direction: column; }
.checkbox-list { display: flex; flex-direction: column; gap: 6px; }
.check-label {
  display: flex; align-items: center; gap: 8px;
  font-size: 0.85rem; color: var(--ink-mid); cursor: pointer;
}
.check-label input { accent-color: var(--gold); cursor: pointer; }
.restore-warning {
  margin: 0 18px 12px;
  font-size: 0.78rem;
  color: #8b6020;
  background: #f5e6c0;
  border-radius: 6px;
  padding: 10px 14px;
  line-height: 1.5;
}
</style>