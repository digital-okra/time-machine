import svelte from "rollup-plugin-svelte";
import resolve from "rollup-plugin-node-resolve";
import commonjs from "rollup-plugin-commonjs";
import livereload from "rollup-plugin-livereload";
import postcss from 'rollup-plugin-postcss';
import { terser } from "rollup-plugin-terser";

const isDev = Boolean(process.env.ROLLUP_WATCH);

export default [
  // Browser bundle
  {
    input: "src/index.js",
    output: {
      sourcemap: true,
      format: "iife",
      name: "app",
      file: "public/bundle.js"
    },
    plugins: [
      svelte({
        hydratable: true,
        emitCss: true,
        css: css => {
          css.write("public/bundle.css");
        }
      }),
      resolve({
        browser: true,
        dedupe: importee => importee === 'svelte' || importee.startsWith('svelte/')
      }),
      commonjs(),
      postcss({
        extract: true,
        minimize: true,
        use: [
          ['sass', {
            includePaths: [
              './theme',
              './node_modules'
            ]
          }]
        ]
      }),
      // App.js will be built after bundle.js, so we only need to watch that.
      // By setting a small delay the Node server has a chance to restart before reloading.
      isDev &&
        livereload({
          watch: "public/App.js",
          delay: 200
        }),
      !isDev && terser()
    ]
  },
  // Server bundle
  {
    input: "src/App.svelte",
    output: {
      sourcemap: false,
      format: "cjs",
      name: "app",
      file: "public/App.js"
    },
    plugins: [
      svelte({
        generate: "ssr"
      }),
      resolve(),
      commonjs(),
      postcss({
        extract: true,
        minimize: true,
        use: [
          ['sass', {
            includePaths: [
              './theme',
              './node_modules'
            ]
          }]
        ]
      }),
      !isDev && terser()
    ]
  }
];
