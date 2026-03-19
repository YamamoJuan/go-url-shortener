# Go URL Shortener

simple url shortener cuma pake go, gin, dan redis.

## Requirements
- Install [GO](https://go.dev/doc/install) dulu.
- Download [redis](https://github.com/tporadowski/redis/releases), extract terus install redis-server.exe
- Text editor pake apa aja tapi gua pake [vscode](https://code.visualstudio.com/download) sih
- udah itu aja


## Features

- Buat short URL
- Redirect to URL aslinya
- Redis storage
- REST API menggunakan Gin


## Run

Start redis

```
redis-server
```

Run app

```
go run main.go
```

Server runs on

```
http://localhost:9808
```

## API

POST /create-short-url

```
{
  "long_url": "https://google.com",
  "user_id": "juan"
}
```

dibuat dengan ikhlas oleh [Yamamo Juan](https://github.com/YamamoJuan)