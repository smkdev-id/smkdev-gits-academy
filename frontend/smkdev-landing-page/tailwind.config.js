module.exports = {
  content: [
    './src/**/*.{js,jsx,ts,tsx}',
    // Add paths to your components here
  ],
  theme: {
    extend: {
      animation: {
        fadeIn: 'fadeIn 2s ease-in forwards',
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: 0 },
          '100%': { opacity: 1 },
        },
      },
      colors: {
        customBlue: '#1e3a8a', // Change to your preferred color
      },
      fontFamily: {
        customFont: ['Helvetica', 'Arial', 'sans-serif'], // Change to your preferred font
      },
    },
  },
  plugins: [],
};


