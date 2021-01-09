.PHONY: all

all: clean run

.PHONY: clean

clean:
	@echo "[✔️] Clean complete!"

.PHONY: run

run:
	@cd ./client && npm install && npm start &
	@cd ./go-server && go run main.go fg &
	@zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties &
	@echo "started zookeeper server"
	@kafka-server-start /usr/local/etc/kafka/server.properties
	@echo "started kafka server"
	@kafka-console-producer --broker-list localhost:9092 --topic test
	@echo "started kafka producer"
	@echo "[✔️] Now Running!"
