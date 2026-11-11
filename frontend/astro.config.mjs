// @ts-check/
import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'astro/config';
import mdx from '@astrojs/mdx';
export default defineConfig({
  vite: {
    plugins: [tailwindcss()],
  },
  integrations: [mdx()]
});


