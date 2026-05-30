<template>
  <div class="section">
    <S type="group" label="Language">
      <S type="row" label="Interface language"
         hint="The language used for menus and UI text.">
        <select class="s-select" v-model="uiLanguage">
          <option value="en">English</option>
          <option value="fr">Français</option>
          <option value="es">Español</option>
          <option value="de">Deutsch</option>
          <option value="pt">Português</option>
          <option value="hi">हिन्दी</option>
          <option value="ta">தமிழ்</option>
        </select>
      </S>
    </S>
    <S type="group" label="Region & Formats">
      <S type="row" label="Date format">
        <select class="s-select" v-model="dateFormat">
          <option value="DD/MM/YYYY">DD/MM/YYYY</option>
          <option value="MM/DD/YYYY">MM/DD/YYYY</option>
          <option value="YYYY-MM-DD">YYYY-MM-DD</option>
        </select>
      </S>
      <S type="row" label="First day of week"
         hint="Affects the weekly reading streak calendar.">
        <select class="s-select" v-model="weekStart">
          <option value="sunday">Sunday</option>
          <option value="monday">Monday</option>
          <option value="saturday">Saturday</option>
        </select>
      </S>
    </S>
    <S type="group" label="Bible Language">
      <S type="row" label="Preferred Bible language"
         hint="Filters available translations to your preferred language.">
        <select class="s-select" v-model="bibleLanguage">
          <option value="en">English</option>
          <option value="fr">French</option>
          <option value="es">Spanish</option>
          <option value="de">German</option>
          <option value="pt">Portuguese</option>
          <option value="all">Show all languages</option>
        </select>
      </S>
    </S>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import S from './SettingPrimitives.vue'
const uiLanguage   = ref('en')
const dateFormat   = ref('DD/MM/YYYY')
const weekStart    = ref('sunday')
const bibleLanguage = ref('en')
</script>
<style scoped>.section{display:flex;flex-direction:column;}</style>