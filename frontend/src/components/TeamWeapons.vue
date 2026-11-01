<script setup lang="ts">
import { computed } from "vue";
import type { SlotItem } from "../models";

const props = defineProps<{
  weapons: SlotItem[];
  editable?: boolean;
}>();

const emit = defineEmits<{
  select: [position: number];
}>();

const mainWeapon = computed(() => props.weapons[1]);
const weapons = computed(() => props.weapons.slice(2, 11));
</script>

<template>
  <div class="container">
    <div class="main">
      <button>
        {{ mainWeapon ? mainWeapon.name : " - " }}
      </button>
    </div>

    <div class="sub">
      <button v-for="(weapon, i) in weapons" :key="i">
        {{ weapon ? weapon : " - " }}
      </button>
    </div>
  </div>
</template>

<style scoped>
.container {
  width: fit-content;
  border-style: solid;
  border-width: 2px;
  border-image: linear-gradient(to bottom right, #95c9e5, #afb2a2) 1;
  background: linear-gradient(to bottom right, #56afb2, #1e3c3c);
  box-shadow: inset 0px 0px 10px 3px black;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px;
}

button {
  width: 110px;
  height: 105px;
}

.main button {
  width: 112px;
  height: 235px;
}

.sub {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 10px;
  justify-content: center;
  align-items: center;
}
</style>
