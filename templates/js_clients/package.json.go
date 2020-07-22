package js_clients

const PACKAGE_JSON = `
{
  "name": "{{.ProjectName}}",
  "version": "1.0.0",
  "description": "Bootstrap {{.ProjectName}} Project",
  "main": "main.js",
  "scripts": {
    "build": "webpack",
    "start": "webpack-dev-server --host 0.0.0.0 --hot --inline"
  },
  "author": "",
  "license": "ISC",
  "dependencies": {
    "@improbable-eng/grpc-web": "^0.11.0",
    "cogment": "^0.3.0-alpha5",
    "google-protobuf": "^3.7.1",
    "grpc-web": "^1.0.4",
    "ts-protoc-gen": "^0.12.0"
  },
  "devDependencies": {
    "webpack": "^4.41.5",
    "webpack-cli": "^3.3.8",
    "webpack-dev-server": "^3.10.1"
  }
}
`
