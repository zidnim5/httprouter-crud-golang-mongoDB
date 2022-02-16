Hi welcome! 👋
<img src="icon.png" align="right" />



# Tech Stack

![Golang](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)&nbsp;
![Git](https://img.shields.io/badge/GIT-%23F05033.svg?&style=flat&logo=git&logoColor=white)&nbsp;
![Redis](https://img.shields.io/badge/REDIS-DC382D.svg?&style=flat&logo=redis&logoColor=white)&nbsp;
![MongoDB](https://img.shields.io/badge/MONGODB-47A248.svg?&style=flat&logo=mongodb&logoColor=white)&nbsp;

# News & Tags

> This repo contains News & Tag service and also apispec & http test.


## Requeirement

- MongoDB
     - Create database name "crud"
     - Craeate collection "news & tags"
- Redis

## Installation

Clone this repository :

```
     git clone https://github.com/zidnim5/httprouter-crud-golang-mongoDB.git
```

Adjust database.go & redis.go in app folder

Execute command below :
 
```
     go mod init golang-mongodb
     go mod tidy
```

Simple way to run this service :

```ssh
     make dev
```


## Project Structure

    .
    ├── app                     # Configuration Files
    ├── controller              # Controller Files
    ├── exception               # Exception handling Files
    ├── helper                  # Helper Files             
    ├── model                   # Model Files
    |      └── domain           
    |      └── web  
    ├── repository              # Repository Files
    ├── service                 # Service Files
    ├── apispec.json            # Service Files
    ├── docker-compose.yml
    ├── dockerfile
    ├── go.mod
    ├── go.sum
    ├── makefile
    ├── tags.http
    ├── news.http
    └── readme.md
    

## Database

This service using "MongoDB & Redis" to stored News & Tags data.


## Endpoint

To connected to this endpoint please ensure your base_url is "localhost:3000" or you can adjust your own port in main.go file



## Http Testing

You can spesific test to available http test below :

- tags.http
- news.http
