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
# Getting Started with Reception Monk Web App

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

The page will reload when you make changes.\
You may also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `npm run eject`

**Note: this is a one-way operation. Once you `eject`, you can't go back!**

If you aren't satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you're on your own.

You don't have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn't feel obligated to use this feature. However we understand that this tool wouldn't be useful if you couldn't customize it when you are ready for it.

### `npm run compile-sass`

This command compiles your Sass files from sass folder and outputs them to src/assets/css folder.

### `npm run dev`
This command run development code of portal app.
