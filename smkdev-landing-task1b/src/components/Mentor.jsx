import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';
import PakSudarYatno from "../assets/expert-mentor/pak-sudaryatno.jpeg";
import PakIbnu from "../assets/expert-mentor/pak-ibnu.jpg";
import PakHelmi from "../assets/expert-mentor/pak-helmi.jpeg";
import PakSamuel from "../assets/expert-mentor/pak-samuel.jpeg";
import BuRachmawati from "../assets/expert-mentor/bu-rachmawati.jpg";

const Mentor = () => {
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
            <div className="px-5 py-12 bg-white">
                <div className="text-start md:text-center">
                    <h1 className="font-bold text-xl text-[#47bb8e]">Expert Mentor</h1>
                    <h1 className="font-bold md:text-3xl text-2xl mt-4">Bimbingan Langsung dari Para Ahli<br/>Untuk Mendorong Kesuksesan Karirmu.</h1>
                </div>
                <Slider {...settings} className="mt-10">
                    

                <div class="w-full max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 py-10">
                    <div class="flex flex-col items-center text-center">
                        <img class="w-24 h-24 mb-3 rounded-full shadow-lg" src={PakSudarYatno} alt="Sudaryatno"/>
                        <h5 class="mb-1 text-xl font-medium text-gray-900 dark:text-white">Sudaryatno</h5>
                        <span class="text-sm text-gray-500 dark:text-gray-400">CTO</span>
                        <span class="text-sm text-gray-500 dark:text-gray-400">GITS Indonesia</span>
                    </div>
                </div>

                <div class="w-full max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 py-10">
                    <div class="flex flex-col items-center text-center">
                        <img class="w-24 h-24 mb-3 rounded-full shadow-lg" src={PakIbnu} alt="Ibnu"/>
                        <h5 class="mb-1 text-xl font-medium text-gray-900 dark:text-white">Ibnu Sina Wardy</h5>
                        <span class="text-sm text-gray-500 dark:text-gray-400">CEO</span>
                        <span class="text-sm text-gray-500 dark:text-gray-400">GITS Indonesia</span>
                    </div>
                </div>

                <div class="w-full max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 py-10">
                    <div class="flex flex-col items-center text-center">
                        <img class="w-24 h-24 mb-3 rounded-full shadow-lg" src={PakHelmi} alt="Helmi"/>
                        <h5 class="mb-1 text-xl font-medium text-gray-900 dark:text-white">Helmi Adi Prasetyo</h5>
                        <span class="text-sm text-gray-500 dark:text-gray-400">Backend Developer</span>
                        <span class="text-sm text-gray-500 dark:text-gray-400">SMKDEV</span>
                    </div>
                </div>

                <div class="w-full max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 py-10">
                    <div class="flex flex-col items-center text-center">
                        <img class="w-24 h-24 mb-3 rounded-full shadow-lg" src={PakSamuel} alt="Samuel"/>
                        <h5 class="mb-1 text-xl font-medium text-gray-900 dark:text-white">Samuel Pandohan Terampil Gultom</h5>
                        <span class="text-sm text-gray-500 dark:text-gray-400">Curriculum Developer</span>
                        <span class="text-sm text-gray-500 dark:text-gray-400">SMKDEV</span>
                    </div>
                </div>

                <div class="w-full max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 py-10">
                    <div class="flex flex-col items-center text-center">
                        <img class="w-24 h-24 mb-3 rounded-full shadow-lg" src={BuRachmawati} alt="Samuel"/>
                        <h5 class="mb-1 text-xl font-medium text-gray-900 dark:text-white">Rachmawati Ari Taurisia</h5>
                        <span class="text-sm text-gray-500 dark:text-gray-400">Curriculum Developer</span>
                        <span class="text-sm text-gray-500 dark:text-gray-400">SMKDEV</span>
                    </div>
                </div>



                </Slider>
            </div>
        </>
    );
}

export default Mentor;