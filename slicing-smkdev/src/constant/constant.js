//import mitra
import digits from "../assets/digits.png"
import arkana from "../assets/arkana.png"
import eudeka from "../assets/eudeka.png"
import gitsid from "../assets/gitsid.png"
import hitopia from "../assets/hitopia.png"
import mantab from "../assets/mantab.png"

//import orientasi
import books from "../assets/icon/books.png"
import books1 from "../assets/icon/open-book.png"
import monitor from "../assets/icon/monitor.png"

//import cardsWhy
import kurikulum from "../assets/icon/kurikulum.png"
import metode from "../assets/icon/metode.png"
import mentor from "../assets/icon/mentor.png"

//import mentor
import ibnu from "../assets/mentor/ibnu-sina.jpg"
import helmi from "../assets/mentor/helmi.png"
import samuel from "../assets/mentor/Samuel.jpg"
import rachmawati from "../assets/mentor/Rachmawati.jpg"
import sudaryatno from "../assets/mentor/sudaryatno.jpeg"

//import artikel
import robot from "../assets/Robot.png"
import panduan from "../assets/panduan.png"
import strategi from "../assets/strategi.png"
import analisis from "../assets/analisis.png"


// NavBar
export const items = [
    { id:1, key:'1', title:"Learn", label:["Bootcamp", "Course", "Expert Class", "Challenges"]}, { id:2, key:'2', title:"Community"}, { id:3, key:'3', title:"Blog"},
]

// Mitra
export const logos = [{
    id: 1,
    src: digits
  },
  {
    id: 2,
    src: arkana
  },
  {
    id: 3,
    src: eudeka
  },
  {
    id: 4,
    src: gitsid
  },
  {
    id: 5,
    src: hitopia
  },
  {
    id: 6,
    src: mantab
  }
]

export const shuffledLogos = [...logos].sort(() => Math.random() - 0.5);
export const shuffledLogos1 = [...logos].sort(() => Math.random() - 1.5);

//Orientasi
export const cardsOrientasi = [
    {id:1, src: books, title: 'Learning Path Industry', desc: 'Learning path disusun berdasarkan kebutuhan dunia industri.'}, 
    {id:2, src: books1, title: 'Kurikulum Komprehensif', desc: 'Kurikulum komprehensif dan senantiasa di perbarui berdasarkan pengalaman di dunia industri'}, 
    {id:3, src: monitor, title: 'Project-Based Learning', desc: 'Materi pembelajaran disusun dengan pendekatan - pendekatan project based learning' }
]

//WhyChoose
export const cardsWhy = [
    {id:1, title: 'Kurikulum', desc: 'Kurikulum berstandar industri global, senantiasa diperbaharui', src: kurikulum},
    {id:2, title: 'Metode', desc: 'Metode belajar project-based learning, dapat menjadi portfolio siswa', src: metode},
    {id:3, title: 'Mentor', desc: 'Dibimbing langsung oleh mentor expert dari industri digital', src: mentor},
]

