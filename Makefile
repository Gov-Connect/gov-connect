.PHONY: all

all: clean dataprep run

.PHONY: clean

clean:
	@echo "[✔️] Clean complete!"

.PHONY: dataprep

dataprep:
	@python3 ./go-server/data-utils/test_data_transform.py
	@echo "[✔️] Data preparation steps complete"
.PHONY: run

run:
	@cd ./client && npm install && npm start &
	@cd ./go-server && go run main.go fg
	@echo "[✔️] Build complete!"
