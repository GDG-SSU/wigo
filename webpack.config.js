module.exports = {
    context: __dirname + "/public/js",
    entry: './entry.js',
    output: {
        path: __dirname + '/public/build',
        filename: 'bundle.js'
    },
    module: {
        loaders: [
            { test: /\.css/, loader: 'style-loader!css-loader' }
        ]
    }
};
