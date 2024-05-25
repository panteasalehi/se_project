# Property Marketplace Platform

Welcome to the Property Marketplace Platform! This project is an online platform for buying and selling properties, built with a Golang API and a React frontend. Docker Compose is used to streamline the build and deployment process.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)

## Introduction

The Property Marketplace Platform is designed to facilitate the seamless buying and selling of properties. Users can list their properties for sale, browse available properties, and connect with sellers.

## Features

- User authentication and authorization
- AD creation
- Browsing and searching properties
- User profiles
- Messaging system between buyers and sellers

## Tech Stack

- **Backend:** Golang
- **Frontend:** React
- **Database:** MySQL
- **Caching:** Redis
- **Containerization:** Docker, Docker Compose

## Prerequisites

Before you begin, ensure you have the following installed:

- Docker: [Install Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Install Docker Compose](https://docs.docker.com/compose/install/)

## Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/panteasalehi/SE_project.git
   cd SE_project

2. **Set up environment variables**

    Create .env files in the / and /backend directories with the necessary environment variables.
    /.env file format:
    ```plain
      # .env
      MYSQL_ROOT_PASSWORD=your_mysql_password
      MYSQL_DATABASE=MelkOnline
      REDIS_PASSWORD=your_redis_password
    ```
    /backend/.env file format:
    ```plain
      # backend/.env
      DB_CONNECTION=user:password@tcp(mysql:3306)/MelkOnline?charset=utf8&parseTime=True&loc=Local
      DB_INIT=false
      REDIS_PASS=your_redis_password
    ```
3. **Build and start the application using Docker Compose**

   ```sh
      docker-compose up --build
  This command will build and start the Golang API, React frontend, MySQL database, and Redis containers.

##  Usage
1.    Access the frontend:
      Open your browser and navigate to http://localhost:3000.

2.    API endpoints:
      The API will be running at http://localhost:8080. You can use tools like Postman or curl to interact with the API.

