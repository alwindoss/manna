import { createRouter, createWebHashHistory } from 'vue-router'
import HomeScreen from '../screens/HomeScreen.vue'
import SettingsScreen from '../screens/SettingsScreen.vue'
import { useUserStore } from '../stores/user'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeScreen,
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