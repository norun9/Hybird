module.exports = {
  extends: ['airbnb', 'airbnb/hooks', 'airbnb-typescript', 'prettier'],
  parserOptions: {
    project: './tsconfig.json',
  },
  rules: {
    // eslint-config-airbnb（eslint)
    'spaced-comment': 'off',
    'no-console': 'off',
    'no-alert': 'off',
    'arrow-body-style': 'off',
    // eslint-config-airbnb（eslint-plugin-import)
    'import/prefer-default-export': 'off',
    // eslint-config-airbnb（eslint-plugin-react）
    'react/require-default-props': 'off',
    'react/prop-types': 'off',
    'react/jsx-props-no-spreading': 'off',
    'react/react-in-jsx-scope': 'off',
    'react/function-component-definition': [2, { namedComponents: 'arrow-function' }],
    // eslint-config-airbnb（eslint-plugin-react-hook）
    'react-hooks/exhaustive-deps': 'warn',
    // eslint-config-airbnb-typescript（eslint-typescript）
    '@typescript-eslint/naming-convention': 'off',
    '@typescript-eslint/no-unused-expressions': ['error', { allowTernary: true }],
    '@typescript-eslint/no-unused-vars': 'off',
  },
}
