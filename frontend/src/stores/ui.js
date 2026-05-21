// src/stores/ui.js
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useUiStore = defineStore('ui', () => {
  // ── Sidebar ──────────────────────────────────
  const sidebarCollapsed = ref(false)
  function toggleSidebar() {
    sidebarCollapsed.value = !sidebarCollapsed.value
  }

  // ── Right panel ──────────────────────────────
  const rightPanelOpen = ref(true)
  function toggleRightPanel() {
    rightPanelOpen.value = !rightPanelOpen.value
  }
  function openRightPanel()  { rightPanelOpen.value = true  }
  function closeRightPanel() { rightPanelOpen.value = false }

  // ── Selected note (shared between Notes view & RightPanel) ──
  const selectedNote = ref(null)
  function setSelectedNote(note) { selectedNote.value = note }
  function clearSelectedNote()   { selectedNote.value = null  }

  // ── Selected verse (shared between Read view & RightPanel) ──
  const highlightedVerses = ref([])
  function toggleVerseHighlight(verseNum) {
    const idx = highlightedVerses.value.indexOf(verseNum)
    if (idx === -1) highlightedVerses.value.push(verseNum)
    else            highlightedVerses.value.splice(idx, 1)
  }
  const isVerseHighlighted = computed(
    () => (num) => highlightedVerses.value.includes(num)
  )

  return {
    sidebarCollapsed, toggleSidebar,
    rightPanelOpen, toggleRightPanel, openRightPanel, closeRightPanel,
    selectedNote, setSelectedNote, clearSelectedNote,
    highlightedVerses, toggleVerseHighlight, isVerseHighlighted,
  }
})