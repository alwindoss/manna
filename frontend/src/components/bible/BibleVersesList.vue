<template>
    <div v-for="verse in props.verses" :key="verse.num" class="verse"
        :class="{ highlighted: ui.isVerseHighlighted(verse.num) }" @click="onClickOfVerse(verse.num)">
        <span class="verse-num">{{ verse.num }}</span>
        <span class="verse-text">{{ verse.text }}</span>
    </div>
</template>
<script setup>
import { useUiStore } from '@/stores/ui';

const ui = useUiStore()

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
  ui.toggleVerseHighlight(num)
  emit('verseNumEvent', num)
  
}

console.log("Verses:", props.verses)
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