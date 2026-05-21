import { createRouter, createWebHashHistory } from 'vue-router'
import HomeScreen from '../screens/HomeScreen.vue'
import SettingsScreen from '../screens/SettingsScreen.vue'
import { useUserStore } from '../stores/user'
import NotesScreen from '../screens/NotesScreen.vue'
import BibleScreen from '../screens/BibleScreen.vue'
import Prototype from '../screens/Prototype.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: Prototype,
  },
  {
    path: '/notes',
    name: 'notes',
    component: NotesScreen,
  },
  {
    path: '/bible',
    name: 'bible',
    component: BibleScreen,
  },
  {
    path: '/settings',
    name: 'settings',
    component: SettingsScreen,
  },
]

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: routes,
})

router.beforeEach(async (to, from) => {
  const userDetails = useUserStore()
  console.log("UserName:", userDetails.name)

  // Otherwise, let them go where they requested (return nothing or true)
  return true
})

export default router