<script setup lang="ts">
import { RouterLink, useRouter } from "vue-router";
import { useFetch } from "../composables/useFetch";
import { usePost } from "../composables/usePost";
import type { Team } from "../models";
import { computed } from "vue";
import ItemsImageList from "../components/ItemsImageList.vue";

const router = useRouter();
const { data, error, loading } = useFetch<Team[]>("/api/teams");
const {
  post,
  data: postData,
  error: postError,
} = usePost<{ team_id: number }>();

const handleCreateBtn = async () => {
  const ok = await post("/api/teams", {});
  if (ok) {
    router.push(`/teams/builder/${postData.value?.team_id}`);
  } else {
    console.error(postError.value);
  }
};

const tableData = computed(() => data.value ?? []);
</script>

<template>
  <button @click="handleCreateBtn">Create new team</button>

  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>

  <div v-else class="container">
    <template v-for="row in tableData" class="row">
      <RouterLink :to="`/teams/builder/${row.id}`">{{ row.name }}</RouterLink>
      <div class="images">
        <div class="sub characters">
          <p>Characters</p>
          <ItemsImageList :items="row.character" category="character" />
        </div>
        <div class="sub summons">
          <p>Summons</p>
          <ItemsImageList :items="row.summon" category="summon" />
        </div>
        <div class="sub weapons">
          <p>Weapons</p>
          <ItemsImageList :items="row.weapon" category="weapon" />
        </div>
      </div>
    </template>
  </div>
</template>

<style scoped>
.container {
  display: grid;
  grid-template-columns: 1fr 3fr;
  align-items: center;
  gap: 10px;
}

.images {
  display: flex;
  flex-direction: column;

  border-radius: 10px;
  background-color: var(--border);
}

.sub {
  display: flex;
  align-items: center;
  padding: 5px;

  &:not(:last-child) {
    border-bottom: 3px solid var(--accent);
  }

  > p {
    width: 100px;
    flex-shrink: 0;
  }
}

.characters {
}

.summons {
}

.weapons {
}
</style>
