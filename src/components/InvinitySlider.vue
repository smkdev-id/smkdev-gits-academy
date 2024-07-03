<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'

const props = defineProps<{ items: string[] }>()
const slider = ref<HTMLDivElement | null>(null)
const offset = ref(0)
const speed = ref(2)

let animationFrame: number | null = null

const displayedItems = [...props.items, ...props.items]

const animate = () => {
  if (!slider.value) return

  offset.value -= speed.value

  if (Math.abs(offset.value) >= slider.value.scrollWidth / 2) {
    offset.value = 0
  }

  animationFrame = requestAnimationFrame(animate)
}

onMounted(() => {
  animationFrame = requestAnimationFrame(animate)
})

onBeforeUnmount(() => {
  if (animationFrame) cancelAnimationFrame(animationFrame)
})
</script>
<template>
  <div class="overflow-hidden relative w-full">
    <div
      ref="slider"
      class="flex transition-transform duration-300 ease-linear"
      :style="{ transform: `translateX(${offset}px)` }"
    >
      <div
        v-for="(item, index) in displayedItems"
        :key="index"
        class="p-4 min-w-[150px] flex-shrink-0"
      >
        <slot :item="item" />
      </div>
    </div>
  </div>
</template>
