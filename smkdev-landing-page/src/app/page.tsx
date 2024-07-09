import { Topnav } from "@/components";
import { Featured, Hero, Mitra, WhyUs } from "@/sections";

export default function Home() {
  return (
    <main className="w-full">
      <Topnav />
      <Hero />
      <Mitra />
      <Featured />
      <WhyUs />
    </main>
  );
}
