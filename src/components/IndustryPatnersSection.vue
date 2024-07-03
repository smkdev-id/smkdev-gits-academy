<script setup lang="ts">
import InvinitySlider from './InvinitySlider.vue'

const images: string[] = [
  'https://picsum.photos/100',
  'https://picsum.photos/100',
  'https://picsum.photos/100',
  'https://picsum.photos/100',
  'https://picsum.photos/100'
]
</script>

<template>
  <div class="text-center px-4 py-8 bg-blue-100">
    <h2>Dipercaya Oleh Mitra Industri</h2>
    <div class="flex overflow-x-auto pt-6">
      <InvinitySlider :items="images">
        <template #default="{ item }">
          <div class="text-center">
            <img :src="item" alt="Company Logo" class="h-16 mx-auto" />
          </div>
        </template>
      </InvinitySlider>
    </div>
  </div>
</template>
<style scoped>
.overflow-x-auto {
  scroll-snap-type: x mandatory;
  scroll-behavior: smooth;
}
</style>
