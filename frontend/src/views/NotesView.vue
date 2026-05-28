<template>
  <div class="view-notes">

    <div class="notes-toolbar">
      <button class="btn-primary" @click="ui.clearSelectedNote">+ New Note</button>
      <input class="search-input" type="text" v-model="query" placeholder="Search notes…" />
    </div>

    <div class="notes-list">
      <div
        v-for="note in filteredNotes"
        :key="note.id"
        class="note-card"
        :class="{ active: ui.selectedNote?.id === note.id }"
        @click="ui.setSelectedNote(note)"
      >
        <div class="note-header">
          <span class="note-ref">{{ note.ref }}</span>
          <span class="note-date">{{ note.date }}</span>
        </div>
        <h4 class="note-title">{{ note.title }}</h4>
        <p class="note-preview">{{ note.preview }}</p>
      </div>

      <p v-if="filteredNotes.length === 0" class="empty-state">
        No notes match "{{ query }}"
      </p>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useUiStore } from '../stores/ui'
import NotesRightPanel from '../components/notes/NotesRightPanel.vue'

const ui    = useUiStore()
const query = ref('')

const allNotes = [
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
    body: "Paul writes from prison, and yet affirms God's providential care. This is not naive optimism — it is a hard-won theological conviction forged in suffering.",
    tags: ['providence', 'suffering', 'Romans'],
  },
  {
    id: 3, ref: 'Genesis 1:1', date: 'May 14',
    title: 'Creatio Ex Nihilo',
    preview: 'Creation from nothing — a foundational doctrine…',
    body: 'The Hebrew "bara" used here is reserved exclusively for divine creative activity. God alone creates from nothing; humans form and shape from existing material.',
    tags: ['creation', 'Hebrew', 'doctrine'],
  },
]

const filteredNotes = computed(() => {
  const q = query.value.toLowerCase()
  if (!q) return allNotes
  return allNotes.filter(
    n => n.title.toLowerCase().includes(q) ||
         n.ref.toLowerCase().includes(q)   ||
         n.preview.toLowerCase().includes(q)
  )
})

onMounted(() => {
  ui.setRightPanel(NotesRightPanel, {
    title: 'Note Detail',
  })

})

function onNoteSelect(note) {
  ui.setSelectedNote(note)
  ui.setRightPanel(NotesRightPanel, {
    title: note.title,
    note,                   // the full note object flows into the panel
  })
}
</script>

<style scoped>
.view-notes { display: flex; flex-direction: column; gap: 16px; }

/* ── Toolbar ────────────────────────────────── */
.notes-toolbar { display: flex; gap: 12px; flex-wrap: wrap; }

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
.btn-primary:hover { background: var(--ink-mid); }

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
.search-input:focus { outline: 2px solid var(--gold); outline-offset: 1px; }

/* ── List ───────────────────────────────────── */
.notes-list { display: flex; flex-direction: column; gap: 12px; }

.note-card {
  background: var(--white);
  border: 1px solid var(--border-light);
  border-radius: var(--radius);
  padding: 16px 18px;
  cursor: pointer;
  transition: border-color var(--transition), box-shadow var(--transition);
}
.note-card:hover { border-color: var(--gold); }
.note-card.active {
  border-color: var(--gold);
  box-shadow: 0 0 0 2px rgba(201, 151, 58, 0.2);
}

.note-header { display: flex; justify-content: space-between; margin-bottom: 6px; }
.note-ref    { font-size: 0.75rem; color: var(--gold); font-weight: 700; font-family: sans-serif; }
.note-date   { font-size: 0.75rem; color: var(--ink-light); font-family: sans-serif; }
.note-title  { font-size: 0.95rem; margin-bottom: 4px; color: var(--ink); }
.note-preview { font-size: 0.82rem; color: var(--ink-light); font-style: italic; }

.empty-state { color: var(--ink-light); font-style: italic; font-size: 0.88rem; }
</style>