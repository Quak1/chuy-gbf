<script setup lang="ts">
const props = defineProps<{
  text: string;
}>();
</script>

<template>
  <UTooltip
    :text="text"
    :delay-duration="0"
    :content="{ side: 'right' }"
    arrow
    :ui="{
      content: 'max-w-xs h-auto',
      text: 'whitespace-normal',
    }"
  >
    <UIcon name="lucide:message-square-text" class="text-blue-300" />
  </UTooltip>
</template>
