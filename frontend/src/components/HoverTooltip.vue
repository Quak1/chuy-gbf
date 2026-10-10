<script setup lang="ts"></script>

<template>
  <div class="tooltip">
    <slot name="outer" />
    <div class="tooltiptext">
      <slot />
    </div>
  </div>
</template>

<style scoped>
.tooltip {
  position: relative;
  display: inline-block;
  cursor: pointer;
}

.tooltiptext {
  visibility: hidden;
  width: 300px;
  max-width: 300px;
  background-color: var(--border);
  border: 2px solid var(--accent);
  box-sizing: border-box;
  color: rgba(255, 255, 255, 0.8);
  text-wrap: wrap;

  position: absolute;
  border-radius: 10px;
  padding: 10px;
  z-index: 1;

  top: -15px;
  left: 105%;
}

.tooltip:hover .tooltiptext {
  visibility: visible;
}
</style>
