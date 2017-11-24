Ping a PostgreSQL database

## build
* linux binary
``` sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./pgping ./main.go
```
* build by platform
``` sh
go build -o ./pgping ./main.go
```

## pgping usage
``` sh
Usage of ./pgping:
  -pgDs string
    	postgres datasource (default "postgres://postgres:0000@postgres:5432/meepshop?sslmode=disable")
```

## use pgping at docker-compose.yml
``` yml
postgres:
  image: postgres:9.6
  ports:
    - "5432:5432"
app:
  image: cwza/app:latest
  links:
    - postgres
  command: > 
    sh -c "
      if [ ! -f ./pgping ]; then wget –q -nv -O pgping https://github.com/cwza/pgping/releases/download/v0.1/pgping; fi
      chmod +x ./pgping
      while ! ./pgping; do
        echo 'Postgres is unavailable.'
        sleep 1
      done
      npm run start"
```