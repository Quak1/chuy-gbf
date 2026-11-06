<script setup lang="ts">
import Navbar from "./components/Navbar.vue";
</script>

<template>
  <UApp>
    <Navbar />
    <UMain>
      <UContainer>
        <RouterView />
      </UContainer>
    </UMain>
  </UApp>
</template>
