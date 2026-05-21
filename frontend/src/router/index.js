// src/router/index.js
import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ReadView from '../views/ReadView.vue'
import NotesView from '../views/NotesView.vue'
import PlaceholderView from '../views/PlaceholderView.vue'


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
    component: ReadView,
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
    component: PlaceholderView,
    meta: { title: 'Bookmarks', subtitle: 'Saved passages', icon: '🔖' },
  },
  {
    path: '/plans',
    name: 'plans',
    component: PlaceholderView,
    meta: { title: 'Reading Plans', subtitle: 'Stay on track', icon: '📅' },
  },
  {
    path: '/search',
    name: 'search',
    component: PlaceholderView,
    meta: { title: 'Search Scripture', subtitle: 'Find any passage', icon: '⌕' },
  },
  {
    path: '/settings',
    name: 'settings',
    component: PlaceholderView,
    meta: { title: 'Settings', subtitle: 'Preferences & display', icon: '⚙' },
  },
  {
    path: '/profile',
    name: 'profile',
    component: PlaceholderView,
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