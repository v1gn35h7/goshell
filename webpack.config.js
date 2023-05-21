const debug = process.env.NODE_ENV !== "production";
const webpack = require("webpack");
const path = require("path");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const { BundleAnalyzerPlugin } = require("webpack-bundle-analyzer");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const CompressionPlugin = require("compression-webpack-plugin");
const TerserPlugin = require("terser-webpack-plugin");

module.exports = {
  entry: "./frontend/src/index.js",
  output: {
    path: path.join(__dirname, "./frontend/dist/"),
    filename: "js/[name].bundle.min.js",
    chunkFilename: "js/[name].bundle.js",
    asyncChunks: true
  },
  devServer: {
    static: {
      directory: path.join(__dirname, './frontend/dist/'),
    },
    port: 3000,
    historyApiFallback: true,
    proxy: {
      '/api': {
        target: 'http://localhost:8080/',
        changeOrigin: true,
      },
    }
  },
  resolve: {
    extensions: [".js", ".jsx"]
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /(node_modules)/,       
        use: {
          loader: "babel-loader",
          options: {
            presets: ["@babel/env", "@babel/preset-react"],
            plugins: [
              "@babel/plugin-proposal-class-properties",
              "@babel/plugin-syntax-dynamic-import",
              "@babel/plugin-proposal-object-rest-spread",
              ["@babel/plugin-proposal-decorators", { legacy: true }]
            ]
          }
        }
      },      
      {
        test: /\.(sa|sc|c)ss$/,
        use: debug
          ? [
              {
                loader: "style-loader"
              },
              {
                loader: "css-loader"
              },
              {
                loader: "sass-loader"
              }
            ]
          : [
              MiniCssExtractPlugin.loader,
              "css-loader",
              "sass-loader"
            ]
      },
      {
        test: /\.(eot|ttf|woff|woff2|otf|svg)$/,
        use: [
          {
            loader: "url-loader",
            options: {
              limit: 100000,
              name: "./frontend/fonts/[name].[ext]"
              // publicPath: '../'
            }
          }
        ]
      },
      {
        test: /\.(gif|png|jpe?g)$/i,
        use: [
          {
            loader: "file-loader",
            options: {
              outputPath: "./frontend/dist/"
            }
          }
        ]
      }
    ]
  },
  plugins: [
    // define NODE_ENV to remove unnecessary code
    // new webpack.DefinePlugin({
    //   "process.env.NODE_ENV": JSON.stringify("production")
    // }),
    //new webpack.optimize.OccurrenceOrderPlugin(),
    //new webpack.optimize.AggressiveMergingPlugin(), // Merge chunks
    // extract imported css into own file
    new MiniCssExtractPlugin({
      // Options similar to the same options in webpackOptions.output
      // both options are optional
      filename: "[name].css",
      chunkFilename: "[id].css"
    }),
    new webpack.LoaderOptionsPlugin({
      minimize: true
    }),
    //new webpack.IgnorePlugin(/^\.\/locale$/, /moment$/),
    new HtmlWebpackPlugin({
      template: "./frontend/templates/react.ejs",
        minify: {
        collapseWhitespace: true,
        removeAttributeQuotes: false
      }
    }), //{
      //template: "./frontend/public/index.html"
      // minify: {
      //   collapseWhitespace: true,
      //   removeAttributeQuotes: false
      // }
    //}),
    new CompressionPlugin({
      test: /\.(html|css|js|gif|svg|ico|woff|ttf|eot)$/,
      exclude: /(node_modules)/
    }),
    
   // new BundleAnalyzerPlugin()
  ],
  optimization: {}
};