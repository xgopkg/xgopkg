module.exports = function override(config, env) {
  const { injectBabelPlugin } = require('react-app-rewired');``
  // do stuff with the webpack config...
  config = injectBabelPlugin(['import', { libraryName: 'antd', libraryDirectory: 'es', style: 'css' }], config);
  return config;
};