build:
	go build -o game .

build-static:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o game .

build-opengl:
	go build -tags "ebitengine_opengl" -o game .