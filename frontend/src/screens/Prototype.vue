<template>
    <div class="app-shell" :class="{ 'sidebar-collapsed': sidebarCollapsed }">

        <!-- ═══════════════════════════════════════════
         LEFT PANEL — Navigation Sidebar
    ════════════════════════════════════════════ -->
        <aside class="sidebar" :class="{ collapsed: sidebarCollapsed }">

            <!-- Top: Logo / Brand -->
            <div class="sidebar-brand">
                <!-- <span class="brand-icon">✦</span> -->
                <span class="brand-icon"><Brain color="#4a3218" /></span>
                <transition name="label-fade">
                    <span v-if="!sidebarCollapsed" class="brand-text">Manna</span>
                </transition>
            </div>

            <!-- Toggle Button -->
            <button class="sidebar-toggle" @click="toggleSidebar"
                :title="sidebarCollapsed ? 'Expand menu' : 'Collapse menu'">
                <span class="toggle-icon" :class="{ rotated: sidebarCollapsed }">‹</span>
            </button>

            <!-- Nav Items -->
            <nav class="sidebar-nav">
                <button v-for="item in navItems" :key="item.id" class="nav-item"
                    :class="{ active: activeView === item.id }" @click="setView(item.id)"
                    :title="sidebarCollapsed ? item.label : ''">
                    <span class="nav-icon">{{ item.icon }}</span>
                    <transition name="label-fade">
                        <span v-if="!sidebarCollapsed" class="nav-label">{{ item.label }}</span>
                    </transition>
                    <transition name="label-fade">
                        <span v-if="!sidebarCollapsed && item.badge" class="nav-badge">{{ item.badge }}</span>
                    </transition>
                </button>
            </nav>

            <!-- Bottom: Settings & Profile -->
            <div class="sidebar-footer">
                <button class="nav-item footer-item" :class="{ active: activeView === 'settings' }"
                    @click="setView('settings')" :title="sidebarCollapsed ? 'Settings' : ''">
                    <span class="nav-icon">⚙</span>
                    <transition name="label-fade">
                        <span v-if="!sidebarCollapsed" class="nav-label">Settings</span>
                    </transition>
                </button>

                <button class="nav-item footer-item profile-item" :class="{ active: activeView === 'profile' }"
                    @click="setView('profile')" :title="sidebarCollapsed ? 'Profile' : ''">
                    <span class="nav-icon profile-avatar">{{ userInitial }}</span>
                    <transition name="label-fade">
                        <span v-if="!sidebarCollapsed" class="nav-label">{{ userName }}</span>
                    </transition>
                </button>
            </div>
        </aside>

        <!-- ═══════════════════════════════════════════
         CENTER PANEL — Main Content
    ════════════════════════════════════════════ -->
        <main class="center-panel">
            <div class="center-header">
                <div class="header-left">
                    <h1 class="page-title">{{ currentView.title }}</h1>
                    <p class="page-subtitle">{{ currentView.subtitle }}</p>
                </div>
                <div class="header-actions">
                    <button class="action-btn" @click="toggleRightPanel"
                        :title="rightPanelOpen ? 'Hide panel' : 'Show panel'">
                        <span>{{ rightPanelOpen ? '▶' : '◀' }}</span>
                    </button>
                </div>
            </div>

            <div class="center-content">

                <!-- HOME VIEW -->
                <div v-if="activeView === 'home'" class="view-home">
                    <div class="verse-of-day">
                        <div class="votd-label">Verse of the Day</div>
                        <blockquote class="votd-text">
                            "For God so loved the world that he gave his one and only Son, that whoever believes in him
                            shall not perish but have eternal life."
                        </blockquote>
                        <cite class="votd-ref">— John 3:16 (NIV)</cite>
                    </div>

                    <div class="home-grid">
                        <div class="home-card" @click="setView('read')">
                            <div class="card-icon">📖</div>
                            <h3>Continue Reading</h3>
                            <p>Genesis · Chapter 3</p>
                            <div class="card-progress">
                                <div class="progress-bar">
                                    <div class="progress-fill" style="width: 12%"></div>
                                </div>
                                <span>12% complete</span>
                            </div>
                        </div>
                        <div class="home-card" @click="setView('notes')">
                            <div class="card-icon">📝</div>
                            <h3>Recent Notes</h3>
                            <p>3 notes this week</p>
                            <div class="card-meta">Last edited · 2h ago</div>
                        </div>
                        <div class="home-card">
                            <div class="card-icon">🔖</div>
                            <h3>Bookmarks</h3>
                            <p>14 saved passages</p>
                            <div class="card-meta">Across 7 books</div>
                        </div>
                        <div class="home-card">
                            <div class="card-icon">📅</div>
                            <h3>Reading Plan</h3>
                            <p>Day 42 of 365</p>
                            <div class="card-progress">
                                <div class="progress-bar">
                                    <div class="progress-fill" style="width: 11.5%"></div>
                                </div>
                                <span>11% complete</span>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- READ BIBLE VIEW -->
                <div v-else-if="activeView === 'read'" class="view-read">
                    <div class="bible-nav">
                        <select class="bible-select" v-model="selectedBook">
                            <option v-for="b in books" :key="b" :value="b">{{ b }}</option>
                        </select>
                        <select class="bible-select" v-model="selectedChapter">
                            <option v-for="c in chapters" :key="c" :value="c">Chapter {{ c }}</option>
                        </select>
                        <select class="bible-select">
                            <option>NIV</option>
                            <option>ESV</option>
                            <option>KJV</option>
                            <option>NLT</option>
                        </select>
                    </div>

                    <div class="scripture-body">
                        <h2 class="scripture-heading">{{ selectedBook }} · Chapter {{ selectedChapter }}</h2>
                        <div class="verses">
                            <div v-for="verse in sampleVerses" :key="verse.num" class="verse"
                                :class="{ highlighted: verse.highlighted }" @click="selectVerse(verse)">
                                <span class="verse-num">{{ verse.num }}</span>
                                <span class="verse-text">{{ verse.text }}</span>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- NOTES VIEW -->
                <div v-else-if="activeView === 'notes'" class="view-notes">
                    <div class="notes-toolbar">
                        <button class="btn-primary" @click="newNote">+ New Note</button>
                        <input class="search-input" type="text" placeholder="Search notes…" />
                    </div>
                    <div class="notes-list">
                        <div v-for="note in sampleNotes" :key="note.id" class="note-card"
                            :class="{ active: selectedNote?.id === note.id }" @click="selectedNote = note">
                            <div class="note-header">
                                <span class="note-ref">{{ note.ref }}</span>
                                <span class="note-date">{{ note.date }}</span>
                            </div>
                            <h4 class="note-title">{{ note.title }}</h4>
                            <p class="note-preview">{{ note.preview }}</p>
                        </div>
                    </div>
                </div>

                <!-- SETTINGS / PROFILE VIEWS -->
                <div v-else class="view-placeholder">
                    <div class="placeholder-icon">{{ currentView.icon }}</div>
                    <h2>{{ currentView.title }}</h2>
                    <p>This section is under construction.</p>
                </div>

            </div>
        </main>

        <!-- ═══════════════════════════════════════════
         RIGHT PANEL — Contextual Info
    ════════════════════════════════════════════ -->
        <transition name="right-panel-slide">
            <aside v-if="rightPanelOpen" class="right-panel">
                <div class="right-header">
                    <h3 class="right-title">{{ rightPanel.title }}</h3>
                    <button class="close-btn" @click="rightPanelOpen = false">✕</button>
                </div>

                <div class="right-content">

                    <!-- Home: Daily Plan -->
                    <template v-if="activeView === 'home'">
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

                    <!-- Read: Commentary / Cross-refs -->
                    <template v-else-if="activeView === 'read'">
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
                                This passage is foundational to understanding the covenant narrative that runs through
                                all of Scripture.
                                Many scholars note the parallel structure between the creation account and the
                                tabernacle instructions.
                            </p>
                        </div>
                        <div class="rp-section">
                            <div class="rp-section-label">Original Language</div>
                            <div class="orig-word" v-for="w in origWords" :key="w.word">
                                <span class="greek">{{ w.word }}</span>
                                <span class="trans">{{ w.meaning }}</span>
                            </div>
                        </div>
                    </template>

                    <!-- Notes: Note editor -->
                    <template v-else-if="activeView === 'notes'">
                        <div v-if="selectedNote" class="note-editor">
                            <div class="editor-ref">{{ selectedNote.ref }}</div>
                            <div class="editor-title">{{ selectedNote.title }}</div>
                            <div class="editor-body">{{ selectedNote.body }}</div>
                            <div class="editor-tags">
                                <span class="tag" v-for="t in selectedNote.tags" :key="t">{{ t }}</span>
                            </div>
                        </div>
                        <div v-else class="no-note-selected">
                            <p>Select a note to view or edit it here.</p>
                        </div>
                    </template>

                    <!-- Other views -->
                    <template v-else>
                        <p style="color: var(--ink-light); font-size: 0.9rem;">No panel content for this section.</p>
                    </template>

                </div>
            </aside>
        </transition>

    </div>
