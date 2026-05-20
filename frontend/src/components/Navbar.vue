<template>
    <nav>
        <ul class="navbar">
            <li>
                <RouterLink to="/">Home</RouterLink>
            </li>
            <li>
                <RouterLink to="/settings">Settings</RouterLink>
            </li>
            <li>
                <button @click="toggleSidePanel">Toggle</button>
            </li>
        </ul>
    </nav>
</template>

<script lang="js" setup>
import { useSettingStore } from '../stores/setting';


const settingStore = useSettingStore();


const toggleSidePanel = () => {
    settingStore.toggleSidePanel()
}
</script>

<style scoped>
.navbar {
    display: flex;
    gap: 3px;
    list-style-type: none;
}
</style>
