<script setup lang="ts">
import MenuIcon from './icons/MenuIcon.vue'
</script>
<template>
  <header class="navbar border-b shadow px-0 md:px-20">
    <div class="dropdown md:hidden">
      <div tabindex="0" role="button" class="btn btn-ghost btn-circle">
        <MenuIcon />
      </div>
      <ul
        tabindex="0"
        class="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow-xl"
      >
        <li><a>Learn</a></li>
        <li><a>Community</a></li>
        <li><a>Blog</a></li>
      </ul>
    </div>
    <div class="flex-1">
      <a class="btn btn-ghost pl-0">
        <img src="/logo/smkdev-logo.png" alt="logo smkdev" width="120" />
      </a>
    </div>
    <div class="flex-none">
      <ul class="menu menu-horizontal px-1 pr-0 hidden md:flex items-center">
        <li><a>Learn</a></li>
        <li><a>Community</a></li>
        <li><a>Blog</a></li>
        <li class="hidden md:block"><a class="btn bg-blue-950 text-white">Register</a></li>
      </ul>
      <ul class="flex md:hidden mr-6">
        <li><a class="btn bg-blue-950 text-white">Register</a></li>
      </ul>
    </div>
  </header>
</template>
