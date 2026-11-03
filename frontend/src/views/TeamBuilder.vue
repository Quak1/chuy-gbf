<script setup lang="ts">
import { computed, watchEffect } from "vue";
import type { ItemCategory, SearchItemModalData, Team } from "../models";
import TeamCharacters from "../components/TeamCharacters.vue";
import TeamSummons from "../components/TeamSummons.vue";
import TeamWeapons from "../components/TeamWeapons.vue";
import { useAPI } from "../composables/useAPI";
import { useRouter } from "vue-router";
import { getEmptyTeam } from "../utils";
import { useModal } from "../composables/useModal";
import SearchItemModal from "../components/SearchItemModal.vue";

const props = defineProps<{
  teamId?: String;
}>();

const router = useRouter();
const { call, data, error, loading, statusCode } = useAPI<Team>();
const { selectedData, openModal, closeModal } = useModal<SearchItemModalData>();

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

const handleSelect = (idx: number, category: ItemCategory) => {
  openModal({ idx, category });
};
const handleSetItem = (itemId: string) => {
  console.log({
    idx: selectedData.value?.idx,
    category: selectedData.value?.category,
    itemId,
  });

  closeModal();
};
</script>

<template>
  <h1>Team Builder</h1>
  <div v-if="loading">Loading...</div>
  <div v-else-if="team">
    <TeamCharacters
      :characters="team.characters"
      @select="(idx) => handleSelect(idx, 'characters')"
    />
    <TeamSummons
      :summons="team.summons"
      @select="(idx) => handleSelect(idx, 'summons')"
    />
    <TeamWeapons
      :weapons="team.weapons"
      @select="(idx) => handleSelect(idx, 'weapons')"
    />
  </div>

  <SearchItemModal
    v-if="selectedData"
    :data="selectedData"
    @close="closeModal"
    @set="handleSetItem"
  />
</template>
