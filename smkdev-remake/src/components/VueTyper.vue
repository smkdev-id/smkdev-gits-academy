<template>
  <div>{{ displayText }}</div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'

interface Props {
  text: string | string[]
  typeSpeed?: number
  deleteSpeed?: number
  delay?: number
  loop?: boolean
}

const props = defineProps<Props>()

const displayText = ref('')
const textArray = computed(() => {
  return typeof props.text === 'string' ? [props.text] : props.text
})

const typeSpeed = props.typeSpeed ?? 100
const deleteSpeed = props.deleteSpeed ?? 50
const delay = props.delay ?? 2000
const loop = props.loop ?? true

let textIndex = 0
let charIndex = 0
let isDeleting = false
let timeout: number | undefined

const typeEffect = () => {
  const currentText = textArray.value[textIndex]
  const currentSpeed = isDeleting ? deleteSpeed : typeSpeed

  if (isDeleting) {
    displayText.value = currentText.substring(0, charIndex - 1)
    charIndex--
  } else {
    displayText.value = currentText.substring(0, charIndex + 1)
    charIndex++
  }

  if (!isDeleting && charIndex === currentText.length) {
    isDeleting = true
    timeout = setTimeout(typeEffect, delay)
  } else if (isDeleting && charIndex === 0) {
    isDeleting = false
    textIndex = (textIndex + 1) % textArray.value.length
    timeout = setTimeout(typeEffect, typeSpeed)
  } else {
    timeout = setTimeout(typeEffect, currentSpeed)
  }
}

watch(
  textArray,
  () => {
    textIndex = 0
    charIndex = 0
    isDeleting = false
    if (timeout) clearTimeout(timeout)
    typeEffect()
  },
  { immediate: true }
)
</script>