</template>

<script setup>
import { Brain } from '@lucide/vue'
import { ref, computed } from 'vue'

/* ── State ─────────────────────────────────── */
const sidebarCollapsed = ref(false)
const rightPanelOpen = ref(true)
const activeView = ref('home')
const selectedNote = ref(null)
const selectedBook = ref('Genesis')
const selectedChapter = ref(1)

const userName = 'James R.'
const userInitial = 'J'

/* ── Navigation ────────────────────────────── */
const navItems = [
    { id: 'home', label: 'Home', icon: '⌂', badge: null },
    { id: 'read', label: 'Read Bible', icon: '📖', badge: null },
    { id: 'notes', label: 'Notes', icon: '✎', badge: '3' },
    { id: 'bookmarks', label: 'Bookmarks', icon: '🔖', badge: null },
    { id: 'plans', label: 'Plans', icon: '📅', badge: null },
    { id: 'search', label: 'Search', icon: '⌕', badge: null },
]

const views = {
    home: { title: 'Good Morning, James', subtitle: 'Continue your journey', icon: '⌂' },
    read: { title: 'Read the Bible', subtitle: 'Explore Scripture', icon: '📖' },
    notes: { title: 'My Notes', subtitle: 'Reflections & insights', icon: '✎' },
    bookmarks: { title: 'Bookmarks', subtitle: 'Saved passages', icon: '🔖' },
    plans: { title: 'Reading Plans', subtitle: 'Stay on track', icon: '📅' },
    search: { title: 'Search Scripture', subtitle: 'Find any passage', icon: '⌕' },
    settings: { title: 'Settings', subtitle: 'Preferences & display', icon: '⚙' },
    profile: { title: 'Profile', subtitle: 'Your account', icon: '👤' },
}
const currentView = computed(() => views[activeView.value] || views.home)

