<template>
  <div class="section">

    <S type="group" label="Default Translation">
      <S type="row" label="Default Bible version"
         hint="Used throughout the app unless overridden per reading plan.">
        <select class="s-select" v-model="defaultTranslation">
          <option v-for="t in installed" :key="t.code" :value="t.code">{{ t.name }}</option>
        </select>
      </S>
    </S>

    <S type="group" label="Installed Translations">
      <div
        v-for="t in installed"
        :key="t.code"
        class="translation-row"
      >
        <div class="trans-info">
          <span class="trans-code">{{ t.code }}</span>
          <span class="trans-name">{{ t.name }}</span>
          <span class="trans-lang">{{ t.language }}</span>
        </div>
        <div class="trans-actions">
          <span v-if="t.code === defaultTranslation" class="badge-default">Default</span>
          <span class="trans-size">{{ t.size }}</span>
          <button class="btn-icon btn-remove" @click="removeTranslation(t.code)" title="Remove">✕</button>
        </div>
      </div>
    </S>

    <S type="group" label="Available to Download"
       hint="Tap a translation to download it for offline use.">
      <div
        v-for="t in available"
        :key="t.code"
        class="translation-row"
      >
        <div class="trans-info">
          <span class="trans-code">{{ t.code }}</span>
          <span class="trans-name">{{ t.name }}</span>
          <span class="trans-lang">{{ t.language }}</span>
        </div>
        <div class="trans-actions">
          <span class="trans-size">{{ t.size }}</span>
          <button
            v-if="t.status === 'idle'"
            class="btn-download"
            @click="download(t)"
          >↓ Download</button>
          <div v-else-if="t.status === 'downloading'" class="download-progress">
            <div class="progress-bar">
              <div class="progress-fill" :style="{ width: t.progress + '%' }"></div>
            </div>
            <span class="progress-pct">{{ t.progress }}%</span>
          </div>
          <span v-else-if="t.status === 'done'" class="badge-done">✓ Done</span>
        </div>
      </div>
    </S>

    <S type="group" label="Add Custom Source"
       hint="Provide a URL to a USFM or SQLite Bible database.">
      <S type="row" label="Source URL">
        <input class="s-input" style="width:240px" v-model="customUrl" placeholder="https://…" />
      </S>
      <S type="action" label="Import from URL"
         hint="The file will be validated before import."
         button-label="Import"
         @click="importCustom" />
    </S>

  </div>
</template>

<script setup>
import { ref } from 'vue'
import S from './SettingPrimitives.vue'

const defaultTranslation = ref('KJV')
const customUrl = ref('')

const installed = ref([
  { code: 'KJV',  name: 'King James Version',          language: 'English', size: '4.2 MB' },
  { code: 'NIV',  name: 'New International Version',   language: 'English', size: '4.8 MB' },
])

const available = ref([
  { code: 'ESV',  name: 'English Standard Version',    language: 'English', size: '4.6 MB', status: 'idle', progress: 0 },
  { code: 'NLT',  name: 'New Living Translation',      language: 'English', size: '4.9 MB', status: 'idle', progress: 0 },
  { code: 'NASB', name: 'New American Standard Bible', language: 'English', size: '4.7 MB', status: 'idle', progress: 0 },
  { code: 'MSG',  name: 'The Message',                 language: 'English', size: '5.1 MB', status: 'idle', progress: 0 },
  { code: 'AMP',  name: 'Amplified Bible',             language: 'English', size: '5.4 MB', status: 'idle', progress: 0 },
  { code: 'LSG',  name: 'Louis Segond',                language: 'French',  size: '4.3 MB', status: 'idle', progress: 0 },
])

function download(t) {
  t.status = 'downloading'
  // Simulate progress — replace with real Wails service call
  const interval = setInterval(() => {
    t.progress += Math.floor(Math.random() * 15) + 5
    if (t.progress >= 100) {
      t.progress = 100
      t.status   = 'done'
      clearInterval(interval)
      setTimeout(() => {
        installed.value.push({ code: t.code, name: t.name, language: t.language, size: t.size })
        available.value = available.value.filter(a => a.code !== t.code)
      }, 800)
    }
  }, 300)
}

function removeTranslation(code) {
  if (code === defaultTranslation.value) return
  installed.value = installed.value.filter(t => t.code !== code)
}

function importCustom() {
  if (!customUrl.value.trim()) return
  console.log('Import from URL:', customUrl.value)
  customUrl.value = ''
}
</script>

<style scoped>
.section { display: flex; flex-direction: column; }

.translation-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  padding: 11px 18px;
}
.translation-row:not(:last-child) { border-bottom: 1px solid var(--border-light); }
.translation-row:hover { background: var(--parchment); }

.trans-info { display: flex; align-items: center; gap: 10px; flex: 1; min-width: 0; }
.trans-code {
  font-size: 0.75rem;
  font-family: sans-serif;
  font-weight: 700;
  color: var(--gold);
  min-width: 40px;
}
.trans-name { font-size: 0.875rem; color: var(--ink); flex: 1; }
.trans-lang { font-size: 0.72rem; color: var(--ink-light); font-style: italic; }

.trans-actions { display: flex; align-items: center; gap: 10px; flex-shrink: 0; }
.trans-size { font-size: 0.72rem; color: var(--ink-light); font-family: sans-serif; }

.badge-default {
  font-size: 0.68rem;
  background: rgba(201,151,58,0.15);
  color: var(--gold);
  border: 1px solid rgba(201,151,58,0.3);
  border-radius: 10px;
  padding: 2px 8px;
  font-family: sans-serif;
}
.badge-done {
  font-size: 0.72rem;
  color: #2d6a4f;
  background: #d4edda;
  border-radius: 10px;
  padding: 2px 8px;
  font-family: sans-serif;
}

.btn-download {
  font-size: 0.78rem;
  padding: 5px 12px;
  background: var(--ink);
  color: var(--parchment);
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-family: inherit;
  transition: background var(--transition);
}
.btn-download:hover { background: var(--ink-mid); }

.btn-icon {
  background: none;
  border: none;
  cursor: pointer;
  font-size: 0.8rem;
  padding: 3px 6px;
  border-radius: 4px;
  transition: background var(--transition);
}
.btn-remove { color: var(--ink-light); }
.btn-remove:hover { background: #fde8e8; color: #8b2e2e; }

.download-progress { display: flex; align-items: center; gap: 8px; }
.progress-bar { width: 80px; height: 5px; background: var(--parchment-dark); border-radius: 3px; overflow: hidden; }
.progress-fill { height: 100%; background: var(--gold); border-radius: 3px; transition: width 0.2s ease; }
.progress-pct { font-size: 0.7rem; color: var(--ink-light); font-family: sans-serif; min-width: 30px; }
</style>