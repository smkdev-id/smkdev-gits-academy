<script setup lang="ts">
const logos = [
  '/logo/DIGITS-logo.png',
  '/logo/Gits-Indonesia-logo.png',
  '/logo/Eudeka-logo.png',
  '/logo/Arkana-logo.png',
  '/logo/Mantab-One-logo.png',
  '/logo/Hitopia-logo.png'
]
</script>

<template>
  <div class="text-center px-4 py-8 bg-blue-100 mt-16">
    <h2 class="text-2xl font-semibold">Dipercaya Oleh Mitra Industri</h2>
    <div class="flex flex-row flex-wrap justify-center items-center gap-4">
      <div v-for="logo in logos" :key="logo">
        <img :src="logo" alt="logo" width="150" />
      </div>
    </div>
  </div>
</template>
