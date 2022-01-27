// 导入compression-webpack-plugin
const CompressionWebpackPlugin = require('compression-webpack-plugin')
// 定义压缩文件类型
const productionGzipExtensions = /\.(js|css|json|txt|html|ico|svg)(\?.*)?$/i

module.exports = {
  transpileDependencies: ['vuetify'],
  assetsDir: 'static',
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = '欢迎来到GinWeb'
      return args
    })
    config.plugin('compressionPlugin').use(new CompressionWebpackPlugin({
      filename: '[path].gz[query]',
      algorithm: 'gzip',
      test: productionGzipExtensions,
      threshold: 10240, // 只有大小大于该值的资源会被处理 10240
      minRatio: 0.8, // 只有压缩率小于这个值的资源才会被处理
      deleteOriginalAssets: false // 删除原文件
    }))
  },
  productionSourceMap: false
}
