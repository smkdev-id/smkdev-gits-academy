/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        "custom-blue": "#1c3965",
        "custom-blue-secondary": "#004fc5",
        "custom-blue-bg": "#f5f8fe",
        "custom-color-font": "#0e1c32",
        "custom-color-font-blue": "#3b6c90",
        "custom-color-bg": " #fcfdff",
        "custom-color-red": "#f9e4e8",
        "custom-color-btn-send": "#17a8e3"
      },
      padding: {
        130: "130px",
      },
      fontFamily: {
        poppins: ["Poppins", "sans-serif"],
      },
      backgroundImage: {
        "smkdev-dashed":
          "url('https://smkdev.storage.googleapis.com/wp/Decoration3.png')",
        "talent-digital":
          "url('https://smkdev.storage.googleapis.com/wp/frame-1000003784-5736d0-575d6d.png')",
      },
    },
  },
  plugins: [],
};
