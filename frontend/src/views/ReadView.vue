<template>
  <div class="view-read">

    <div class="bible-nav">
      <select class="bible-select" v-model="selectedBook">
        <option v-for="b in books" :key="b" :value="b">{{ b }}</option>
      </select>
      <select class="bible-select" v-model="selectedChapter">
        <option v-for="c in chapters" :key="c" :value="c">Chapter {{ c }}</option>
      </select>
      <select class="bible-select" v-model="selectedTranslation">
        <option>NIV</option>
        <option>ESV</option>
        <option>KJV</option>
        <option>NLT</option>
      </select>
    </div>

    <div class="scripture-body">
      <h2 class="scripture-heading">{{ selectedBook }} · Chapter {{ selectedChapter }}</h2>
      <div class="verses">
        <div
          v-for="verse in sampleVerses"
          :key="verse.num"
          class="verse"
          :class="{ highlighted: ui.isVerseHighlighted(verse.num) }"
          @click="ui.toggleVerseHighlight(verse.num)"
        >
          <span class="verse-num">{{ verse.num }}</span>
          <span class="verse-text">{{ verse.text }}</span>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useUiStore } from '../stores/ui'

const ui = useUiStore()

const selectedBook        = ref('Genesis')
const selectedChapter     = ref(1)
const selectedTranslation = ref('NIV')

const books    = ['Genesis','Exodus','Psalms','Proverbs','Matthew','John','Romans','Revelation']
const chapters = Array.from({ length: 50 }, (_, i) => i + 1)

const sampleVerses = [
  { num: 1,  text: 'In the beginning God created the heavens and the earth.' },
  { num: 2,  text: 'Now the earth was formless and empty, darkness was over the surface of the deep, and the Spirit of God was hovering over the waters.' },
  { num: 3,  text: 'And God said, "Let there be light," and there was light.' },
  { num: 4,  text: 'God saw that the light was good, and he separated the light from the darkness.' },
  { num: 5,  text: 'God called the light "day," and the darkness he called "night." And there was evening, and there was morning—the first day.' },
  { num: 6,  text: 'And God said, "Let there be a vault between the waters to separate water from water."' },
  { num: 7,  text: 'So God made the vault and separated the water under the vault from the water above it. And it was so.' },
  { num: 8,  text: 'God called the vault "sky." And there was evening, and there was morning—the second day.' },
  { num: 9,  text: 'And God said, "Let the water under the sky be gathered to one place, and let dry ground appear." And it was so.' },
  { num: 10, text: 'God called the dry ground "land," and the gathered waters he called "seas." And God saw that it was good.' },
]
</script>

<style scoped>
.view-read { display: flex; flex-direction: column; gap: 20px; }

/* ── Nav bar ────────────────────────────────── */
.bible-nav {
  display: flex;
  gap: 10px;
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
.bible-select:focus { outline: 2px solid var(--gold); outline-offset: 1px; }

/* ── Scripture ──────────────────────────────── */
.scripture-heading {
  font-size: 1.4rem;
  color: var(--ink);
  margin-bottom: 20px;
  padding-bottom: 14px;
  border-bottom: 1px solid var(--border-light);
  letter-spacing: 0.02em;
}
.verse {
  display: flex;
  gap: 14px;
  padding: 10px 12px;
  border-radius: 6px;
  cursor: pointer;
  line-height: 1.8;
  transition: background var(--transition);
}
.verse:hover       { background: var(--parchment-mid); }
.verse.highlighted { background: rgba(201, 151, 58, 0.15); }

.verse-num {
  font-size: 0.72rem;
  font-family: sans-serif;
  color: var(--gold);
  font-weight: 700;
  min-width: 20px;
  padding-top: 5px;
  flex-shrink: 0;
}
.verse-text { font-size: 1rem; color: var(--ink-mid); }
</style>