## Added Docker Network
```
$ docker network create --driver bridge kekasi-network
```

## INSTALL REDIS
```
$ docker run -d --name redis-700 -p 6379:6379 redis:7.0.0 

$ docker network connect kekasi-network redis-700
```


## INSTALL MariaDB
```
$ docker run  -d --name mariadb-10-4-30 -p 3306:3306 --env MARIADB_USER=kekasigen --env MARIADB_PASSWORD=ArdityaKekasi --env MARIADB_ROOT_PASSWORD=Kekasi.Co.ID mariadb:10.4.30 
```
```
$ docker network connect kekasi-network mariadb-10-4-30
```
```
$ docker exec -it mariadb-10-4-30 bash
$ mysql -u root -pKekasi.Co.ID
$ CREATE DATABASE restapi_socketio;
```
## Build GO & Run
```
$ docker build . -t restapi-socketio:0.3 --no-cache
```
```
$ docker run -dit -p 8989:8989 --name restapi-socketio restapi-socketio:0.3
```
```
$ docker network connect kekasi-network restapi-socketio
```
```
$ docker restart restapi-socketio
```

## Configure .env
```
APP_TIMEOUT=10
APP_PORT=8989
APP_MODE=development

// MySQL Configuration
DB_HOST=mariadb-10-4-30
DB_DRIVER=mysql
DB_USER=root
DB_PASSWORD=Kekasi.Co.ID
DB_NAME=restapi_socketio
DB_PORT=3306

// Redis Configuration
REDIS_ADDRESS=redis-700
REDIS_PORT=6379
REDIS_PASSWORD=

ALLOW_ORIGIN=http://localhost:3000,http://localhost:3001

VENDOR_PRODUCT_URL=https://dummyjson.com/products

SOCK_NS_NOTIFIKASI=/notifikasi
SOCK_EVENT_NOTIFIKASI=keluarga
```

## RUN Go
```
$ go run .\app\main.go
```

## GENERATE SWAGGER
```
$ swag init -g app/main.go --output swagger/
```


## URL SWAGGER
http://localhost:8989/swagger/index.html

## POSTMAN COLLECTION RESTAPI
https://documenter.getpostman.com/view/2450580/2s8Z6zyWhb

## Socket Client (Postman)
- Tab Setting 
    - Client version : v2
    - Handshake Parh : /kekasigen

- Tab Event (Create)
    - Event Name : join, Listen on connect : enable
    - Event Name : keluarga, Listen on connect : enable


## Testing Connection Socket (Postman)
- request type : 
    - JSON
- message : 
    - {   "id_keluarga": 13, "orang_tua":null }
- choose event "join" and click "send"

## ATTENTION
The table will be automatically created. when running the app. Make sure the database has been created and can be connected. 

## SQL File
[ restapi_socketio.sql ](restapi_socketio.sql)