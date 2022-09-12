# test-case-backend-majoo

## Setup Database Mysql
```bash
CREATE DATABASE db_majoo;
```

## Setup Connection
```bash
Url setting db: database/database.go:10
```

## Run Project
```bash
$ go run main.go
```

## Endpoints

### Endpoint Login
```bash
curl -X POST \
  http://localhost:8080/login \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: c96090c2-6133-36e3-3f44-55974d503aed' \
  -d '{
  "user_name": "admin1",
  "password": "admin1"
}'
```

### Endpoint Get Report By Merchant
```bash
curl -X GET \
  'http://localhost:8080/report/merchants?fromDate=2021-11-01&toDate=2021-11-10&skip=0&take=10' \
  -H 'authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImFkbWluMSIsInVzZXJfbmFtZSI6ImFkbWluMSIsImV4cCI6MTY2MzIwODA0NH0.Pfp_EsscyWhyn9QuZwlMPGOTKwbbOCMctEzZzXLmjlY' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: f41acd3e-15ae-30e5-3644-e448cdbfd9e7'
```

### Endpoint Get Report By Outlet
```bash
curl -X GET \
  'http://localhost:8080/report/outlets?fromDate=2021-11-01&toDate=2021-11-10&skip=0&take=10' \
  -H 'authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImFkbWluMSIsInVzZXJfbmFtZSI6ImFkbWluMSIsImV4cCI6MTY2MzIwODA0NH0.Pfp_EsscyWhyn9QuZwlMPGOTKwbbOCMctEzZzXLmjlY' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 35a947e0-754f-2028-1bfa-1d7ca7c6ae19'
```
