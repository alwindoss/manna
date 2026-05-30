<template>
  <div class="settings-shell">

    <!-- ── Left category nav ─────────────────── -->
    <nav class="settings-nav">
      <div class="nav-search">
        <span class="search-icon">⌕</span>
        <input v-model="searchQuery" class="nav-search-input" placeholder="Search settings…" @input="onSearch" />
      </div>

      <div class="nav-groups">
        <div v-for="group in filteredGroups" :key="group.id" class="nav-group">
          <div class="nav-group-label">{{ group.label }}</div>
          <button v-for="item in group.items" :key="item.id" class="nav-item"
            :class="{ active: activeSection === item.id }" @click="navigate(item.id)">
            <span class="nav-item-icon">{{ item.icon }}</span>
            <span class="nav-item-label">{{ item.label }}</span>
            <span v-if="item.badge" class="nav-item-badge" :class="item.badgeType">
              {{ item.badge }}
            </span>
          </button>
        </div>
      </div>

      <div class="nav-version">
        <span class="version-text">Manna v0.1.0</span>
      </div>
    </nav>

    <!-- ── Right content panel ───────────────── -->
    <div class="settings-content">
      <div class="content-header">
        <div class="header-breadcrumb">
          <span class="breadcrumb-root">Settings</span>
          <span class="breadcrumb-sep">›</span>
          <span class="breadcrumb-current">{{ currentSection?.label }}</span>
        </div>
        <p class="content-description">{{ currentSection?.description }}</p>
      </div>

      <div class="content-body">
        <Transition name="section-fade" mode="out-in">
          <component :is="currentSection?.component" :key="activeSection" />
        </Transition>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'

import BiblesSection from '@/components/settings/BiblesSection.vue'
import AppearanceSection from '@/components/settings/AppearanceSection.vue'
import ReadingSection from '@/components/settings/ReadingSection.vue'
import CrossRefsSection from '@/components/settings/CrossRefsSection.vue'
import OfflineSection from '@/components/settings/OfflineSection.vue'
import SyncSection from '@/components/settings/SyncSection.vue'
import AccountSection from '@/components/settings/AccountSection.vue'
import NotificationsSection from '@/components/settings/NotificationsSection.vue'
import BackupSection from '@/components/settings/BackupSection.vue'
import LanguageSection from '@/components/settings/LanguageSection.vue'

import { useUiStore } from '@/stores/ui'

const ui = useUiStore()

// ── Settings registry ─────────────────────────
// To add a new section in future: add one entry here + create the component.
const groups = [
  {
    id: 'bible',
    label: 'Bible',
    items: [
      {
        id: 'bibles',
        label: 'Translations',
        icon: '📖',
        description: 'Download, manage and set your default Bible translations.',
        component: BiblesSection,
      },
      {
        id: 'reading',
        label: 'Reading Display',
        icon: '✦',
        description: 'Font size, verse formatting, red letter and layout options.',
        component: ReadingSection,
      },
      {
        id: 'crossrefs',
        label: 'Cross References',
        icon: '⇌',
        description: 'Configure cross-reference sources and study tools.',
        component: CrossRefsSection,
      },
    ],
  },
  {
    id: 'app',
    label: 'Application',
    items: [
      {
        id: 'appearance',
        label: 'Appearance',
        icon: '◑',
        description: 'Theme, colours and typography preferences.',
        component: AppearanceSection,
      },
      {
        id: 'language',
        label: 'Language & Region',
        icon: '⊕',
        description: 'Interface language, date format and region settings.',
        component: LanguageSection,
      },
      {
        id: 'notifications',
        label: 'Notifications',
        icon: '◎',
        description: 'Daily reminders, reading goals and Hub activity alerts.',
        component: NotificationsSection,
      },
    ],
  },
  {
    id: 'data',
    label: 'Data',
    items: [
      {
        id: 'offline',
        label: 'Offline & Storage',
        icon: '⬡',
        description: 'Manage locally cached data, downloaded content and storage usage.',
        component: OfflineSection,
        badge: '2.4 GB',
        badgeType: 'neutral',
      },
      {
        id: 'sync',
        label: 'Sync & Manna Hub',
        icon: '⟳',
        description: 'Sync your data to Manna Hub and manage connected devices.',
        component: SyncSection,
      },
      {
        id: 'backup',
        label: 'Backup & Restore',
        icon: '⊞',
        description: 'Export your notes, highlights and bookmarks or restore from a backup.',
        component: BackupSection,
      },
    ],
  },
  {
    id: 'account',
    label: 'Account',
    items: [
      {
        id: 'account',
        label: 'Account & Profile',
        icon: '◉',
        description: 'Your Manna Hub account, display name, avatar and bio.',
        component: AccountSection,
      },
    ],
  },
]

const activeSection = ref('bibles')
const searchQuery = ref('')

const allItems = computed(() => groups.flatMap(g => g.items))

const currentSection = computed(() =>
  allItems.value.find(i => i.id === activeSection.value)
)

