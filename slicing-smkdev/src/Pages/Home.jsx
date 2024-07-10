import React from 'react';
import Layout from '../Layout';
import StyleHome from './Home.module.css';
import Hero1 from '../assets/belajar-coding.jpg';
import Boy from '../assets/boy.jpg';
import Boy2 from '../assets/boy2.jpg';
import Girl from '../assets/girl.png';
import ImgKurikulum from '../assets/kurikulum.png';
import ImgMetode from '../assets/metode.png';
import ImgMentor from '../assets/mentor.png';

function Home() {
  return (
    <>
      <Layout>
        {/* Hero Section */}
        <section className={`${StyleHome.hero}`}>
          <div className="container">
            <div className="row align-items-center">
              <div className="col-sm-6">
                <h1> Jadilah Talenta Digital</h1>
                <h2>Masa Depan Indonesia </h2>
                <p>
                  Belajar langsung dengan expert dari industri dengan kurikulum{' '}
                  <br />
                  komperhensif berbasis project-based learning
                </p>
              </div>
              <div className="col-sm-6 d-sm-flex ">
                <img
                  src={Hero1}
                  alt="hero1"
                  className={`${StyleHome.hero1img}`}
                />
              </div>
            </div>
          </div>
        </section>
        {/* orientasi */}
        <section className={`${StyleHome.orientasi}`}>
          <div className="row justify-content-center my-5">
            <div className="w-auto">
              <div className="border-bottom border-2 border-warning text-center">
                <h2>Orientasi Belajar SMKDEV</h2>
              </div>
              <div className="mt-3 text-center">
                <p>
                  Dapatkan pengalaman belajar berorientasi pengalaman kerja yang
                  dapat mengantarkan Anda <br /> menjadi talenta yang dibutuhkan
                  oleh industri digital terkini.
                </p>
              </div>
            </div>
          </div>
          <div className="row justify-content-center flex-sm-row flex-column">
            <div className="col-sm-3 d-flex flex-column align-items-center">
              <div className="rounded-4 bg-primary d-flex justify-content-center align-items-center py-5">
                <img src={Boy} alt="learningPath" className="w-75" />
              </div>
              <h5 className="mt-3">Learning Path Industri</h5>
              <p>Learning path disusun berdasarkan kebutuhan dunia industri.</p>
            </div>
            <div className="col-sm-3 d-flex flex-column align-items-center">
              <div className="rounded-4 bg-warning d-flex justify-content-center align-items-center py-5">
                <img src={Girl} alt="kurikulumKomprehensif" className="w-75" />
              </div>
              <h5 className="mt-3">Kurikulum Komprehensif</h5>
              <p>
                Kurikulum komprehensif dan senantiasa di perbarui berdasarkan
                pengalaman di dunia industri
              </p>
            </div>
            <div className="col-sm-3 d-flex flex-column align-items-center">
              <div className="rounded-4 bg-success d-flex justify-content-center align-items-center py-5">
                <img src={Boy2} alt="project" className="w-75" />
              </div>
              <h5 className="mt-3">Project-Based Learning</h5>
              <p>
                Materi pembelajaran disusun dengan pendekatan - pendekatan
                project based learning
              </p>
            </div>
          </div>
        </section>
        {/* Mengapa SMKDEV */}
        <section className="why">
          <div className="container">
            <div className="row justify-content-center my-5">
              <div className="w-auto">
                <div className="border-bottom border-2 border-warning text-center">
                  <h2>MENGAPA HARUS MEMILIH BELAJAR DI SMKDEV</h2>
                </div>
                <div className="mt-3 text-center">
                  <p>
                    Pengalaman belajarmu jauh lebih menyenangkan dengan
                    perpaduan belajar teori dan capstone project yang dapat
                    menjadi <br />
                    portofolio pribadimu. Tidak lupa, Kamu juga dapat belajar
                    langsung dari expert mentor dari dunia industri.
                  </p>
                </div>
              </div>
            </div>
            <div className="row align-items-sm-center flex-column-reverse flex-sm-row my-5">
              <div className="col-sm-6">
                <h3>Kurikulum</h3>
                <p>
                  Kurikulum berstandar industri global, senantiasa diperbaharui
                  <br />
                  Lorem ipsum dolor, sit amet consectetur adipisicing elit.
                  Molestias magnam temporibus minus quis ratione nam aliquam
                  maiores saepe eveniet dolorem inventore nisi in deleniti iste,
                  ab consectetur earum harum odio!
                </p>
              </div>
              <div className="col-sm-6">
                <div className="ps-sm-5">
                  <img
                    src={ImgKurikulum}
                    alt="kurikulim"
                    className="w-100 rounded-4"
                  />
                </div>
              </div>
            </div>
            <div className="row align-items-sm-center flex-column flex-sm-row my-5">
              <div className="col-sm-6">
                <div className="pe-sm-5">
                  <img
                    src={ImgMetode}
                    alt="metode"
                    className="w-100 rounded-4"
                  />
                </div>
              </div>
              <div className="col-sm-6">
                <h3>Metode</h3>
                <p>
                  Metode belajar project-based learning, dapat menjadi portfolio
                  siswa
                  <br />
                  Lorem ipsum dolor, sit amet consectetur adipisicing elit.
                  Molestias magnam temporibus minus quis ratione nam aliquam
                  maiores saepe eveniet dolorem inventore nisi in deleniti iste,
                  ab consectetur earum harum odio!
                </p>
              </div>
            </div>
            <div className="row align-items-sm-center flex-column-reverse flex-sm-row my-5">
              <div className="col-sm-6">
                <h3>Mentor</h3>
                <p>
                  Dibimbing langsung oleh mentor expert dari industri digital{' '}
                  <br />
                  Lorem ipsum dolor, sit amet consectetur adipisicing elit.
                  Molestias magnam temporibus minus quis ratione nam aliquam
                  maiores saepe eveniet dolorem inventore nisi in deleniti iste,
                  ab consectetur earum harum odio!
                </p>
              </div>
              <div className="col-sm-6">
                <div className="ps-sm-5">
                  <img
                    src={ImgMentor}
                    alt="mentor"
                    className="w-100 rounded-4"
                  />
                </div>
              </div>
            </div>
          </div>
        </section>
      </Layout>
    </>
  );
}

export default Home;
