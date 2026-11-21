import js from "@eslint/js";
import { defineConfig } from "eslint/config";
import eslintConfigPrettier from "eslint-config-prettier/flat";
import tseslint from "typescript-eslint";

export default defineConfig(
  { ignores: ["dist/"] },
  js.configs.recommended,
  tseslint.configs.recommended,
  tseslint.configs.stylistic,
  eslintConfigPrettier,
);
