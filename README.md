# Account Service

This service covers user account storage

## Running the application

Start the application on port 18080 (or whatever the PORT variable is set to).

### Using Local

```bash
go run main.go
```

### Using Docker Compose

```bash
docker-compose up
```

### Using Vagrant and Docker Compose

```bash
vagrant up
```

## Try it!


### Health Api
```bash
curl -XGET localhost:18080/account/health
{
  "service": "account",
  "status": "OK",
  "time": "2018-01-14 01:59:12.984661547 +0900 JST m=+358.329539774"
}
```

### Register Api
```bash
curl -XPOST localhost:18080/account/register -d '{"username": "user1", "email": "user@dummy.com", "password": "hoge"}'
{
  "account": {
    "id": "1515862496333013291",
    "username": "user1",
    "email": "user@dummy.com"
  }
}
```

### Login Api
```bash
curl -XPOST localhost:18080/account/login -d '{"username": "user1", "password": "hoge"}'
{
  "account": {
    "id": "1515862496333013291",
    "username": "user1",
    "email": "user@dummy.com"
  }
}
```
