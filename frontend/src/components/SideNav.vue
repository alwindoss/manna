<template>
    <div class="main-content-column-left" v-if="settingStore.settings.leftMenuVisible">
        <RouterLink class="menu-item" to="/">
            <span>
                <House />
            </span>
            <span>Home</span>
        </RouterLink>
        <RouterLink class="menu-item" to="/bible">
            <span>
                <BookOpenText />
            </span>
            <span>Read Bible</span>
        </RouterLink>
        <RouterLink class="menu-item" to="/notes">
            <span>
                <NotebookText />
            </span>
            <span>Notes</span>
        </RouterLink>
        <RouterLink class="menu-item" to="/settings">
            <span>
                <Settings />
            </span>
            <span>Settings</span>
        </RouterLink>
    </div>
</template>
<script lang="js" setup>
import { BookOpenText, House, NotebookText, Settings } from '@lucide/vue';
import { useSettingStore } from '../stores/setting';

const settingStore = useSettingStore()
</script>

<style lang="css" scoped>
.main-content-column-left {
    /* background-color: yellow; */
    width: 15%;          /* Lock the side panel width */
    flex-shrink: 0;        /* Prevent it from being squeezed */
    display: flex;
    flex-direction: column;
}

.menu-item {
    margin: 3px;
    /* top | right | bottom | left */
    padding: 10px 0 10px 10px; 
    border: 1px solid #3b82f6;
    border-radius: 5px;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
    font-size: 16px;
    background-color: #5f8fde;
    color: #ffffff;
    cursor: pointer;
    transition: all 0.2s ease;
    text-align: center;
    display: flex;
    gap: 10px;
    text-decoration: none;
}
</style>