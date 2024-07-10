<script setup lang="ts">
type Mentor = {
  name: string
  title: string
  corp: string
  image: string
}

const mentors: Mentor[] = [
  {
    name: 'Ibnu Sina Wardy',
    title: 'CEO',
    corp: 'GITS Indonesia',
    image: '/mentors/ibnu-sina-wardy.jpg'
  },
  {
    name: 'Sudaryatno',
    title: 'CTO',
    corp: 'GITS Indonesia',
    image: '/mentors/sudaryatno.jpeg'
  },
  {
    name: 'Helmi Adi Prasetyo',
    title: 'Backend Developer',
    corp: 'SMKDEV',
    image: '/mentors/helmi-adi-prasetyo.png'
  },
  {
    name: 'Rachmawati A. Taurisia',
    title: 'Curriculum Developer',
    corp: 'SMKDEV',
    image: '/mentors/rachmawati-ari-taurisia.jpg'
  },
  {
    name: 'Samuel Pandohan',
    title: 'Curriculum Developer',
    corp: 'SMKDEV',
    image: '/mentors/samuel-pandohan-terampil-gultom.jpg'
  }
]
</script>
<template>
  <div class="flex flex-col justify-center items-center py-8">
    <h1 class="text-3xl font-semibold text-blue-950 text-center mb-4">Mentor SMKDEV</h1>
    <div class="flex flex-row flex-wrap justify-center items-center gap-10">
      <div
        v-for="mentor in mentors"
        :key="mentor.name"
        class="flex flex-col gap-2 items-center mt-4"
      >
        <img :src="mentor.image" alt="Mentor" width="200" height="200" class="rounded-full" />
        <p class="text-xl font-medium">{{ mentor.name }}</p>
        <div class="divider text-black my-0"></div>
        <p>{{ mentor.title }}</p>
        <p>{{ mentor.corp }}</p>
      </div>
    </div>
  </div>
</template>
