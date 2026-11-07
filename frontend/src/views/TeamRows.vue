<script setup lang="ts">
import { RouterLink, useRouter } from "vue-router";
import { useFetch } from "../composables/useFetch";
import { usePost } from "../composables/usePost";
import { useItems } from "../composables/useItems";
import type { Team } from "../models";
import { computed, ref } from "vue";
import ItemsImageList from "../components/ItemsImageList.vue";
import { flattenItemsObject, isIn } from "../utils";

const router = useRouter();
const { data, error, loading } = useFetch<Team[]>("/api/teams");
const {
  post,
  data: postData,
  error: postError,
} = usePost<{ team_id: number }>();
const { items, loading: itemsLoading } = useItems();

const handleCreateBtn = async () => {
  const ok = await post("/api/teams", {});
  if (ok) {
    router.push(`/teams/builder/${postData.value?.team_id}`);
  } else {
    console.error(postError.value);
  }
};

const weapons = ref<{ name: string; id: string }[]>([]);
const includeWeapons = ref(true);
const summons = ref([]);
const includeSummons = ref(true);
const characters = ref([]);
const includeCharacters = ref(true);
const selectItems = computed(() => {
  if (!items.value) return { characters: [], summons: [], weapons: [] };
  return {
    characters: flattenItemsObject(items.value.character),
    summons: flattenItemsObject(items.value.summon),
    weapons: flattenItemsObject(items.value.weapon),
  };
});

const tableData = computed(() => {
  if (!data.value) {
    return [];
  }

  const filtered = data.value.filter((team) => {
    const characterFilter = isIn(
      includeCharacters.value,
      characters.value,
      team.character,
    );
    const summonFilter = isIn(includeSummons.value, summons.value, team.summon);
    const weaponFilter = isIn(includeWeapons.value, weapons.value, team.weapon);
    return characterFilter && summonFilter && weaponFilter;
  });

  return filtered;
});
</script>

<template>
  <div class="flex flex-wrap gap-2 my-5">
    <UButton @click="handleCreateBtn">Create new team</UButton>

    <FilterSelector
      v-model="characters"
      v-model:include="includeCharacters"
      :items="selectItems.characters"
      :loading="itemsLoading"
    />
    <FilterSelector
      v-model="weapons"
      v-model:include="includeWeapons"
      :items="selectItems.weapons"
      :loading="itemsLoading"
    />
    <FilterSelector
      v-model="summons"
      v-model:include="includeSummons"
      :items="selectItems.summons"
      :loading="itemsLoading"
    />
  </div>

  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>

  <div
    v-else
    class="grid grid-cols-[1fr_3fr] items-center gap-2 bg-muted rounded-md p-3"
  >
    <template v-for="row in tableData" class="row">
      <RouterLink :to="`/teams/builder/${row.id}`" class="text-lg underline">{{
        row.name
      }}</RouterLink>

      <div class="flex flex-col bg-accented rounded-sm p-3 pt-1">
        <USeparator label="Characters" position="start" />
        <ItemsImageList :items="row.character" category="character" />

        <USeparator label="Summons" position="start" />
        <ItemsImageList :items="row.summon" category="summon" />

        <USeparator label="Weapon" position="start" />
        <ItemsImageList :items="row.weapon" category="weapon" />
      </div>
    </template>
  </div>
</template>
