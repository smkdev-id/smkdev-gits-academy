import MitraSection from "./components/MitraSection";
import HeroSection from "./components/HeroSection";
import Navbar from "./components/Navbar";
import OrientasiSection from "./components/OrientasiSection";
import KurikulumSection from "./components/KurikulumSection";
import TalentaDigital from "./components/TalentaDigitalSection";
import ApaKataMerekaSection from "./components/ApaKataMerekaSection";
import ExpertMentorSection from "./components/ExpertMentorSection";
import Artikel from "./components/Article";
import Footer from "./components/Footer";

function App() {
  return (
    <main>
      <Navbar />
      <HeroSection />
      <MitraSection />
      <OrientasiSection />
      <KurikulumSection />
      <TalentaDigital />
      <ApaKataMerekaSection />
      <ExpertMentorSection />
      <Artikel />
      <Footer />
    </main>
  );
}

export default App;
