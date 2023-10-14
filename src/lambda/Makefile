TARGET = target
TMP = tmp

zip: main
	zip $(TMP)/$(TARGET).zip $(TMP)/$(TARGET)

main:
	GOARCH=amd64 GOOS=linux go build -o $(TMP)/$(TARGET) $(TARGET).go
