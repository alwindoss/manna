<template>
  <div class="section">
    <S type="group" label="Daily Reading">
      <S type="row" label="Daily reading reminder"
         hint="A prompt to read your daily plan chapter.">
        <label class="s-toggle">
          <input type="checkbox" v-model="dailyReminder" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="Reminder time" :disabled="!dailyReminder">
        <input class="s-input" type="time" v-model="reminderTime" style="width:120px"
               :disabled="!dailyReminder" />
      </S>
      <S type="row" label="Streak reminder"
         hint="Reminds you before midnight if your streak is at risk.">
        <label class="s-toggle">
          <input type="checkbox" v-model="streakReminder" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
    </S>

    <S type="group" label="Manna Hub Activity">
      <S type="row" v-for="n in hubNotifs" :key="n.id"
         :label="n.label" :hint="n.hint">
        <label class="s-toggle">
          <input type="checkbox" v-model="n.enabled" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
    </S>

    <S type="group" label="Delivery">
      <S type="row" label="In-app notifications">
        <label class="s-toggle">
          <input type="checkbox" v-model="inApp" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="System / OS notifications">
        <label class="s-toggle">
          <input type="checkbox" v-model="systemNotif" />
          <span class="s-toggle-track"><span class="s-toggle-thumb"></span></span>
        </label>
      </S>
      <S type="row" label="Notification sound">
        <select class="s-select" v-model="sound">
          <option value="none">None</option>
          <option value="soft">Soft chime</option>
          <option value="bell">Bell</option>
        </select>
      </S>
    </S>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import S from './SettingPrimitives.vue'
const dailyReminder  = ref(true)
const reminderTime   = ref('07:00')
const streakReminder = ref(true)
const inApp          = ref(true)
const systemNotif    = ref(true)
const sound          = ref('soft')
const hubNotifs = ref([
  { id: 'likes',    label: 'Likes on your posts',      hint: 'When someone likes your post or note.',  enabled: true  },
  { id: 'comments', label: 'Comments',                 hint: 'Replies to your posts and notes.',       enabled: true  },
  { id: 'follows',  label: 'New followers',            hint: 'When someone follows you on the Hub.',   enabled: true  },
  { id: 'mentions', label: 'Mentions',                 hint: 'When someone tags you in a post.',       enabled: true  },
  { id: 'plans',    label: 'Reading plan invitations', hint: 'When someone invites you to a plan.',    enabled: false },
])
</script>
<style scoped>.section{display:flex;flex-direction:column;}</style>