const rightPanelTitles = {
    home: 'Today', read: 'Study Tools', notes: 'Note Detail',
    bookmarks: 'Details', plans: 'Progress', search: 'Filters',
    settings: '', profile: '',
}
const rightPanel = computed(() => ({ title: rightPanelTitles[activeView.value] || 'Info' }))

/* ── Actions ───────────────────────────────── */
function toggleSidebar() { sidebarCollapsed.value = !sidebarCollapsed.value }
function toggleRightPanel() { rightPanelOpen.value = !rightPanelOpen.value }
function setView(id) { activeView.value = id }
function newNote() { selectedNote.value = null }
function selectVerse(verse) { verse.highlighted = !verse.highlighted }

/* ── Sample Data ───────────────────────────── */
const books = ['Genesis', 'Exodus', 'Psalms', 'Proverbs', 'Matthew', 'John', 'Romans', 'Revelation']
const chapters = Array.from({ length: 50 }, (_, i) => i + 1)

const sampleVerses = ref([
    { num: 1, text: 'In the beginning God created the heavens and the earth.', highlighted: false },
    { num: 2, text: 'Now the earth was formless and empty, darkness was over the surface of the deep, and the Spirit of God was hovering over the waters.', highlighted: false },
    { num: 3, text: 'And God said, "Let there be light," and there was light.', highlighted: true },
    { num: 4, text: 'God saw that the light was good, and he separated the light from the darkness.', highlighted: false },
    { num: 5, text: 'God called the light "day," and the darkness he called "night." And there was evening, and there was morning—the first day.', highlighted: false },
    { num: 6, text: 'And God said, "Let there be a vault between the waters to separate water from water."', highlighted: false },
    { num: 7, text: 'So God made the vault and separated the water under the vault from the water above it. And it was so.', highlighted: false },
    { num: 8, text: 'God called the vault "sky." And there was evening, and there was morning—the second day.', highlighted: false },
    { num: 9, text: 'And God said, "Let the water under the sky be gathered to one place, and let dry ground appear." And it was so.', highlighted: false },
    { num: 10, text: 'God called the dry ground "land," and the gathered waters he called "seas." And God saw that it was good.', highlighted: false },
])

