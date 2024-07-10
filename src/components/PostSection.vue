<script setup lang="ts">
type Post = {
  title: string
  cover: string
  description: string
}

const posts: Post[] = [
  {
    title: 'Prompt Engineering: Keterampilan, Prospek Karir Dan Kursus AI',
    cover: '/posts/prompt-engineering.png',
    description:
      'Kenapa prompt engineering penting untuk AI generatif seperti ChatGPT dan Google Gemini? Pelajari keterampilan, prospek...'
  },
  {
    title: 'Panduan Pemula Menguasai AI Tools (Alat AI)',
    cover: '/posts/menguasai-tools-ai.jpeg',
    description:
      'Pelajari cara menguasai alat AI dengan panduan pemula ini. Temukan langkah-langkah, alat populer, dan tips untuk sukses dalam memanfaatkan teknologi'
  },
  {
    title: 'Strategi menurunkan Churn Rate untuk Pertumbuhan Bisnis',
    cover: '/posts/strategi-efektif-churnrate.png',
    description:
      'Memahami Churn Rate dan Strategi menurunkannya Apa itu Churn Rate? Churn rate adalah persentase pelanggan yang berhenti menggunakan produk atau...'
  }
]
</script>
<template>
  <div class="distance">
    <h2 class="text-center text-3xl font-semibold text-blue-950">Artikel Pilihan</h2>
    <div class="flex flex-wrap mt-8 gap-4">
      <div
        class="card bg-base-100 w-96 border-none rounded-none shadow-lg"
        v-for="post in posts"
        :key="post.title"
      >
        <figure>
          <img
            :src="post.cover"
            :alt="post.title"
            class="w-full h-48 object-cover object-center rounded-none"
          />
        </figure>
        <div class="card-body">
          <h2 class="card-title">{{ post.title }}</h2>
          <p>{{ post.description }}</p>
          <div class="card-actions justify-end">
            <button class="btn btn-primary">Read More</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
