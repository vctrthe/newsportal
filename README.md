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