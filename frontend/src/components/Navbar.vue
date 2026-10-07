<script setup lang="ts">
import { computed } from "vue";
import { useUser } from "../composables/useUser";

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
    <RouterLink v-if="user?.role === 'admin'" to="/items"
      >Items Admin</RouterLink
    >
    <button v-if="user" @click="logout">Logout</button>
  </nav>
</template>

<style scoped>
nav {
  margin: 20px 0;
  display: flex;
  justify-content: center;
  gap: 10px;
}

a,
button {
  color: var(--accent-light);
  text-decoration: none;
  padding: 8px 10px;
  border: 1px solid var(--border);
  border-radius: 10px;
  background-color: transparent;
  font-size: 1em;
  display: flex;
  align-items: center;
  line-height: 1;
}

.router-link-exact-active {
  color: var(--accent-light);
  font-weight: bold;
  border-bottom: 2px solid var(--accent-light);
}
</style>
