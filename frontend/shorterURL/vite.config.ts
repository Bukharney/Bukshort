import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  server: {
    proxy: {
      "/shorten": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
      "/:shortURL": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  },
});
