const path = require('path');

module.exports = {
	entry: {
		login: './src/ts/login.ts',  // Entry point for the login page
		signup: './src/ts/signup.ts',  // Entry point for the login page
		main: './src/ts/main.ts'     // Entry point for the main page
	},
	output: {
		filename: '[name].bundle.js',  // Output filename
		path: path.resolve(__dirname, 'static/js'),  // Output directory
	},
	module: {
		rules: [
			{
				test: /\.(ts|js)x?$/,  // Regex to handle both TypeScript and JavaScript files
				use: 'ts-loader',
				exclude: /node_modules/,
			}
		],
	},
	resolve: {
		extensions: ['.tsx', '.ts', '.js'],  // Resolve these file types
	},
	mode: 'production',
	devtool: false
};
