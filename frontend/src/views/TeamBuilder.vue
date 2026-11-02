<script setup lang="ts">
import { computed, watchEffect } from "vue";
import type { Team } from "../models";
import TeamCharacters from "../components/TeamCharacters.vue";
import TeamSummons from "../components/TeamSummons.vue";
import TeamWeapons from "../components/TeamWeapons.vue";
import { useAPI } from "../composables/useAPI";
import { useRouter } from "vue-router";
import { getEmptyTeam } from "../utils";

const props = defineProps<{
  teamId: String;
}>();

const router = useRouter();
const { call, data, error, loading, statusCode } = useAPI<Team>();

const team = computed(() => {
  return props.teamId !== undefined ? data.value : getEmptyTeam();
});

watchEffect(() => {
  if (props.teamId !== undefined) {
    call(`/api/teams/${props.teamId}`, {});
  }
});

watchEffect(() => {
  if (error.value && statusCode.value === 404) {
    router.replace({ name: "NotFound" });
  }
});

const handleSelect = (idx: number) => {
  console.log({ idx });
};
</script>

<template>
  <h1>Team Builder</h1>
  <div v-if="team">
    <TeamCharacters :characters="team.characters" @select="handleSelect" />
    <TeamSummons :summons="team.summons" @select="handleSelect" />
    <TeamWeapons :weapons="team.weapons" @select="handleSelect" />
  </div>
</template>
