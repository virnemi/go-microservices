
TARGET := client

API		   := ./api
GRPC		   := ./grpc
READER	   := ./reader
GOARGS	   := GO111MODULE=off GOPATH=$(GOPATH)

all:
	@echo "[BUILD] $(TARGET).exe"
	@$(GOARGS) go get -d -v .\...
	@$(GOARGS) go build -o $(TARGET).exe $^

racewin:
	@echo "[BUILD(DATA RACE DETECTION)] $(TARGET).exe"
	@$(GOARGS) go build -race -v -o $(TARGET).exe $^

linux:
	@echo "[LINUX BUILD] $(GOPATH)/$(TARGET)"
	@$(GOARGS) go get -d -v ./...
	@$(GOARGS) GOOS=linux go build -o $(GOPATH)/$(TARGET) $^

race:
	@echo "[BUILD(DATA RACE DETECTION)] $(TARGET)"
	@$(GOARGS) go build -race -v -o $(TARGET) $^

format:
	@echo "[FORMAT] *.go files"
	@find . -name "*.go" -exec go fmt {} \;

clean:
	@echo "[CLEAN] $(TARGET)"
	@rm -rf $(TARGET) $(TARGET).exe

test:
	@echo "[TEST] $(TARGET) files"
	@$(GOARGS) go get -d -v .\...
	@$(GOARGS) go test -v $(API)
	@$(GOARGS) go test -v $(GRPC)
	@$(GOARGS) go test -v $(READER)

test_race:
	@echo "[TEST RACE] $(TARGET) files"
	@$(GOARGS) go get -d -v .\...
	@$(GOARGS) go test -race -v $(API)
	@$(GOARGS) go test -race -v $(GRPC)
	@$(GOARGS) go test -race -v $(READER)

.PHONY: clean format linux
