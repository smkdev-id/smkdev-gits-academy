import { Footer, Topnav } from "@/components";
import {
  Article,
  Contact,
  Featured,
  Hero,
  Mentor,
  Mitra,
  TalentDigital,
  Testimoni,
  WhyUs,
} from "@/sections";

export default function Home() {
  return (
    <main className="w-full">
      <Topnav />

      <Hero />
      <Mitra />
      <Featured />
      <WhyUs />
      <TalentDigital />
      <Testimoni />
      <Mentor />
      <Article />
      <Contact />

      <Footer />
    </main>
  );
}
