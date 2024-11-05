
# Party Plot Booking Application

This repository contains a Party Plot Booking Application developed in Go. It demonstrates how to build a web application with a backend in Go that connects to MongoDB, along with a front end built using React.

## Table of Contents
- [History](#history)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Why Go?](#why-go)
- [Credits](#credits)

## History

This project was inspired by the full course on building web applications with Go by [Nana Jenshia](https://www.youtube.com/@TechWorldwithNana). The original repository for this course can be found [here](https://gitlab.com/nanuchi/go-full-course-youtube).

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Go**: You need to have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).
- **MongoDB**: A MongoDB instance is required. You can use [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) for a cloud solution or set up a local instance.
- **Node.js**: For the front-end React application, you need Node.js installed. You can download it from [the official Node.js website](https://nodejs.org/).

## Getting Started

To clone this repository, run the following command in your terminal:

```bash
git remote add origin https://github.com/BhautikVekariya21/goLang.git
```

After cloning, navigate to the project directory:

```bash
cd go
```

### Setting Up the Backend

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Create a `.env` file in the backend directory and add your MongoDB URI:
   ```env
   MONGODB_URI=mongodb+srv://<username>:<password>@cluster0.vwvev.mongodb.net/<dbname>?retryWrites=true&w=majority
   PORT=8081
   ```

3. Install Go dependencies:
   ```bash
   go mod tidy
   ```

4. Run the backend server:
   ```bash
   go run main.go
   ```

### Setting Up the Frontend

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install Node.js dependencies:
   ```bash
   npm install
   ```

3. Start the React application:
   ```bash
   npm start
   ```

## Usage

Once both the backend and frontend are running, you can access the application at `http://localhost:8080`.

## Why Go?

- **Simplicity and Efficiency**: Go's syntax is simple, making it easy to learn and use. Its concurrency model allows for efficient processing of tasks.
- **Performance**: Go is compiled to machine code, which makes it faster than interpreted languages.
- **Strong Standard Library**: Go comes with a rich standard library that helps with building networked and web applications.
- **Robustness**: The static typing and garbage collection features help developers write reliable and maintainable code.

## Credits

This project was inspired by the tutorial series from [Nana Jenshia](https://www.youtube.com/@TechWorldwithNana). Special thanks to TechWorld with Nana for their guidance and resources throughout this project.

Feel free to contribute to this project or provide feedback!

```

### Notes:
- Replace `<username>`, `<password>`, and `<dbname>` in the `.env` section with the actual credentials for your MongoDB.
- You can add more details or customize any part of this README to fit your project better. Let me know if you need any changes or additional sections!