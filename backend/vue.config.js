module.exports = {
    // publicPath: '/wxzf/dist/', // 打包文件 静态文件 相对路径
    devServer: {
        open: true,
        //host: 'localhost',
        host: '0.0.0.0',
        port: 8081,
        https: false,
        hotOnly: false,
        proxy: {
            // 配置跨域
            '/backend': {
                target: 'http://192.168.1.102:8090/backend',
                ws: true,
                changOrigin: true,
                pathRewrite: {
                    '^/backend': ''
                }
            }
            // '/tmsug': {
            //   target: 'https://suggest.taobao.com/sug',
            //   secure: false,
            //   ws: true,
            //   changOrigin: true,
            //   pathRewrite: {
            //     '^/tmsug': ''
            //   }
            // }
        },
        before: app => { }
    }
}