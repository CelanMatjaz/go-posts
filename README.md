# Go-posts

Simple fullstack app written in Go where users can post posts. Project was meant primarily for learning Go.

## Run 

To run the app, first rename `example.env` to `.env` and fill in the needed values.

To run migrations, run:

```
go run ./cmd/migrations/migrations.go
```

Next, run the following command to start the app: 

```
go run ./cmd/go-posts/main.go
```

You can then access the page on `http://localhost:<PROVIDED_PORT>`.