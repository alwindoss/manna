<script setup>
import Navbar from './components/Navbar.vue';
import SideNav from './components/SideNav.vue';

</script>

<template>
  <header>
    <Navbar></Navbar>
  </header>
  <main>
    <div class="main-content">
      <!-- Left Side Panel -->
      <SideNav></SideNav>


      <!-- Middle Panel -->
      <div class="main-content-column-middle">
        <RouterView />
      </div>

      <!-- Right Side Panel -->
      <div class="main-content-column-right"></div>
    </div>
  </main>
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
  width: 240px;
  /* Lock the side panel width */
  flex-shrink: 0;
  /* Prevent it from being squeezed */
}
</style>
