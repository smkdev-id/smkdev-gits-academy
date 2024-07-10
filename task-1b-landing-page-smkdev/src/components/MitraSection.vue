<script setup>
import { Vue3Marquee } from 'vue3-marquee';

</script>

<template>
    <div style="background-color: aliceblue;">
        <div class="container py-3">
            <h4 class="text-center fw-semibold mb-3">Dipercaya Oleh Mitra Industri</h4>
            <Vue3Marquee id="marquee-slider" :pauseOnHover="true" :duration="15">
                <img height="50" src="@/components/assets/images/mitra.png" />
            </Vue3Marquee>
        </div>
    </div>
</template>