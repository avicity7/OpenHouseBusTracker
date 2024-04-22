/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}', './node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      fontFamily: {
        public: ['Public Sans', 'sans-serif']
      }
    },
  },
  plugins: [],
  darkMode: 'class'
}

