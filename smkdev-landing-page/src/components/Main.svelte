<script>
    import Step from "./Step.svelte";
    import { onMount } from 'svelte';
    // Import Swiper bundle with all modules installed
    import Swiper from 'swiper/bundle';

    // Import styles bundle
    import 'swiper/css/bundle';

    let slogan = [
        {
            name: "Jadilah Talenta Digital",
            description: "Belajar langsung dengan expert dari industri dengan kurikulum komperhensif berbasis project-based learning"
        },
        {
            name: "Why SMKDEV",
            description: "Pengalaman belajarmu jauh lebih menyenangkan dengan perpaduan belajar teori dan capstone project yang dapat menjadi portofolio pribadimu. Tidak lupa, Kamu juga dapat belajar langsung dari expert mentor dari dunia industri."
        }
    ];

    let steps = [
        {
            name: "Learning Path Industri",
            icon: "fa-solid fa-book fa-xs",
            description:
                "Learning path disusun berdasarkan kebutuhan dunia industri.",
        },
        {
            name: "Kurikulum Komprehensif",
            icon: "fa-solid fa-business-time fa-xs",
            description:
                "Kurikulum komprehensif dan senantiasa di perbarui berdasarkan pengalaman di dunia industri",
        },
        {
            name: "Project-Based Learning",
            icon: "fa-solid fa-laptop fa-xs",
            description:
                "Materi pembelajaran disusun dengan pendekatan - pendekatan project based learning",
        },
    ];

    let methodSmkdev = [
        {
            title: "Kurikulum",
            icon: "fa-solid fa-globe fa-2x",
            description: "Kurikulum berstandar industri global, senantiasa diperbaharui"
        },
        {
            title: "Metode",
            icon: "fa-solid fa-pencil fa-2x",
            description: "Metode belajar project-based learning, dapat menjadi portfolio siswa"
        },
        {
            title: "Mentor",
            icon: "fa-solid fa-person-chalkboard",
            description: "Dibimbing langsung oleh mentor expert dari industri digital"
        }
    ]
    let talents = [
        {
            title: "Kurikulum",
            icon: "fa-solid fa-globe fa-2x",
            description: "Kurikulum berstandar industri global, senantiasa diperbaharui"
        },
        {
            title: "Metode",
            icon: "fa-solid fa-pencil fa-2x",
            description: "Metode belajar project-based learning, dapat menjadi portfolio siswa"
        },
        {
            title: "Mentor",
            icon: "fa-solid fa-person-chalkboard",
            description: "Dibimbing langsung oleh mentor expert dari industri digital"
        }
    ]

    onMount(() => {
        const swiper = new Swiper(".centered-slide-carousel", {
        centeredSlides: true,
        paginationClickable: true,
        loop: true,
        spaceBetween: 30,
        slideToClickedSlide: true,
        pagination: {
            el: ".centered-slide-carousel .swiper-pagination",
            clickable: true,
        },
        breakpoints: {
            
            1028: {
            slidesPerView: 5,
            spaceBetween: 10
            },
            0: {
            slidesPerView: 3,
            spaceBetween: 0
            }
        }
      });
    });
</script>

