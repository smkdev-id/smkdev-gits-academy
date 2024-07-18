<template>
  <div class="distance flex flex-col gap-8 lg:flex-row mx-auto p-4 bg-blue-50 pt-8 mt-10">
    <div>
      <h1 class="text-3xl font-bold mb-4">Kontak SMKDEV</h1>
      <p class="mb-4">
        Berapa investasi untuk menggunakan SMKDEV? Anda bisa mulai dengan
        <span class="font-bold">GRATIS</span>.
      </p>
      <p class="mb-4">Silahkan isi form disebelah untuk informasi lebih lanjut.</p>
    </div>
    <div>
      <form @submit.prevent="submitForm" class="max-w-md">
        <div class="form-control mb-4">
          <label class="label" for="name">
            <span class="label-text">Nama Anda *</span>
          </label>
          <input type="text" id="name" v-model="form.name" class="input input-bordered" required />
        </div>

        <div class="form-control mb-4">
          <label class="label" for="email">
            <span class="label-text">Alamat Email *</span>
          </label>
          <input
            type="email"
            id="email"
            v-model="form.email"
            class="input input-bordered"
            required
          />
        </div>

        <div class="form-control mb-4">
          <label class="label" for="whatsapp">
            <span class="label-text">Nomor WhatsApp *</span>
          </label>
          <input
            type="tel"
            id="whatsapp"
            v-model="form.whatsapp"
            class="input input-bordered"
            required
          />
        </div>

        <div class="form-control mb-4">
          <label class="label" for="institution">
            <span class="label-text">Nama Satuan Pendidikan/Instansi *</span>
          </label>
          <input
            type="text"
            id="institution"
            v-model="form.institution"
            class="input input-bordered"
            required
          />
        </div>

        <div class="form-control mb-4">
          <label class="label" for="message">
            <span class="label-text">Pesan *</span>
          </label>
          <textarea
            id="message"
            v-model="form.message"
            class="textarea textarea-bordered"
            rows="4"
            required
          ></textarea>
        </div>

        <button type="submit" class="btn btn-primary">Kirim</button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const form = ref({
  name: '',
  email: '',
  whatsapp: '',
  institution: '',
  message: ''
})

const submitForm = () => {
  console.log('Form submitted:', form)
}
</script>
