<script setup lang="ts">
import { computed, ref, watchEffect, watch } from "vue";
import type {
  ItemCategory,
  SearchItemModalData,
  Team,
  WeaponSkill,
} from "../models";
import TeamCharacters from "../components/TeamCharacters.vue";
import TeamSummons from "../components/TeamSummons.vue";
import TeamWeapons from "../components/TeamWeapons.vue";
import { useAPI } from "../composables/useAPI";
import { useRouter } from "vue-router";
import { getEmptyTeam } from "../utils";
import SearchItemModal from "../components/SearchItemModal.vue";
import { usePost } from "../composables/usePost";
import { useItems } from "../composables/useItems";
import EditTeamModal from "../components/EditTeamModal.vue";
import { useFetch } from "../composables/useFetch";

const props = defineProps<{
  teamId: String;
}>();

const router = useRouter();
const {
  data: teamData,
  // error: teamError,
  loading: teamLoading,
  statusCode,
} = useFetch<Team>(`/api/teams/${props.teamId}`);
const { call: apiCall } = useAPI();
const { post, error } = usePost();
const { items } = useItems();
const team = computed(() => (teamData.value ? teamData.value : getEmptyTeam()));

watchEffect(() => {
  if (statusCode.value === 404) {
    router.replace({ name: "NotFound" });
  }
});

const isModalOpen = ref(false);
const selectedData = ref<SearchItemModalData | null>(null);

watch(isModalOpen, () => {
  if (!isModalOpen.value) selectedData.value = null;
});

const handleSelect = (idx: number, category: ItemCategory) => {
  selectedData.value = { idx, category };
  isModalOpen.value = true;
};

const handleSetItem = async (itemId: string) => {
  if (selectedData.value === null) return;

  const { category, idx: position } = selectedData.value;

  if (itemId === "clear") {
    clearItem(category, position);
  } else {
    setItem(itemId, category, position);
  }

  isModalOpen.value = false;
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
    console.error(selectedData.value);
    console.error(error.value);
    alert(`error`);
  } else if (items.value) {
    const newItem = {
      ...items.value[category][itemId],
      position: position,
      skills: new Array(4).fill(null),
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
    console.error(selectedData.value);
    console.error(error.value);
    alert(`error`);
    return;
  }

  team.value[category][position] = null;
};

const handleUpdateTeam = (name: string, url: string) => {
  team.value.name = name;
  team.value.url = url;
};

const handleDeleteTeam = async () => {
  if (!confirm("Delete this team?")) {
    return;
  }

  const ok = await apiCall(`/api/teams/${team.value.id}`, { method: "DELETE" });
  if (!ok) {
    console.error(error);
    alert("Error deleting this team. Please try again later!");
  } else router.push({ name: "TeamsTable" });
};

const isSkillModalOpen = ref(false);
const selectedSkill = ref<{
  slot: number;
  position: number;
  itemId: string;
} | null>(null);

const handleSelectSkill = (slot: number, position: number) => {
  if (!team.value) return;

  const weaponId = team.value.weapon[slot]?.id;
  if (!weaponId) return;

  selectedSkill.value = { slot, position, itemId: weaponId };
  isSkillModalOpen.value = true;
};

const handleSetSkill = async (skillId: string) => {
  if (selectedSkill.value === null) return;

  const { position, slot } = selectedSkill.value;

  if (skillId === "clear") {
    clearSkill(slot, position);
  } else {
    setSkill(skillId, slot, position);
  }

  isSkillModalOpen.value = false;
};

const {
  call: skillCall,
  error: skillError,
  data: skillData,
} = useAPI<WeaponSkill>();

const setSkill = async (skillId: string, slot: number, position: number) => {
  const ok = await skillCall(
    `/api/teams/${props.teamId}/weapons/${slot}/skills/${position}`,
    {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ skill_id: skillId }),
    },
  );

  if (!ok) {
    console.error(selectedSkill.value);
    console.error(skillError.value);
    alert(`error`);
  } else if (items.value) {
    if (team.value.weapon[slot])
      team.value.weapon[slot].skills[position] = skillData.value;
  }
};

const clearSkill = async (slot: number, position: number) => {
  const ok = await skillCall(
    `/api/teams/${props.teamId}/weapons/${slot}/skills/${position}`,
    {
      method: "DELETE",
    },
  );

  if (!ok) {
    console.error(selectedSkill.value);
    console.error(skillError.value);
    alert(`error`);
    return;
  }

  if (team.value.weapon[slot]) team.value.weapon[slot].skills[position] = null;
};
</script>

<template>
  <div class="my-3">
    <div class="flex items-center gap-2">
      <h1 class="text-xl">{{ team.name }}</h1>
      <EditTeamModal v-if="team" :team="team" @update="handleUpdateTeam" />
      <UButton @click="handleDeleteTeam" color="error" variant="outline"
        >Delete</UButton
      >
    </div>

    <div class="text-sm text-muted">
      <span> Source: </span>
      <a v-if="team.url" :href="team.url">{{ team.url }}</a>
    </div>
  </div>

  <div v-if="teamLoading">Loading...</div>

  <div v-else-if="team" class="flex flex-wrap gap-2 items-start">
    <TeamCharacters
      :characters="team.character"
      @select="(idx) => handleSelect(idx, 'character')"
    />
    <TeamWeapons
      :weapons="team.weapon"
      @select="(idx) => handleSelect(idx, 'weapon')"
      @selectSkillPos="handleSelectSkill"
    />
    <TeamSummons
      :summons="team.summon"
      @select="(idx) => handleSelect(idx, 'summon')"
    />
  </div>

  <SearchItemModal
    v-model:open="isModalOpen"
    v-if="selectedData"
    :data="selectedData"
    @set="handleSetItem"
  />
  <SearchSkillModal
    v-model:open="isSkillModalOpen"
    v-if="selectedSkill"
    :weaponId="selectedSkill.itemId"
    @set="handleSetSkill"
  />
</template>
