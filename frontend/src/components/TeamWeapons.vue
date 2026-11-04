<script setup lang="ts">
import { computed } from "vue";
import type { SlotItem } from "../models";
import { sliceTeam } from "../utils";
import ItemImage from "./ItemImage.vue";
import EmptyItemImage from "./EmptyItemImage.vue";

const props = defineProps<{
  weapons: SlotItem[];
  editable?: boolean;
}>();

const emit = defineEmits<{
  select: [position: number];
}>();

const mainWeapon = computed(() => sliceTeam(props.weapons, 1, 2)[0]);
const weapons = computed(() => sliceTeam(props.weapons, 2, 11));
</script>

<template>
  <div class="container">
    <div class="main">
      <img src="/img/Main_Weapon_label.png" />
      <button @click="emit('select', mainWeapon.position)">
        <ItemImage
          v-if="'type' in mainWeapon"
          :item="mainWeapon"
          portrait
          highRes
          disabled
        />
        <EmptyItemImage v-else category="weapon" portrait />
      </button>
    </div>

    <div class="sub">
      <button
        v-for="(weapon, i) in weapons"
        :key="i"
        @click="emit('select', weapon.position)"
      >
        <ItemImage v-if="'type' in weapon" :item="weapon" disabled />
        <EmptyItemImage v-else category="weapon" />
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
  padding: 0;
  border: none;
}

.main {
  display: flex;
  flex-direction: column;
  align-items: center;

  img {
    display: block;
    width: 120px;
  }
  button {
    width: 112px;
    height: 235px;
  }
}

.sub {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 10px;
  justify-content: center;
  align-items: center;
}
</style>
