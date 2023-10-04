# User Authentication API

## User Authentication REST API Design

**Design Decisions**:
1. **Modular Code Structure**: The code is organized into distinct packages for better readability and maintainability.
   - `server`: Handles all matters related to the HTTP service and its middlewares.
   - `controllers`: Contains logic for handling requests and responses.
   - `repositories`: Functions to interact directly with the database.
   - `services`: Handles business logic.

2. **JWT Authentication**: Utilized JWT (JSON Web Tokens) for user authentication.

3. **Middleware**: Introduced middleware for JWT user authentication, granting controlled access to protected routes.

4. **Viper for Configuration**: Viper library is employed for configuration, allowing for flexible configuration setups and support for diverse configuration formats.

## Libraries Choices:
- **Gin**: A web framework used to create the REST API.
- **Viper**: Library for handling configurations with support for formats such as JSON, TOML, YAML, etc.
- **JWT-go**: Library to handle JSON Web Tokens (JWT) in Go.
- **PostgreSQL**: The chosen database for the project.
- **lib/pq**: The library for interfacing Go applications with PostgreSQL databases.

## Challenges Faced:
1. **Code Structuring**: Deciding how to structure the code in a logical and maintainable manner.
2. **Authentication**: Setting up an authentication system with JWT while ensuring its security.
3. **Error Handling**: Establishing how errors are handled and propagated through various application layers.
4. **Configuration**: Organizing the configuration in a manner that's easily changeable without modifying the source code.

## Running the Application:
1. Set up the database configurations and other information in the configuration file (`config.toml`).
2. Ensure PostgreSQL database is up and running and properly configured.
3. Run the application by executing the `go run main.go` command in the root directory of the project.

## Using the API:
Use tools like Postman or similar to access the API endpoints. Examples of testing endpoints include user registration, login, and accessing resources protected with JWT authentication tokens.
