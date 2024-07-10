"use client";

import React from "react";
import { InfiniteMovingCards } from "./ui/InfiniteMovingCards";

export function TestimonialsScroll() {
  return (
    <div className="mt-40">
      <h1 className="text-3xl font-bold justify-center flex">
        Apa Kata <span className="text-[#004FC5] ml-2">Mereka?</span>
      </h1>
      <h1 className="text-xl px-20 flex justify-center text-center pt-4">
        Mereka telah merasakan serunya belajar skill digital di SMKDEV. Apakah
        kamu ingin seperti mereka?
      </h1>
      <div className="h-[30rem] rounded-md flex flex-col antialiased bg-white items-center justify-center relative overflow-hidden">
        <InfiniteMovingCards
          items={testimonials}
          direction="right"
          speed="slow"
        />
      </div>
    </div>
  );
}

const testimonials = [
  {
    quote:
      "Merupakan pengalaman yang sangat menyenangkan bagi saya, mencoba hal baru dan memecahkan suatu masalah. Selain itu, saya juga mendapatkan kesempatan untuk menjalin banyak hubungan dan koneksi dengan orang-orang yang memiliki minat serupa. Dengan ini, saya dapat berbagi pengetahuan dan pengalaman saya dengan orang lain, serta menerima wawasan dan ilmu baru melalui berbagai perspektif yang berbeda.",
    name: "Muhammad Iqbal Pasha Al Farabi",
    title: "Teknik Infomatika ITENAS",
  },
  {
    quote:
      "Saya merasa sangat bahagia dan bangga karena berhasil menjadi pemenang dalam SMKDEV Coding Challenge. Selama kompetisi ini, menurut saya, pengalaman ini telah memperluas wawasan dan keterampilannya dalam pemrograman. Saya berterima kasih kepada tim penyelenggara SMKDEV karena memberikan kesempatan ini dan memberikan tantangan yang menginspirasi serta mengapresiasi kerja keras peserta lain. SMKDEV Coding Challenge adalah pengalaman berharga yang telah memotivasi saya untuk terus belajar dan berkembang dalam dunia pemrograman.",
    name: "Shevabey Rahman",
    title: "Sistem Informasi Universitas Ahmad Dahlan",
  },
  {
    quote:
      "Pada program SMKDEV Scholarship ini memberi saya kesempatan untuk berkenalan dan berinteraksi langsung dengan pengajar profesional di bidang data analyst. Pada program ini juga saya mendapatkan wawasan tentang berbagai peluang karier dalam analisis data dan teknologi informasi dan membantu saya merencanakan masa depan karir di data analyst.",
    name: "Guruh Maulana",
    title: "STMIK IKMI Cirebon",
  },
  {
    quote:
      "Scholarship SMKDEV itu kece banget, walaupun aku dapat Scholarship 100%, kualitas kelasnya gak sembarangan. Kelas meet yang diadain malem bikin aku gak bingung bagi waktu sama kegiatan lainnya. Terus, kak Samuel juga ramah banget. Gak kayak yang bikin aku jadi 'segan bertanya'. Materinya juga oke, bikin otakku berasa pening nyusun kode buat submission. Terima kasih, SMKDEV!",
    name: "Tiara Febrianie",
    title: "SMKN 11 Malang",
  },
  {
    quote:
      "SMKDEV Scholarship ini sangat menarik, karena memiliki penyampaian yang seru dan mudah dipahami. Selain itu juga di dalam bootcamp ini selalu ada di sela-sela kegiatan mengadakan mini games berhadiah",
    name: "Rizal Maulana",
    title: "SMK NEGERI 1 CIMAHI",
  },
  {
    quote:
      "Seru! Bisa ketemu orang-orang hebat yang udah bikin banyak project keren. Bisa saling sharing pengalaman ngoding juga. Pokoknya seru deh!",
    name: "Hasban Fardani",
    title: "Siswa SMKN 11 Bandung",
  },
];
