<script setup>
import Navbar from './components/Navbar.vue';
import SideNav from './components/SideNav.vue';

</script>

<template>
  <RouterView />
</template>

<style lang="css" scoped>
.main-content {
  display: flex;
}



.main-content-column-middle {
  background-color: green;
  flex-grow: 5;
}

.main-content-column-right {
  background-color: blueviolet;
  width: 15%;
  /* Lock the side panel width */
  flex-shrink: 0;
  /* Prevent it from being squeezed */
}
</style>
