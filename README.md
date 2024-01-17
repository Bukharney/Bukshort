# Bukshort

A simple and efficient URL shortener web application built with Go for the backend and Svelte for the frontend. This project allows users to shorten long URLs into more manageable and shareable links.

## Table of Contents

- [Features](#features)
- [Documentation](#documentation)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Contributing](#contributing)
- [License](#license)

## Features

- Shorten long URLs into concise and user-friendly links.
- Easily share shortened links via social media or other platforms.
- Clean and responsive user interface.
- Secure and efficient handling of short URLs.

## Documentation

The API documentation for Bukshort can be found at [https://gbukshort.bukharney.tech/docs](https://gbukshort.bukharney.tech/docs/index.html)

## Prerequisites

Before you get started, make sure you have the following prerequisites installed on your system:

- [Go](https://golang.org/doc/install) - The Go programming language.
- [Node.js](https://nodejs.org/en/download/) - JavaScript runtime for building the Svelte frontend.
- [pnpm](https://pnpm.io/installation) - Node.js package manager.
- [Git](https://git-scm.com/downloads) - Version control system (if not already installed).

## Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/Bukharney/Bukshort.git
   ```

2. Change to the project directory:

   ```bash
   cd Bukshort
   ```

3. Install the frontend dependencies:

   ```bash
   cd frontend
   pnpm install
   ```

4. Build the frontend:

   ```bash
   pnpm run build
   ```

5. Return to the project directory:

   ```bash
   cd ..
   ```

6. Build the backend:

   ```bash
   go build
   ```

## Contributing

If you would like to contribute to this project, please follow these steps:

1. Fork the repository.

2. Create a new branch for your feature or bug fix:

   ```bash
   git checkout -b feature/your-feature
   ```

3. Make your changes and commit them:

   ```bash
   git commit -m "Your commit message"
   ```

4. Push your changes to your forked repository:

   ```bash
   git push origin feature/your-feature
   ```

5. Create a pull request on the original repository.

## License

This project is licensed under the [MIT License](LICENSE). You are free to use and modify this code for your own purposes.

---
