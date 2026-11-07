<script setup lang="ts">
const props = defineProps<{
  items: any;
  loading: boolean;
  label: string;
}>();

const model = defineModel<any[] | undefined>();
const toggleModel = defineModel<boolean | undefined>("include");
</script>

<template>
  <div class="flex flex-col justify-center items-center gap-1">
    <div class="flex items-center gap-3">
      <p>{{ label }}:</p>
      <USwitch
        v-model="toggleModel"
        :label="toggleModel ? 'Include' : 'Exclude'"
        :ui="{ label: 'text-muted' }"
      />
    </div>
    <USelectMenu
      v-model="model"
      :items="items"
      :loading="loading"
      variant="subtle"
      virtualize
      multiple
      clear
      class="w-60"
      :placeholder="label"
    />
  </div>
</template>
