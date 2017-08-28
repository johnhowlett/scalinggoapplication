# GO horizontal scaling apps

Kurs von pluralsight by Mike Van Sickle  go hrizontal scaning apps

## Initial

docker network
```
docker network create --subnet=172.18.0.0/16 mynet
```
docker image dataservice
```
docker build -f Docker-dataservice -t ps/dataservice .
```
docker image web
```
docker build -f Dockerfile-web -t ps/web .
```
Start dataservice
```
docker run --name dataservice --ip=172.18.0.10 --net=mynet -P -p 4000:4000 --rm -it ps/dataservice 
```
Start web
```
docker run --name web --ip=172.18.0.11 --net=mynet -P -p 3000:3000 --rm -it -- ps/web --dataservice=http://172.18.0.10:4000
```

Für Docker for Mac musste ich immer wieder beim start das -p port:port damit ich via FireFox http://localhost:port auf den Serverice komme.

## Tag 1.0
In diesem Tag ist der Inital Code von der Webapplication dieser ist inkl. Dockerfiles das Netwerk muss vorgängig erstellt werden.

## Tag 2.0
In diesem Tag wurde das Gzip eingeführt es wurde ein eigener Handler derstellt je nach dem Browser ob dieser gzip unstützt wird der entsprechende.


