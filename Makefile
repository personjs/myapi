NAME := myapi

# Detect Windows vs Unix-based OS
ifdef OS
    ifeq ($(OS),Windows_NT)
        DETECTED_OS := Windows
        RM := del /s /q
        RMDIR := rmdir /s /q
        EXE := .exe
    else
        DETECTED_OS := Unix
        RM := rm -rf
        RMDIR := rm -rf
        EXE :=
    endif
else
    DETECTED_OS := Unix
    RM := rm -rf
    RMDIR := rm -rf
    EXE :=
endif

.PHONY: build docker run test clean

build:
	go build -o bin/$(NAME)$(EXE) cmd/main.go

docker:
	docker build -t $(NAME) .

run:
	./bin/$(NAME)$(EXE)

run-docker:
	docker run --rm -p 8080:8080 -it $(NAME)

test:
	go test ./...

clean:
	$(RMDIR) bin 2> NUL || true
