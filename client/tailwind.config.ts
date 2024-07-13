import type { Config } from 'tailwindcss'

const config: Config = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
    './src/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {
      colors: {
        'gray-56': 'rgb(56, 58, 66)',
        'gray-50': 'rgb(50, 51, 56)',
        'gray-43': 'rgb(43, 45, 49)',
        'gray-48': 'rgb(48, 49, 54)',
        'gray-30': 'rgb(30, 31, 34)',
        'gray-light': 'rgb(220, 222, 225)',
        'gray-border-1': 'rgb(54, 54, 60)',
        'gray-border-2': 'rgb(60, 61, 68)',
        'gray-border-3': 'rgb(64, 65, 71)',
      },
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'gradient-conic': 'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
      },
    },
  },
  plugins: [],
}
export default config
