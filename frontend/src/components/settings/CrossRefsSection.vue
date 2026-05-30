<!-- CrossRefsSection.vue -->
<template>
  <div class="section">
    <S type="group" label="Cross Reference Sources">
      <S type="row" v-for="src in sources" :key="src.id"
         :label="src.name" :hint="src.description">
        <label class="s-toggle">
          <input type="checkbox" v-model="src.enabled" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
    </S>
    <S type="group" label="Display">
      <S type="row" label="Max cross-references shown per verse">
        <select class="s-select" v-model="maxRefs">
          <option value="3">3</option>
          <option value="5">5</option>
          <option value="10">10</option>
          <option value="0">All</option>
        </select>
      </S>
      <S type="row" label="Show cross-ref preview text"
         hint="Display a snippet of the referenced verse inline.">
        <label class="s-toggle">
          <input type="checkbox" v-model="showPreview" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="Show Strong's numbers"
         hint="Display original language word numbers for study.">
        <label class="s-toggle">
          <input type="checkbox" v-model="strongsNumbers" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="Show original language"
         hint="Display Hebrew (OT) and Greek (NT) words in the study panel.">
        <label class="s-toggle">
          <input type="checkbox" v-model="showOriginal" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
    </S>
    <S type="group" label="Commentary">
      <S type="row" label="Enable commentary panel">
        <label class="s-toggle">
          <input type="checkbox" v-model="commentaryEnabled" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="Commentary source" :disabled="!commentaryEnabled">
        <select class="s-select" v-model="commentarySource" :disabled="!commentaryEnabled">
          <option value="matthew_henry">Matthew Henry</option>
          <option value="jamieson">Jamieson-Fausset-Brown</option>
          <option value="gill">John Gill</option>
        </select>
      </S>
    </S>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import S from './SettingPrimitives.vue'
const maxRefs           = ref('5')
const showPreview       = ref(true)
const strongsNumbers    = ref(false)
const showOriginal      = ref(true)
const commentaryEnabled = ref(true)
const commentarySource  = ref('matthew_henry')
const sources = ref([
  { id: 'treasury',  name: 'Treasury of Scripture Knowledge', description: 'Classic comprehensive cross-reference system.', enabled: true  },
  { id: 'nave',      name: "Nave's Topical Bible",            description: 'Topically organised reference system.',         enabled: false },
  { id: 'torrey',    name: "Torrey's New Topical Textbook",   description: 'Systematic topical references.',                enabled: false },
  { id: 'openscriptures', name: 'OpenScriptures',            description: 'Community-maintained reference dataset.',       enabled: true  },
])
</script>
<style scoped>.section{display:flex;flex-direction:column;}</style>