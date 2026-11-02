<script setup lang="ts">
import { computed } from "vue";
import type { SlotItem } from "../models";
import { sliceTeam } from "../utils";

const props = defineProps<{
  summons: SlotItem[];
  editable?: boolean;
}>();

const emit = defineEmits<{
  select: [position: number];
}>();

const mainSummon = computed(() => sliceTeam(props.summons, 1, 2)[0]);
const summons = computed(() => sliceTeam(props.summons, 2, 6));
const subSummons = computed(() => sliceTeam(props.summons, 6, 8));
</script>

<template>
  <div class="container">
    <div class="top">
      <button @click="emit('select', mainSummon.position)">
        {{ mainSummon.name }}
      </button>
      <div class="cards">
        <button
          v-for="(summon, i) in summons"
          :key="i"
          @click="emit('select', summon.position)"
        >
          {{ summon.name }}
        </button>
      </div>
    </div>

    <div class="bot">
      <div class="cards">
        <button
          v-for="(sub, i) in subSummons"
          :key="i"
          @click="emit('select', sub.position)"
        >
          {{ sub.name }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  width: fit-content;
  border-style: solid;
  border-width: 2px;
  border-image: linear-gradient(to bottom right, #95c9e5, #afb2a2) 1;
  margin: 10px 0;
}

button {
  width: 130px;
  height: 100px;
}

.top,
.bot {
  padding: 10px;
  box-shadow: inset 0px 0px 10px 3px black;
}

.top {
  display: flex;
  background: linear-gradient(to bottom right, #91754e, #3a2719);
  gap: 10px;

  > button {
    width: 140px;
    height: 240px;
  }
}

.bot {
  display: flex;
  justify-content: end;
  background: linear-gradient(to bottom right, #814c30, #6e3a21);
}

.cards {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 10px;
  justify-content: center;
  align-items: center;
}
</style>
