<script setup>
import Hero from '@/components/Hero.vue';
import Superiority from '@/components/Superiority.vue';
import TalentaDigitalCycle from '@/components/TalentaDigitalCycle.vue';
import MitraSection from '@/components/MitraSection.vue';
import MentorSection from '@/components/MentorSection.vue';

</script>

<template>
  <main style="padding-top: 100px;">
    <div class="container">
      <Hero/>
    </div>
    <div class="mt-3">
      <Superiority/>
    </div>
    <div class="mt-3">
      <TalentaDigitalCycle/>
    </div>
    <div style="background-color: aliceblue;">
      <div class="mt-3">
        <MitraSection/>
      </div>
      <div class="mt-3">
        <MentorSection/>
      </div>
    </div>
  </main>
</template>
