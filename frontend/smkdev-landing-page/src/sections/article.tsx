import { ArticleCard, MediaQuery } from "@/components";
import { NextPage } from "next";

interface Props {}

const Article: NextPage<Props> = ({}) => {
  const ARTICLE = [
    {
      title: "Prompt Engineering: Keterampilan, Prospek Karir Dan Kursus AI",
      desc: "Kenapa prompt engineering penting untuk AI generatif seperti ChatGPT dan Google Gemini? Pelajari keterampilan, prospe",
      tag: "general",
      img: "https://smkdev.storage.googleapis.com/wp/Visualization-Of-Human-and-AI-Collaboration.png",
      url: "#",
    },
    {
      title: "Mengoptimalkan ROI dengan Strategi Analisis Data yang Efektif",
      desc: "Dalam era digital yang semakin maju, data telah menjadi salah satu aset paling berharga bagi perusahaan. Namun, nilai sejati",
      tag: "",
      img: "https://smkdev.storage.googleapis.com/wp/strategi-efektif-churnrate.png",
      url: "#",
    },
    {
      title: "Strategi menurunkan Churn Rate untuk Pertumbuhan Bisnis",
      desc: "Memahami Churn Rate dan Strategi menurunkannya Apa itu Churn Rate? Churn rate adalah persentase pelanggan yang berhenti menggunakan produk atau",
      tag: "data science",
      img: "https://smkdev.storage.googleapis.com/wp/Visualization-Of-Human-and-AI-Collaboration.png",
      url: "#",
    },
    {
      title: "Panduan Pemula Menguasai AI Tools (Alat AI)",
      desc: "Pelajari cara menguasai alat AI dengan panduan pemula ini. Temukan langkah-langkah, alat populer, dan tips untuk sukses dalam memanfaatkan teknologi",
      tag: "",
      img: "https://smkdev.storage.googleapis.com/wp/strategi-efektif-churnrate.png",
      url: "#",
    },
  ];
  return (
    <div className="flex w-full justify-center py-20">
      <MediaQuery className="flex flex-col items-center justify-center gap-y-10">
        <h1 className="text-3xl font-semibold text-primary">
          Artikel Pilihan SMKDEV
        </h1>

        <div className="gird-cols-1 grid w-full gap-6 lg:grid-cols-2">
          {ARTICLE.map((i, index) => {
            return (
              <ArticleCard
                key={index}
                title={i.title}
                tag={i.tag}
                desc={i.desc}
                img={i.img}
                url={i.url}
              />
            );
          })}
        </div>
      </MediaQuery>
    </div>
  );
};

export default Article;