const filteredGroups = computed(() => {
  const q = searchQuery.value.toLowerCase().trim()
  if (!q) return groups
  return groups
    .map(g => ({
      ...g,
      items: g.items.filter(i => i.label.toLowerCase().includes(q)),
    }))
    .filter(g => g.items.length > 0)
})

function navigate(id) {
  activeSection.value = id
  searchQuery.value = ''
}

function onSearch() {
  const q = searchQuery.value.toLowerCase().trim()
  if (!q) return
  const match = allItems.value.find(i => i.label.toLowerCase().includes(q))
  if (match) activeSection.value = match.id
}

onMounted(() => {
  ui.setRightPanelDisabled(true)
})
onUnmounted(() => {
  ui.setRightPanelDisabled(false)
})
</script>

<style scoped>
.settings-shell {
  display: grid;
  grid-template-columns: 220px 1fr;
  height: 100%;
  overflow: hidden;
  background: var(--parchment);
}

/* ── Nav ──────────────────────────────────── */
.settings-nav {
  background: var(--white);
  border-right: 1px solid var(--border);
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.nav-search {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 14px 14px 10px;
  border-bottom: 1px solid var(--border-light);
}

.search-icon {
  color: var(--ink-light);
  font-size: 1rem;
  flex-shrink: 0;
}

.nav-search-input {
  flex: 1;
  border: none;
  background: var(--parchment-mid);
  border-radius: 6px;
  padding: 6px 10px;
  font-family: inherit;
  font-size: 0.82rem;
  color: var(--ink);
  outline: none;
}

.nav-search-input:focus {
  box-shadow: 0 0 0 2px var(--gold);
}

.nav-search-input::placeholder {
  color: var(--ink-light);
}

.nav-groups {
  flex: 1;
  overflow-y: auto;
  padding: 8px 8px 0;
}

.nav-group {
  margin-bottom: 4px;
}

.nav-group-label {
  font-size: 0.63rem;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: var(--ink-light);
  font-family: sans-serif;
  font-weight: 700;
  padding: 10px 8px 4px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 100%;
  padding: 8px 10px;
  border: none;
  background: none;
  border-radius: 7px;
  cursor: pointer;
  font-family: inherit;
  font-size: 0.875rem;
  color: var(--ink-mid);
  text-align: left;
  transition: background var(--transition), color var(--transition);
  position: relative;
}

.nav-item:hover {
  background: var(--parchment-mid);
  color: var(--ink);
}

.nav-item.active {
  background: rgba(201, 151, 58, 0.12);
  color: var(--ink);
  font-weight: 600;
}

.nav-item.active::before {
  content: '';
  position: absolute;
  left: 0;
  top: 20%;
  bottom: 20%;
  width: 3px;
  background: var(--gold);
  border-radius: 0 2px 2px 0;
}

.nav-item-icon {
  font-size: 0.95rem;
  width: 20px;
  text-align: center;
  flex-shrink: 0;
}

.nav-item-label {
  flex: 1;
}

.nav-item-badge {
  font-size: 0.65rem;
  font-family: sans-serif;
  padding: 2px 6px;
  border-radius: 10px;
}

.nav-item-badge.neutral {
  background: var(--parchment-dark);
  color: var(--ink-mid);
}

.nav-item-badge.warning {
  background: #f5e6c0;
  color: #8b6020;
}

.nav-item-badge.success {
  background: #d4edda;
  color: #2d6a4f;
}

.nav-version {
  padding: 12px 16px;
  border-top: 1px solid var(--border-light);
}

.version-text {
  font-size: 0.72rem;
  color: var(--ink-light);
  font-family: sans-serif;
}

/* ── Content ──────────────────────────────── */
.settings-content {
  display: flex;
  flex-direction: column;
  overflow: hidden;
  min-width: 0;
}

.content-header {
  padding: 20px 32px 16px;
  border-bottom: 1px solid var(--border-light);
  flex-shrink: 0;
}

.header-breadcrumb {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 4px;
}

.breadcrumb-root {
  font-size: 0.78rem;
  color: var(--ink-light);
}

.breadcrumb-sep {
  font-size: 0.78rem;
  color: var(--ink-light);
}

.breadcrumb-current {
  font-size: 0.78rem;
  color: var(--gold);
  font-weight: 600;
}

.content-description {
  font-size: 0.82rem;
  color: var(--ink-light);
  font-style: italic;
  margin: 0;
}

.content-body {
  flex: 1;
  overflow-y: auto;
  padding: 28px 32px;
}

/* ── Section transition ───────────────────── */
.section-fade-enter-active,
.section-fade-leave-active {
  transition: opacity 0.15s ease, transform 0.15s ease;
}

.section-fade-enter-from {
  opacity: 0;
  transform: translateX(8px);
}

.section-fade-leave-to {
  opacity: 0;
  transform: translateX(-8px);
}

/* ── Responsive ───────────────────────────── */
@media (max-width: 768px) {
  .settings-shell {
    grid-template-columns: 1fr;
    grid-template-rows: auto 1fr;
  }

  .settings-nav {
    border-right: none;
    border-bottom: 1px solid var(--border);
    max-height: 200px;
  }
}
</style>