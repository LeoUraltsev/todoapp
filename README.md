# todoapp

## Данное приложение создано для изучения backend рзработки на языке Go

### Запуск приложения

1. Создать файл конфигурации в корне проета "config.yml"
    Пример config файла:
        is_debug: true (false)
        listen
            port: 8080
            host: 0.0.0.0

2. Запустить Build приложения: 
    Windows: go build -o build/app.exe cmd/main/main.go 
    Linux / Mac: go build -o build/app cmd/main/main.go 

3. Запустить проложение Run:
    Windows: ./build/app.exe
    Linux: ./build/app

