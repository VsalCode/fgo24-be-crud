# CRUD Users Backend API

This is a RESTful API for managing user data, built with Go, Gin Gonic, and PostgreSQL. The API supports CRUD operations for user management.

## API Endpoints Documentation

The following table lists the available API endpoints for managing users:

| Method | Endpoint                  | Description                |
|--------|---------------------------|----------------------------|
| GET    | `/users`                  | Get All Users              |
| POST   | `/users`                  | Create User                |
| GET    | `/users/:id`              | Detail User                |
| PATCH  | `/users/:id`              | Update User                |
| DELETE | `/users/:id`              | Delete User                |

## Installation

1. **Clone the repository**:
   ```
   git clone https://github.com/VsalCode/fgo24-be-crud.git
   cd fgo24-be-crud
   ```

2. **Create a .env file**:
   Add the necessary environment variables in the `.env` file (e.g., database configuration).

3. **Install dependencies**:
   ```
   go mod tidy
   ```
   Additionally, install Swagger dependencies:
   ```
   go get github.com/swaggo/swag/cmd/swag
   go get github.com/swaggo/gin-swagger
   go get github.com/swaggo/files
   ```

4. **Generate Swagger Documentation**:
   Ensure your Go code includes Swagger annotations (e.g., `// @Summary`, `// @Description`) for each endpoint. Run the following command to generate Swagger JSON:
   ```
   swag init
   ```
   This generates a `docs` folder with `swagger.json` and `swagger.yaml`.

5. **Run the application**:
   ```
   go run main.go
   ```
   The API will be available at `http://localhost:8080` (or the port specified in your `.env` file).

6. **Access Swagger UI**:
   Open your browser and navigate to:
   ```
   http://localhost:8080/docs/index.html
   ```
   This displays the Swagger UI, where you can view and test the API endpoints interactively.

## How To Contribute
Pull requests are welcome! For major changes, please open an issue first to discuss your proposed changes.

## License
This project is licensed under the [MIT](https://opensource.org/license/mit) License