<template>
  <nav class="navbar navbar-fixed">
    <div class="container">
      <a href="/" class="logo">
        <img src="../assets/smkdev.png" alt="SMKDEV" />
      </a>
      <div class="nav-links-wrapper">
        <ul class="nav-links">
          <li>
            <a href="/learn" class="nav-link nav-link-hover">Learn</a>
          </li>
          <li>
            <a href="/community" class="nav-link nav-link-hover">Community</a>
          </li>
          <li>
            <a href="/blog" class="nav-link nav-link-hover">Blog</a>
          </li>
        </ul>
        <ul>
          <a href="/dashboard" class="nav-link btn btn-primary">Dashboard</a>
        </ul>
      </div>
    </div>
  </nav>

  <div id="app">
    <Homepage />
    <section class="hero">
      <div class="hero-text">
        <h1>Jadilah <span class="bold-text">Talenta Digital</span></h1>
        <h1><span class="red-white animate-right">Masa Depan Indonesia</span></h1>
        <p>Belajar langsung dengan expert dari industri dengan kurikulum <br> komprehensif berbasis project-based
          learning
        </p>
      </div>
      <div class="hero-image">
        <img src="../assets/Talent.jpg" alt="Hero Image" />
      </div>
    </section>
  </div>

  <section class="top-features">
    <div class="intro">
      <h2>Orientasi Belajar SMKDEV</h2>
      <p>
        Dapatkan pengalaman belajar berorientasi pengalaman kerja yang dapat
        mengantarkan <br> Anda menjadi talenta yang dibutuhkan oleh industri
        digital terkini.
      </p>
    </div>
    <div class="features">
      <div class="feature animate-hover">
        <img src="../assets/learningpath.jpg" alt="Learning Path Industri" />
        <h3>Learning Path Industri</h3>
        <p>Learning path disusun berdasarkan kebutuhan dunia industri.</p>
      </div>
      <div class="feature animate-hover">
        <img src="../assets/kurikulum.jpg" alt="Kurikulum Komprehensif" />
        <h3>Kurikulum Komprehensif</h3>
        <p>
          Kurikulum komprehensif dan senantiasa diperbarui berdasarkan
          pengalaman di dunia industri.
        </p>
      </div>
      <div class="feature animate-hover">
        <img src="../assets/project.jpg" alt="Project-Based Learning" />
        <h3>Project-Based Learning</h3>
        <p>
          Materi pembelajaran disusun dengan pendekatan - pendekatan
          project based learning.
        </p>
      </div>
    </div>
  </section>


  <section class="talenta-digital">
    <div class="container">
      <h2 class="title">Talenta Digital SMKDEV</h2>
      <div class="talenta-content">
        <div class="image">
          <img src="../assets/talentdigital.png" alt="People Image" />
        </div>
      </div>
    </div>
  </section>


  <section class="testimonial">
    <div class="container">
      <h2>Apa Kata Mereka?</h2>
      <div class="testimonial-cards">
        <div class="card">
          <p>
            Saya merasa sangat bahagia dan bangga karena berhasil menjadi
            pemenang dalam SMKDEV Coding Challenge. Selama kompetisi ini,
            menurut saya, pengalaman ini telah memperluas wawasan dan
            keterampilannya dalam pemrograman. Saya berterima kasih kepada
            tim penyelenggara SMKDEV karena memberikan kesempatan ini dan
            memberikan tantangan yang menginspirasi serta mengapresiasi kerja
            keras peserta lain.
          </p>
          <h4>Shevabey Rahman</h4>
          <h5>Sistem Informasi Universitas Ahmad Dahlan</h5>
        </div>

        <div class="card">
          <p>
            Pada program SMKDEV Scholarship ini memberi saya kesempatan
            untuk berkenalan dan berinteraksi langsung dengan pengajar
            profesional di bidang data analyst. Pada program ini juga saya
            mendapatkan wawasan tentang berbagai peluang karier dalam
            analisis data dan teknologi informasi dan membantu saya
            merencanakan masa depan karir di data analyst.
          </p>
          <h4>Guruh Maulana</h4>
          <h5>STMIK IKM Cirebon</h5>
        </div>

        <div class="card">
          <p>
            Scholarship SMKDEV itu kece banget, walaupun aku dapat
            Scholarship 100%, kualitas kelasnya gak sembarangan. Kelas meet
            yang diadain malem bikin aku gak bingung bagi waktu sama
            kegiatan lainnya. Terus, kak Samuel juga ramah banget. Gak
            kayak yang bikin aku jadi segan bertanya. Materinya juga oke,
            bikin otakku berasa pening nyusun kode buat submission. Terima
            kasih, SMKDEV!
          </p>
          <h4>Tiara Febrianie</h4>
          <h5>SMKN 11 Malang</h5>
        </div>
      </div>
    </div>
  </section>

  <section class="mentor">
    <div class="container">
      <h2>Expert Mentor</h2>
      <p>SMKDEV menghadirkan expert dari industri untuk mendampingi proses belajarmu.</p>

      <div class="mentor-cards">
        <div class="card">
          <img src="../assets/samuel.jpg" alt="Mentor 1" />
          <h3>Samuel Pandohan Terampil Gultom</h3>
          <p>Curriculum Developer</p>
          <p>SMKDEV</p>
        </div>

        <div class="card">
          <img src="../assets/rachmawati.jpg" alt="Mentor 2" />
          <h3>Rachmawati Ari Taurisia</h3>
          <p>Curriculum Developer</p>
          <p>SMKDEV</p>
        </div>

        <div class="card">
          <img src="../assets/sudaryatno.jpg" alt="Mentor 3" />
          <h3>Sudaryatno</h3>
          <p>CTO</p>
          <p>GITS Indonesia</p>
        </div>

        <div class="card">
          <img src="../assets/ibnu.jpg" alt="Mentor 4" />
          <h3>Ibnu Sina Wardy</h3>
          <p>CEO</p>
          <p>GITS Indonesia</p>
        </div>

        <div class="card">
          <img src="../assets/helmi.jpg" alt="Mentor 5" />
          <h3>Helmi Adi Prasetyo</h3>
          <p>Backend Developer</p>
          <p>SMKDEV</p>
        </div>
      </div>
    </div>
  </section>


  <section class="articles">
    <div class="container">
      <h2>Artikel Pilihan SMKDEV</h2>
      <div class="article-cards">
        <div class="card">
          <img src="../assets/promt.jpg" alt="Artikel 1" />
          <h3>Prompt Engineering: Keterampilan, Prospek Karir Dan Kursus AI</h3>
          <p>Kenapa prompt engineering penting untuk AI generatif seperti ChatGPT dan Google Gemini?
            Pelajari keterampilan, prospek karir, dan kursus AI yang dapat membantu Anda sukses dalam industri ini.</p>
          <a href="#" class="btn btn-primary">Baca Lebih Lanjut</a>
        </div>

        <div class="card">
          <img src="../assets/ai.jpg" alt="Artikel 2" />
          <h3>Panduan Pemula Menguasai AI Tools (Alat AI)</h3>
          <p>Pelajari cara menguasai alat AI dengan panduan pemula ini. Temukan langkah-langkah, alat populer, dan tips
            untuk sukses dalam memanfaatkan teknologi AI.</p>
          <a href="#" class="btn btn-primary">Baca Lebih Lanjut</a>
        </div>

        <div class="card">
          <img src="../assets/churn.jpg" alt="Artikel 3" />
          <h3>Strategi menurunkan Churn Rate untuk Pertumbuhan Bisnis</h3>
          <p>Memahami Churn Ratedan strategi menurunkannya. Apa itu Churn Rate ? Churn rate adalah persentase pelanggan
            yang berhenti menggunakan produk atau jasa Anda.</p>
          <a href="#" class="btn btn-primary">Baca Lebih Lanjut</a>
        </div>

        <div class="card">
          <img src="../assets/data.jpg" alt="Artikel 4" />
          <h3>Mengoptimalkan ROI dengan Strategi Analisis Data yang Efektif</h3>
          <p>Dalam era digital yang semakin maju, data telah menjadi salah satu aset paling berhargan bagi perusahaan.
            Namun, nilai sejati dari data terletak pada bagaimana Anda mengolah dan memanfaatkannya dengan efektif. </p>
          <a href="#" class="btn btn-primary">Baca Lebih Lanjut</a>
        </div>
      </div>
    </div>
  </section>

  <footer class="footer">
    <div class="container">
      <div class="footer-content">
        <div class="namecompany">
          <div class="logo">
            <img src="../assets/smkdev.png" alt="SMKDEV" />
          </div>
          <div class="copyright">
            <p>PT Eureka Merdeka Indonesia. All Rights Reserved.</p>
          </div>
        </div>
        <div class="social-links">
          <a href="#" target="_blank"><i class="fab fa-youtube"></i> YouTube</a>
          <a href="#" target="_blank"><i class="fab fa-linkedin-in"></i> LinkedIn</a>
          <a href="#" target="_blank"><i class="fab fa-instagram"></i> Instagram</a>
          <a href="#" target="_blank"><i class="fab fa-tiktok"></i> TikTok</a>
          <a href="#" target="_blank"><i class="fab fa-telegram-plane"></i> Telegram</a>
          <a href="#" target="_blank"><i class="fab fa-whatsapp"></i> WhatsApp</a>
        </div>
      </div>
    </div>
  </footer>

