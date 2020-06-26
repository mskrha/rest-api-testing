BUILD = rest-api-testing
SRCS = main.go

all: clean format $(BUILD)

clean:
	rm -f $(BUILD)

format:
	go fmt

$(BUILD): clean format
	go build -o $(@) $(SRCS)
