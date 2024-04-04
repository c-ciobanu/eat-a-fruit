run:
	wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: go run cmd/main.go

build:
	go build -o main cmd/main.go