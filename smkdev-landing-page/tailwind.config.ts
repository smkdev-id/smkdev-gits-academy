import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/sections/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: "#1c3965",
        secondary: "#00a92f",
        bg2: "#f3f8fb",
      },
      backgroundColor: {
        primary: "#1c3965",
        secondary: "#00a92f",
        bg2: "#f3f8fb",
      },
    },
  },
  plugins: [],
};
export default config;
