dir_static:=./static

all:
	swag init --output $(dir_static) --outputTypes yaml
	# build binary
	go build -o bin/app main.go

clean:
	rm -f bin/app