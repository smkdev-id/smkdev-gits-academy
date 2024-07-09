<script setup lang="ts">
import NavBar from './components/NavBar.vue'
import HeroSection from './components/HeroSection.vue'
import IndustryPatnersSection from './components/IndustryPatnersSection.vue'
import OrientationLearningSection from './components/OrientationLearningSection.vue'
import WhyChooseUsSection from './components/WhyChooseUsSection.vue'
</script>
<template>
  <NavBar />
  <main class="min-h-screen relative">
    <HeroSection />
    <IndustryPatnersSection />
    <OrientationLearningSection />
    <WhyChooseUsSection />
  </main>
</template>
