// src/LandingPage.jsx
import React from 'react';
import logo from './assets/smkdev.png';
import belajarCoding from './assets/belajar-coding.jpg';
import gitsid from './assets/gitsid.png';
import eudeka from './assets/eudeka.png';
import talenta from './assets/talenta.png';
import book from './assets/book.png';
import laptop from './assets/laptop.png';
import document from './assets/document.png';
import pencil from './assets/pencil.png';
import kuri from './assets/kuri.png';
import presentation from './assets/presentation.png';

const LandingPage = () => {
  return (
    <div>
      <Header/>
      <main>
        <Menu/>

        <Mitra/>

        <Orientasi/>

        <Belajar/>

        <Talenta/>

      </main>
      <Footer/>
    </div>
  );
};

const Header = () => {
  return (
    <header>
      <nav className="navbar">
        <div className="navbar-logo">
          <img src={logo} alt="Logo" className='logo' />
        </div>
        <ul className="navbar-menu">
          <li><a href="#learn">Learn</a></li>
          <li><a href="#community">Community</a></li>
          <li><a href="#blog">Blog</a></li>
          <li><a href='#dashboard'> Dashboard</a></li>  
        </ul>
      </nav>
    </header>
  );
};

const Footer = () => {
  return (
    <footer>
        <p>Creating High-Caliber Digital Talent</p>
        <p>Summarecon Bandung, Jl. Magna Barat Blok MD No.02, Rancabolang, Kec. Gedebage, Kota Bandung, Jawa Barat 40295</p>
        <p>© 2024 SMKDEV – PT Eureka Merdeka Indonesia. All Rights Reserved.</p>
    </footer>
  );
};

// src/Menu.jsx

const Menu = () => {
  return (
    <div className="menu-item">
      <div className="menu-description">
        <h1>Jadilah Talenta Digital Masa Depan Indonesia</h1>
        <p>Belajar langsung dengan expert dari industri dengan kurikulum komperhensif berbasis project-based learning</p>
      </div>
      <div className="menu-image">
        <img src={belajarCoding} alt="Menu Image" />
      </div>
    </div>
  );
};


const Mitra = () => {
  return (

    <div className='mitra'>
      <div className='mitra-description'>
        <h3>Dipercaya Oleh industri</h3> 
      </div>
      <div className='mitra-image'>
        <img src={gitsid} alt="mitra image" />
        <img src={eudeka} alt="mitra image" />
      </div>
    </div>
  );
};

const Orientasi = () => {
  return (
    <div className="orientasi">
      <div className="orientasi-description">
        <h1>Orientasi Belajar SMKDEV</h1>
        <p>Dapatkan pengalaman belajar berorientasi pengalaman kerja yang dapat mengantarkan Anda menjadi talenta yang dibutuhkan oleh industri digital terkini.</p>
      </div>
      <div className="orientasi-cards">
        <div className="card">
          <img src={book} alt='book'/>
          <h3>Learning Path Industri</h3>
          <p>Learning path disusun berdasarkan kebutuhan dunia industri.</p>
        </div>
        <div className="card">
          <img src={document} alt='document'/>
          <h3>Kurikulum Komprehensif</h3>
          <p>Kurikulum komprehensif dan senantiasa di perbarui berdasarkan pengalaman di dunia industri</p>
        </div>
        <div className="card">
          <img src={laptop} alt='laptop'/>
          <h3>Project-Based Learning</h3>
          <p>Materi pembelajaran disusun dengan pendekatan - pendekatan project based learning</p>
        </div>
      </div>
    </div>
  );
};

const Belajar = () => {
  return (
    <div className="belajar">
      <div className="belajar-cards">
        <div className="card">
          <img src={kuri} alt='kuri'/>
          <h3>Kurikulum</h3>
          <p>Kurikulum berstandar industri global, senantiasa diperbaharui</p>
        </div>
        <div className="card">
          <img src={pencil} alt='pencil'/>
          <h3>Metode</h3>
          <p>Metode belajar project-based learning, dapat menjadi portfolio siswa</p>
        </div>
        <div className="card">
          <img src={presentation} alt='presentation'/>
          <h3>Mentor</h3>
          <p>Dibimbing langsung oleh mentor expert dari industri digital</p>
        </div>
      </div>
      <div className="belajar-description">
        <h1>MENGAPA HARUS MEMILIH BELAJAR DI SMKDEV</h1>
        <p>Pengalaman belajarmu jauh lebih menyenangkan dengan perpaduan belajar teori dan capstone project yang dapat menjadi portofolio pribadimu. Tidak lupa, Kamu juga dapat belajar langsung dari expert mentor dari dunia industri.</p>
      </div>
    </div>
  );
};

const Talenta = () => {
  return (
    <div className="talenta">
      <div className="talenta-header">
        <h1>Talenta Digital SMKDEV</h1>
        <img src={talenta} alt="Talenta Image" />
        <h1>Apa Kata Mereka?</h1>
        <p>Mereka telah merasakan serunya belajar skill digital di SMKDEV. Apakah kamu ingin seperti mereka?</p>
      </div>
      <div className="talenta-columns">
        <div className="column">
          <p>Seru! Bisa ketemu orang-orang hebat yang udah bikin banyak project keren. Bisa saling sharing pengalaman ngoding juga. Pokoknya seru deh!</p>
          <p><b>Hasban Fardani</b></p>
          <p>Siswa SMKN 11 Bandung</p>
        </div>
        <div className="column">
        <p>Kebetulan aku switch career dari guru ngaji, terus aku baca perkembangan sekarang tuh karir yang dibutuhkan di dunia IT itu gak jauh dari programming ataupun olah data. Karena kab data itu “gold”nya zaman now.Nah kebetulan SMKDEV ada Bootcamp dan webinar Data Analyst. Saya tertariknya sama materi Chrun yang ternyata bisa jadi dasar pengambilan keputusan dalam menentukan strategi untuk meningkatkan pangsa pasar atau meningkatkan loyalty customer kita.</p>
          <p><b>Rahmatulloh</b></p>
          <p>Pengajar Pondok Al Fatih Depok</p>
        </div>
        <div className="column">
        <p>Sejak bergabung dengan komunitas SMKDEV saya mendapat pengalaman yang banyak terlebih lagi melalui banyak event seperti community bonding talk yang membantu mengembangkan soft skill saya, dengan adanya coding challenge juga membuat saya bersemangat untuk meningkatkan skill pemrograman khususnya problem solving dengan lebih baik lagi, terimakasih SMKDEV dan tim</p>
          <p><b>Karel Trisnanto Utomo</b></p>
          <p>Frontend Web Developer
          POLITEKNIK HARAPAN BERSAMA</p>
        </div>
      </div>
    </div>
  );
};

export {LandingPage};