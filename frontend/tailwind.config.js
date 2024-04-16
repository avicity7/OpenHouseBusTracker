/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}', './src/app.html'],
  theme: {
    extend: {
      fontFamily: {
        public: ['Public Sans', 'sans-serif']
      }
    },
  },
  plugins: [],
}

