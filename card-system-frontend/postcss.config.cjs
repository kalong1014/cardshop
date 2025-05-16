// postcss.config.cjs
module.exports = {
  plugins: {
    // 注意：这里使用 @tailwindcss/postcss 而不是直接引用 tailwindcss
    '@tailwindcss/postcss': {}, 
    autoprefixer: {},
  },
};