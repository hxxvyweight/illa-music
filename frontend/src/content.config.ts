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

const venues = defineCollection({
  loader: glob({
    pattern: "**/*.{md,mdx}",
    base: new URL("./content/venues", import.meta.url)
  }),
  schema: z.object({
    name: z.string(),
    location: z.string().optional(),
    slug: z.string().optional(),
    tier: z.string().optional(),
  }),
});

const genres = defineCollection({
  loader: glob({
    pattern: "**/*.{md,mdx}",
    base: new URL("./content/genres", import.meta.url)
  }),
  schema: z.object({
    name: z.string(),
    location: z.string().optional(),
    slug: z.string().optional(),
    tier: z.string().optional(),
  }),
});

const labels = defineCollection({
  loader: glob({
    pattern: "**/*.{md,mdx}",
    base: new URL("./content/labels", import.meta.url)
  }),
  schema: z.object({
    name: z.string(),
    location: z.string().optional(),
    slug: z.string().optional(),
    tier: z.string().optional(),
  }),
});

export const collections = { artists, venues, genres, labels };
