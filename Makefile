
.PHONY: build

BINARY_NAME=devpage

# build builds the tailwind css sheet, and compiles the binary into a usable thing.
build:
	go mod tidy && \
   	templ generate && \
	go generate && \
	go build -ldflags="-w -s" -o bin/${BINARY_NAME}

run: build
	@./bin/devpage

css:
	npx tailwindcss -i view/css/app.css -o public/styles.css  
clean:
	go clean

dev: css
	templ generate --watch --cmd="go generate" &\
	templ generate --watch --cmd="go run ."