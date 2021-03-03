
TARGET := server

REPO		   := ./repository
SERVICE	   := ./service
GOARGS	   := GO111MODULE=off GOPATH=$(GOPATH)

all:
	@echo "[BUILD] $(TARGET).exe"
	@$(GOARGS) go get -d -v .\...
	@$(GOARGS) go build -o $(TARGET).exe $^

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
	@$(source .env)
	@$(GOARGS) go get -d -v ./...
	@$(GOARGS) go test -v $(REPO)
	@$(GOARGS) go test -v $(SERVICE)

test_race:
	@echo "[TEST] $(TARGET) files"
	@$(source .env)
	@$(GOARGS) go get -d -v ./...
	@$(GOARGS) go test -race -v $(REPO)
	@$(GOARGS) go test -race -v $(SERVICE)


.PHONY: clean format linux
