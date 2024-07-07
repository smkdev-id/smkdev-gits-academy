<script>
    import Step from "./Step.svelte"
    let slogan = [
        {
            name: "Jadilah Talenta Digital",
            runTxt: "Masa Depan Indonesia",
            description: "Belajar langsung dengan expert dari industri dengan kurikulum komperhensif berbasis project-based learning"
        }
    ]
</script>

<main class="flex flex-col flex-1 p-4 ">
    <section
    id="introPage"
    class="grid grid-cols-1 lg:grid-cols-2 gap-10 py-8 sm:py-14"
>
    <div
        class="flex flex-col lg:justify-center text-center lg:text-left gap-6 md:gap-8 lg:gap-10"
    >
        <h4 class="font-semibold text-4xl sm:text-4xl ">
            {slogan[0].name}
            <span class="poppins text-violet-400">{slogan[0].runTxt}</span>
        </h4>
        <p class="text-base sm:text-lg md:text-lg">
            {slogan[0].description}
        </p>
    </div>
    <div class="relative grid place-items-end">
        <img
            src={"images/profile.png"}
            alt="Zetane Engine"
            class="object-cover z-[2] max-h-[70vh]"
        />
    </div>
</section>
</main>