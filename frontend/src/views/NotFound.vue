<script setup lang="ts"></script>

<template>
  <UError
    :error="{
      status: 404,
      statusText: 'Page not found',
      message: 'The page you are looking for does not exist.',
    }"
    :clear="{
      to: '/',
    }"
  />
</template>
