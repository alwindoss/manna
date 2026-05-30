<template>
  <div class="section">

    <S type="group" label="Typography">
      <S type="row" label="Font size" hint="Applies to all Bible reading views.">
        <div class="font-size-control">
          <button class="size-btn" @click="fontSize = Math.max(12, fontSize - 1)">−</button>
          <span class="size-value">{{ fontSize }}px</span>
          <button class="size-btn" @click="fontSize = Math.min(28, fontSize + 1)">+</button>
        </div>
      </S>
      <S type="row" label="Reading font" hint="Font used for Scripture text.">
        <select class="s-select" v-model="readingFont">
          <option value="Georgia">Georgia (Serif)</option>
          <option value="'Palatino Linotype'">Palatino (Serif)</option>
          <option value="'Courier New'">Courier (Mono)</option>
          <option value="sans-serif">System Sans</option>
        </select>
      </S>
      <S type="row" label="Line spacing">
        <select class="s-select" v-model="lineHeight">
          <option value="1.5">Compact</option>
          <option value="1.8">Normal</option>
          <option value="2.2">Relaxed</option>
        </select>
      </S>
    </S>

    <S type="group" label="Verse Display">
      <S type="row" label="Show verse numbers"
         hint="Display verse numbers inline with text.">
        <label class="s-toggle">
          <input type="checkbox" v-model="showVerseNumbers" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="Verse numbers position"
         :disabled="!showVerseNumbers">
        <select class="s-select" v-model="verseNumPosition" :disabled="!showVerseNumbers">
          <option value="inline">Inline (superscript)</option>
          <option value="margin">Left margin</option>
        </select>
      </S>
      <S type="row" label="Each verse on new line"
         hint="Displays each verse as a separate paragraph.">
        <label class="s-toggle">
          <input type="checkbox" v-model="versePerLine" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="Show paragraph breaks"
         hint="Preserve paragraph markers from the original text.">
        <label class="s-toggle">
          <input type="checkbox" v-model="showParagraphs" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
    </S>

    <S type="group" label="Red Letter">
      <S type="row" label="Red letter text"
         hint="Highlight the words of Jesus Christ in red.">
        <label class="s-toggle">
          <input type="checkbox" v-model="redLetter" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="Red letter colour" :disabled="!redLetter">
        <div class="color-picker-row">
          <div
            v-for="c in redLetterColors"
            :key="c"
            class="color-swatch"
            :class="{ selected: redLetterColor === c }"
            :style="{ background: c }"
            @click="redLetter && (redLetterColor = c)"
          />
        </div>
      </S>
    </S>

    <!-- Live preview -->
    <S type="group" label="Preview">
      <div
        class="reading-preview"
        :style="{
          fontFamily: readingFont,
          fontSize: fontSize + 'px',
          lineHeight: lineHeight,
        }"
      >
        <span v-if="showVerseNumbers && verseNumPosition === 'inline'" class="verse-num-preview">3 </span>
        <span :style="{ color: redLetter ? redLetterColor : 'inherit' }">
          "For God so loved the world that he gave his one and only Son,
          that whoever believes in him shall not perish but have eternal life."
        </span>
        <span v-if="showVerseNumbers && verseNumPosition === 'inline'" class="verse-num-preview"> 4 </span>
        <span> For God did not send his Son into the world to condemn the world,
          but to save the world through him.</span>
      </div>
    </S>

  </div>
</template>

<script setup>
import { ref } from 'vue'
import S from './SettingPrimitives.vue'

const fontSize       = ref(16)
const readingFont    = ref('Georgia')
const lineHeight     = ref('1.8')
const showVerseNumbers = ref(true)
const verseNumPosition = ref('inline')
const versePerLine   = ref(false)
const showParagraphs = ref(true)
const redLetter      = ref(true)
const redLetterColor = ref('#8b2e2e')
const redLetterColors = ['#8b2e2e','#c0392b','#e74c3c','#c0522b','#9b2d20']
</script>

<style scoped>
.section { display: flex; flex-direction: column; }

.font-size-control {
  display: flex;
  align-items: center;
  gap: 10px;
}
.size-btn {
  width: 28px; height: 28px;
  border: 1px solid var(--border);
  border-radius: 6px;
  background: var(--parchment-mid);
  cursor: pointer;
  font-size: 1rem;
  line-height: 1;
  color: var(--ink);
  transition: background var(--transition);
}
.size-btn:hover { background: var(--parchment-dark); }
.size-value { font-size: 0.85rem; color: var(--ink); min-width: 36px; text-align: center; font-family: sans-serif; }

.color-picker-row { display: flex; gap: 8px; }
.color-swatch {
  width: 22px; height: 22px;
  border-radius: 50%;
  cursor: pointer;
  border: 2px solid transparent;
  transition: border-color var(--transition), transform var(--transition);
}
.color-swatch:hover { transform: scale(1.15); }
.color-swatch.selected { border-color: var(--gold); transform: scale(1.15); }

.reading-preview {
  padding: 18px 20px;
  color: var(--ink-mid);
  background: var(--parchment);
  border-radius: 6px;
  margin: 4px;
  line-height: 1.8;
  transition: font-size 0.15s ease, line-height 0.15s ease;
}
.verse-num-preview {
  font-size: 0.7em;
  color: var(--gold);
  font-weight: 700;
  font-family: sans-serif;
  vertical-align: super;
}
</style>