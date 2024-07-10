<script>
    export let blog;
</script>
      <article class="mx-auto my-4 flex lg:mx-0 mx-2 flex-col overflow-hidden rounded-2xl border border-gray-300 bg-white text-gray-900 transition hover:translate-y-2 hover:shadow-lg">
        <a href="{blog.linkDetailBlog}">
          <img src="{blog.samllImg}" class="h-56 w-full object-cover" alt="" />
          <div class="flex-auto px-6 py-5">
            <span class="mb-2 flex items-center text-sm font-semibold">
              {blog.category}</span
            >
            <h3 class="text-base sm:text-md md:text-lg font-bold">{blog.title}</h3>
            <p class="mb-4 text-sm lg:text-[15px] font-light">{blog.description}.</p>
            <span class="inline-block cursor-pointer select-none rounded-md border border-violet-900 bg-violet-900 px-2 py-1 text-center align-middle text-[12px] font-semibold leading-normal text-white no-underline shadow-sm">Baca Lebih Lanjut &rarr;</span>
          </div>
        </a>
      </article>
  