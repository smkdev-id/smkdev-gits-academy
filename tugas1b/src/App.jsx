import FeatureList from "./components/FeatureList"
import Feature from "./components/Feature"
import HeroApps from "./components/Hero"
import NavbarApps from "./components/Navbar"
import Benefits from "./components/Benefits"
import BenefitList from "./components/BenefitList"
import MiniBanner from "./components/MiniBanner"

function App() {
  return (
    <>
      <NavbarApps />
      <HeroApps />
      <Feature />
      <FeatureList />
      <MiniBanner />
      <Benefits />
      <BenefitList />
    </>
  )
}

export default App
