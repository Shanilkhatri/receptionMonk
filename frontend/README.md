# Reception Monk Web App


## Used Scripts

In the project directory, you can run:

### `npm install`

This command is used to install dependencies listed in a project's package.json file.  If we run npm install without specifying any package names, npm will install all the dependencies listed in your package.json file. Then npm will create a node_modules directory in project where the packages are stored.

### `npm run compile-sass`

This is custom script defined in npm run compile-sass. It is not a built-in npm script. This command compiles Sass files into CSS. 

### `npm start`

This command Runs the application in the development mode.
Open [http://localhost:3000] to view it in your browser.

npm start depends on how the "start" script is configured in package.json. It's used for starting development servers, running build processes, and initiating other necessary tasks.

### `npm run build`

This command is used to trigger a build process that generates production-ready assets from source code. It builds the application to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

This build is optimized, minified and ready to be deployed to a web server.

**Note: It's important that you should include your node_modules directory in your project's .gitignore file to prevent it from being committed to Git.**
