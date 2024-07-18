import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';
import AiImage from "../assets/article/ai.png";
import StrategiImage from "../assets/article/strategi.png";


const Article = () => {
    const settings = {
        dots: true,
        infinite: true,
        speed: 500,
        slidesToShow: 3,
        slidesToScroll: 1,
        autoplay: true,
        autoplaySpeed: 3000,
        responsive: [
            {
                breakpoint: 1024,
                settings: {
                    slidesToShow: 2,
                    slidesToScroll: 1,
                    infinite: true,
                    dots: true,
                },
            },
            {
                breakpoint: 600,
                settings: {
                    slidesToShow: 1,
                    slidesToScroll: 1,
                },
            },
        ],
    };
    return (
        <>
            <div className='text-start md:text-center px-5 py-12'>
                <h1 className='font-bold text-xl text-[#47bb8e]'>Artikel Pilihan SMKDEV</h1>
                <h1 className='font-bold md:text-3xl text-2xl mt-4'>Bekali Masa Depan dengan Keterampilan Digital.</h1>
                <Slider {...settings} className="mt-10">
                    
                <div class="max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 text-start">
                    <a href="#">
                        <img class="rounded-t-lg" src={AiImage} alt="" />
                    </a>
                    <div class="p-5">
                        <a href="#">
                            <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">Prompt Engineering: Keterampilan, Prospek Karir Dan Kursus AI</h5>
                        </a>
                        <p class="mb-3 font-normal text-gray-700 dark:text-gray-400">Kenapa <a target="_blank" href="https://www.smk.dev/tag/prompt-engineering/">prompt engineering</a> penting untuk AI generatif seperti ChatGPT dan Google Gemini? Pelajari keterampilan, prospe.</p>
                        <a href="#" class="inline-flex items-center px-3 py-2 text-sm font-medium text-center text-white bg-gray-900 rounded-md shadow-md shadow-gray-900/10 transition-all hover:shadow-lg hover:shadow-gray-900/20 focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none">
                            Baca Lebih Lanjut
                            <svg class="rtl:rotate-180 w-3.5 h-3.5 ms-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5h12m0 0L9 1m4 4L9 9"/>
                            </svg>
                        </a>
                    </div>
                </div>

                <div class="max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 text-start">
                    <a href="#">
                        <img class="rounded-t-lg" src={AiImage} alt="" />
                    </a>
                    <div class="p-5">
                        <a href="#">
                            <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">Panduan Pemula Menguasai AI Tools (Alat AI)</h5>
                        </a>
                        <p class="mb-3 font-normal text-gray-700 dark:text-gray-400">Pelajari cara menguasai alat AI dengan panduan pemula ini. Temukan langkah-langkah, alat populer, dan tips untuk sukses dalam memanfaatkan teknologi.</p>
                        <a href="#" class="inline-flex items-center px-3 py-2 text-sm font-medium text-center text-white bg-gray-900 rounded-md shadow-md shadow-gray-900/10 transition-all hover:shadow-lg hover:shadow-gray-900/20 focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none">
                            Baca Lebih Lanjut
                            <svg class="rtl:rotate-180 w-3.5 h-3.5 ms-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5h12m0 0L9 1m4 4L9 9"/>
                            </svg>
                        </a>
                    </div>
                </div>

                <div class="max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 text-start">
                    <a href="#">
                        <img class="rounded-t-lg" src={StrategiImage} alt="" />
                    </a>
                    <div class="p-5">
                        <a href="#">
                            <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">Strategi menurunkan Churn Rate untuk Pertumbuhan Bisnis</h5>
                        </a>
                        <p class="mb-3 font-normal text-gray-700 dark:text-gray-400">Memahami <a target="_blank" href="https://www.smk.dev/tag/churn-rate/">Churn Rate</a> dan Strategi menurunkannya Apa itu <a target="_blank" href="https://www.smk.dev/tag/churn-rate/">Churn Rate</a>? <a target="_blank" href="https://www.smk.dev/tag/churn-rate/">Churn rate</a> adalah persentase pelanggan yang berhenti menggunakan produk atau.</p>
                        <a href="#" class="inline-flex items-center px-3 py-2 text-sm font-medium text-center text-white bg-gray-900 rounded-md shadow-md shadow-gray-900/10 transition-all hover:shadow-lg hover:shadow-gray-900/20 focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none">
                            Baca Lebih Lanjut
                            <svg class="rtl:rotate-180 w-3.5 h-3.5 ms-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5h12m0 0L9 1m4 4L9 9"/>
                            </svg>
                        </a>
                    </div>
                </div>

                <div class="max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 text-start">
                    <a href="#">
                        <img class="rounded-t-lg" src={StrategiImage} alt="" />
                    </a>
                    <div class="p-5">
                        <a href="#">
                            <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">Mengoptimalkan ROI dengan Strategi Analisis Data yang Efektif</h5>
                        </a>
                        <p class="mb-3 font-normal text-gray-700 dark:text-gray-400">Dalam <a target="_blank" href="https://www.smk.dev/tag/era-digital/">era digital</a> yang semakin maju, data telah menjadi salah satu aset paling berharga bagi perusahaan. Namun, nilai sejati.</p>
                        <a href="#" class="inline-flex items-center px-3 py-2 text-sm font-medium text-center text-white bg-gray-900 rounded-md shadow-md shadow-gray-900/10 transition-all hover:shadow-lg hover:shadow-gray-900/20 focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none">
                            Baca Lebih Lanjut
                            <svg class="rtl:rotate-180 w-3.5 h-3.5 ms-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5h12m0 0L9 1m4 4L9 9"/>
                            </svg>
                        </a>
                    </div>
                </div>
                </Slider>
            </div>
        </>
    )
}

export default Article;