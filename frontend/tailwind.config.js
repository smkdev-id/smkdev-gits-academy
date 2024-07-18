/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class', // or 'media'
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        mitra: '#1E293B',
        orientasi: '#222831',
        mentor: '#1C2129',
        testy: '#0F172A',
        'testy-start': '#0F172A',
        'testy-end': '#1E293B',
        cardmentor: '#13295A',
        blog: '#373A40',
        mengapa: '#0F0E0E',
        // LIGHT
        bgHeroLight: {
          DEFAULT: '#ffffff',
          start: '#ffffff', 
          end: '#cccccc', 
        },
      },
      fontFamily: {
        'poppins': ['Poppins', 'sans-serif'],
        'playwrite': ['"Playwrite GB S"', 'sans-serif'],
      },
    },
  },
  plugins: [
    require('@tailwindcss/line-clamp'),
    require('daisyui'),
  ],
}
