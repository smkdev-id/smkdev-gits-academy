<script setup lang="ts">
import ReadBookIcon from './icons/ReadBookIcon.vue'
import PencilIcon from './icons/PencilIcon.vue'
import MentorIcon from './icons/MentorIcon.vue'
</script>
<template>
  <div
    class="flex flex-col-reverse md:flex-row distance justify-center items-center gap-8 py-8 w-full"
  >
    <div class="flex flex-col w-full lg:w-1/2 gap-4">
      <div class="flex gap-6 items-center border rounded-xl shadow-lg p-4 px-6 h-32">
        <ReadBookIcon class="w-10 h-10 fill-blue-600" />
        <div>
          <h3 class="font-medium text-2xl text-blue-950 mb-2">Kurikulum</h3>
          <p>Kurikulum berstandar industri global, senantiasa diperbaharui</p>
        </div>
      </div>
      <div>
        <div class="flex gap-6 items-center border rounded-xl shadow-lg p-4 px-6 h-32">
          <PencilIcon class="w-10 h-10 fill-orange-600" />
          <div>
            <h3 class="font-medium text-2xl text-blue-950 mb-2">Metode Belajar</h3>
            <p>Metode belajar project-based learning, dapat menjadi portfolio siswa</p>
          </div>
        </div>
      </div>
      <div>
        <div class="flex gap-6 items-center border rounded-xl shadow-lg p-4 px-6 h-32">
          <MentorIcon class="w-10 h-10 fill-green-600" />
          <div>
            <h3 class="font-medium text-2xl text-blue-950 mb-2">Mentor</h3>
            <p>Dibimbing langsung oleh mentor expert dari industri digital</p>
          </div>
        </div>
      </div>
    </div>
    <div class="w-full lg:w-1/2">
      <h2 class="text-start text-3xl font-semibold text-blue-950">
        MENGAPA HARUS MEMILIH BELAJAR DI <span class="text-blue-600">SMKDEV?</span>
      </h2>
      <p class="mt-4 text-gray-600 px-8 text-lg">
        Pengalaman belajarmu jauh lebih menyenangkan dengan perpaduan belajar teori dan capstone
        project yang dapat menjadi portofolio pribadimu. Tidak lupa, Kamu juga dapat belajar
        langsung dari expert mentor dari dunia industri.
      </p>
    </div>
  </div>
</template>
