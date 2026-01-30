// Utilities
import { defineStore } from "pinia";

export const useAppStore = defineStore("app", {
  state: () => ({
    //
  }),
});

export const useAuthStore = defineStore("auth", {
  state: () => ({
    user: null as null | { id: number; name: string; email: string },
  }),
});
