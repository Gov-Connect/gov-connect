.PHONY: all

all: clean dataprep run

.PHONY: clean

clean:
	-rm -r ./go-server/riot-index
	@echo "[✔️] Clean complete!"

.PHONY: dataprep

dataprep:
	@python3 ./go-server/data-utils/test_data_transform.py
	@echo "[✔️] Data preparation steps complete"
.PHONY: run

run:
	@cd ./client && sudo npm install && sudo npm run build &
	@echo "creating build..."
	-sudo rm -rf /var/www/html/* && sudo mv ./client/build/* /var/www/html/ && npm install -g serve && serve -s build &
	@echo "serving build..."
	@cd ./go-server && go run main.go fg &
	nohup geany > /dev/null
	disown
	@echo "[✔️] Build complete!"
