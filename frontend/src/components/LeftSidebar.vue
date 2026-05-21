<template>
  <aside class="sidebar" :class="{ collapsed: ui.sidebarCollapsed }">

    <!-- Brand -->
    <div class="sidebar-brand">
      <span class="brand-icon"><Wheat color="#a9a094" /></span>
      <!-- <span class="brand-icon">✦</span> -->
      <Transition name="label-fade">
        <span v-if="!ui.sidebarCollapsed" class="brand-text">Manna</span>
        <!-- <span v-if="!ui.sidebarCollapsed" class="brand-text">Verbum</span> -->
      </Transition>
    </div>

    <!-- Collapse toggle -->
    <button
      class="sidebar-toggle"
      @click="ui.toggleSidebar"
      :title="ui.sidebarCollapsed ? 'Expand menu' : 'Collapse menu'"
    >
      <span class="toggle-icon" :class="{ rotated: ui.sidebarCollapsed }">‹</span>
    </button>

    <!-- Primary nav -->
    <nav class="sidebar-nav">
      <RouterLink
        v-for="item in navItems"
        :key="item.name"
        :to="{ name: item.name }"
        class="nav-item"
        :title="ui.sidebarCollapsed ? item.label : ''"
        custom
        v-slot="{ navigate, isActive }"
      >
        <button
          class="nav-item"
          :class="{ active: isActive }"
          @click="navigate"
        >
          <span class="nav-icon">{{ item.icon }}</span>
          <Transition name="label-fade">
            <span v-if="!ui.sidebarCollapsed" class="nav-label">{{ item.label }}</span>
          </Transition>
          <Transition name="label-fade">
            <span v-if="!ui.sidebarCollapsed && item.badge" class="nav-badge">{{ item.badge }}</span>
          </Transition>
        </button>
      </RouterLink>
    </nav>

    <!-- Footer: Settings + Profile -->
    <div class="sidebar-footer">
      <RouterLink :to="{ name: 'settings' }" custom v-slot="{ navigate, isActive }">
        <button
          class="nav-item footer-item"
          :class="{ active: isActive }"
          :title="ui.sidebarCollapsed ? 'Settings' : ''"
          @click="navigate"
        >
          <span class="nav-icon">⚙</span>
          <Transition name="label-fade">
            <span v-if="!ui.sidebarCollapsed" class="nav-label">Settings</span>
          </Transition>
        </button>
      </RouterLink>

      <RouterLink :to="{ name: 'profile' }" custom v-slot="{ navigate, isActive }">
        <button
          class="nav-item footer-item profile-item"
          :class="{ active: isActive }"
          :title="ui.sidebarCollapsed ? 'Profile' : ''"
          @click="navigate"
        >
          <span class="nav-icon profile-avatar">{{ userInitial }}</span>
          <Transition name="label-fade">
            <span v-if="!ui.sidebarCollapsed" class="nav-label">{{ userName }}</span>
          </Transition>
        </button>
      </RouterLink>
    </div>

  </aside>
</template>

<script setup>
import { Wheat } from '@lucide/vue';
import { useUiStore } from '../stores/ui';


const ui = useUiStore()

const userName    = 'James R.'
const userInitial = 'J'

const navItems = [
  { name: 'home',      label: 'Home',       icon: '⌂',  badge: null },
  { name: 'read',      label: 'Read Bible', icon: '📖', badge: null },
  { name: 'notes',     label: 'Notes',      icon: '✎',  badge: '3'  },
  { name: 'bookmarks', label: 'Bookmarks',  icon: '🔖', badge: null },
  { name: 'plans',     label: 'Plans',      icon: '📅', badge: null },
  { name: 'search',    label: 'Search',     icon: '⌕',  badge: null },
]
</script>

<style scoped>
/* ── Sidebar shell ──────────────────────────── */
.sidebar {
  grid-column: 1;
  background: var(--sidebar-bg);
  display: flex;
  flex-direction: column;
  width: var(--sidebar-w);
  min-width: var(--sidebar-w);
  transition: width var(--transition), min-width var(--transition);
  overflow: hidden;
  position: relative;
  z-index: 10;
  border-right: 1px solid #2e2318;
}
.sidebar.collapsed {
  width: var(--sidebar-w-min);
  min-width: var(--sidebar-w-min);
}

