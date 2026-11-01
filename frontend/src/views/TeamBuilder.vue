<script setup lang="ts">
import { watch } from "vue";
import { useItems } from "../composables/useItems";
import { useFetch } from "../composables/useFetch";
import type { Team } from "../models";
import TeamCharacters from "../components/TeamCharacters.vue";
import TeamSummons from "../components/TeamSummons.vue";
import TeamWeapons from "../components/TeamWeapons.vue";

const { items, error, loading } = useItems();
const { data } = useFetch<Team>("/api/teams/1");

watch(items, () => {
  console.log("=========================================================");
  console.log(items.value);
  console.log(error.value);
  console.log(loading.value);
  console.log("=========================================================");
  console.log(data.value);
});
</script>

<template>
  <h1>Team Builder</h1>
  <div v-if="data">
    <TeamCharacters :characters="data.characters" />
    <TeamSummons :summons="data.summons" />
    <TeamWeapons :weapons="data.weapons" />
  </div>
</template>
