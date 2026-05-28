<template>
  <div class="view-read">

    <div class="bible-nav">
      <select class="bible-select" @change="updateChaptersAndVerses" v-model="selectedBook">
        <option v-for="b in books" :key="b" :value="b">{{ b }}</option>
      </select>
      <select class="bible-select" @change="updateChaptersAndVerses" v-model="selectedChapter">
        <option v-for="c in chapters" :key="c" :value="c">Chapter {{ c }}</option>
      </select>
      <select class="bible-select" @change="updateChaptersAndVerses" v-model="selectedTranslation">
        <option v-for="(translation, tindex) in listOfTranslationsAvailable" :key="tindex">{{ translation }}</option>

        <!-- <option>KJV</option>
        <option>NIV</option>
        <option>ESV</option>
        <option>NLT</option> -->
      </select>
    </div>

    <div class="scripture-body">
      <h2 class="scripture-heading">{{ selectedBook }} · Chapter {{ selectedChapter }}</h2>
      <div class="verses">
        <BibleVersesList :verses="versesInReadView"></BibleVersesList>
      </div>
    </div>

  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useUiStore } from '@/stores/ui'
import { GetBooksOfTheBible, GetCountOfChaptersInTheBook, GetListOfTranslationsAvailable, GetVerses, ShowError, ShowWarning } from '../../bindings/github.com/dailymanna/manna/internal/bible/bibleservice'
import { Application } from '@wailsio/runtime'
import BibleVersesList from '@/components/bible/BibleVersesList.vue'

const ui = useUiStore()

const selectedBook = ref('Genesis')
const selectedChapter = ref(1)
const selectedTranslation = ref('KJV')
const defaultTranslation = ref('KJV')
const listOfTranslationsAvailable = ref([])

// const books    = ['Genesis','Exodus','Psalms','Proverbs','Matthew','John','Romans','Revelation']
var books = ref([])
// const chapters = Array.from({ length: 50 }, (_, i) => i + 1)
var chapters = ref([])

var numOfChapters = ref(0)

const versesInReadView = ref([])

const updateChaptersAndVerses = () => {
  console.log("Selected Translation:", selectedTranslation.value)
  GetCountOfChaptersInTheBook(selectedBook.value).then((data) => {
    numOfChapters.value = data
    chapters.value = Array.from({ length: data }, (_, i) => i + 1)
    GetVerses(selectedTranslation.value, selectedBook.value, selectedChapter.value).then((verses) => {
      versesInReadView.value = verses
    }).catch((err) => {
      console.log("Error in getting verses:", err)
      ShowWarning("Bible version not found", `The bible version ${selectedTranslation.value} is not available. Switching to default(${defaultTranslation.value}) version.`)
      selectedTranslation.value = 'KJV'
    })
  }).catch((err) => {
    ShowError("Error occured", err)
  })
}

const fetchReadViewData = () => {

  // Fetch all the available translations
  GetListOfTranslationsAvailable().then((translations) => {
    listOfTranslationsAvailable.value = translations
  }).catch((err) => {
    ShowError("Error occured", err)
  })

  // Fetch all the books of the bible
  GetBooksOfTheBible().then((data) => {
    books.value = data
    selectedBook.value = data[0]
    updateChaptersAndVerses()
  }).catch((err) => {
    ShowError("Error occured", err)
  })
}

onMounted(() => {
  fetchReadViewData()
})
</script>

<style scoped>
.view-read {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

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

.bible-select:focus {
  outline: 2px solid var(--gold);
  outline-offset: 1px;
}

/* ── Scripture ──────────────────────────────── */
.scripture-heading {
  font-size: 1.4rem;
  color: var(--ink);
  margin-bottom: 20px;
  padding-bottom: 14px;
  border-bottom: 1px solid var(--border-light);
  letter-spacing: 0.02em;
}


</style>