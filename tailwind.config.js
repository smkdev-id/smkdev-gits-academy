/* eslint-disable no-undef */
/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts}'],
  theme: {
    extend: {}
  },
  plugins: [require('daisyui')],
  daisyui: {
    themes: [
      {
        education: {
          primary: '#22d3ee',
          secondary: '#6366f1',
          accent: '#fbbf24',
          neutral: '#e5e7eb',
          info: '#2563eb',
          success: '#16a34a',
          warning: '#d97706',
          error: '#dc2626',
          'base-100': '#ffffff',
          '--rounded-box': '0.5rem',
          '--rounded-btn': '0.375rem',
          '--rounded-badge': '1.9rem',
          '--animation-btn': '0.25s',
          '--animation-input': '0.2s',
          '--btn-focus-scale': '0.95',
          '--border-btn': '2px',
          '--tab-border': '2px',
          '--tab-radius': '0.5rem'
        }
      }
    ]
  },
  darkMode: ['class', '[data-theme="education-dark"]']
}
