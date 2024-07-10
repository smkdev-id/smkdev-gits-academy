<script>
    import Step from "./Step.svelte";
    import { onMount } from 'svelte';
    import Testimoni from "./Testimoni.svelte";
    import Mentor from './Mentor.svelte';
    import Blog from "./Blog.svelte";
    import Contact from "./Contact.svelte";
    import Footer from "./Footer.svelte"
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
    ];
    let talents = [
        {
            description: "Daftar ke program pelatihan SMKDEV"
        },
        {
            description: "Mengikuti semua program pelatihan sampai tuntas"
        },
        {
            description: "Terdaftar pada Talent Pool SMKDEV"
        },
        {
            description: "Mengikuti program magang/kerja di Partner Industri SMKDEV"
        },
        {
            description: "Saya telah menjadi talenta digital handal"
        }
    ];
    let testimonies = [
        {
            name: "Guruh Maulana",
            job: "STMIK IKMI Cirebon",
            message: "Pada program SMKDEV Scholarship ini memberi saya kesempatan untuk berkenalan dan berinteraksi langsung dengan pengajar profesional di bidang data analyst. Pada program ini juga saya mendapatkan wawasan tentang berbagai peluang karier dalam analisis data dan teknologi informasi dan membantu saya merencanakan masa depan karir di data analyst.",
        },
        {
            name: "Tiara Febrianie",
            job: "SMKN 11 Malang",
            message: "Scholarship SMKDEV itu kece banget, walaupun aku dapat Scholarship 100%, kualitas kelasnya gak sembarangan. Kelas meet yang diadain malem bikin aku gak bingung bagi waktu sama kegiatan lainnya. Terus, kak Samuel juga ramah banget. Gak kayak yang bikin aku jadi 'segan bertanya'. Materinya juga oke, bikin otakku berasa pening nyusun kode buat submission. Terima kasih, SMKDEV!",
        },
        {
            name: "Rizal Maulana",
            job: "SMK NEGERI 1 CIMAHI",
            message: "SMKDEV Scholarship ini sangat menarik, karena memiliki penyampaian yang seru dan mudah dipahami. Selain itu juga di dalam bootcamp ini selalu ada di sela-sela kegiatan mengadakan mini games berhadiah",
        },
        {
            name: "Rahmatulloh",
            job: "Pengajar Pondok Al Fatih Depok",
            message: "Kebetulan aku switch career dari guru ngaji, terus aku baca perkembangan sekarang tuh karir yang dibutuhkan di dunia IT itu gak jauh dari programming ataupun olah data. Karena kab data itu “gold”nya zaman now. Nah kebetulan SMKDEV ada Bootcamp dan webinar Data Analyst. Saya tertariknya sama materi Chrun yang ternyata bisa jadi dasar pengambilan keputusan dalam menentukan strategi untuk meningkatkan pangsa pasar atau meningkatkan loyalty customer kita.",
        },
        {
            name: "Hasban Fardani",
            job: "Siswa SMKN 11 Bandung",
            message: "Seru! Bisa ketemu orang-orang hebat yang udah bikin banyak project keren. Bisa saling sharing pengalaman ngoding juga. Pokoknya seru deh!",
        },
        {
            name: "Karel Trisnanto Utomo",
            job: "Frontend Web Developer",
            message: "Sejak bergabung dengan komunitas SMKDEV saya mendapat pengalaman yang banyak terlebih lagi melalui banyak event seperti community bonding talk yang membantu mengembangkan soft skill saya, dengan adanya coding challenge juga membuat saya bersemangat untuk meningkatkan skill pemrograman khususnya problem solving dengan lebih baik lagi, terimakasih SMKDEV dan tim",
        },
        {
            name: "Asep Dwi Saputra",
            job: "Teknik Informatika STMIK Widya Utama",
            message: "Menjadi bagian keluarga SMKDev membantu mengembangkan keterampilan pemrograman saya (melalui event Coding Challenge misalnya). Selain itu, dukungan pemecahan masalah teknis, diskusi bersama, berbagi pengalaman (di Comunity Bounding misalnya) dan kesempatan berkolaborasi dalam sebuah proyek menjadi hal menarik yang tidak bisa saya lewatkan.",
        },
        {
            name: "Muhammad Iqbal Pasha Al Farabi",
            job: "Teknik Infomatika ITENAS",
            message: "Merupakan pengalaman yang sangat menyenangkan bagi saya, mencoba hal baru dan memecahkan suatu masalah. Selain itu, saya juga mendapatkan kesempatan untuk menjalin banyak hubungan dan koneksi dengan orang-orang yang memiliki minat serupa. Dengan ini, saya dapat berbagi pengetahuan dan pengalaman saya dengan orang lain, serta menerima wawasan dan ilmu baru melalui berbagai perspektif yang berbeda."
        },
    ];
    let mentors = [
        {
            name: "Ibnu Sina Wardy",
            position: "CEO",
            job: "GITS.ID",
            img: "/images/ibnuS.png",
            twitter: "#",
            github: "#",
            linkedin: "#",
        },
        {
            name: "Helmi Adi Prasetyo",
            position: "SMKDEV",
            job: "Backend Developer",
            img: "/images/helmi.png",
            twitter: "#",
            github: "#",
            linkedin: "#",
        },
        {
            name: "Samuel Pandohan T G",
            position: "SMKDEV",
            job: "Curriculum Developer",
            img: "/images/samuel.png",
            twitter: "#",
            github: "#",
            linkedin: "#",
        },
        {
            name: "Rachmawati Ari T",
            position: "SMKDEV",
            job: "Curriculum Developer",
            img: "/images/rachmawati.png",
            twitter: "#",
            github: "#",
            linkedin: "#",
        },
        {
            name: "Sudaryatno",
            position: "CTO",
            job: "GITS.ID",
            img: "/images/sudaryatno.png",
            twitter: "#",
            github: "#",
            linkedin: "#",
        },
    ];
    let blogs = [
        {
            linkDetailBlog : "#",
            samllImg: "/images/smallImg1.png",
            category : "General",
            title : "Prompt Engineering: Keterampilan, Prospek Karir Dan Kursus AI",
            description : "Kenapa prompt engineering penting untuk AI generatif seperti ChatGPT dan Google Gemini? Pelajari keterampilan, prospek"
        },
        {
            linkDetailBlog : "#",
            samllImg: "/images/smallImg1.png",
            category : "General",
            title : "Panduan Pemula Menguasai AI Tools (Alat AI)",
            description : "Pelajari cara menguasai alat AI dengan panduan pemula ini. Temukan langkah-langkah, alat populer, dan tips untuk sukses dalam memanfaatkan teknologi"
        },
        {
            linkDetailBlog : "#",
            samllImg: "/images/smallImg1.png",
            category : "Data Sains",
            title : "Strategi menurunkan Churn Rate untuk Pertumbuhan Bisnis",
            description : "Memahami Churn Rate dan Strategi menurunkannya Apa itu Churn Rate? Churn rate adalah persentase pelanggan yang berhenti menggunakan produk atau"
        },
        {
            linkDetailBlog : "#",
            samllImg: "/images/smallImg1.png",
            category : "General",
            title : "Mengoptimalkan ROI dengan Strategi Analisis Data yang Efektif",
            description : "  Dalam era digital yang semakin maju, data telah menjadi salah satu aset paling berharga bagi perusahaan. Namun, nilai sejati"
        },
    ]
    onMount(() => {
        const swiper = new Swiper(".mitra-carousel", {
        centeredSlides: true,
        paginationClickable: true,
        loop: true,
        spaceBetween: 30,
        slideToClickedSlide: true,
        pagination: {
            el: ".mitra-carousel .swiper-pagination",
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
        const swiperTesti = new Swiper(".testi-carousel", {
        centeredSlides: true,
        paginationClickable: true,
        loop: true,
        spaceBetween: 30,
        slideToClickedSlide: true,
        pagination: {
            el: ".testi-carousel .swiper-pagination",
            clickable: true,
        },
        breakpoints: {
            
            1028: {
            slidesPerView: 3,
            spaceBetween: 10
            },
            0: {
            slidesPerView: 1,
            spaceBetween: 0
            }
        }
      });
        const swiperMentor = new Swiper(".mentor-carousel", {
        centeredSlides: true,
        paginationClickable: true,
        loop: true,
        spaceBetween: 30,
        slideToClickedSlide: true,
        pagination: {
            el: ".mentor-carousel .swiper-pagination",
            clickable: true,
        },
        breakpoints: {
            
            1028: {
            slidesPerView: 3,
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
                <div class="swiper mitra-carousel swiper-container relative ">
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
    <!-- Talenta -->
    <section
    id="about"
    class="py-6 lg:py-8 flex flex-col gap-2 sm:gap-10 md:gap-6 px-10 lg:px-[300px]">
    <div class="flex flex-col gap-2 text-center relative  before:absolute before:top-0 before:left-0 before:w-2/3 before:h-1.5 before:bg-violet-100 after:absolute after:bottom-0 after:right-0 after:w-2/3 after:h-1.5 after:bg-violet-100 py-4">
        <h3 class="font-semibold text-xl sm:text-2xl md:text-4xl">
            Talenta Digital <span class="text-violet-400">SMKDEV</span> 
        </h3>
    </div>
    <div class="lg:flex lg:items-center ">
        <div class="hidden lg:flex lg:items-center lg:w-1/2 lg:justify-center">
            <img class="w-[20rem] h-[20rem] object-cover xl:w-[28rem] xl:h-[28rem] rounded-xl" src="/images/ilustration1.png" alt="">
        </div>
        <div class="w-full lg:space-y-4 lg:w-1/2 grid lg:grid-cols-1 grid-cols-2 lg:gap-0 gap-4">
            {#each talents as talent, index}
            <div class="md:flex md:items-start md:-mx-4  ">
                <div class="md:mx-4 shadow-md w-full p-3 rounded-xl">
                    <h1 class="text-sm sm:text-md md:text-lg">{talent.description}</h1>
                </div>
            </div>
            {/each}
        </div>
        
    </div>
    </section>
    <!-- Testimoni -->
    <section class="py-3 lg:py-5 flex flex-col gap-2 bg-violet-50" id="testimoni">
    <div class="flex flex-col items-center gap-2 text-center px-6 mb-6">
        <h3 class="font-semibold text-3xl sm:text-4xl mb-3">
            What <span class="poppins text-violet-400">they</span> said <span class="text-violet-400" id="cursor">
                ?
            </span>
        </h3>
        <h2 class="text-sm sm:text-md md:text-lg">
            Mereka telah merasakan serunya belajar skill digital di SMKDEV. Apakah kamu ingin seperti mereka?
        </h2>
    </div>
    <div class="flex justify-center bg-violet-50 px-6 lg:px-[300px]">
        <div class="w-full relative">
            <div class="swiper testi-carousel swiper-container">
                <div class="swiper-wrapper h-max">
                    {#each testimonies as testi }
                    <div class="swiper-slide">
                    <Testimoni testimoni = {testi}>
                    </Testimoni>
                    </div>
                    {/each}
                </div>
            </div>
        </div>    
    </div>
    </section>
    <!-- Mentor -->
    <section class="py-3 lg:py-5 flex flex-col gap-2 " id="Mentor">
    <div class="flex flex-col items-center gap-2 text-center px-6 mb-6">
        <h3 class="font-semibold text-3xl sm:text-4xl mb-3">
            Expert Mentor
        </h3>
        <h2 class="text-sm sm:text-md md:text-lg">
            SMKDEV menghadirkan expert dari industri untuk mendampingi proses belajarmu.
        </h2>
    </div>
    <div class="flex justify-center px-6 lg:px-[300px]">
        <div class="w-full relative">
            <div class="swiper mentor-carousel swiper-container">
                <div class="swiper-wrapper h-max">
                    {#each mentors as mentor }
                    <div class="swiper-slide">
                    <Mentor mentor = {mentor}>
                    </Mentor>
                    </div>
                    {/each}
                </div>
            </div>
        </div>    
    </div>
    </section>
    <!-- Article -->
    <section class="py-10 lg:pt-30 flex flex-col gap-8 px-6 lg:px-[300px]" id="blog">
        <h3 class="font-semibold text-3xl sm:text-4xl text-center">
            Artikel Pilihan <span class="poppins text-violet-400">SMKDEV</span> 
        </h3>
        <div class="mx-auto grid max-w-screen-lg justify-center px-4 grid-cols-2 sm:gap-4 sm:px-8 md:grid-cols-3">
            {#each blogs as blog }
                <Blog blog = {blog}>
                </Blog>
            {/each}
        </div>
        <a href="#" class="blueShadow relative overflow-hidden px-10 py-2 mx-auto group rounded-md bg-violet-900 text-slate-950">
            <div
                class="absolute top-0 right-full w-full h-full bg-violet-400 opacity-20 group-hover:translate-x-full z-0 duration-200"
            />
            <h4 class="relative text-sm z-9 text-white text-center">NEXT</h4>
    </a>
    </section>
    <!-- Contact -->
     <section class="py-10 lg:py-5 px-6 lg:px-[300px] bg-violet-50">
        <Contact>
        </Contact>
     </section>
</main>

