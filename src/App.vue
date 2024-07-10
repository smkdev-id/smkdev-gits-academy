<script setup lang="ts">
import NavBar from './components/NavBar.vue'
import HeroSection from './components/HeroSection.vue'
import IndustryPatnersSection from './components/IndustryPatnersSection.vue'
import OrientationLearningSection from './components/OrientationLearningSection.vue'
import WhyChooseUsSection from './components/WhyChooseUsSection.vue'
import DigitalTalentSection from './components/DigitalTalentSection.vue'
import TestimoniSection from './components/TestimoniSection.vue'
import MentorsSection from './components/MentorsSection.vue'
import PostSection from './components/PostSection.vue'
import ContactSection from './components/ContactSection.vue'
import FooterSection from './components/FooterSection.vue'
</script>
<template>
  <NavBar />
  <main class="min-h-screen relative">
    <HeroSection />
    <IndustryPatnersSection />
    <OrientationLearningSection />
    <WhyChooseUsSection />
    <DigitalTalentSection />
    <TestimoniSection />
    <MentorsSection />
    <PostSection />
    <ContactSection />
    <FooterSection />
  </main>
</template>
