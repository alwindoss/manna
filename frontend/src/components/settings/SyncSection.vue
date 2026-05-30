<!-- ============================================================
     SyncSection.vue
============================================================ -->
<template>
  <div class="section">
    <S type="group" label="Manna Hub Connection">
      <div class="hub-status" :class="isConnected ? 'connected' : 'disconnected'">
        <span class="status-dot"></span>
        <div class="status-info">
          <span class="status-title">{{ isConnected ? 'Connected to Manna Hub' : 'Not connected' }}</span>
          <span class="status-sub">{{ isConnected ? 'Syncing as james@example.com' : 'Sign in to sync your data across devices.' }}</span>
        </div>
        <button class="btn-hub" @click="isConnected = !isConnected">
          {{ isConnected ? 'Disconnect' : 'Sign in' }}
        </button>
      </div>
    </S>

    <template v-if="isConnected">
      <S type="group" label="What to Sync">
        <S type="row" v-for="item in syncItems" :key="item.id"
           :label="item.label" :hint="item.hint">
          <label class="s-toggle">
            <input type="checkbox" v-model="item.enabled" />
            <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
          </label>
        </S>
      </S>

      <S type="group" label="Sync Behaviour">
        <S type="row" label="Auto-sync on app open">
          <label class="s-toggle">
            <input type="checkbox" v-model="autoSync" />
            <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
          </label>
        </S>
        <S type="row" label="Sync over Wi-Fi only">
          <label class="s-toggle">
            <input type="checkbox" v-model="wifiOnly" />
            <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
          </label>
        </S>
        <S type="action" label="Sync now"
           hint="Last synced: 2 minutes ago"
           button-label="Sync now"
           @click="syncNow" />
      </S>

      <S type="group" label="Connected Devices">
        <div v-for="d in devices" :key="d.id" class="device-row">
          <span class="device-icon">{{ d.icon }}</span>
          <div class="device-info">
            <span class="device-name">{{ d.name }}</span>
            <span class="device-last">{{ d.last }}</span>
          </div>
          <span v-if="d.current" class="device-current">This device</span>
          <button v-else class="btn-icon btn-remove" @click="removeDevice(d.id)">✕</button>
        </div>
      </S>
    </template>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import S from './SettingPrimitives.vue'
const isConnected = ref(true)
const autoSync    = ref(true)
const wifiOnly    = ref(false)
const syncItems   = ref([
  { id: 'notes',     label: 'Notes',           hint: 'All your study notes.',              enabled: true  },
  { id: 'highlights',label: 'Highlights',      hint: 'Verse highlights and colours.',      enabled: true  },
  { id: 'bookmarks', label: 'Bookmarks',       hint: 'Saved verse bookmarks.',             enabled: true  },
  { id: 'plans',     label: 'Reading plans',   hint: 'Plans and daily progress.',          enabled: true  },
  { id: 'history',   label: 'Reading history', hint: 'Used for streak and analytics.',     enabled: false },
])
const devices = ref([
  { id: 1, icon: '💻', name: "James's MacBook",    last: 'Active now',     current: true  },
  { id: 2, icon: '🖥',  name: 'Home Desktop',       last: 'Last seen 1h ago', current: false },
  { id: 3, icon: '📱', name: 'iPhone 15 Pro',      last: 'Last seen 2d ago', current: false },
])
function syncNow()        { console.log('Sync now') }
function removeDevice(id) { devices.value = devices.value.filter(d => d.id !== id) }
</script>
<style scoped>
.section { display: flex; flex-direction: column; }
.hub-status {
  display: flex; align-items: center; gap: 14px; padding: 16px 18px;
}
.status-dot { width: 10px; height: 10px; border-radius: 50%; flex-shrink: 0; }
.connected .status-dot    { background: #2d6a4f; box-shadow: 0 0 0 3px rgba(45,106,79,0.2); }
.disconnected .status-dot { background: var(--ink-light); }
.status-info { flex: 1; display: flex; flex-direction: column; gap: 2px; }
.status-title { font-size: 0.88rem; color: var(--ink); font-weight: 600; }
.status-sub   { font-size: 0.75rem; color: var(--ink-light); }
.btn-hub {
  padding: 6px 14px; border-radius: 7px; border: 1px solid var(--border);
  background: var(--parchment-mid); font-family: inherit; font-size: 0.82rem;
  cursor: pointer; color: var(--ink); transition: background var(--transition);
}
.btn-hub:hover { background: var(--parchment-dark); }
.device-row {
  display: flex; align-items: center; gap: 12px; padding: 11px 18px;
}
.device-row:not(:last-child) { border-bottom: 1px solid var(--border-light); }
.device-icon { font-size: 1.2rem; }
.device-info { flex: 1; display: flex; flex-direction: column; gap: 2px; }
.device-name { font-size: 0.88rem; color: var(--ink); }
.device-last { font-size: 0.72rem; color: var(--ink-light); font-family: sans-serif; }
.device-current { font-size: 0.72rem; color: var(--gold); font-family: sans-serif; padding: 2px 8px; background: rgba(201,151,58,0.1); border-radius: 10px; }
.btn-icon { background: none; border: none; cursor: pointer; font-size: 0.8rem; padding: 3px 6px; border-radius: 4px; }
.btn-remove { color: var(--ink-light); }
.btn-remove:hover { background: #fde8e8; color: #8b2e2e; }
</style>