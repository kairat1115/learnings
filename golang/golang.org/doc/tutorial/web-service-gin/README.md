# Link

https://golang.org/doc/tutorial/web-service-gin

# build
```
go build
```

# run

## Unix
```
./web-service-gin
```

## Windows
```
.\web-service-gin.exe
```

# Change `curl` command, depending OS

## Unix
```
curl
```

## Windows
```
curl.exe
```

# Get Albums
```
curl "localhost:8080/albums"
```

# Get Albums by ID
```
curl "localhost:8080/albums/2"
```

# Post to albums
curl "localhost:8080/albums" --include --header "Content-Type: application/json" --data "@postAlbums.json"

