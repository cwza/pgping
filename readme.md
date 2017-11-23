Ping a PostgreSQL database

## pgping usage
``` sh
Usage of ./pgping:
  -pgDs string
    	postgres datasource (default "postgres://postgres:0000@postgres:5432/meepshop?sslmode=disable")
```

## use pgping at docker-compose.yml
``` yml
app:
  image: cwza/app:latest
  links:
    - postgres
  command: > 
    sh -c "
      wget –q -nv -O pgping https://github.com/cwza/pgping/releases/download/v0.1/pgping
      chmod +x ./pgping
      while ! ./pgping; do
        echo 'Postgres is unavailable.'
        sleep 1
      done
      npm run start"
```