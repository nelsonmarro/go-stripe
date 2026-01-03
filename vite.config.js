import { resolve } from 'path';
import { defineConfig } from 'vite';
import { globSync } from 'glob';
import path from 'node:path';

// Dynamically find all .js files in the assets/js directory
const entryPoints = globSync('assets/js/**/*.{js,ts}').reduce((acc, file) => {
    // Use the filename (without extension) as the entry point name
    const name = path.basename(file, path.extname(file));
    acc[name] = resolve(__dirname, file);
    return acc;
}, {});

export default defineConfig({
  build: {
    // Set the output directory
    outDir: 'static/js',
    rollupOptions: {
      // Define the entry points
      input: entryPoints,
      output: {
        // Ensure the output filename matches the entry point name
        entryFileNames: `[name].js`,
        chunkFileNames: `chunks/[name].js`,
        assetFileNames: `assets/[name].[ext]`
      }
    },
  },
});
