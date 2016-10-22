#!/bin/bash
docker build -t my-dotnet-app .
docker run -it --rm --name my-running-app my-dotnet-app
#docker run -it --rm --name my-running-app my-dotnet-app
#docker run -it --name my-running-app my-dotnet-app bash
