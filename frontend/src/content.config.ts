import { defineCollection } from 'astro:content';
import { glob } from 'astro/loaders';
import { z } from 'zod';

const artists = defineCollection({
  loader: glob({ 
    pattern: "**/*.{md,mdx}", 
    base: new URL("./content/artists", import.meta.url) 
  }),
  schema: z.object({
    title: z.string().optional(),
  }),
});

export const collections = { artists };
