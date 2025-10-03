# Имя выходного файла
OUTPUT=./bin/math.dll

# Путь к исходникам (где лежит main.go с //export функциями)
SRC=./

# Кросс-компилятор
CC=x86_64-w64-mingw32-gcc

# Переменные окружения для Go
GOENV=GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=$(CC)

.PHONY: all dll clean

all: dll

dll:
	$(GOENV) go build -o $(OUTPUT) -buildmode=c-shared $(SRC)

clean:
	rm -f $(OUTPUT) ../bin/math.h