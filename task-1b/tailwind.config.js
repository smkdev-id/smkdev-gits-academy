/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/index.html"],
  theme: {
    container: {
      center: true,
      padding: "16px",
    },
    extend: {
      colors: {
        primary: "#fab107",
        secondary: "#0cbf06",
        dark: "#1e293b",
      },
      screens: {
        "2xl": "1320px",
      },
    },
  },
  plugins: [],
};
