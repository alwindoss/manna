// src/stores/ui.js
import { defineStore } from 'pinia'
import { ref, shallowRef, computed, markRaw } from 'vue'

export const useUiStore = defineStore('ui', () => {
  // ── Sidebar ──────────────────────────────────
  const sidebarCollapsed = ref(false)
  function toggleSidebar() {
    sidebarCollapsed.value = !sidebarCollapsed.value
  }

  // ── Right panel ──────────────────────────────

  const rightPanelContext = shallowRef({
    component: null,   // which sub-component / template to show
    props: {}          // data to pass into it
  })

  function setRightPanel(component, props = {}) {
    rightPanelContext.value = { component: markRaw(component), props }
  }
  function clearRightPanel() {
    rightPanelContext.value = { component: null, props: {} }
  }

  // Some screens should not show any side panel, like Settings and Profile
  const disableRightPanel = ref(false)
  // function setRightPanelDisabled(val) {
  //   disableRightPanel.value = val
  //   if (val) rightPanelOpen.value = false  // close it immediately when disabled
  // }
  function setRightPanelDisabled(val) {
    if (val) {
      _panelStateBeforeDisable = rightPanelOpen.value  // save
      rightPanelOpen.value = false
    } else {
      rightPanelOpen.value = _panelStateBeforeDisable  // restore
    }
    disableRightPanel.value = val
  }

  let _panelStateBeforeDisable = true


  const rightPanelOpen = ref(true)
  function toggleRightPanel() {
    rightPanelOpen.value = !rightPanelOpen.value
  }
  function openRightPanel() { rightPanelOpen.value = true }
  function closeRightPanel() { rightPanelOpen.value = false }

  // ── Selected note (shared between Notes view & RightPanel) ──
  const selectedNote = ref(null)
  function setSelectedNote(note) { selectedNote.value = note }
  function clearSelectedNote() { selectedNote.value = null }

  // ── Selected verse (shared between Read view & RightPanel) ──
  const highlightedVerse = ref(null)
  function selectVerseHighlight(verseNum) {
    highlightedVerse.value = verseNum
  }
  const isVerseHighlighted = computed(
    () => (num) => highlightedVerse.value === num
  )

  return {
    sidebarCollapsed, toggleSidebar,
    rightPanelOpen, toggleRightPanel, openRightPanel, closeRightPanel,
    selectedNote, setSelectedNote, clearSelectedNote,
    highlightedVerse, selectVerseHighlight, isVerseHighlighted,
    rightPanelContext, setRightPanel, clearRightPanel,
    disableRightPanel, setRightPanelDisabled,
  }
})