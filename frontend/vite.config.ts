import vue from "@vitejs/plugin-vue";
import ui from "@nuxt/ui/vite";
import { defineConfig } from "vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    ui({
      ui: {
        colors: { primary: "sky", neutral: "mist" },
        button: { defaultVariants: { variant: "soft" } },
        badge: { defaultVariants: { variant: "soft" } },
        input: { defaultVariants: { variant: "soft" } },
        select: { defaultVariants: { variant: "soft" } },
        textarea: { defaultVariants: { variant: "soft" } },
        selectMenu: { defaultVariants: { variant: "soft" } },
        inputMenu: { defaultVariants: { variant: "soft" } },
        inputNumber: { defaultVariants: { variant: "soft" } },
        inputTags: { defaultVariants: { variant: "soft" } },
        inputDate: { defaultVariants: { variant: "soft" } },
        inputTime: { defaultVariants: { variant: "soft" } },
        pinInput: { defaultVariants: { variant: "soft" } },
      },
    }),
  ],
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  },
  build: {
    outDir: "../dist",
    emptyOutDir: true,
  },
});
