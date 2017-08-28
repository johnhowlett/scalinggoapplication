# GO horizontal scaling apps

Kurs von pluralsight by Mike Van Sickle  go hrizontal scaning apps

## Initial

docker network
```
docker network create --subnet=172.18.0.0/16 mynet
```
docker image dataservice
```
docer build -f Docker-dataservice -t ps/dataservice .
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
