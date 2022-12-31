## INSTALL REDIS
docker run --name redis-700 -p 6379:6379 -d redis:7.0.0

## CREATE DATABASE ( )
restapi-socketio

## Configura .ENV
APP_TIMEOUT=10
APP_PORT=8989
APP_MODE=development

// MySQL Configuration
DB_HOST=localhost
DB_DRIVER=mysql
DB_USER=root
DB_PASSWORD=ardityakekasi
DB_NAME=restapi-socketio
DB_PORT=3306

// Redis Configuration
REDIS_ADDRESS=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

ALLOW_ORIGIN=http://localhost:3000,http://localhost:3001

VENDOR_PRODUCT_URL=https://dummyjson.com/products

SOCK_NS_NOTIFIKASI=/notifikasi
SOCK_EVENT_NOTIFIKASI=keluarga

## RUN API
go run .\app\main.go

## GENERATE SWAGGER
swag init -g app/main.go --output swagger/


## URL SWAGGER
http://localhost:8989/swagger/index.html

## POSTMAN COLLECTION RESTAPI
https://documenter.getpostman.com/view/2450580/2s8Z6zyWhb

## Socket Client (Postman)
Tab Setting 
- Client version : v2
- Handshake Parh : /kekasigen

Tab Event (Create)
- Event Name : join, Listen on connect : enable
- Event Name : keluarga, Listen on connect : enable


## Testing Connection Socket (Postman)
- request type : JSON
- message : {   "id_keluarga": 13, "orang_tua":null }
- choose event "join" and klik "send"

## ATTENTION
The table will be automatically created. when running the app. Make sure the database has been created and can be connected. 

## SQL File
restapi_socketio.sql