BIN=../bin/

build:
	cd game2; go build -o $(BIN)game2 .

build-static:
	cd game2; CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BIN)game2-st .

build-opengl:
	cd game2; go build -tags "ebitengine_opengl" -o $(BIN)game2-gl .