import React from "react";
import NavigationBar from "../components/NavigationBar";

const GuestLayout = ({ children }) => {
  return (
    <>
      <NavigationBar />
      {children}
    </>
  );
};

export default GuestLayout;
