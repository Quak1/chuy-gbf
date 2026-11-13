<script setup lang="ts">
import { computed, watchEffect } from "vue";
import { useFetch } from "../composables/useFetch";
import type { WeaponSkill } from "../models";

const props = defineProps<{
  weaponId: string;
}>();

const emit = defineEmits<{
  set: [skillId: string];
}>();

const open = defineModel<boolean>("open");

const apiUrl = computed(() => `/api/items/${props.weaponId}/skills`);
const { data, fetchData } = useFetch<WeaponSkill[]>(apiUrl.value);
watchEffect(() => fetchData(apiUrl.value));
</script>

<template>
  <UModal v-model:open="open" title="Skills">
    <template #body>
      <div class="flex justify-end">
        <UButton
          @click="emit('set', 'clear')"
          color="error"
          icon="i-lucide-trash"
          >Remove skill</UButton
        >
      </div>

      <div class="flex flex-col gap-1 my-2">
        <div
          v-for="skill in data"
          @click="emit('set', skill.id)"
          class="flex items-center gap-2 bg-muted rounded-sm capitalize cursor-pointer min-h-[50px]"
          :key="skill.id"
        >
          <SkillImage :skill="skill" />
          <div>{{ skill.name }}</div>
        </div>
      </div>
    </template>
  </UModal>
</template>
