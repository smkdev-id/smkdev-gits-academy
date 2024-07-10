import Navbar from "@/components/ui/Navigation";
import { HeroScroll } from "@/components/HeroScroll";
import { WhyCard } from "@/components/WhyCard";
import Image from "next/image";
import CompanySwiper from "@/components/CompanySwiper";
import { OrientationCard } from "@/components/OrientationCard";
import { TestimonialsScroll } from "@/components/TestimonialsScroll";
import { MentorCard } from "@/components/MentorCard";
import { ArticelBento } from "@/components/ArticelBento";
import Contact from "@/components/Contact";
import { Footer } from "@/components/Footer";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between">
      <Navbar />
      <HeroScroll />
      <div className="w-full bg-slate-100 pb-4">
        <h1 className="flex justify-center mt-8 text-xl font-semibold ">
          Trusted by Leading Industry Partners
        </h1>
        <CompanySwiper />
      </div>
      <OrientationCard />
      <WhyCard />
      <div className="flex flex-col justify-center p-20 mt-20">
        <h1 className="text-3xl font-bold justify-center flex">
          Talenta Digital SMKDEV
        </h1>
        <Image
          src="/professional_5_steps.png"
          width="1000"
          height="100"
          alt="decoration"
        />
      </div>
      <TestimonialsScroll />
      <MentorCard />
      <ArticelBento />
      <Contact />
      <Footer />
    </main>
  );
}
