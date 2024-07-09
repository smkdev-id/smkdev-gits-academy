<script>
    export let testimoni;
</script>

<figure class="flex flex-col items-center justify-center p-8 mx-4 lg:mx-0 text-center bg-white border-b border-gray-200 rounded-t-lg md:rounded-t-none md:rounded-ss-lg md:border-e dark:bg-white">
    <blockquote class="max-w-2xl mx-auto mb-4 text-dark  lg:mb-8 dark:text-drak">
        <p class="my-4">"{testimoni.message}"</p>
    </blockquote>
    <figcaption class="flex items-center justify-center ">
        <img class="rounded-full w-9 h-9" src="/images/profile.png" alt="{testimoni.name}">
        <div class="space-y-0.5 font-medium text-left rtl:text-right ms-3">
            <div class="font-semibold ">{testimoni.name}</div>
            <div class="text-sm  ">{testimoni.job}</div>
        </div>
    </figcaption>    
</figure>