# Estructura Projecto GO UnoTV

# Config

go mod init github.com/etcsilver/go-skeleton.git

# Posibles Dependencias
go get -u github.com/gin-gonic/gin
go get -u github.com/rs/zerolog/log
go get -u github.com/joho/godotenv
go get -u gopkg.in/mgo.v2
go get -u github.com/zsais/go-gin-prometheus
go get -u github.com/swaggo/swag/cmd/swag


`LOGS`
```LOG

```

# DOC
https://swagger.unotv.com


* **Ejemplo curl:**

```
### Pruebas Locales
curl -v -X POST \
  http://localhost:9037/method_name \
  -H 'content-type: application/json' \
  -d '{ "id": "46169" }'

```

export PATH=$(go env GOPATH)/bin:$PAT
go get -u github.com/swaggo/swag/cmd/swag
swag init -g cmd/main.go -o /Users/fernando/containers/swagger/unotv/doc --ot json --instanceName apiname