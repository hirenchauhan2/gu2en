import commonjs from '@rollup/plugin-commonjs';
import eslint from '@rollup/plugin-eslint';
import terser from '@rollup/plugin-terser';

/**
 * @type {import('rollup').RollupOptions}
 */
const config = [
  {
    input: './src/main.js',
    plugins: [eslint(), commonjs()],
    output: [
      {
        file: 'dist/gu2en.cjs',
        format: 'cjs',
      },
      {
        file: 'dist/gu2en.esm.js',
        format: 'es',
      },
      {
        file: 'dist/gu2en.umd.min.js',
        format: 'umd',
        name: 'gu2en',
        sourcemap: 'inline',
        plugins: [terser()],
      },
    ],
  },
];

export default config;
