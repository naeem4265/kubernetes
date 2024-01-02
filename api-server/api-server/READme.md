# API Server for Book Management

## Introduction

This project is a simple API server written in Go that provides functionality for authenticating users and managing a collection of books (referred to as albums). It supports user sign-in, sign-out, and CRUD operations (Create, Read, Update, Delete) for albums. The server uses JWT (JSON Web Tokens) for user authentication.

## Features

1. **User Authentication:**
   - **Sign In:** Users can authenticate by sending a POST request to `/signin`.
   - **Sign Out:** Users can sign out by sending a GET request to `/signout`.

2. **Album Management:**
   - **Get Albums:** Retrieve a list of all albums by sending a GET request to `/albums`.
   - **Get Album by ID:** Retrieve a specific album by its ID with a GET request to `/albums/{id}`.
   - **Insert Album:** Add a new album with a POST request to `/albums`.
   - **Update Album:** Modify an existing album with a PUT request to `/albums/{id}`.
   - **Delete Album:** Remove an album by its ID with a DELETE request to `/albums/{id}`.

3. **Authentication Middleware:**
   - The server uses a middleware function to check the validity of JWT tokens for secured routes.
   - Token verification is implemented using a secret key (`JWTKey`).

## Usage

1. **Sign In:**
   - Send a POST request to `/signin` with appropriate credentials.
   - If successful, a JWT token will be provided in the response.

2. **Access Secured Routes:**
   - Include the JWT token in the request header for secured routes.

3. **Album Endpoints:**
   - Access album-related endpoints (`/albums`) to perform CRUD operations.

4. **Sign Out:**
   - Send a GET request to `/signout` to invalidate the current token.

## Dependencies

- [github.com/go-chi/chi/v5](https://github.com/go-chi/chi): Lightweight and fast HTTP router for Go.
- [github.com/golang-jwt/jwt/v4](https://github.com/golang-jwt/jwt): JSON Web Token implementation for Go.

## How to Run

1. Clone the repository:

   ```bash
   git clone <repository_url>
   cd <repository_directory>
   ```

2. Install dependencies:

   ```bash
   go get -u ./...
   ```

3. Build and run the server:

   ```bash
   go build
   ./<executable_name>
   ```

   The server will be running on `http://localhost:8080`.

## Configuration

- The server uses a default port (`8080`), but you can modify it in the `main` function if needed.

## Conclusion

This API server provides a foundation for building a book management system. Developers can extend and customize it to suit their specific requirements for book-related applications. Feel free to enhance and contribute to the project for additional features and improvements.