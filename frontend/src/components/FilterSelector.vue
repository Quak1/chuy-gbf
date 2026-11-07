<script setup lang="ts">
const props = defineProps<{
  items: any;
  loading: boolean;
}>();

const model = defineModel<any[] | undefined>();
const toggleModel = defineModel<boolean | undefined>("include");
</script>

<template>
  <div class="flex flex-col justify-center items-center gap-1">
    <USelectMenu
      v-model="model"
      :items="items"
      :loading="loading"
      variant="subtle"
      virtualize
      multiple
      clear
      class="w-60"
    />
    <USwitch
      v-model="toggleModel"
      :label="toggleModel ? 'Include' : 'Exclude'"
    />
  </div>
</template>
