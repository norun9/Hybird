import eslint from '@eslint/js'
import typeScriptESLintParser from '@typescript-eslint/parser'
import eslintConfigPrettier from 'eslint-config-prettier'
import pluginImport from 'eslint-plugin-import'
import pluginReact from 'eslint-plugin-react'
import pluginUnusedImport from 'eslint-plugin-unused-imports'
import globals from 'globals'
import tseslint from 'typescript-eslint'

export default [
  eslint.configs.recommended,
  ...tseslint.configs.recommended,
  eslintConfigPrettier,
  {
    ignores: ['**/build/**'],
  },
  {
    files: ['**/*.ts', '**/*.tsx'],
  },
  {
    languageOptions: {
      globals: {
        ...globals.browser,
        ...globals.es2021,
        ...globals.node,
      },
      ecmaVersion: 'latest',
      sourceType: 'module',
      parser: typeScriptESLintParser,
    },
    settings: {
      react: {
        version: '18.2.0',
      },
    },
  },
  {
    plugins: {
      react: pluginReact,
      import: pluginImport,
      'unused-imports': pluginUnusedImport,
    },
    rules: {
      'no-undef': 'error',
      'no-unused-vars': 'off',
      'unused-imports/no-unused-imports': 'error',
      'react/react-in-jsx-scope': 'off',
    },
  },
]
