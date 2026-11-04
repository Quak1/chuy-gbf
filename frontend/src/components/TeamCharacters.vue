<script setup lang="ts">
import { computed } from "vue";
import type { SlotItem } from "../models";
import { sliceTeam } from "../utils";
import ItemImage from "./ItemImage.vue";
import EmptyItemImage from "./EmptyItemImage.vue";

const props = defineProps<{
  characters: SlotItem[];
  editable?: boolean;
}>();

const emit = defineEmits<{
  select: [position: number];
}>();

const sections = computed(() => [
  { name: "Main", characters: sliceTeam(props.characters, 1, 5) },
  { name: "Sub", characters: sliceTeam(props.characters, 5, 7) },
]);
</script>

<template>
  <div class="container">
    <div
      v-for="section in sections"
      class="cards"
      :class="section.name.toLowerCase()"
    >
      <span>{{ section.name }}</span>
      <button
        v-for="(char, i) in section.characters"
        :key="i"
        @click="$emit('select', char.position)"
      >
        <ItemImage v-if="'type' in char" :item="char" portrait disabled />
        <EmptyItemImage v-else category="character" portrait />
      </button>
    </div>
  </div>
</template>

<style scoped>
button {
  width: 72px;
  height: 150px;
  padding: 0;
  border: none;
}

.container {
  padding: 15px;
  display: flex;
  gap: 20px;
  width: fit-content;

  background: linear-gradient(to bottom, #4391a7, #173033);
  box-shadow:
    inset 0 0 2px 3px #e5ffff,
    inset 0 0 6px 5px #59eaff,
    1px 3px 4px #77776f;
}

.cards {
  background: linear-gradient(to bottom, #093d4d, #001e28);
  border: solid 1px #356c79;
  padding: 10px;
  display: flex;
  gap: 5px;
  position: relative;

  span {
    position: absolute;
    top: -8px;
    left: 0px;

    font-size: 16px;
    -webkit-text-stroke: 5px black;
    paint-order: stroke fill;
    font-family: Arial, Helvetica, sans-serif;
    user-select: none;
  }

  &.main {
    color: #a3d9e6;
  }

  &.sub {
    color: #f5e7c1;
  }
}
</style>
