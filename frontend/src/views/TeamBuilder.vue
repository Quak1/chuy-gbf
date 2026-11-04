<script setup lang="ts">
import { ref, watchEffect } from "vue";
import type { ItemCategory, SearchItemModalData, Team } from "../models";
import TeamCharacters from "../components/TeamCharacters.vue";
import TeamSummons from "../components/TeamSummons.vue";
import TeamWeapons from "../components/TeamWeapons.vue";
import { useAPI } from "../composables/useAPI";
import { useRouter } from "vue-router";
import { getEmptyTeam } from "../utils";
import { useModal } from "../composables/useModal";
import SearchItemModal from "../components/SearchItemModal.vue";
import { usePost } from "../composables/usePost";
import { useItems } from "../composables/useItems";

const props = defineProps<{
  teamId: String;
}>();

const router = useRouter();
const {
  call,
  data: teamData,
  error: teamError,
  loading: teamLoading,
  statusCode,
} = useAPI<Team>();
const { selectedData, openModal, closeModal } = useModal<SearchItemModalData>();
const { post, error } = usePost();
const { items } = useItems();

const team = ref(getEmptyTeam());

watchEffect(async () => {
  await call(`/api/teams/${props.teamId}`, {});
  if (teamData.value) {
    team.value = teamData.value;
  }
});

watchEffect(() => {
  if (teamError.value && statusCode.value === 404) {
    router.replace({ name: "NotFound" });
  }
});

const handleSelect = (idx: number, category: ItemCategory) => {
  openModal({ idx, category });
};

const handleSetItem = async (itemId: string) => {
  if (selectedData.value === null) return;

  const category = selectedData.value.category;
  const position = selectedData.value.idx;

  if (itemId === "clear") {
    clearItem(category, position);
  } else {
    setItem(itemId, category, position);
  }

  closeModal();
};

const setItem = async (
  itemId: string,
  category: ItemCategory,
  position: number,
) => {
  const ok = await post(`/api/teams/${props.teamId}/slot`, {
    category: category,
    position: position,
    item_id: itemId,
  });

  if (!ok) {
    console.log(selectedData.value);
    console.log(error.value);
    alert(`error`);
  } else if (items.value) {
    const newItem = {
      ...items.value[category][itemId],
      position: position,
    };

    team.value[category][position] = newItem;
  }
};

const clearItem = async (category: ItemCategory, position: number) => {
  const ok = await post(`/api/teams/${props.teamId}/slot/clear`, {
    category: category,
    position,
  });

  if (!ok) {
    console.log(selectedData.value);
    console.log(error.value);
    alert(`error`);
    return;
  }

  team.value[category][position] = null;
};
</script>

<template>
  <h1>{{ team.name }}</h1>
  <a v-if="team.url" :href="team.url">{{ team.url }}</a>

  <div v-if="teamLoading">Loading...</div>

  <div v-else-if="team">
    <TeamCharacters
      :characters="team.character"
      @select="(idx) => handleSelect(idx, 'character')"
    />
    <TeamSummons
      :summons="team.summon"
      @select="(idx) => handleSelect(idx, 'summon')"
    />
    <TeamWeapons
      :weapons="team.weapon"
      @select="(idx) => handleSelect(idx, 'weapon')"
    />
  </div>

  <SearchItemModal
    v-if="selectedData"
    :data="selectedData"
    @close="closeModal"
    @set="handleSetItem"
  />
</template>