/* ── Brand ──────────────────────────────────── */
.sidebar-brand {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 22px 18px 16px;
  border-bottom: 1px solid rgba(255,255,255,0.07);
  overflow: hidden;
  white-space: nowrap;
}
.brand-icon {
  font-size: 1.3rem;
  color: var(--gold);
  flex-shrink: 0;
  width: 28px;
  text-align: center;
}
.brand-text {
  font-size: 1.2rem;
  font-weight: 700;
  letter-spacing: 0.08em;
  color: var(--parchment);
  font-family: 'Georgia', serif;
}

/* ── Toggle button ──────────────────────────── */
.sidebar-toggle {
  position: absolute;
  top: 20px;
  right: 10px;
  background: rgba(255,255,255,0.06);
  border: 1px solid rgba(255,255,255,0.1);
  border-radius: 50%;
  width: 26px;
  height: 26px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: var(--sidebar-text);
  font-size: 1rem;
  line-height: 1;
  transition: background var(--transition);
  z-index: 5;
}
.sidebar-toggle:hover { background: rgba(255,255,255,0.14); }

.toggle-icon {
  display: inline-block;
  transition: transform var(--transition);
  font-size: 1.1rem;
}
.toggle-icon.rotated { transform: rotate(180deg); }

/* ── Nav ────────────────────────────────────── */
.sidebar-nav {
  flex: 1;
  padding: 8px 10px;
  display: flex;
  flex-direction: column;
  gap: 2px;
  overflow-y: auto;
  overflow-x: hidden;
  margin-top: 4px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 8px;
  border: none;
  background: transparent;
  color: var(--sidebar-text);
  border-radius: 8px;
  cursor: pointer;
  font-family: inherit;
  font-size: 0.9rem;
  letter-spacing: 0.02em;
  text-align: left;
  width: 100%;
  overflow: hidden;
  white-space: nowrap;
  transition: background var(--transition), color var(--transition);
  position: relative;
  text-decoration: none;
}
.nav-item:hover { background: var(--sidebar-hover); color: var(--parchment); }
.nav-item.active {
  background: rgba(139, 111, 62, 0.25);
  color: var(--gold-light);
}
.nav-item.active::before {
  content: '';
  position: absolute;
  left: 0; top: 20%; bottom: 20%;
  width: 3px;
  background: var(--gold);
  border-radius: 0 2px 2px 0;
}

.nav-icon {
  font-size: 1rem;
  width: 24px;
  text-align: center;
  flex-shrink: 0;
}
.nav-label { flex: 1; }

.nav-badge {
  font-size: 0.7rem;
  background: var(--red-accent);
  color: white;
  border-radius: 10px;
  padding: 1px 7px;
  font-family: sans-serif;
  flex-shrink: 0;
}

/* ── Footer ─────────────────────────────────── */
.sidebar-footer {
  padding: 10px 10px 16px;
  border-top: 1px solid rgba(255,255,255,0.07);
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.profile-avatar {
  background: var(--gold);
  color: var(--sidebar-bg);
  border-radius: 50%;
  width: 26px;
  height: 26px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.8rem;
  font-weight: 700;
  font-family: sans-serif;
  flex-shrink: 0;
}

/* ── Transitions ────────────────────────────── */
.label-fade-enter-active,
.label-fade-leave-active { transition: opacity 0.18s ease; }
.label-fade-enter-from,
.label-fade-leave-to     { opacity: 0; }

/* ── Mobile: becomes a bottom tab bar ─────────  */
@media (max-width: 640px) {
  .sidebar {
    grid-column: 1;
    grid-row: 2;
    width: 100% !important;
    min-width: 0 !important;
    height: 60px;
    flex-direction: row;
    align-items: center;
    border-right: none;
    border-top: 1px solid #2e2318;
    overflow: visible;
  }
  .sidebar-brand,
  .sidebar-toggle,
  .sidebar-footer { display: none; }

  .sidebar-nav {
    flex-direction: row;
    padding: 0;
    margin: 0;
    gap: 0;
    justify-content: space-around;
    align-items: center;
    width: 100%;
    overflow: visible;
    flex: unset;
  }
  .nav-item {
    flex-direction: column;
    gap: 2px;
    padding: 8px 12px;
    font-size: 0.65rem;
    border-radius: 8px;
    width: auto;
  }
  .nav-item::before { display: none; }
  .nav-icon { font-size: 1.1rem; width: auto; }
  .nav-label { font-size: 0.6rem; white-space: nowrap; }
  .nav-badge { display: none; }
}
</style>