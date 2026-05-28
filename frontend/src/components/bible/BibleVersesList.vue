<template>
  <div>
    <div v-for="verse in props.verses" :key="verse.num" class="verse"
        :class="{ highlighted: ui.isVerseHighlighted(verse.num) }" 
        @click="onClickOfVerse(verse.num)"
        @contextmenu="onContextMenuOfVerse($event, verse)">
        <span class="verse-num">{{ verse.num }}</span>
        <span class="verse-text">{{ verse.text }}</span>
    </div>
    <ReadBibleContextMenu
      ref="contextMenuRef"
      @copy-verse="handleCopyVerse"
      @highlight-verse="handleHighlightVerse"
      @share="handleShare"
      @add-bookmark="handleAddBookmark"
      @create-note="handleCreateNote"
      @report-verse="handleReportVerse"
    />
  </div>
</template>
<script setup>
import { ref } from 'vue';
import { useUiStore } from '@/stores/ui';
import ReadBibleContextMenu from './ReadBibleContextMenu.vue';

const ui = useUiStore();
const contextMenuRef = ref(null);

const props = defineProps({
    verses: {
    type: Array,
    required: true,
    default: () => []
  }
})

const emit = defineEmits([
  'verseNumEvent',
])

const onClickOfVerse = (num) => {
  const verseNum = num
  ui.selectVerseHighlight(verseNum)
  emit('verseNumEvent', verseNum)
}

const onContextMenuOfVerse = (event, verse) => {
  ui.selectVerseHighlight(verse.num)
  contextMenuRef.value?.openMenu(event, verse)
}

const handleCopyVerse = (verse) => {
  console.log('Copy verse:', verse);
}

const handleHighlightVerse = (verse) => {
  console.log('Highlight verse:', verse);
}

const handleShare = (verse) => {
  console.log('Share verse:', verse);
}

const handleAddBookmark = (verse) => {
  console.log('Add bookmark:', verse);
}

const handleCreateNote = (verse) => {
  console.log('Create note:', verse);
}

const handleReportVerse = (verse) => {
  console.log('Report verse:', verse);
}

</script>
<style lang="css" scoped>
.verse {
  display: flex;
  gap: 14px;
  padding: 10px 12px;
  border-radius: 6px;
  cursor: pointer;
  line-height: 1.8;
  transition: background var(--transition);
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
</style>