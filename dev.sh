echo "Deleting old container"
docker kill berrydb > /dev/null
docker rm berrydb > /dev/null

echo "Starting new container"
docker run --name berrydb -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=berrydb -p 5432:5432 -d postgres  > /dev/null

sleep 5

nodemon -e go --signal SIGTERM --exec make run