# Image-Processor
This project goal is to handle image-processing with simple handler. 

The source code layout were separated into 4 group based on separation of concern. Those groups are:
- Framework 
  - This layer is consist of infrastucture of the project such as databases, logger, error, utils function, .etc

- Application 
  - This layer is focusing as presentation layer to the client such as http, gRPC, websocket, .etc. In terms of http all of the handler function will be define in this layer

- Usecase
  - This layer is focusing on business logic of the application like retrieve data, transform it, and manipulating the data output  

- Entity
  - This layer is consist of all data sources of the projects like query to the databases, fetching data from external resources through API, .etc. 


# How to run (Locally)

- clone the project and cd to the project
```
git clone https://github.com/ihsanul14/image-processor.git
cd /path/to/project
```

- install additional dependency. Follow this link :
  - gocv: https://gocv.io/getting-started/
  - ffmpeg: https://ffmpeg.org/download.html/

- running the application
```
go run .
```

- access the application to http://localhost:30001

# How to run (Docker)

- clone the project and cd to the project
```
git clone https://github.com/ihsanul14/image-processor.git
cd /path/to/project
```

- install docker. Follow this link :
  - docker: https://docs.docker.com/engine/install

- verify the installation
```
docker version
```

- running the application
```
docker build . -t image-processor
docker run -v /local/path:/go/src/image-processor/framework/output -p 30001:30001 image-processor # replace the local path to your preferred destination path
```

- access the application to http://localhost:30001

# Run Unit Test

```
go test ./... --coverprofile=coverage.out
```

# Test the Application

[API Documentation](./image-processor.postman_collection.json)