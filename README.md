# newsportal

Web-Based Company News Portal.

## Project Structure
```
newsportal/
├───backend/
│    ├───cmd/
│    │       root.go                                        -> Config Initialization and Executor
│    │       start.go                                       -> Server Runner
│    │
│    ├───config/
│    │       cloudflareR2.go                                -> Cloudflare Object Storage Config Loader
│    │       config.go                                      -> Environment Table Reader
│    │       database.go                                    -> Database Configuration
│    │
│    ├───database/
│    │   ├───migrations/                                    -> Database Migration File (using go-migrate)
│    │   │       000001_create_users_table.down.sql         -> Drop Table Users
│    │   │       000001_create_users_table.up.sql           -> Create Table Users
│    │   │       000002_create_categories_table.down.sql    -> Drop Table Categories
│    │   │       000002_create_categories_table.up.sql      -> Create Table Categories
│    │   │       000003_create_contents_table.down.sql      -> Drop Table Contents
│    │   │       000003_create_contents_table.up.sql        -> Create Table Contents
│    │   │
│    │   └───seeds/                                         -> Database Seed File
│    │           user_seeder.go                             -> Users table Seeder
│    │
│    ├───docs/                                              -> API Documentation using Swagger
│    │       swagger.json                                   -> Definition Table for Swagger
│    │
│    ├───internal/
│    │   ├───adapter/
│    │   │   ├───cloudflare/                                -> Cloudflare Interaction
│    │   │   │       r2.go
│    │   │   │
│    │   │   ├───handler/                                   -> Main Application Handler
│    │   │   │   │   auth_handler.go
│    │   │   │   │   category_handler.go
│    │   │   │   │   content_handler.go
│    │   │   │   │   user_handler.go
│    │   │   │   │
│    │   │   │   ├───request/                               -> User Input Request
│    │   │   │   │       auth_request.go
│    │   │   │   │       category_request.go
│    │   │   │   │       content_request.go
│    │   │   │   │       file_upload_request.go
│    │   │   │   │       user_request.go
│    │   │   │   │
│    │   │   │   └───response/                              -> Application's Response
│    │   │   │           auth_response.go
│    │   │   │           category_response.go
│    │   │   │           content_response.go
│    │   │   │           default_response.go
│    │   │   │           user_response.go
│    │   │   │
│    │   │   └───repository/                                -> Application's Repository
│    │   │           auth_repository.go
│    │   │           category_repository.go
│    │   │           content_repository.go
│    │   │           user_repository.go
│    │   │
│    │   ├───app/                                           -> Application's Endpoint List and Application's Initiator
│    │   │       app.go
│    │   │
│    │   └───core/
│    │       ├───domain/
│    │       │   ├───entity/                                -> Bridge for Request and Model
│    │       │   │       auth_entity.go
│    │       │   │       category_entity.go
│    │       │   │       content_entity.go
│    │       │   │       file_upload_entity.go
│    │       │   │       jwt_entity.go
│    │       │   │       middleware_entity.go
│    │       │   │       page_entity.go
│    │       │   │       user_entity.go
│    │       │   │
│    │       │   └───model/                                 -> Database Struct
│    │       │           category_model.go
│    │       │           content_model.go
│    │       │           user_model.go
│    │       │
│    │       └───service/                                   -> Application's Usecase (business logic)
│    │               auth_service.go
│    │               category_service.go
│    │               content_service.go
│    │               user_service.go
│    │
│    ├───lib/                                               -> Reusable functions
│    │   ├───auth/
│    │   │       jwt.go                                     -> JWT Generator and Validator
│    │   │
│    │   ├───conv/
│    │   │       conv.go                                    -> Slug generator, Generate Password Hash, Check Password Hash, and Data Type Converter
│    │   │
│    │   ├───middleware/
│    │   │       middleware.go                              -> JWT Checker
│    │   │
│    │   ├───pagination/
│    │   │       error.go                                   -> Error response for invalid pagination
│    │   │       pagination.go                              -> Global pagination method
│    │   │
│    │   └───validatorLib/
│    │           validatorLib.go                            -> Struct validator
│    │
│    ├───temp/
│    │   └───content/                                       -> Place for temporary uploaded images
│    │
│    │   .air.toml.example                                  -> Air default configuration example
│    │   .env.example                                       -> Environment table
│    │   go.mod                                             -> Go module
│    │   go.sum                                             -> Go module checksum
│    │   main.go                                            -> Main entrypoint of the application
│
│   .gitignore
│   README.md
│
├───[WIP] frontend/
```

## Packages needed for this project
### Backend
#### Air - Hot Reload for Golang
```bash
go install github.com/air-verse/air@latest
```

#### ZeroLog
```bash
go get -u github.com/rs/zerolog/log
```

#### Gorm & PostgreSQL Driver
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

#### Viper & Cobra
```bash
go get -u github.com/spf13/viper
go get -u github.com/spf13/cobra
```

#### Golang-Migrate for PostgreSQL
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

#### AWS
```bash
go get -u github.com/aws/aws-sdk-go-v2/aws
go get -u github.com/aws/aws-sdk-go-v2/config
go get -u github.com/aws/aws-sdk-go-v2/service/dynamodb
go get -u github.com/aws/aws-sdk-go-v2/service/s3
```

#### Golang-JWT
```bash
go get -u github.com/golang-jwt/jwt/v5
```

#### GoFiber V2
```bash
go get -u github.com/gofiber/fiber/v2
```

#### GoFiber Swagger (API Documentation)
```bash
go get github.com/gofiber/contrib/swagger
```

## How to Use
> Development edition
1. Rename `.air.toml.example` as `.air.toml` for Hot-Reload functionality (after installing the Air CLI)
2. **IMPORTANT!** Rename `.env.example` as `.env` for Application Environment! (application wont work without the environment table)
3. Add the value of Database, JWT Secret Key, JWT Issuer, and Cloudflare Object Storage credentials.

> Production edition
1. Change the `APP_ENV` to `production`, Swagger watched the `APP_ENV` variable
2. Pretty much the same thing as Development edition, minus renaming the Air configuration file

### Launch the application with Hot Reload (Recommended only for development)
```bash
air
```

### Launch the application without Hot Reload
```bash
go run main.go
```