const sampleNotes = ref([
    {
        id: 1, ref: 'John 3:16', date: 'May 19',
        title: 'The Nature of Divine Love',
        preview: 'This verse captures the entire gospel in a single sentence…',
        body: 'This verse captures the entire gospel in a single sentence. The Greek word "agape" here denotes a selfless, sacrificial love — not merely affection but intentional action on behalf of another.\n\nNote the universal scope: "the world" — not the righteous, not Israel only, but all of humanity.',
        tags: ['love', 'gospel', 'Greek'],
    },
    {
        id: 2, ref: 'Romans 8:28', date: 'May 17',
        title: 'All Things Work Together',
        preview: 'Sovereignty and suffering held in tension…',
        body: 'Paul writes from prison, and yet affirms God\'s providential care. This is not naive optimism — it is a hard-won theological conviction forged in suffering.',
        tags: ['providence', 'suffering', 'Romans'],
    },
    {
        id: 3, ref: 'Genesis 1:1', date: 'May 14',
        title: 'Creatio Ex Nihilo',
        preview: 'Creation from nothing — a foundational doctrine…',
        body: 'The Hebrew "bara" used here is reserved exclusively for divine creative activity. God alone creates from nothing; humans form and shape from existing material.',
        tags: ['creation', 'Hebrew', 'doctrine'],
    },
])

const todaysPlan = [
    { ref: 'Genesis 3–4', done: true },
    { ref: 'Psalm 23', done: true },
    { ref: 'Matthew 5:1–12', done: false },
]
const quickLinks = ['Prayer Journal', 'Concordance', 'Bible Maps', 'Lexicon']

const crossRefs = [
    { ref: 'John 1:1', snippet: '"In the beginning was the Word…"' },
    { ref: 'Heb 11:3', snippet: 'By faith we understand the universe was created…' },
    { ref: 'Col 1:16', snippet: 'For in him all things were created…' },
]
const origWords = [
    { word: 'בְּרֵאשִׁית', meaning: 'b\'reshit — In the beginning' },
    { word: 'בָּרָא', meaning: 'bara — created (divine act)' },
    { word: 'אֱלֹהִים', meaning: 'Elohim — God (plural majesty)' },
]
</script>

