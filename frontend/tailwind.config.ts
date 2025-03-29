import type { Config } from 'tailwindcss'

const config: Config = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {
      maxWidth: {
        container: '1440px',
        contentContainer: '1140px',
        containerSmall: '1024px',
        containerXs: '768px',
      },

      screens: {
        xs: '320px',
        sm: '375px',
        sml: '500px',
        md: '667px',
        mdl: '768px',
        lg: '960px',
        lgl: '1024px',
        xl: '1280px',
      },
    },
    colors: {
      primary: '#00B207',
      subprimary: '#74E291',
      green100: '#618062',
      blue: '#387ADF',
      orange: '#FF8A00',
      black: '#333333',
      white: '#FFFFFF',
      gray100: '#C4C4C4',
      gray200: '#F2F2F2',
      gray300: '#E6E6E6',
      gray500: '#808080',
      error: '#FF004D',
    },
  },
  plugins: [],
}
export default config