<main class="flex flex-col flex-1">
    <!-- Hero -->
    <section
        id="introPage"
        class="grid grid-cols-1 lg:grid-cols-2 gap-30 py-8 sm:py-14 px-6 lg:px-[300px]"
    >
        <div
            class="flex flex-col lg:justify-center text-center lg:text-left gap-6 md:gap-8 lg:gap-10"
        >
            <h4 class="font-semibold text-4xl sm:text-4xl">
                {slogan[0].name}
                <span class="poppins text-violet-400">
                    <span id="typewriter"></span><span id="cursor">|</span>
                </span>
            </h4>
            <p class="text-base sm:text-lg md:text-lg">
                {slogan[0].description}
            </p>
        </div>
        <div class="relative grid place-items-center">
            <img
                src={"images/profile.png"}
                alt="Zetane Engine"
                class="object-cover z-[1] max-h-[70vh] lg:h-max h-screen"
            />
        </div>
    </section>
    <!-- Mitra -->
    <section class="py-3 lg:py-5 flex flex-col gap-2 bg-violet-50" id="mitra">
        <div class="flex flex-col items-center gap-2 text-center">
            <h6 class="font-semibold text-base sm:text-lg md:text-lg">
                Dipercayai Oleh Mitra Industri
            </h6>
        </div>
        <div class="flex justify-center bg-violet-50 px-6 lg:px-[300px]">
            <div class="w-full relative">
                <div class="swiper centered-slide-carousel swiper-container relative ">
                    <div class="swiper-wrapper h-max relative">
                        <div class="swiper-slide">
                            <div class="bg-violet-50 rounded-2xl h-36 flex justify-center items-center">
                                <img src="https://git.gits.id/uploads/-/system/appearance/logo/1/Logo-Main__1_.png" alt="GITS" class="max-h-10"/>
                            </div>
                        </div>
                        <div class="swiper-slide">
                            <div class="bg-violet-50 rounded-2xl h-36 flex justify-center items-center">
                                <img src="https://eudeka.storage.googleapis.com/wp/logo-eudeka-512x512-1-300x300.png" alt="Eudeka" class="max-h-24"/>
                            </div>
                        </div>
                        <div class="swiper-slide">
                            <div class="bg-violet-50 rounded-2xl h-36 flex justify-center items-center">
                                <img src="images/arkana.png" alt="Arkana" class="h-8"/>
                            </div>
                        </div>
                        <div class="swiper-slide">
                            <div class="bg-violet-50 rounded-2xl h-36 flex justify-center items-center">
                                <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSd-YWd6j6ZfTCJ7X8pv_n96QnZQtd8qAwwfQ&s" alt="violantis" class="max-h-7"/>
                            </div>
                        </div>
                        <div class="swiper-slide">
                            <div class="bg-violet-50 rounded-2xl h-36 flex justify-center items-center">
                                <img src="https://hitopia.id/wp-content/uploads/2023/04/Logo-Deep-Hitopia-Brown-min-Copy.png" alt="Hiptopia" class="max-h-16"/>
                            </div>
                        </div>
                        <div class="swiper-slide">
                            <div class="bg-violet-50 rounded-2xl h-36 flex justify-center items-center">
                                <img src="https://mantab.one/wp-content/uploads/2024/05/logo.png" alt="Mantabone" class="max-h-12"/>
                            </div>
                        </div>
                    </div>
                    <div class="swiper-pagination "></div>
                </div>
            </div>    
        </div>
    </section>
    <!-- Kurikulum SMKDEV -->
    <section class="py-10 lg:py-28 flex flex-col gap-16 px-6 lg:px-[300px]" id="learning">
        <div class="flex flex-col gap-2 text-center">
            <h3 class="font-semibold text-3xl sm:text-4xl">
                Orientasi Belajar <span class="poppins text-violet-400">SMKDEV</span> 
            </h3>
            <h6 class="text-base sm:text-lg md:text-lg px-6 lg:px-[150px]">
                Dapatkan pengalaman belajar berorientasi pengalaman kerja yang dapat mengantarkan Anda menjadi talenta yang dibutuhkan oleh industri digital terkini.
            </h6>
        </div>
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-12 lg:gap-10">
            <Step step={steps[0]}>
                <p>
                    {steps[0].description}
                </p>
            </Step>
            <Step step={steps[1]}>
                <p>
                    {steps[1].description}
                </p>
            </Step>
            <Step step={steps[2]}>
                <p>
                    {steps[2].description}
                </p>
            </Step>
        </div>
    </section >
    <!-- Why SMKDEV? -->
    <section class="grid grid-cols-1 lg:grid-cols-2 gap-25 py-8 sm:py-10 lg:-top-10 top-0  px-6 lg:px-[300px] relative" id="smkdev">
        <div class="relative flex flex-col rounded-xl bg-white bg-clip-border text-gray-700 shadow-md lg:mr-10 order-2 sm:order-1" id="card">
            <div class="m-auto px-4">
                <div class="divide-y divide-gray-200">
                    <div class="flex items-center justify-between pb-5 pt-3 last:pb-0">
                        <div class="flex items-center gap-x-3">
                            <div class="relative inline-block rounded-md bg-violet-100 flex items-center justify-center h-11 w-11">
                                <i class="{methodSmkdev[0].icon} fa-2x "></i>
                            </div>
                            <div>
                                <h6 class="block text-xl font-semibold text-blue-gray-900 antialiased">
                                    {methodSmkdev[0].title}
                                </h6>
                                <p class="block font-sans text-sm text-gray-700 antialiased">
                                    {methodSmkdev[0].description}
                                </p>
                            </div>
                        </div>
                    </div>
                    <div class="flex items-center justify-between pb-5 pt-3 last:pb-0">
                        <div class="flex items-center gap-x-3">
                            <div class="relative inline-block rounded-md bg-violet-100 flex items-center justify-center h-11 w-11">
                                <i class="{methodSmkdev[1].icon} fa-2x "></i>
                            </div>
                            <div>
                                <h6 class="block text-xl font-semibold text-blue-gray-900 antialiased">
                                    {methodSmkdev[1].title}
                                </h6>
                                <p class="block font-sans text-sm text-gray-700 antialiased">
                                    {methodSmkdev[1].description}
                                </p>
                            </div>
                        </div>
                    </div>
                    <div class="flex items-center justify-between mb-3 pt-3 last:pb-0">
                        <div class="flex items-center gap-x-3">
                            <div class="relative inline-block rounded-md bg-violet-100 flex items-center justify-center h-11 w-11">
                                <i class="{methodSmkdev[2].icon} fa-2x "></i>
                            </div>
                            <div>
                                <h6 class="block text-xl font-semibold text-blue-gray-900 antialiased">
                                    {methodSmkdev[2].title}
                                </h6>
                                <p class="block font-sans text-sm text-gray-700 antialiased">
                                    {methodSmkdev[2].description}
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="flex flex-col lg:justify-center text-center lg:mb-0 mb-7 lg:text-left gap-6 md:gap-8 lg:gap-10 order-1 sm:order-2" id="text">
            <h4 class="font-semibold text-4xl sm:text-4xl">
                {slogan[1].name}
                <span class="poppins text-violet-400" id="cursor">
                    ?
                </span>
            </h4>
            <p class="text-base sm:text-lg md:text-lg">
                {slogan[1].description}
            </p>
        </div>
    </section>    
    <!-- About -->
    <section
    id="about"
    class="py-8 lg:py-10 pt-7 lg:pt-9 flex flex-col gap-10 sm:gap-20 md:gap-24 relative lg:px-[300px]">
    <div class="flex flex-col gap-2 text-center relative  before:absolute before:top-0 before:left-0 before:w-2/3 before:h-1.5 before:bg-violet-400 after:absolute after:bottom-0 after:right-0 after:w-2/3 after:h-1.5 after:bg-violet-400 py-4">
        <h3 class="font-semibold text-xl sm:text-2xl md:text-4xl">
            Talenta Digital <span class="text-violet-400">SMKDEV</span> 
        </h3>
    </div>
    <div class="flex flex-col gap-20 w-full bg-indigo-100 grid place-items-center">
        {#each talents as talent, index}
            <div class="flex gap-2 sm:gap-5">
                <p
                    class="poppins text-xl sm:text-2xl md:text-3xl text-slate-500 font-semibold"
                >
                    {index + 1}
                </p>
                <div class="flex flex-col gap-2 sm:gap-5">
                    <h3 class="text-2xl sm:text-3xl md:text-5xl">
                        {talent.name}
                    </h3>
                    <p>{talent.description}</p>
                </div>
            </div>
        {/each}
    </div>
</section>
</main>

