<script setup lang="ts">
import { computed } from "vue";
import type { SlotWeapon } from "../models";
import ItemImage from "./ItemImage.vue";
import EmptyItemImage from "./EmptyItemImage.vue";

const props = defineProps<{
  weapons: SlotWeapon[];
  editable?: boolean;
}>();

const emit = defineEmits<{
  select: [position: number];
  selectSkillPos: [slot: number, position: number];
}>();

const mainWeapon = computed(() => props.weapons.slice(1, 2)[0]);
const weapons = computed(() => props.weapons.slice(2, 11));
</script>

<template>
  <div class="container">
    <div class="main">
      <img src="/img/Main_Weapon_label.png" class="block w-[120px]" />

      <button @click="emit('select', 1)" class="w-[112px] h-[235px]">
        <ItemImage
          v-if="mainWeapon"
          :item="mainWeapon"
          portrait
          highRes
          disabled
        />
        <EmptyItemImage v-else category="weapon" portrait />
      </button>

      <div
        class="flex justify-center items-center gap-1 h-[40px] bg-[#19373a] rounded-xs mt-0.5 w-full"
      >
        <template v-if="mainWeapon">
          <button
            v-for="idx in 3"
            @click="$emit('selectSkillPos', 1, idx)"
            class="block w-[32px] h-[32px] bg-ink rounded-xs overflow-hidden"
          >
            <SkillImage :skill="mainWeapon.skills[idx]" />
          </button>
        </template>
      </div>
    </div>

    <div class="sub">
      <div v-for="(weapon, i) in weapons" :key="i">
        <button @click="emit('select', i + 2)" class="block w-[120px]">
          <ItemImage v-if="weapon" :item="weapon" disabled />
          <EmptyItemImage v-else category="weapon" />
        </button>

        <div
          class="flex justify-center items-center gap-1 h-[40px] bg-[#19373a] rounded-xs mt-0.5"
        >
          <template v-if="weapon">
            <button
              v-for="idx in 3"
              @click="$emit('selectSkillPos', i + 2, idx)"
              class="block w-[32px] h-[32px] bg-ink rounded-xs overflow-hidden"
            >
              <SkillImage :skill="weapon.skills[idx]" />
            </button>
          </template>
        </div>
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
  background: linear-gradient(to bottom right, #56afb2, #1e3c3c);
  box-shadow: inset 0px 0px 10px 3px black;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px;
}

button {
  padding: 0;
  border: none;
}

.main {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.sub {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 10px;
  justify-content: center;
  align-items: center;
}
</style>
