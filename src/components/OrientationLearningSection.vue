<script setup lang="ts">
import BookIcon from './icons/BookIcon.vue'
import BookOpenIcon from './icons/BookOpenIcon.vue'
import LaptopDocumentIcon from './icons/LaptopDocumentIcon.vue'

type Data = {
  title: string
  description: string
  icon: any
}

const data: Data[] = [
  {
    title: 'Project-Based Learning',
    description:
      'Kurikulum komprehensif dan senantiasa di perbarui berdasarkan pengalaman di dunia industri',
    icon: BookOpenIcon
  },
  {
    title: 'Learning Path Industri',
    description: 'Learning path disusun berdasarkan kebutuhan dunia industri.',
    icon: BookIcon
  },
  {
    title: 'Learning Path Global',
    description: 'Learning path disusun berdasarkan kebutuhan dunia industri.',
    icon: LaptopDocumentIcon
  }
]
</script>
<template>
  <div class="distance">
    <div>
      <h2 class="text-center text-3xl text-blue-950">Orientasi Belajar SMKDEV</h2>
      <div class="flex justify-center">
        <p class="text-lg mt-4 text-center max-w-prose">
          Dapatkan pengalaman belajar berorientasi pengalaman kerja yang dapat mengantarkan Anda
          menjadi talenta yang dibutuhkan oleh industri digital terkini.
        </p>
      </div>
    </div>
    <div class="flex flex-col lg:flex-row lg:gap-6 gap-0 items-center justify-center">
      <div
        v-for="item in data"
        :key="item.title"
        class="flex flex-col gap-4 mt-8 p-4 pt-0 pl-0 rounded-2xl w-full lg:w-1/3 h-52 shadow-lg"
      >
        <div class="flex gap-2 items-center">
          <div class="bg-green-100 rounded-2xl rounded-tr-none rounded-bl-none p-4">
            <component :is="item.icon" class="icon-class" />
          </div>
          <h3 class="text-xl font-medium text-blue-950 ml-2">{{ item.title }}</h3>
        </div>
        <p class="px-4 text-gray-500">{{ item.description }}</p>
      </div>
    </div>
  </div>
</template>
