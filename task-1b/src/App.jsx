import { CardArticle } from "./components/CardArticle";
import { CardBenefit } from "./components/CardBenefit";
import { CardLearningMethod } from "./components/CardLearningMethod";
import { CardMentor } from "./components/CardMentor";
import { CardTestimony } from "./components/CardTestimony";
import { Footer } from "./components/Footer";
import { Navigation } from "./components/Navigation";
import { TextAnimation } from "./components/TextAnimation";
import ScrollCarousel from "scroll-carousel-react";

export default function App() {
  const mitra = ["digits", "gits", "eudeka", "hitopia", "arkana", "mantab-one"];
  return (
    <div className="w-full">
      <Navigation />
      <main className="w-full bg-white">
        <div className="w-full h-full sm:h-screen px-[20px] sm:px-[30px] lg:px-[70px] xl:px-[140px] pt-[15vh] pb-[10vh] sm:py-0 sm:pb-0 flex flex-col sm:flex-row items-center">
          <div className="w-full sm:w-[55%]">
            <p className="text-[42px] sm:text-[50px] font-semibold">
              Jadilah Talenta Digital{" "}
              <span className="text-[#004FC5]">
                <TextAnimation />
              </span>
            </p>
            {/* <p className="text-[42px] sm:text-[50px] text-[#004FC5] font-semibold"><TextAnimation/></p> */}
            <p className="text-[16px] sm:text-[20px] text-[#172F54] mt-[10px] sm:mt-[30px]">
              Belajar langsung dengan expert dari industri dengan kurikulum
              komperhensif berbasis project-based learning
            </p>
          </div>
          <div className="w-[80%] sm:w-[45%] mt-[50px] flex justify-center">
            <img src="image-header.svg" alt="" />
          </div>
        </div>

        <div className="h-[300px] flex flex-col items-center justify-center bg-[#F3F8FB]">
          <p className="text-[21px] text-[#1C3965] font-medium">
            Dipercaya Oleh Mitra Industri
          </p>
          <ScrollCarousel autoplay autoplaySpeed={8}>
            <div className="flex mt-[30px] gap-[20px] ">
              {mitra.map((mitra) => {
                return (
                  <img src={`${mitra}.svg`} alt="" className="cursor-pointer" />
                );
              })}
            </div>
          </ScrollCarousel>
        </div>

        <div className="h-full sm:h-screen px-[20px] sm:px-[30px] lg:px-[70px] xl:px-[140px] py-[10vh] sm:py-0 flex flex-col justify-center items-center">
          <div className="w-full sm:w-[70%] text-center">
            <p className="text-[28px] sm:text-[30px] text-[#1C3965] font-semibold">
              Orientasi Belajar SMKDEV
            </p>
            <p className="text-[16px] sm:text-[18px] mt-[10px]">
              Dapatkan pengalaman belajar berorientasi pengalaman kerja yang
              dapat mengantarkan Anda menjadi talenta yang dibutuhkan oleh
              industri digital terkini.
            </p>
          </div>
          <div className="flex flex-col sm:flex-row gap-[25px] sm:gap-[2%] mt-[50px]">
            <CardLearningMethod
              img="Learning-method1.svg"
              title="Learning Path Industri"
              desc="Learning path disusun berdasarkan kebutuhan dunia industri."
            />
            <CardLearningMethod
              img="Learning-method2.svg"
              title="Kurikulum Komprehensif"
              desc="Kurikulum komprehensif dan senantiasa di perbarui berdasarkan pengalaman di dunia industri"
            />
            <CardLearningMethod
              img="Learning-method3.svg"
              title="Project-Based Learning"
              desc="Materi pembelajaran disusun dengan pendekatan - pendekatan project based learning"
            />
          </div>
        </div>

        <div className="h-full sm:h-screen w-full px-[20px] sm:px-[30px] lg:px-[70px] xl:px-[140px] py-[10vh] sm:py-0 flex flex-col-reverse sm:flex-row items-center justify-between bg-[#F3F8FB]">
          <div className="w-full sm:w-[48%] flex flex-col gap-[20px]">
            <CardBenefit
              img="kurikulum.svg"
              title="Kurikulum"
              desc="Kurikulum berstandar industri global, senantiasa diperbaharui"
            />
            <CardBenefit
              img="metode.svg"
              title="Metode"
              desc="Metode belajar project-based learning, dapat menjadi portfolio siswa"
            />
            <CardBenefit
              img="mentor.svg"
              title="Mentor"
              desc="Dibimbing langsung oleh mentor expert dari industri digital"
            />
          </div>
          <div className="w-full sm:w-[48%] mb-[50px] sm:mb-0">
            <p className="text-[28px] sm:text-[30px] text-[#1C3965] font-semibold">
              MENGAPA HARUS MEMILIH BELAJAR DI{" "}
              <span className="text-[#004FC5]">SMKDEV</span>
            </p>
            <p className="mt-[10px]">
              Pengalaman belajarmu jauh lebih menyenangkan dengan perpaduan
              belajar teori dan capstone project yang dapat menjadi portofolio
              pribadimu. Tidak lupa, Kamu juga dapat belajar langsung dari
              expert mentor dari dunia industri.
            </p>
          </div>
        </div>

        <div className="w-full lg:h-screen px-[20px] sm:px-[30px] lg:px-[70px] xl:px-[140px] py-[10vh] lg:py-0 flex flex-col items-center justify-center">
          <p className="mb-[20px] text-[20px] sm:text-[24px] text-[#1C3965] font-semibold">
            Talenta Digital SMKDEV
          </p>
          <img src="talenta-digital.svg" alt="" />
        </div>

        <div className="px-[20px] sm:px-[30px] lg:px-[70px] xl:px-[140px]">
          <p className="text-[28px] sm:text-[30px] text-[#0E1C32] font-semibold text-center">
            Apa Kata Mereka?
          </p>
          <p className="text-[16px] sm:text-[18px] text-[#0E1C32] text-center mt-[10px] mb-[30px]">
            Mereka telah merasakan serunya belajar skill digital di SMKDEV.
            Apakah kamu ingin seperti mereka?
          </p>
          <div className="w-full flex justify-between overflow-x-auto scrollbar-hide">
            <div className="min-w-full sm:min-w-min md:w-[31.5%]">
              <CardTestimony
                name="Rahmatulloh"
                job="Pengajar Pondok Al Fatih Depok"
                desc="Kebetulan aku switch career dari guru ngaji, terus aku baca perkembangan sekarang tuh karir yang dibutuhkan di dunia IT itu gak jauh dari programming ataupun olah data. Karena kab data itu “gold”nya zaman now. Nah kebetulan SMKDEV ada Bootcamp dan webinar Data Analyst. Saya tertariknya sama materi Chrun yang ternyata bisa jadi dasar pengambilan keputusan dalam menentukan strategi untuk meningkatkan pangsa pasar atau meningkatkan loyalty customer kita."
              />
            </div>
            <div className="min-w-full sm:min-w-min md:w-[31.5%]">
              <CardTestimony
                name="Karel Trisnanto Utomo"
                job="Frontend Web Developer"
                desc="Sejak bergabung dengan komunitas SMKDEV saya mendapat pengalaman yang banyak terlebih lagi melalui banyak event seperti community bonding talk yang membantu mengembangkan soft skill saya, dengan adanya coding challenge juga membuat saya bersemangat untuk meningkatkan skill pemrograman khususnya problem solving dengan lebih baik lagi, terimakasih SMKDEV dan tim"
              />
            </div>
            <div className="min-w-full sm:min-w-min md:w-[31.5%]">
              <CardTestimony
                name="Asep Dwi Saputra"
                job="Teknik Informatika STMIK Widya Utama"
                desc="Menjadi bagian keluarga SMKDEV membantu mengembangkan keterampilan pemrograman saya (melalui event Coding Challenge misalnya). Selain itu, dukungan pemecahan masalah teknis, diskusi bersama, berbagi pengalaman (di Comunity Bounding misalnya) dan kesempatan berkolaborasi dalam sebuah proyek menjadi hal menarik yang tidak bisa saya lewatkan."
              />
            </div>
          </div>
        </div>

        <div className="w-full h-screen flex flex-col justify-center items-center px-[20px] sm:px-[30px] lg:px-[70px] xl:px-[140px]">
          <p className="text-[28px] sm:text-[30px] text-[#0E1C32] font-semibold">
            Expert Mentor
          </p>
          <p className="text-[16px] sm:text-[18px] text-[#0E1C32] mt-[10px] mb-[60px] text-center">
            SMKDEV menghadirkan expert dari industri untuk mendampingi proses
            belajarmu.
          </p>
          <div className="w-full flex sm:justify-between overflow-x-auto scrollbar-hide">
            <div className="min-w-full sm:min-w-min sm:w-[18%]">
              <CardMentor
                name="Ibnu Sina Wardy"
                job="CEO"
                company="GITS Indonesia"
                img="mentor-ibnu.svg"
              />
            </div>
            <div className="min-w-full sm:min-w-min sm:w-[18%]">
              <CardMentor
                name="Helmi Adi Prasetyo"
                job="Backend Developer"
                company="SMKDEV"
                img="mentor-helmi.png"
              />
            </div>
            <div className="min-w-full sm:min-w-min sm:w-[18%]">
              <CardMentor
                name="Samuel Pandohan Terampil Gultom"
                job="Curriculum Developer"
                company="SMKDEV"
                img="mentor-samuel.svg"
              />
            </div>
            <div className="min-w-full sm:min-w-min sm:w-[18%]">
              <CardMentor
                name="Rachmawati Ari Taurisia"
                job="Curriculum Developer"
                company="SMKDEV"
                img="mentor-rachmawati.svg"
              />
            </div>
            <div className="min-w-full sm:min-w-min sm:w-[18%]">
              <CardMentor
                name="Sudaryatno"
                job="CTO"
                company="GITS Indonesia"
                img="mentor-sudaryatno.svg"
              />
            </div>
          </div>
        </div>

        <div className="w-full h-screen px-[20px] sm:px-[30px] lg:px-[70px] xl:px-[140px] flex flex-col justify-center items-center">
          <p className="text-[28px] sm:text-[30px] text-[#0E1C32] font-semibold">
            Artikel Pilihan SMKDEV
          </p>
          <div className="w-full mt-[30px] flex gap-[20px] sm:justify-between overflow-x-auto scrollbar-hide">
            <div className="min-w-full sm:min-w-min lg:w-[23%]">
              <CardArticle
                title="Prompt Engineering: Keterampilan, Prospek Karir Dan Kursus AI"
                desc="Pada era di mana Artificial Intelligence (AI) semakin menyebar ke berbagai aspek kehidupan kita, salah satu elemen kunci kesuksesan implementasi AI adalah rekayasa prompt atau prompt engineering. Namun, apa sebenarnya rekayasa prompt dan mengapa hal ini begitu penting dalam konteks teknologi AI? AI memungkinkan komputer meniru kecerdasan manusia untuk memecahkan masalah. AI banyak digunakan dalam aplikasi seperti asisten digital, kendaraan otonom, dan alat generatif seperti ChatGPT. "
                img="artikel-prompt-engineering.png"
              />
            </div>
            <div className="min-w-full sm:min-w-min lg:w-[23%]">
              <CardArticle
                title="Panduan Pemula Menguasai AI Tools (Alat AI)"
                desc="Kecerdasan buatan, atau AI, telah dengan cepat berkembang dalam beberapa tahun terakhir dan semakin banyak diterapkan di berbagai industri. Saat ini, tidak hanya para teknisi dapat menemukan alat AI, tetapi juga pemula yang tertarik untuk menggunakannya di tempat kerja mereka atau dalam kehidupan sehari-hari. Pada artikel ini, saya akan menguraikan tentang cara-cara belajar dan menggunakan alat AI bagi para pemula – serangkaian rekomendasi yang bisa mengubah pekerjaan tanpa memahami AI menjadi pekerjaan dengan menggunakan AI."
                img="artikel-AI.png"
              />
            </div>
            <div className="min-w-full sm:min-w-min lg:w-[23%]">
              <CardArticle
                title="Strategi Menurunkan Churn Rate untuk Pertumbuhan Bisnis"
                desc="Memahami Churn Rate  dan Strategi menurunkannya Apa itu Churn Rate? Churn rate adalah persentase pelanggan yang berhenti menggunakan produk atau layanan Anda dalam periode tertentu. Semakin tinggi churn rate, semakin banyak pelanggan yang hilang. Hal ini dapat berdampak negatif pada pendapatan dan reputasi bisnis Anda."
                img="artikel-churn-rate.png"
              />
            </div>
            <div className="min-w-full sm:min-w-min lg:w-[23%]">
              <CardArticle
                title="Mengoptimalkan ROI dengan Strategi Analisis Data yang Efektif"
                desc="Dalam era digital yang semakin maju, data telah menjadi salah satu aset paling berharga bagi perusahaan. Namun, nilai sejati dari data tersebut hanya dapat direalisasikan melalui analisis yang efektif. Salah satu tujuan utama dari analisis data adalah untuk meningkatkan Return on Investment (ROI), yaitu mengoptimalkan hasil dari investasi yang dilakukan. Dalam artikel ini, kita akan menjelaskan bagaimana strategi analisis data yang efektif dapat membantu dalam mengoptimalkan ROI."
                img="artikel-analisis-data.png"
              />
            </div>
          </div>

          <button className="text-white text-[14px] bg-[#1C3965] py-[10px] px-[20px] mt-[20px] rounded-[10px]">
            Lihat Semua Artikel
          </button>
        </div>
      </main>
      <Footer />
    </div>
  );
}
