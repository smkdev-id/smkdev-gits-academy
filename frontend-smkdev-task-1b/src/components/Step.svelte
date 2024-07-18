<script>
    export let step;
</script>

<a
    href={step.href}
    target="_blank"
    class="p-4 sm:p-6 md:p-8 flex flex-col gap-4 rounded-lg border border-solid border-violet-700 cursor-pointer group hover:border-violet-400 duration-200 "
>
    <div
        class="text-violet-200 grid place-items-center px-4 text-5xl md:text-6xl -mt-6 sm:-mt-6 md:-mt-8 lg:-mt-10 mx-auto duration-200 "
    >
        <i class={step.icon} />
    </div>
    <h3 class="font-medium text-lg sm:text-1xl md:text-2xl text-center mt-4">
        {step.name}
    </h3>
    <slot />
    <div class="flex-1 flex justify-betweeen gap-4 items-end">
        <div
            class="ml-auto cursor-pointer hover:text-slate-950 duration-200 relative after:absolute after:top-0 after:h-0 after:right-full after:bg-white after:w-full after:h-full after:duration-200 hover:after:translate-x-full after:z-[-1] overflow-hidden"
        >
            <p class="z-4">Go to &rarr;</p>
        </div>
    </div>
</a>