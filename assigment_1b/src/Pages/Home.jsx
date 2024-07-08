import React from "react";
import GuestLayout from "../layout/GuestLayout";
import LayoutHero from "../layout/LayoutHero";
import LayoutMitra from "../layout/LayoutMitra";
import LayoutOrientasi from "../layout/LayoutOrientasi";
import LayoutMengapa from "../layout/LayoutMengapa";
import LayoutFooter from "../layout/LayoutFooter";
import LayoutContact from "../layout/LayoutContact";
import LayoutMentor from "../layout/LayoutMentor";
import LayoutTesty from "../layout/LayoutTesty";
import LayoutBlog from "../layout/LayoutBlog";
import LayoutCountdown from "../layout/LayoutCountdown";
import LayoutTalenta from "../layout/LayoutTalenta";
import NavigationBar from "../components/NavigationBar";

const Home = ({ toggleTheme }) => {
  return (
    <>
      <GuestLayout>
        <NavigationBar toggleTheme={toggleTheme} />
        <LayoutHero />
        <LayoutMitra />
        <LayoutOrientasi />
        <LayoutMengapa />
        <LayoutTalenta />
        <LayoutMentor />
        <LayoutTesty />
        <LayoutBlog />
        <LayoutCountdown />
        <LayoutContact />
        <LayoutFooter />
      </GuestLayout>
    </>
  );
};

export default Home;
