jet -dsn="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -schema=public -path=.gen
cd cmd/urbathon-backend-2023
echo "build started"
go build
./urbathon-backend-2023 -srvAddr=webapi -srvPort=8080
