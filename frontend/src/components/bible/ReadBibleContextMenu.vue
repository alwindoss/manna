<template>
  <div
    v-if="isVisible"
    class="context-menu"
    :style="{ top: position.y + 'px', left: position.x + 'px' }"
    @mouseleave="closeMenu"
  >
    <div class="context-menu-item" @click="handleCopyVerse">
      <span class="menu-icon"><Copy /></span>
      <span>Copy verse</span>
    </div>
    <div class="context-menu-item" @click="handleHighlightVerse">
      <span class="menu-icon"><Highlighter /></span>
      <span>Highlight</span>
    </div>
    <div class="context-menu-item" @click="handleShare">
      <span class="menu-icon"><Share2 /></span>
      <span>Share</span>
    </div>
    <div class="context-menu-divider"></div>
    <div class="context-menu-item" @click="handleAddBookmark">
      <span class="menu-icon"><Star /></span>
      <span>Add bookmark</span>
    </div>
    <div class="context-menu-item" @click="handleCreateNote">
      <span class="menu-icon"><NotebookPen /></span>
      <span>Create note</span>
    </div>
    <div class="context-menu-divider"></div>
    <div class="context-menu-item" @click="handleReport">
      <span class="menu-icon">🚩</span>
      <span>Report verse</span>
    </div>
  </div>
</template>

<script setup>
import { Copy, Highlighter, NotebookPen, Share2, Star } from '@lucide/vue';
import { ref } from 'vue';

const isVisible = ref(false);
const position = ref({ x: 0, y: 0 });
const selectedVerse = ref(null);

const emit = defineEmits([
  'copy-verse',
  'highlight-verse',
  'share',
  'add-bookmark',
  'create-note',
  'report-verse',
]);

const openMenu = (event, verse) => {
  event.preventDefault();
  selectedVerse.value = verse;
  position.value = {
    x: event.clientX,
    y: event.clientY,
  };
  isVisible.value = true;
};

const closeMenu = () => {
  isVisible.value = false;
};

const handleCopyVerse = () => {
  emit('copy-verse', selectedVerse.value);
  closeMenu();
};

const handleHighlightVerse = () => {
  emit('highlight-verse', selectedVerse.value);
  closeMenu();
};

const handleShare = () => {
  emit('share', selectedVerse.value);
  closeMenu();
};

const handleAddBookmark = () => {
  emit('add-bookmark', selectedVerse.value);
  closeMenu();
};

const handleCreateNote = () => {
  emit('create-note', selectedVerse.value);
  closeMenu();
};

const handleReport = () => {
  emit('report-verse', selectedVerse.value);
  closeMenu();
};

// Expose the openMenu function to parent components
defineExpose({
  openMenu,
  closeMenu,
});
</script>

<style lang="css" scoped>
.context-menu {
  position: fixed;
  background: #f5f1e8;
  border: 1px solid #c9973a;
  border-radius: 6px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
  z-index: 1000;
  min-width: 160px;
  overflow: hidden;
}

.context-menu-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 12px;
  cursor: pointer;
  transition: background 0.2s ease;
  color: #3a3a3a;
  font-size: 0.9rem;
}

.context-menu-item:hover {
  background: #e8dcc8;
}

.context-menu-item:active {
  background: #dccfb8;
}

.menu-icon {
  font-size: 1rem;
}

.context-menu-divider {
  height: 1px;
  background: #c9973a;
  opacity: 0.3;
  margin: 4px 0;
}
</style>
