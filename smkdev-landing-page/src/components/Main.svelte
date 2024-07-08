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
        }
    ];

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
    <section class="py-3 lg:py-5 flex flex-col gap-2 bg-violet-100" id="mitra">
        <div class="flex flex-col items-center gap-2 text-center z-[12]">
            <h6 class="font-semibold text-base sm:text-lg md:text-lg">
                Dipercayai Oleh Mitra Industri
            </h6>
        </div>
        <div class="flex justify-center bg-violet-100 px-6 lg:px-[300px]">
             <!--HTML CODE-->
            <div class="w-full relative">
                <div class="swiper centered-slide-carousel swiper-container relative ">
                    <div class="swiper-wrapper h-max relative">
                        <div class="swiper-slide">
                            <div class="bg-violet-100 rounded-2xl h-36 flex justify-center items-center">
                                <img src="https://git.gits.id/uploads/-/system/appearance/logo/1/Logo-Main__1_.png" alt="GITS" class="max-h-10"/>
                            </div>
                        </div>
                        <div class="swiper-slide">
                            <div class="bg-violet-100 rounded-2xl h-36 flex justify-center items-center">
                                <img src="https://eudeka.storage.googleapis.com/wp/logo-eudeka-512x512-1-300x300.png" alt="Eudeka" class="max-h-24"/>
                            </div>
                        </div>
                        <div class="swiper-slide">
                            <div class="bg-violet-100 rounded-2xl h-36 flex justify-center items-center">
                                <img src="images/arkana.png" alt="Arkana" class="h-8"/>
                            </div>
                        </div>
                        <div class="swiper-slide">
                            <div class="bg-violet-100 rounded-2xl h-36 flex justify-center items-center">
                                <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSd-YWd6j6ZfTCJ7X8pv_n96QnZQtd8qAwwfQ&s" alt="violantis" class="max-h-7"/>
                            </div>
                        </div>
                        <div class="swiper-slide">
                            <div class="bg-violet-100 rounded-2xl h-36 flex justify-center items-center">
                                <img src="https://hitopia.id/wp-content/uploads/2023/04/Logo-Deep-Hitopia-Brown-min-Copy.png" alt="Hiptopia" class="max-h-16"/>
                            </div>
                        </div>
                        <div class="swiper-slide">
                            <div class="bg-violet-100 rounded-2xl h-36 flex justify-center items-center">
                                <img src="https://mantab.one/wp-content/uploads/2024/05/logo.png" alt="Mantabone" class="max-h-12"/>
                            </div>
                        </div>
                    </div>
                    <div class="swiper-pagination "></div>
                </div>
            </div>    
        </div>
    </section>
</main>

