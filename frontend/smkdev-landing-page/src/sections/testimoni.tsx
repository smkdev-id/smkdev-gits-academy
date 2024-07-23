import { MediaQuery, PinterestGrid } from "@/components";
import { NextPage } from "next";

interface Props {}

const Testimoni: NextPage<Props> = ({}) => {
  const dataTestimoni = [
    {
      text: "Menjadi bagian keluarga SMKDev membantu mengembangkan keterampilan pemrograman saya (melalui event Coding Challenge misalnya). Selain itu, dukungan pemecahan masalah teknis, diskusi bersama, berbagi pengalaman (di Comunity Bounding misalnya) dan kesempatan berkolaborasi dalam sebuah proyek menjadi hal 'menarik' yang tidak bisa saya lewatkan.",
      nama: "Asep Dwi Saputra",
      asal: "Teknik Informatika STMIK Widya Utama",
    },
    {
      text: "Merupakan pengalaman yang sangat menyenangkan bagi saya, mencoba hal baru dan memecahkan suatu masalah. Selain itu, saya juga mendapatkan kesempatan untuk menjalin banyak hubungan dan koneksi dengan orang-orang yang memiliki minat serupa. Dengan ini, saya dapat berbagi pengetahuan dan pengalaman saya dengan orang lain, serta menerima wawasan dan ilmu baru melalui berbagai perspektif yang berbeda.",
      nama: "Muhammad Iqbal Pasha Al Farabi",
      asal: "Teknik Infomatika ITENAS",
    },
    {
      text: "Pada program SMKDEV Scholarship ini memberi saya kesempatan untuk berkenalan dan berinteraksi langsung dengan pengajar profesional di bidang data analyst. Pada program ini juga saya mendapatkan wawasan tentang berbagai peluang karier dalam analisis data dan teknologi informasi dan membantu saya merencanakan masa depan karir di data analyst.",
      nama: "Guruh Maulana",
      asal: "STMIK IKMI Cirebon",
    },
    {
      text: "Scholarship SMKDEV itu kece banget, walaupun aku dapat Scholarship 100%, kualitas kelasnya gak sembarangan. Kelas meet yang diadain malem bikin aku gak bingung bagi waktu sama kegiatan lainnya. Terus, kak Samuel juga ramah banget. Gak kayak yang bikin aku jadi 'segan bertanya'. Materinya juga oke, bikin otakku berasa pening nyusun kode buat submission. Terima kasih, SMKDEV!",
      nama: "Tiara Febrianie",
      asal: "SMKN 11 Malang",
    },
    {
      text: "Saya merasa sangat bahagia dan bangga karena berhasil menjadi pemenang dalam SMKDEV Coding Challenge. Selama kompetisi ini, menurut saya, pengalaman ini telah memperluas wawasan dan keterampilannya dalam pemrograman. Saya berterima kasih kepada tim penyelenggara SMKDEV karena memberikan kesempatan ini dan memberikan tantangan yang menginspirasi serta mengapresiasi kerja keras peserta lain. <br/><br/> SMKDEV Coding Challenge adalah pengalaman berharga yang telah memotivasi saya untuk terus belajar dan berkembang dalam dunia pemrograman.",
      nama: "Shevabey Rahman",
      asal: "Sistem Informasi Universitas Ahmad Dahlan",
    },
    {
      text: "Seru! Bisa ketemu orang-orang hebat yang udah bikin banyak project keren. Bisa saling sharing pengalaman ngoding juga. Pokoknya seru deh!",
      nama: "Hasban Fardani",
      asal: "Siswa SMKN 11 Bandung",
    },
    {
      text: "SMKDEV Scholarship ini sangat menarik, karena memiliki penyampaian yang seru dan mudah dipahami. Selain itu juga di dalam bootcamp ini selalu ada di sela-sela kegiatan mengadakan mini games berhadiah",
      nama: "Rizal Maulana",
      asal: "SMK NEGERI 1 CIMAHI",
    },
    {
      text: "Kebetulan aku switch career dari guru ngaji, terus aku baca perkembangan sekarang tuh karir yang dibutuhkan di dunia IT itu gak jauh dari programming ataupun olah data. Karena kab data itu “gold”nya zaman now. <br/><br/> Nah kebetulan SMKDEV ada Bootcamp dan webinar Data Analyst. Saya tertariknya sama materi Chrun yang ternyata bisa jadi dasar pengambilan keputusan dalam menentukan strategi untuk meningkatkan pangsa pasar atau meningkatkan loyalty customer kita.",
      nama: "Rahmatulloh",
      asal: "Pengajar Pondok Al Fatih Depok",
    },
    {
      text: "Sejak bergabung dengan komunitas SMKDEV saya mendapat pengalaman yang banyak terlebih lagi melalui banyak event seperti community bonding talk yang membantu mengembangkan soft skill saya, dengan adanya coding challenge juga membuat saya bersemangat untuk meningkatkan skill pemrograman khususnya problem solving dengan lebih baik lagi, terimakasih SMKDEV dan tim",
      nama: "Karel Trisnanto Utomo",
      asal: "Frontend Web Developer POLITEKNIK HARAPAN BERSAMA",
    },
  ];

  return (
    <div className="flex w-full justify-center py-20">
      <MediaQuery className="flex flex-col items-center justify-center gap-y-5">
        <div className="mb-10 flex w-full flex-col items-center gap-y-3">
          <h1 className="text-center text-3xl font-semibold capitalize text-primary">
            Apa Kata Mereka?
          </h1>
          <p className="w-[80%] text-center text-xl font-medium">
            Mereka telah merasakan serunya belajar skill digital di SMKDEV.
            Apakah kamu ingin seperti mereka?
          </p>
        </div>

        <div className="w-full">
          <PinterestGrid dataTestimoni={dataTestimoni} />
        </div>
      </MediaQuery>
    </div>
  );
};

export default Testimoni;
