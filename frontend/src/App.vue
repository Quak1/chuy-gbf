<script setup lang="ts"></script>

<template>
  <nav>
    <RouterLink to="/">Home</RouterLink>
    <RouterLink to="/hello">Hello</RouterLink>
    <RouterLink to="/login">Login</RouterLink>
  </nav>
  <RouterView />
</template>
