import { Topnav } from "@/components";
import { Featured, Hero, Mitra } from "@/sections";

export default function Home() {
  return (
    <main className="w-full">
      <Topnav />
      <Hero />
      <Mitra />
      <Featured />
    </main>
  );
}
