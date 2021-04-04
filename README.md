# Overview
This repository contains a REST API written with Go and a React web-app for demonstration purposes. The API is a simple counter-service which can increment by +1, decrement by -1 and return the current counter-value.




#Installation & Setup
- Make sure you have [Node.js](https://nodejs.org/en/ "Node.js") and [Go](https://golang.org/ "Go") installed.
- Install [Docker](https://www.docker.com/products/docker-desktop "Docker") if you want to use containers.
- Pull this repository and you should see two folders. The **Go** folder contains the REST API and unit tests. The **react-client** folder contains the web-application.
- Open a terminal in the **react-client** folder and run: `npm install`

#Deployment & Usage
This project supports two deployment scenarios:
1) Local deployment
2) Deployment with Docker containers

###1) Local Deployment
**Deploy REST API**
- Go inside the project and open the **Go** folder.
- Install all dependencies: `go get ./...`
- Then open a terminal inside and start the API server: `go run restAPI.go`
- The server will terminate once you close the terminal or when you press **CTRL+C**.

**Deploy React web-app**
- Go inside the project and open the **react-client** folder.
- Then open a terminal inside and start the web-app: `npm run start`
- The web-app will terminate once you close the terminal or when you press **CTRL+C**.

###2) Docker Deployment
**Build and deploy REST API container**
- Go inside the project and open the **Go** folder.
- Then open a terminal inside and build a Docker image: `docker build -t api:dev .`
- Start the container with: `docker run -p 8000:8000 api:dev`
- The server will terminate once you shut down the container.

**Build and deploy React web-app container**
- Go inside the project and open the **react-client** folder.
- Then open a terminal inside and build a Docker image: `docker build -t react:dev .`
- Start the container with: `docker run -p 3000:3000 react:dev`
- The web-app will terminate once you shut down the container.

# API Tests
###1) Unit Tests
- For the unit tests, the API-server must be running.
- Go inside the project and open the **Go** folder.
- Install all dependencies: `go get ./...`
- Then open a terminal inside and start the API server: `go run restAPI.go`
- Afterwards open another terminal and run the unit tests with: `go test`
- Following tests will be executed: TestApiHome, TestIncrement, TestMaxLimit, TestDecrement, TestMinLimit, TestResetCounter, TestReturnValue.

###2) Testing with Postman
- Install  [Postman](https://www.postman.com/downloads/"Postman").
- Make sure the API-server is up and running.
- Go inside the project and open the **Go** folder.
- Then open a terminal inside and start the API server: `go run restAPI.go`
Avaiable endpoints & examples:
- POST: http://localhost:8000/increment
Body: `{"value":110}`
Expected response: 111
- POST: http://localhost:8000/decrement
Body: `{"value":110}`
Expected response: 109
- POST: http://localhost:8000/value
Body: `{"value":110}`
Expected response: 110
- GET: http://localhost:8000/reset
Expected response: 0

###3) Testing with the web-app
- Go inside the project and open the **react-client** folder.
- Then open a terminal inside and start the web-app: `npm run start`
- You should be automatically redirected to http://localhost:3000
- Use the buttons to test. You will see the counter change accordingly.

##End
