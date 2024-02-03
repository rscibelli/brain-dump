## Flash drive stuff

mount flash drive: `sudo mount /dev/sda1 /mnt/usb/`
https://forums.raspberrypi.com/viewtopic.php?t=48958
`lsusb`?

## GO stuff

build react project: `npm run build`
build go project: `env GOOS=linux GOARCH=arm GOARM=7 go build`

## Docker stuff

docker compose logs: `docker compose logs -f`
run docker compose: `docker compose up -d`

## Mongo shell commands

- `use local` [or whatever db you want]
- `db.createCollection("posts")`
- `db.posts.insert({name: "Robert Scibelli", date: "date", title: "first post", post: "hello there"})`
- `db.posts.insertOne({name: "Robert Scibelli", date: "date", title: "first post", post: "hello there"})`

sample data: `{name: "Robert Scibelli", date: "date", title: "first post", post: "hello there"}`