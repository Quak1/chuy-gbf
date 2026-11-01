<script setup lang="ts">
import { computed } from "vue";
import type { SlotItem } from "../models";

const props = defineProps<{
  characters: SlotItem[];
  editable?: boolean;
}>();

const emit = defineEmits<{
  select: [position: number];
}>();

const mainCharacters = computed(() => props.characters.slice(1, 5));
const subCharacters = computed(() => props.characters.slice(5, 7));
</script>

<template>
  <div class="container">
    <div class="cards main">
      <span>Main</span>
      <button v-for="(char, i) in mainCharacters" :key="i">
        {{ char ? char.name : " - " }}
      </button>
    </div>
    <div class="cards sub">
      <span>Sub</span>
      <button v-for="(char, i) in subCharacters" :key="i">
        {{ char ? char.name : " - " }}
      </button>
    </div>
  </div>
</template>

<style scoped>
button {
  width: 72px;
  height: 150px;
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
