echo building...

templ generate

go build -o ./api web/main/main.go

echo starting server...

./api