//Testimoni
export const testiCards = [
    {id:1, name:"Hasban Fardani", major:"Siswa SMKN 11 Bandung", desc:"Seru! Bisa ketemu orang-orang hebat yang udah bikin banyak project keren. Bisa saling sharing pengalaman ngoding juga. Pokoknya seru deh!"},
    {id:2, name:"Rahmatulloh", major:"Pengajar Pondok Al Fatih Depok", desc:'Kebetulan aku switch career dari guru ngaji, terus aku baca perkembangan sekarang tuh karir yang dibutuhkan di dunia IT itu gak jauh dari programming ataupun olah data. Karena kab data itu “gold”nya zaman now.  Nah kebetulan SMKDEV ada Bootcamp dan webinar Data Analyst. Saya tertariknya sama materi Chrun yang ternyata bisa jadi dasar pengambilan keputusan dalam menentukan strategi untuk meningkatkan pangsa pasar atau meningkatkan loyalty customer kita.'},
    {id:3, name:"Karel Trisnanto Utomo", major:"POLITEKNIK HARAPAN BERSAMA", desc:"Sejak bergabung dengan komunitas SMKDEV saya mendapat pengalaman yang banyak terlebih lagi melalui banyak event seperti community bonding talk yang membantu mengembangkan soft skill saya, dengan adanya coding challenge juga membuat saya bersemangat untuk meningkatkan skill pemrograman khususnya problem solving dengan lebih baik lagi, terimakasih SMKDEV dan tim"},
    {id:4, name:"Asep Dwi Saputra", major:"Teknik Informatika STMIK Widya Utama", desc:'Menjadi bagian keluarga SMKDev membantu mengembangkan keterampilan pemrograman saya (melalui event Coding Challenge misalnya). Selain itu, dukungan pemecahan masalah teknis, diskusi bersama, berbagi pengalaman (di Comunity Bounding misalnya) dan kesempatan berkolaborasi dalam sebuah proyek menjadi hal "menarik" yang tidak bisa saya lewatkan.'},
    {id:5, name:"Muhammad Iqbal Pasha Al Farabi", major:"Teknik Infomatika ITENAS", desc:"Merupakan pengalaman yang sangat menyenangkan bagi saya, mencoba hal baru dan memecahkan suatu masalah. Selain itu, saya juga mendapatkan kesempatan untuk menjalin banyak hubungan dan koneksi dengan orang-orang yang memiliki minat serupa. Dengan ini, saya dapat berbagi pengetahuan dan pengalaman saya dengan orang lain, serta menerima wawasan dan ilmu baru melalui berbagai perspektif yang berbeda."},
    {id:6, name:"Shevabey Rahman", major:"Sistem Informasi Universitas Ahmad Dahlan", desc:"Saya merasa sangat bahagia dan bangga karena berhasil menjadi pemenang dalam SMKDEV Coding Challenge. Menurut saya, pengalaman ini telah memperluas wawasan dan keterampilannya dalam pemrograman. SMKDEV Coding Challenge adalah pengalaman berharga yang telah memotivasi saya untuk terus belajar dan berkembang dalam dunia pemrograman."},
    {id:7, name:"Guruh Maulana", major:"STMIK IKMI Cirebon", desc:"Pada program SMKDEV Scholarship ini memberi saya kesempatan untuk berkenalan dan berinteraksi langsung dengan pengajar profesional di bidang data analyst. Pada program ini juga saya mendapatkan wawasan tentang berbagai peluang karier dalam analisis data dan teknologi informasi dan membantu saya merencanakan masa depan karir di data analyst."},
    {id:8, name:"Tiara Febrianie", major:"SMKN 11 Malang", desc:"Scholarship SMKDEV itu kece banget, walaupun aku dapat Scholarship 100%, kualitas kelasnya gak sembarangan. Kelas meet yang diadain malem bikin aku gak bingung bagi waktu sama kegiatan lainnya. Terus, kak Samuel juga ramah banget. Gak kayak yang bikin aku jadi 'segan bertanya'. Materinya juga oke, bikin otakku berasa pening nyusun kode buat submission. Terima kasih, SMKDEV!"},
    {id:9, name:"Rizal Maulana", major:"SMK NEGERI 1 CIMAHI", desc:"SMKDEV Scholarship ini sangat menarik, karena memiliki penyampaian yang seru dan mudah dipahami. Selain itu juga di dalam bootcamp ini selalu ada di sela-sela kegiatan mengadakan mini games berhadiah"},
]

//Mentor
export const mentors = [
    {id:1, src: ibnu, name: "Ibnu Sina Wardy", role: "CEO", Workplace: "GITS Indonesia"},
    {id:2, src: helmi, name: "Helmi Adi Prasetyo", role: "Backend Developer", Workplace: "SMKDEV"},
    {id:3, src: samuel, name: "Samuel Pandohan Terampil Gultom", role: "Curriculum Developer", Workplace: "SMKDEV"},
    {id:4, src: rachmawati, name: "Rachmawati Ari Taurisia", role: "Curriculum Developer", Workplace: "SMKDEV"},
    {id:5, src: sudaryatno, name: "Sudaryatno", role: "CTO", Workplace: "GITS Indonesia"},
  ]

//Artikel
export  const artikels = [
    {id:1, src: robot, title: 'Prompt Engineering: Keterampilan, Prospek Karir Dan Kursus AI', desc: `Kenapa prompt engineering penting untuk AI generatif seperti ChatGPT dan Google Gemini? Pelajari keterampilan, prospe`, link: 'https://www.smk.dev/prompt-engineering-karir-kursus-ai-generatif/'},
    {id:2, src: panduan, title: 'Panduan Pemula Menguasai AI Tools (Alat AI)', desc: 'Pelajari cara menguasai alat AI dengan panduan pemula ini. Temukan langkah-langkah, alat populer, dan tips untuk sukses dalam memanfaatkan teknologi', link: 'https://www.smk.dev/panduan-pemula-menguasai-ai-tools-alat-ai/'},
    {id:3, src: strategi, title: 'Strategi menurunkan Churn Rate untuk Pertumbuhan Bisnis', desc: 'Memahami Churn Rate dan Strategi menurunkannya Apa itu Churn Rate? Churn rate adalah persentase pelanggan yang berhenti menggunakan produk atau', link: 'https://www.smk.dev/pahami-churn-rate-kunci-kesuksesan-di-digital-marketplace/'},
    {id:4, src: analisis, title: 'Mengoptimalkan ROI dengan Strategi Analisis Data yang Efektif', desc: 'Dalam era digital yang semakin maju, data telah menjadi salah satu aset paling berharga bagi perusahaan. Namun, nilai sejati', link: 'https://www.smk.dev/mengoptimalkan-roi-dengan-strategi-analisis-data-yang-efektif/'}
]

