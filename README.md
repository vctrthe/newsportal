# newsportal

Web-Based Company News Portal.

## Project Structure
```
newsportal/
    backend/                    -> Built using Go
        cmd/                    -> starting application/entrypoint
        config/                 -> configuration file
        database/               -> db migration file (using go-migrate)
        internal/               -> application core
            adapter/            -> HTTP handler, repository layer
                repository/     -> db interaction
                handler/        -> HTTP handler
                cloudlfare/     -> Cloudflare R2 interaction
            app/                -> application initiator and bridge to another layer, endpoint housing
                app.go
            core/
                domain/         -> entity/model
                service/        -> usecase (main complex business logic)
        lib/                    -> reuseable functions
            conv/
                conv.go         -> string-int or vice versa conversion, or other reuseable functions
            jwt/                -> generate and validate JWT
    
    [WIP] frontend/
```

### Packages needed for this project
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