<style scoped>
/* ═══════════════════════════════════════════════════════════
   DESIGN TOKENS
════════════════════════════════════════════════════════════ */
:root {
    /* Warm parchment palette */
    --parchment: #f7f2e8;
    --parchment-mid: #ede6d5;
    --parchment-dark: #d9cfb8;
    --sidebar-bg: #1c1610;
    --sidebar-text: #c8b99a;
    --sidebar-hover: rgba(200, 185, 154, 0.12);
    --sidebar-active: #8b6f3e;
    --ink: #2c1f0e;
    --ink-mid: #5a4230;
    --ink-light: #8b7355;
    --gold: #c9973a;
    --gold-light: #e8c06a;
    --red-accent: #8b2e2e;
    --border: #d5c9b0;
    --border-light: #e8e0cc;
    --white: #ffffff;

    /* Sizing */
    --sidebar-w: 220px;
    --sidebar-w-min: 64px;
    --right-panel-w: 300px;
    --header-h: 64px;
    --radius: 10px;
    --transition: 0.28s cubic-bezier(0.4, 0, 0.2, 1);
}

/* ═══════════════════════════════════════════════════════════
   RESET / BASE
════════════════════════════════════════════════════════════ */
*,
*::before,
*::after {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
}

.app-shell {
    display: grid;
    grid-template-columns: var(--sidebar-w) 1fr var(--right-panel-w);
    grid-template-rows: 100vh;
    height: 100vh;
    overflow: hidden;
    background: var(--parchment);
    font-family: 'Georgia', 'Times New Roman', serif;
    color: var(--ink);
    transition: grid-template-columns var(--transition);
}

.app-shell.sidebar-collapsed {
    grid-template-columns: var(--sidebar-w-min) 1fr var(--right-panel-w);
}

/* ═══════════════════════════════════════════════════════════
   SIDEBAR
════════════════════════════════════════════════════════════ */
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

/* Brand */
.sidebar-brand {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 22px 18px 16px;
    border-bottom: 1px solid rgba(255, 255, 255, 0.07);
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

/* Toggle */
.sidebar-toggle {
    position: absolute;
    top: 20px;
    right: 10px;
    background: rgba(255, 255, 255, 0.06);
    border: 1px solid rgba(255, 255, 255, 0.1);
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
    transition: background var(--transition), transform var(--transition);
    z-index: 5;
}

.sidebar-toggle:hover {
    background: rgba(255, 255, 255, 0.14);
}

.toggle-icon {
    display: inline-block;
    transition: transform var(--transition);
    font-size: 1.1rem;
}

.toggle-icon.rotated {
    transform: rotate(180deg);
}

/* Nav */
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
}

.nav-item:hover {
    background: var(--sidebar-hover);
    color: var(--parchment);
}

