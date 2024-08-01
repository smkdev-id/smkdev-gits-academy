import { Footer, Navbar } from "../components";
import {
  About,
  Explore,
  GetStarted,
  Hero,
  Insights,
  WhatsNew,
} from "../sections";

const Page = () => (
  <div className="bg-secondary-white overflow-hidden">
    <Navbar />
    <Hero />
    <About />
    <Explore />
    <GetStarted />
    <WhatsNew />
    <Insights />
    <Footer />
  </div>
);

export default Page;
