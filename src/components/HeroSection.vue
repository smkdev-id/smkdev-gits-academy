<script setup lang="ts">
import VueTyper from './VueTyper.vue'
import GraphicsIcon from './icons/GraphicsIcon.vue'
</script>
<template>
  <div
    class="distance flex flex-col md:flex-row justify-center items-center gap-14 text-start py-16 lg:h-[90vh]"
  >
    <div class="w-full md:w-1/2 mt-12">
      <h1 class="text-4xl font-medium text-blue-950 leading-tight">
        Jadilah Talenta Digital
        <VueTyper class="text-blue-800" :text="['Masa Depan Indonesia', 'Global']" />
      </h1>
      <p class="mt-6 font-medium text-lg">
        Belajar langsung dengan expert dari industri dengan kurikulum komperhensif berbasis
        project-based learning
      </p>
    </div>
    <div class="mt-6 w-full md:w-1/2 px-6 relative">
      <GraphicsIcon class="absolute -z-10 -left-2 md:left-10 -top-4" />
      <img
        src="@/assets/hero-image.jpg"
        alt="learn new skills ilustration"
        class="rounded-[30px] md:m-6 md:ml-16 mt-10 z-10 w-full md:w-[400px]"
        width="450"
      />
    </div>
  </div>
</template>
