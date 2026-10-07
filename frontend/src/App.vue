<script setup lang="ts">
import { computed } from "vue";
import { useUser } from "./composables/useUser";

const { user, logout } = useUser(false);
const itemsURL = computed(() => {
  return `/users/${user.value?.id}/items`;
});
</script>

<template>
  <nav>
    <RouterLink to="/">Home</RouterLink>
    <RouterLink v-if="user" :to="itemsURL">Items</RouterLink>
    <RouterLink v-if="!user" to="/login">Login</RouterLink>
    <RouterLink v-if="user" to="/items">Items Admin</RouterLink>
    <button v-if="user" @click="logout">Logout</button>
  </nav>
  <RouterView />
</template>

<style>
nav > * {
  margin: 10px;
}
</style>