.nav-item.active {
    background: rgba(139, 111, 62, 0.25);
    color: var(--gold-light);
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

.nav-icon {
    font-size: 1rem;
    width: 24px;
    text-align: center;
    flex-shrink: 0;
}

.nav-label {
    flex: 1;
}

.nav-badge {
    font-size: 0.7rem;
    background: var(--red-accent);
    color: white;
    border-radius: 10px;
    padding: 1px 7px;
    font-family: sans-serif;
    flex-shrink: 0;
}

/* Footer */
.sidebar-footer {
    padding: 10px 10px 16px;
    border-top: 1px solid rgba(255, 255, 255, 0.07);
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

/* ═══════════════════════════════════════════════════════════
   CENTER PANEL
════════════════════════════════════════════════════════════ */
.center-panel {
    grid-column: 2;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    background: var(--parchment);
    min-width: 0;
}

.center-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 28px;
    height: var(--header-h);
    border-bottom: 1px solid var(--border);
    background: var(--parchment);
    flex-shrink: 0;
}

.page-title {
    font-size: 1.3rem;
    font-weight: 700;
    color: var(--ink);
    letter-spacing: 0.01em;
}

.page-subtitle {
    font-size: 0.8rem;
    color: var(--ink-light);
    margin-top: 1px;
    font-style: italic;
}

.header-actions {
    display: flex;
    gap: 8px;
}

.action-btn {
    background: none;
    border: 1px solid var(--border);
    border-radius: 6px;
    padding: 6px 10px;
    cursor: pointer;
    color: var(--ink-mid);
    font-size: 0.8rem;
    transition: background var(--transition), color var(--transition);
}

.action-btn:hover {
    background: var(--parchment-mid);
    color: var(--ink);
}

.center-content {
    flex: 1;
    overflow-y: auto;
    padding: 28px;
}

/* ═══════════════════════════════════════════════════════════
   HOME VIEW
════════════════════════════════════════════════════════════ */
.verse-of-day {
    background: linear-gradient(135deg, #a9a094 0%, #4a3218 100%);
    border-radius: var(--radius);
    padding: 28px 32px;
    margin-bottom: 28px;
    position: relative;
    overflow: hidden;
}

.verse-of-day::before {
    content: '❝';
    position: absolute;
    top: -10px;
    left: 16px;
    font-size: 7rem;
    color: rgba(201, 151, 58, 0.12);
    line-height: 1;
    font-family: Georgia, serif;
}

.votd-label {
    font-size: 0.72rem;
    letter-spacing: 0.14em;
    text-transform: uppercase;
    color: var(--gold);
    margin-bottom: 14px;
    font-family: sans-serif;
}

.votd-text {
    font-size: 1.05rem;
    line-height: 1.75;
    color: var(--parchment);
    font-style: italic;
    border: none;
    padding: 0;
    margin-bottom: 14px;
    position: relative;
}

.votd-ref {
    color: var(--gold-light);
    font-size: 0.85rem;
    font-style: normal;
}

.home-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    gap: 16px;
}

.home-card {
    background: var(--white);
    border: 1px solid var(--border-light);
    border-radius: var(--radius);
    padding: 20px;
    cursor: pointer;
    transition: transform var(--transition), box-shadow var(--transition);
}

.home-card:hover {
    transform: translateY(-2px);
    box-shadow: 0 8px 24px rgba(44, 31, 14, 0.1);
}

.card-icon {
    font-size: 1.6rem;
    margin-bottom: 10px;
}

.home-card h3 {
    font-size: 0.95rem;
    color: var(--ink);
    margin-bottom: 4px;
}

.home-card p {
    font-size: 0.8rem;
    color: var(--ink-light);
    font-style: italic;
}

.card-progress {
    margin-top: 12px;
}

.progress-bar {
    height: 4px;
    background: var(--parchment-dark);
    border-radius: 2px;
    overflow: hidden;
    margin-bottom: 4px;
}

.progress-fill {
    height: 100%;
    background: var(--gold);
    border-radius: 2px;
    transition: width 0.6s ease;
}

.card-progress span,
.card-meta {
    font-size: 0.72rem;
    color: var(--ink-light);
    font-family: sans-serif;
}

/* ═══════════════════════════════════════════════════════════
   READ BIBLE VIEW
════════════════════════════════════════════════════════════ */
.bible-nav {
    display: flex;
    gap: 10px;
    margin-bottom: 24px;
    flex-wrap: wrap;
}

.bible-select {
    appearance: none;
    background: var(--white);
    border: 1px solid var(--border);
    border-radius: 7px;
    padding: 8px 14px;
    font-family: inherit;
    font-size: 0.88rem;
    color: var(--ink);
    cursor: pointer;
}

.bible-select:focus {
    outline: 2px solid var(--gold);
    outline-offset: 1px;
}

.scripture-heading {
    font-size: 1.4rem;
    color: var(--ink);
    margin-bottom: 24px;
    padding-bottom: 14px;
    border-bottom: 1px solid var(--border-light);
    letter-spacing: 0.02em;
}

.verses {
    display: flex;
    flex-direction: column;
    gap: 0;
}

.verse {
    display: flex;
    gap: 14px;
    padding: 10px 12px;
    border-radius: 6px;
    cursor: pointer;
    transition: background var(--transition);
    line-height: 1.8;
}

.verse:hover {
    background: var(--parchment-mid);
}

.verse.highlighted {
    background: rgba(201, 151, 58, 0.15);
}

.verse-num {
    font-size: 0.72rem;
    font-family: sans-serif;
    color: var(--gold);
    font-weight: 700;
    min-width: 20px;
    padding-top: 5px;
    flex-shrink: 0;
}

.verse-text {
    font-size: 1rem;
    color: var(--ink-mid);
}

/* ═══════════════════════════════════════════════════════════
   NOTES VIEW
════════════════════════════════════════════════════════════ */
.notes-toolbar {
    display: flex;
    gap: 12px;
    margin-bottom: 20px;
    flex-wrap: wrap;
}

.btn-primary {
    background: var(--ink);
    color: var(--parchment);
    border: none;
    border-radius: 7px;
    padding: 9px 18px;
    font-family: inherit;
    font-size: 0.88rem;
    cursor: pointer;
    transition: background var(--transition);
}

.btn-primary:hover {
    background: var(--ink-mid);
}

.search-input {
    flex: 1;
    min-width: 160px;
    padding: 9px 14px;
    border: 1px solid var(--border);
    border-radius: 7px;
    font-family: inherit;
    font-size: 0.88rem;
    background: var(--white);
    color: var(--ink);
}

.search-input:focus {
    outline: 2px solid var(--gold);
    outline-offset: 1px;
}

.notes-list {
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.note-card {
    background: var(--white);
    border: 1px solid var(--border-light);
    border-radius: var(--radius);
    padding: 16px 18px;
    cursor: pointer;
    transition: border-color var(--transition), box-shadow var(--transition);
}

.note-card:hover {
    border-color: var(--gold);
}

.note-card.active {
    border-color: var(--gold);
    box-shadow: 0 0 0 2px rgba(201, 151, 58, 0.2);
}

.note-header {
    display: flex;
    justify-content: space-between;
    margin-bottom: 6px;
}

.note-ref {
    font-size: 0.75rem;
    color: var(--gold);
    font-weight: 700;
    font-family: sans-serif;
}

.note-date {
    font-size: 0.75rem;
    color: var(--ink-light);
    font-family: sans-serif;
}

.note-title {
    font-size: 0.95rem;
    margin-bottom: 4px;
    color: var(--ink);
}

.note-preview {
    font-size: 0.82rem;
    color: var(--ink-light);
    font-style: italic;
}

/* ═══════════════════════════════════════════════════════════
   VIEW PLACEHOLDER
════════════════════════════════════════════════════════════ */
.view-placeholder {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
    color: var(--ink-light);
    text-align: center;
    gap: 14px;
}

.placeholder-icon {
    font-size: 3rem;
}

.view-placeholder h2 {
    font-size: 1.2rem;
    color: var(--ink-mid);
}

/* ═══════════════════════════════════════════════════════════
   RIGHT PANEL
════════════════════════════════════════════════════════════ */
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

.cross-ref {
    display: flex;
    flex-direction: column;
    gap: 3px;
    padding: 8px 0;
    border-bottom: 1px solid var(--border-light);
}

.cross-ref:last-child {
    border-bottom: none;
}

.cross-ref-tag {
    font-size: 0.75rem;
    font-weight: 700;
    color: var(--gold);
    font-family: sans-serif;
}

.cross-ref-text {
    font-size: 0.82rem;
    color: var(--ink-light);
    font-style: italic;
}

.commentary-text {
    font-size: 0.88rem;
    line-height: 1.7;
    color: var(--ink-mid);
}

.orig-word {
    display: flex;
    gap: 14px;
    align-items: baseline;
    padding: 6px 0;
}

.greek {
    font-size: 1rem;
    color: var(--ink);
}

.trans {
    font-size: 0.8rem;
    color: var(--ink-light);
    font-family: sans-serif;
    font-style: italic;
}

/* Note editor in right panel */
.note-editor {
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.editor-ref {
    font-size: 0.75rem;
    color: var(--gold);
    font-weight: 700;
    font-family: sans-serif;
    letter-spacing: 0.08em;
}

.editor-title {
    font-size: 1.05rem;
    font-weight: 700;
    color: var(--ink);
}

.editor-body {
    font-size: 0.88rem;
    line-height: 1.75;
    color: var(--ink-mid);
    white-space: pre-wrap;
}

.editor-tags {
    display: flex;
    gap: 6px;
    flex-wrap: wrap;
}

.tag {
    background: var(--parchment-mid);
    border: 1px solid var(--border);
    border-radius: 20px;
    padding: 3px 10px;
    font-size: 0.72rem;
    color: var(--ink-mid);
    font-family: sans-serif;
}

.no-note-selected {
    color: var(--ink-light);
    font-size: 0.88rem;
    font-style: italic;
}

/* ═══════════════════════════════════════════════════════════
   TRANSITIONS
════════════════════════════════════════════════════════════ */
.label-fade-enter-active,
.label-fade-leave-active {
    transition: opacity 0.18s ease, width 0.18s ease;
}

.label-fade-enter-from,
.label-fade-leave-to {
    opacity: 0;
}

.right-panel-slide-enter-active,
.right-panel-slide-leave-active {
    transition: width var(--transition), opacity var(--transition);
    overflow: hidden;
}

.right-panel-slide-enter-from,
.right-panel-slide-leave-to {
    width: 0;
    opacity: 0;
}

.right-panel-slide-enter-to,
.right-panel-slide-leave-from {
    width: var(--right-panel-w);
    opacity: 1;
}

/* ═══════════════════════════════════════════════════════════
   RESPONSIVE
════════════════════════════════════════════════════════════ */

/* Tablet: collapse right panel by default */
@media (max-width: 1024px) {
    :root {
        --right-panel-w: 260px;
    }

    .right-panel {
        position: fixed;
        right: 0;
        top: 0;
        bottom: 0;
        z-index: 100;
        box-shadow: -4px 0 20px rgba(44, 31, 14, 0.15);
    }

    .app-shell {
        grid-template-columns: var(--sidebar-w) 1fr;
    }

    .app-shell.sidebar-collapsed {
        grid-template-columns: var(--sidebar-w-min) 1fr;
    }
}

/* Mobile: sidebar becomes bottom nav, right panel full overlay */
@media (max-width: 640px) {
    .app-shell {
        grid-template-columns: 1fr;
        grid-template-rows: 1fr 60px;
        height: 100dvh;
    }

    .app-shell.sidebar-collapsed {
        grid-template-columns: 1fr;
    }

    .sidebar {
        grid-column: 1;
        grid-row: 2;
        width: 100% !important;
        min-width: 0 !important;
        height: 60px;
        flex-direction: row;
        align-items: center;
        justify-content: space-around;
        border-right: none;
        border-top: 1px solid #2e2318;
        overflow: visible;
    }

    .sidebar-brand,
    .sidebar-toggle,
    .sidebar-footer {
        display: none;
    }

    .sidebar-nav {
        flex-direction: row;
        padding: 0;
        margin: 0;
        gap: 0;
        justify-content: space-around;
        align-items: center;
        width: 100%;
        overflow: visible;
    }

    .nav-item {
        flex-direction: column;
        gap: 2px;
        padding: 8px 12px;
        font-size: 0.65rem;
        border-radius: 8px;
        width: auto;
    }

    .nav-item::before {
        display: none;
    }

    .nav-icon {
        font-size: 1.1rem;
        width: auto;
    }

    .nav-label {
        font-size: 0.6rem;
        white-space: nowrap;
    }

    .nav-badge {
        display: none;
    }

    .center-panel {
        grid-column: 1;
        grid-row: 1;
    }

    .center-content {
        padding: 16px;
    }

    .center-header {
        padding: 0 16px;
    }

    .right-panel {
        width: 100% !important;
        max-width: 100%;
    }

    .home-grid {
        grid-template-columns: 1fr 1fr;
    }
}
</style>