.PHONY: all

all: clean run

.PHONY: clean

clean:
	@echo "[✔️] Clean complete!"

.PHONY: run

run:
	@cd ./client && npm install && npm start &
	@cd ./go-server && go run main.go fg
	@echo "[✔️] Build complete!"
