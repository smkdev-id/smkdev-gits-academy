"use client";
import { ChakraProvider, extendTheme } from "@chakra-ui/react";
import React, { ReactNode } from "react";

const theme = extendTheme({
  components: {},
  colors: {
    primary: "#1c3965",
    secondary: "#00a92f",
    bg2: "#f3f8fb",
  },
  borders: {
    primary: `1px solid #1c3965`,
    secondary: `1px solid #00a92f`,
  },
});
interface CustomChakraProps {
  children: ReactNode;
}

const Providers: React.FC<CustomChakraProps> = ({ children }) => {
  return <ChakraProvider theme={theme}>{children}</ChakraProvider>;
};

export default Providers;