</template>




<script>

export default {
  name: 'HomePage',
  mounted() {
    window.addEventListener('scroll', this.handleScroll);
  },
  beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll);
  },
  methods: {
    handleScroll() {
      const navbar = document.querySelector('.navbar');
      const scrollPosition = window.scrollY;
      if (scrollPosition >= 100) {
        navbar.classList.add('navbar-fixed');
      } else {
        navbar.classList.remove('navbar-fixed');
      }
    }
  }
};


</script>


<style scoped>
.navbar {
  background-color: #fff;
  padding: 1rem 0;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  max-width: 1200px;
  margin: 0 auto;
  padding-left: 40px;
  /* contoh: geser ke kiri 40px */
}

.logo img {
  height: 30px;
}

.nav-links-wrapper {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.nav-links {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  gap: 20px;
}

.nav-link {
  color: #333;
  text-decoration: none;
  padding: 1rem 1.5rem;
  transition: background-color 0.3s ease;
}

.nav-link:hover {
  background-color: rgb(242, 242, 242)2;
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.nav-link-hover:hover {
  background-color: #007bff;
  color: #ffffff;
  /* Warna biru */
}

.btn {
  padding: 1rem 2rem;
  border-radius: 20px;
  font-weight: bold;
}

.btn-primary {
  background-color: #007bff;
  color: #fff;
}

.btn-primary:hover {
  background-color: #0056b3;
}

.navbar-fixed {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 100;
}

.hero {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 50px;
  background-color: #fff;
  margin-bottom: 50px;
  padding-top: 100px;
}

.hero-text {
  max-width: 500px;
}

.hero-text p {
  line-height: 1.5;
  /* Misalnya, line-height 1.5 */
}

.hero-text h1 {
  font-size: 3em;
  margin-bottom: 10px;
}

.hero-text h1 .bold-text {
  font-weight: bold;
}

.hero-text h1 .red-white {
  color: #ff0000;
  /* Merah */
  text-shadow: 0 0 5px #ffffff;
  /* Putih */
}

.hero-image {
  width: 500px;
}

.hero-image img {
  width: 100%;
  height: auto;
  /* Menambahkan bayangan */
  border-radius: 10px;
  /* Menambahkan border radius */
  box-shadow: 8px 8px 8px 8px rgba(48, 48, 48, 0.2);
}

.hero-text h1:nth-child(2) {
  /* Menarget h1 kedua */
  margin-top: 10px;
  /* Misalnya, jarak 10px */


}

.animate-right {
  /* Durasi 3 detik, efek ease-in-out */
  animation: slide-right 3s ease-in-out infinite;
}

@keyframes slide-right {
  from {
    transform: translateX(-100%);
    /* Awal di sebelah kiri */
  }

  to {
    transform: translateX(0);
    /* Akhir di tengah */
  }
}



.top-features {
  display: flex;
  flex-direction: column;
  /* Mengatur elemen menjadi kolom */
  justify-content: space-between;
  padding: 50px;
  margin: 0 auto;
  max-width: 1200px;
  gap: 20px;
  /* Atur jarak sesuai keinginan */
}

.features {
  display: flex;
  justify-content: space-between;
  flex-direction: row;
  /* Mengatur elemen menjadi baris */
  gap: 20px;
}

.intro {
  text-align: center;
  margin-bottom: 20px;
  /* Tambah jarak antara intro dan fitur */
}

.feature {
  background-color: #fff;
  padding: 20px;
  border-radius: 20px;
  box-shadow: 8px 8px 8px 8px rgba(0, 0, 0, 0.1);
  width: 30%;
  text-align: center;
}

.animate-hover {
  transition: background-color 0.3s ease-in-out;
  /* Durasi 0.3 detik, efek ease-in-out */
}


.animate-hover:hover {
  background-color: #007bff;
  /* Warna biru */
}

.feature img {
  width: 100px;
  height: 100px;
  margin-bottom: 20px;
}

.feature h3 {
  font-size: 1.5em;
  margin-bottom: 10px;
}

.feature p {
  font-size: 0.9em;
  line-height: 1.5;
}



.talenta-digital {
  text-align: center;
  padding: 50px 0;
  background-color: #fff;
}

.talenta-digital .container {
  flex-direction: column;
}

.title {
  font-size: 2em;
  font-weight: bold;
  margin-bottom: 20px;
  color: #5daafb;
}

.talenta-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 30px;
}


.image {
  width: auto;
  /* Sesuaikan lebar gambar */
  margin: 0 auto;

}

.image img {
  width: 60%;
}


.testimonial {
  padding: 50px 0;
  background-color: #fff;
}

.testimonial .container {
  flex-direction: column;
  width: 1000px;
}

.testimonial h2 {
  text-align: center;
  font-size: 2em;
  margin-bottom: 20px;
  color: #5daafb;
}

.testimonial p {
  text-align: justify;
  margin-bottom: 20px;
  color: #000000;
}

.testimonial h4 {
  text-align: center;
}

.testimonial h5 {
  text-align: center;
}

.testimonial-cards {
  display: flex;
  flex-direction: row;
  /* Mengatur elemen menjadi kolom */
  justify-content: space-between;
  gap: 30px;
  margin-top: 30px;
}

.card {
  background-color: #f2f2f2;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 100%;
  /* Lebar 100% agar card mengisi seluruh lebar */
}



.card p {
  font-size: 0.9em;
  line-height: 1.5;
  margin-bottom: 10px;
}

.card h4 {
  font-size: 1.2em;
  margin-bottom: 5px;
}

.card:hover {
  background-color: #007bff;
  /* Warna biru saat di-hover */
  color: #fff;
  /* Warna teks putih saat di-hover */
}

.mentor-cards .card:hover {
  background-color: #007bff;
  color: #000000;
}


.mentor {
  padding: 50px 0;
  background-color: #fff;
}

.mentor h2 {
  text-align: center;
  font-size: 2em;
  margin-bottom: 20px;
  color: #5daafb;
}


.mentor p {
  text-align: center;
  margin-bottom: 10px;
  color: #666;
}


.mentor-cards {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  /* Mengatur agar card berpindah baris */
  gap: 50px;
  margin-top: 30px;
}

.mentor .container {
  flex-direction: column;
}

.mentor .card {
  background-color: #f2f2f2;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 8px 8px 8px 8px rgba(0, 0, 0, 0.1);
  width: calc(20% - 10px);
  /* Menghitung lebar agar card rapi */
  text-align: center;
}


.card img {
  width: 150px;
  height: 150px;
  border-radius: 50%;
  margin-bottom: 15px;
}

.card h3 {
  font-size: 1.2em;
  margin-bottom: 5px;
}

.mentor .card p {
  font-size: 0.9em;
  line-height: 1.5;
  margin-bottom: 10px;
  color: #000000;
}

.mentor .card p:hover {
  font-size: 0.9em;
  line-height: 1.5;
  margin-bottom: 10px;
  color: #ffffff;
}




.articles {
  padding: 50px 0;
  background-color: #fff;
}

.articles .container {
  flex-direction: column;
}

.articles h2 {
  text-align: center;
  font-size: 2em;
  margin-bottom: 20px;
  color: #5daafb;
}

.articles p {
  text-align: center;
  margin-bottom: 20px;
  color: #666;
}

.article-cards {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin-top: 30px;
  width: 1000px;
}

.articles .card {
  background-color: #f2f2f2;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: calc(25% - 10px);
  /* Lebar 25% agar rapi */
  text-align: center;
}

.articles .card:hover {
  background-color: #3858fb;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: calc(25% - 10px);
  /* Lebar 25% agar rapi */
  text-align: center;
  color: black;
}

.articles .card img {
  width: 100%;
  height: auto;
  margin-bottom: 15px;
}

.articles .card h3 {
  font-size: 1.2em;
  margin-bottom: 10px;
}

.articles .card p {
  font-size: 0.9em;
  line-height: 1.5;
  margin-bottom: 15px;
  color: #000000;
  text-align: justify;
}

.articles .btn {
  padding: 1rem 1.5rem;
  border-radius: 20px;
  font-weight: bold;
  text-decoration: none;
  color: #fff;
  background-color: #007bff;
  /* Biru */
  transition: background-color 0.3s ease;
}

.btn:hover {
  background-color: #0056b3;
  /* Biru lebih gelap */
}



.footer {
  background-color: #f2f2f2;
  padding: 30px 0;
  color: #333;
}

.footer .container {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
}

.footer-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.logo {
  margin-right: 30px;
}

.logo img {
  height: 30px;
}

.namecompany {
  display: flex;
  align-items: center;
  justify-content: space-between;
  /* Menambahkan space-between */
}

.social-links {
  display: flex;
  align-items: center;
  gap: 20px;
}

.social-links a {
  color: #333;
  text-decoration: none;
}

.social-links i {
  margin-right: 5px;
}

.copyright {
  text-align: right;
}
</style>