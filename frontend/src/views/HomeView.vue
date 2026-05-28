<template>
  <div class="view-home">

    <div class="verse-of-day">
      <div class="votd-label">Verse of the Day</div>
      <blockquote class="votd-text">
        "For God so loved the world that he gave his one and only Son, that whoever
        believes in him shall not perish but have eternal life."
      </blockquote>
      <cite class="votd-ref">— John 3:16 (NIV)</cite>
    </div>

    <div class="home-grid">
      <div class="home-card" @click="router.push({ name: 'read' })">
        <div class="card-icon">📖</div>
        <h3>Continue Reading</h3>
        <p>Genesis · Chapter 3</p>
        <div class="card-progress">
          <div class="progress-bar"><div class="progress-fill" style="width: 12%"></div></div>
          <span>12% complete</span>
        </div>
      </div>

      <div class="home-card" @click="router.push({ name: 'notes' })">
        <div class="card-icon">📝</div>
        <h3>Recent Notes</h3>
        <p>3 notes this week</p>
        <div class="card-meta">Last edited · 2h ago</div>
      </div>

      <div class="home-card" @click="router.push({ name: 'bookmarks' })">
        <div class="card-icon">🔖</div>
        <h3>Bookmarks</h3>
        <p>14 saved passages</p>
        <div class="card-meta">Across 7 books</div>
      </div>

      <div class="home-card" @click="router.push({ name: 'plans' })">
        <div class="card-icon">📅</div>
        <h3>Reading Plan</h3>
        <p>Day 42 of 365</p>
        <div class="card-progress">
          <div class="progress-bar"><div class="progress-fill" style="width: 11.5%"></div></div>
          <span>11% complete</span>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { onMounted } from 'vue';
import { useRouter } from 'vue-router'
import HomeRightPanel from '@/components/home/HomeRightPanel.vue';
import { useUiStore } from '@/stores/ui.js';
const router = useRouter()
const ui = useUiStore()

onMounted(() => {
  ui.setRightPanel(HomeRightPanel, {
    title: 'Today',
    // book: selectedBook.value,
    // chapter: selectedChapter.value,
    // selectedVerse: verse,   // live data from center → right
  })
})

</script>

<style scoped>
.view-home { display: flex; flex-direction: column; gap: 24px; }

/* ── Verse of the Day ───────────────────────── */
.verse-of-day {
  background: linear-gradient(135deg, #2c1f0e 0%, #4a3218 100%);
  border-radius: var(--radius);
  padding: 28px 32px;
  position: relative;
  overflow: hidden;
}
.verse-of-day::before {
  content: '❝';
  position: absolute;
  top: -10px; left: 16px;
  font-size: 7rem;
  color: rgba(201, 151, 58, 0.12);
  line-height: 1;
  font-family: Georgia, serif;
}
.votd-label {
  font-size: 0.72rem;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: var(--gold);
  margin-bottom: 14px;
  font-family: sans-serif;
}
.votd-text {
  font-size: 1.05rem;
  line-height: 1.75;
  color: var(--parchment);
  font-style: italic;
  border: none;
  padding: 0;
  margin-bottom: 14px;
  position: relative;
}
.votd-ref { color: var(--gold-light); font-size: 0.85rem; font-style: normal; }

/* ── Grid ───────────────────────────────────── */
.home-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
}
.home-card {
  background: var(--white);
  border: 1px solid var(--border-light);
  border-radius: var(--radius);
  padding: 20px;
  cursor: pointer;
  transition: transform var(--transition), box-shadow var(--transition);
}
.home-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(44, 31, 14, 0.1);
}
.card-icon { font-size: 1.6rem; margin-bottom: 10px; }
.home-card h3 { font-size: 0.95rem; color: var(--ink); margin-bottom: 4px; }
.home-card p  { font-size: 0.8rem; color: var(--ink-light); font-style: italic; }

.card-progress { margin-top: 12px; }
.progress-bar {
  height: 4px;
  background: var(--parchment-dark);
  border-radius: 2px;
  overflow: hidden;
  margin-bottom: 4px;
}
.progress-fill {
  height: 100%;
  background: var(--gold);
  border-radius: 2px;
}
.card-progress span,
.card-meta { font-size: 0.72rem; color: var(--ink-light); font-family: sans-serif; }

@media (max-width: 480px) {
  .home-grid { grid-template-columns: 1fr 1fr; }
}
</style>