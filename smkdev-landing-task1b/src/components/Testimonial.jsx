import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';

const Testimonial = () => {
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
        <div className="px-5 py-12">
            <div className="md:text-center text-start">
                <h1 className="font-bold text-xl text-[#47bb8e]">Cerita Sukses Mereka</h1>
                <h1 className="font-bold md:text-3xl text-2xl mt-4">Mereka telah merasakan belajar di SMKDEV.<br/>Siapkah kamu bergabung?</h1>
                <Slider {...settings} className="mt-10">
                    <div className="p-2">
                        <div className="bg-white rounded-lg shadow-md p-6 text-start">
                            <p className="text-gray-600">Seru! Bisa ketemu orang-orang hebat yang udah bikin banyak project keren. Bisa saling sharing pengalaman ngoding juga. Pokoknya seru deh!</p>
                            <h2 className="text-[15px] font-bold text-gray-800 mt-2">Hasban Fardani</h2>
                            <h2 className="text-[15px] text-gray-800">SMKN 11 Bandung</h2>
                            
                        </div>
                    </div>

                    <div className="p-2">
                        <div className="bg-white rounded-lg shadow-md text-start p-6">
                            <p className="text-gray-600">Sejak bergabung dengan komunitas SMKDEV saya mendapat pengalaman yang banyak terlebih lagi melalui banyak event seperti community bonding talk yang membantu mengembangkan soft skill saya, dengan adanya coding challenge juga membuat saya bersemangat untuk meningkatkan skill pemrograman khususnya problem solving dengan lebih baik lagi, terimakasih SMKDEV dan tim</p>
                            <h2 className="text-[15px] font-bold text-gray-800 mt-2">Karel Trisnanto Utomo</h2>
                            <h2 className="text-[15px] text-gray-800 mt-2">Frontend Web Developer<br/>POLITEKNIK HARAPAN BERSAMA</h2>
                        </div>
                    </div>

                    <div className="p-2">
                        <div className="bg-white rounded-lg shadow-md text-start p-6">
                            <p className="text-gray-600">Kebetulan aku switch career dari guru ngaji, terus aku baca perkembangan sekarang tuh karir yang dibutuhkan di dunia IT itu gak jauh dari programming ataupun olah data. Karena kab data itu “gold”nya zaman now. Nah kebetulan SMKDEV ada Bootcamp dan webinar Data Analyst. Saya tertariknya sama materi Chrun yang ternyata bisa jadi dasar pengambilan keputusan dalam menentukan strategi untuk meningkatkan pangsa pasar atau meningkatkan loyalty customer kita</p>
                            <h2 className="text-[15px] font-bold text-gray-800 mt-2">Rahmatulloh</h2>
                            <h2 className="text-[15px] text-gray-800 mt-2">Pengajar Pondok Al Fatih Depok</h2>
                        </div>
                    </div>

                    <div className="p-2">
                        <div className="bg-white rounded-lg shadow-md text-start p-6">
                            <p className="text-gray-600">Menjadi bagian keluarga SMKDev membantu mengembangkan keterampilan pemrograman saya (melalui event Coding Challenge misalnya). Selain itu, dukungan pemecahan masalah teknis, diskusi bersama, berbagi pengalaman (di Comunity Bounding misalnya) dan kesempatan berkolaborasi dalam sebuah proyek menjadi hal "menarik" yang tidak bisa saya lewatkan.</p>
                            <h2 className="text-[15px] font-bold text-gray-800 mt-2">Asep Dwi Saputra</h2>
                            <h2 className="text-[15px] text-gray-800 mt-2">Teknik Informatika STMIK Widya Utama</h2>
                        </div>
                    </div>

                    <div className="p-2">
                        <div className="bg-white rounded-lg shadow-md text-start p-6">
                            <p className="text-gray-600">Merupakan pengalaman yang sangat menyenangkan bagi saya, mencoba hal baru dan memecahkan suatu masalah. Selain itu, saya juga mendapatkan kesempatan untuk menjalin banyak hubungan dan koneksi dengan orang-orang yang memiliki minat serupa. Dengan ini, saya dapat berbagi pengetahuan dan pengalaman saya dengan orang lain, serta menerima wawasan dan ilmu baru melalui berbagai perspektif yang berbeda.</p>
                            <h2 className="text-[15px] font-bold text-gray-800 mt-2">Muhammad Iqbal Pasha Al Farabi</h2>
                            <h2 className="text-[15px] text-gray-800 mt-2">Teknik Infomatika ITENAS</h2>
                        </div>
                    </div>

                    <div className="p-2">
                        <div className="bg-white rounded-lg shadow-md text-start p-6">
                            <p className="text-gray-600">Saya merasa sangat bahagia dan bangga karena berhasil menjadi pemenang dalam SMKDEV Coding Challenge. Selama kompetisi ini, menurut saya, pengalaman ini telah memperluas wawasan dan keterampilannya dalam pemrograman. Saya berterima kasih kepada tim penyelenggara SMKDEV karena memberikan kesempatan ini dan memberikan tantangan yang menginspirasi serta mengapresiasi kerja keras peserta lain.SMKDEV Coding Challenge adalah pengalaman berharga yang telah memotivasi saya untuk terus belajar dan berkembang dalam dunia pemrograman.</p>
                            <h2 className="text-[15px] font-bold text-gray-800 mt-2">Shevabey Rahman</h2>
                            <h2 className="text-[15px] text-gray-800 mt-2">Sistem Informasi Universitas Ahmad Dahlan</h2>
                        </div>
                    </div>

                    <div className="p-2">
                        <div className="bg-white rounded-lg shadow-md text-start p-6">
                            <p className="text-gray-600">Pada program SMKDEV Scholarship ini memberi saya kesempatan untuk berkenalan dan berinteraksi langsung dengan pengajar profesional di bidang data analyst. Pada program ini juga saya mendapatkan wawasan tentang berbagai peluang karier dalam analisis data dan teknologi informasi dan membantu saya merencanakan masa depan karir di data analyst.</p>
                            <h2 className="text-[15px] font-bold text-gray-800 mt-2">Guruh Maulana</h2>
                            <h2 className="text-[15px] text-gray-800 mt-2">STMIK IKMI Cirebon</h2>
                        </div>
                    </div>

                    <div className="p-2">
                        <div className="bg-white rounded-lg shadow-md text-start p-6">
                            <p className="text-gray-600">Scholarship SMKDEV itu kece banget, walaupun aku dapat Scholarship 100%, kualitas kelasnya gak sembarangan. Kelas meet yang diadain malem bikin aku gak bingung bagi waktu sama kegiatan lainnya. Terus, kak Samuel juga ramah banget. Gak kayak yang bikin aku jadi 'segan bertanya'. Materinya juga oke, bikin otakku berasa pening nyusun kode buat submission. Terima kasih, SMKDEV!</p>
                            <h2 className="text-[15px] font-bold text-gray-800 mt-2">Tiara Febrianie</h2>
                            <h2 className="text-[15px] text-gray-800 mt-2">SMKN 11 Malang</h2>
                        </div>
                    </div>

                    <div className="p-2">
                        <div className="bg-white rounded-lg shadow-md text-start p-6">
                            <p className="text-gray-600">SMKDEV Scholarship ini sangat menarik, karena memiliki penyampaian yang seru dan mudah dipahami. Selain itu juga di dalam bootcamp ini selalu ada di sela-sela kegiatan mengadakan mini games berhadiah</p>
                            <h2 className="text-[15px] font-bold text-gray-800 mt-2">Rizal Maulana</h2>
                            <h2 className="text-[15px] text-gray-800 mt-2">SMKN 1 Cimahi</h2>
                        </div>
                    </div>
                </Slider>
            </div>
        </div>
    );
};

export default Testimonial;
