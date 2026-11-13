<script setup lang="ts">
import type { WeaponSkill } from "../models";

const props = defineProps<{
  skill: WeaponSkill | null;
}>();
</script>

<template>
  <img
    v-if="skill"
    :src="`/files/${skill.img}`"
    :alt="skill.name"
    :title="skill.name"
  />
</template>
