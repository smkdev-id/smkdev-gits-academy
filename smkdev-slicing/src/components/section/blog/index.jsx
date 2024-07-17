import Header from "../ui/Header";
import Card from "./Card";

const Blog = () => {
  const articel = [
    {
      subject: "Promp Engineering: Keterampilan, Prospek Karir Dan Kursus AI",
      body: "Kenapa prompt engineering penting untuk AI generatif seperti ChatGPT dan Google Gemini? Pelajari keterampilan",
      link: "https://i.pinimg.com/736x/7c/8a/5c/7c8a5cc61e919dc25eb1a794460b735d.jpg",
      category: "General",
    },
    {
      subject: "Panduan Pemula Menguasai AI Tools (Alat AI)",
      body: "Pelajari cara menguasai alat AI dengan panduan pemula ini. Temukan langkah-langkah, alat populer, dan tips untuk sukses dalam memanfaatkan teknologi",
      link: "https://i.pinimg.com/564x/63/9f/63/639f6304fa34796c10508511e082e4e9.jpg",
      category: "Artificial Intelligent",
    },
    {
      subject: "Strategi menurunkan Churn Rate untuk Pertumbuhan Bisnis",
      body: "Memahami Churn Rate dan Strategi menurunkannya Apa itu Churn Rate? Churn rate adalah persentase pelanggan yang berhenti menggunakan produk",
      link: "https://i.pinimg.com/564x/90/cc/a5/90cca5043b39b58866eec10dcc12864e.jpg",
      category: "Data Science",
    },
    {
      subject: "Mengoptimalkan ROI dengan Strategi Analisis Data yang Efektif",
      body: "Dalam era digital yang semakin maju, data telah menjadi salah satu aset paling berharga bagi perusahaan. Namun, nilai sejati",
      link: "https://i.pinimg.com/564x/96/4f/26/964f26df472674dc168554a3b66d88a7.jpg",
      category: "Artificial Intelligent",
    },
  ];

  return (
    <div className="py-10">
      <Header title="Artikel Pilihan SMKDEV" />
      <div className="flex flex-wrap w-full gap-6 pt-5 lg:flex-nowrap md:flex-nowrap">
        {articel.map((data, i) => {
          return (
            <Card
              key={i}
              title={data.subject}
              body={data.body}
              img={data.link}
              category={data.category}
            />
          );
        })}
      </div>

      <div className="flex items-center justify-center pt-7">
        <button className="text-white bg-gray-900 btn hover:bg-black/80 gap-x-2">
          Lihat Semua Artikel 👀
        </button>
      </div>
    </div>
  );
};

export default Blog;
