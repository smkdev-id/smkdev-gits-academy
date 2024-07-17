/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    screens: {
      sm: '480px',
      md: '768px',
      lg: '976px',
      xl: '1440px'
    },
    extend: {
      colors: {
        accent: 'rgba(153, 102, 255, 0.7)',
        dark: 'rgba(31,41,55,255)',
        // grey: 'rgba(246,248,253,255)',
      }
    },
  },
  plugins: [
    require('daisyui'),
    'prettier-plugin-tailwindcss',
  ],
}

