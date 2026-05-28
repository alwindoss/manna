// src/router/index.js
import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import NotesView from '@/views/NotesView.vue'
import PlaceholderView from '@/views/PlaceholderView.vue'
import ReadBibleView from '@/views/ReadBibleView.vue'
import PlansView from '@/views/PlansView.vue'
import SearchView from '@/views/SearchView.vue'
import SettingsView from '@/views/SettingsView.vue'
import AccountView from '@/views/AccountView.vue'
import BookmarksView from '@/views/BookmarksView.vue'


const routes = [
  {
    path: '/',
    redirect: '/home',
  },
  {
    path: '/home',
    name: 'home',
    component: HomeView,
    meta: { title: 'Good Morning, James', subtitle: 'Continue your journey', icon: '⌂' },
  },
  {
    path: '/read',
    name: 'read',
    component: ReadBibleView,
    meta: { title: 'Read the Bible', subtitle: 'Explore Scripture', icon: '📖' },
  },
  {
    path: '/notes',
    name: 'notes',
    component: NotesView,
    meta: { title: 'My Notes', subtitle: 'Reflections & insights', icon: '✎' },
  },
  {
    path: '/bookmarks',
    name: 'bookmarks',
    component: BookmarksView,
    meta: { title: 'Bookmarks', subtitle: 'Saved passages', icon: '🔖' },
  },
  {
    path: '/plans',
    name: 'plans',
    component: PlansView,
    meta: { title: 'Reading Plans', subtitle: 'Stay on track', icon: '📅' },
  },
  {
    path: '/search',
    name: 'search',
    component: SearchView,
    meta: { title: 'Search Scripture', subtitle: 'Find any passage', icon: '⌕' },
  },
  {
    path: '/settings',
    name: 'settings',
    component: SettingsView,
    meta: { title: 'Settings', subtitle: 'Preferences & display', icon: '⚙' },
  },
  {
    path: '/profile',
    name: 'profile',
    component: AccountView,
    meta: { title: 'Profile', subtitle: 'Your account', icon: '👤' },
  },
  // Catch-all → home
  {
    path: '/:pathMatch(.*)*',
    redirect: '/home',
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router