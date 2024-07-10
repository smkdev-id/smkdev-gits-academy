import { AnimatedCard } from "@/components/AnimatedCard";
import Navbar from "@/components/ui/Navigation";
import { HeroScroll } from "@/components/HeroScroll";
import SwiperCompany from "@/components/SwiperCompany";
import { Why } from "@/components/Why";

export default function Home() {
  const words = ["Global", "Masa Depan Indonesia", "beautiful", "modern"];

  return (
    <main className="flex min-h-screen flex-col items-center justify-between">
      <Navbar />
      <HeroScroll />
      <div className="w-full bg-slate-100 pb-4">
        <h1 className="flex justify-center mt-8 text-xl font-semibold ">
          Trusted by Leading Industry Partners
        </h1>
        <SwiperCompany />
      </div>
      <AnimatedCard />
      <Why />
    </main>
  );
}
