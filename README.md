# todoapp

## Данное приложение создано для изучения backend рзработки на языке Go

### Запуск приложения

1. Создать файл конфигурации в корне проета "config.yml"   
    Пример config файла:   
    ```
        is_debug: true
        listen
            port: 8080
            host: 0.0.0.0
    ```
2. Запустить Build приложения:   
    Windows: ``` go build -o build/app.exe cmd/main/main.go ```   
    Linux / Mac: ``` go build -o build/app cmd/main/main.go ```

3. Запустить проложение Run:   
    Windows: ``` ./build/app.exe ``` 
    Linux: ``` ./build/app ```


### Маршруты в приложении:
    GET /tasks - получения всех задач
    
    GET /tasks/:id - получение задачи по идентификатору
    
    POST /tasks - добавление новой задачи
    BODY: {
    "title":"title"
    "description":"description"
    }

    PUT /tasks/:id" - полное обновление задачи
    BODY: {
    "title":"title"
    "description":"description"
    }

    PATCH /tasks/:id - частичное обновление задачи
    BODY: {
    "title":"title"
    }
    
    DELETE /tasks/:id - удаление задачи

    

### Используемые библиотеки:
**Логирование:** [ZAP](https://github.com/uber-go/zap)   
**Конфигурация:** [Clean Env](https://github.com/ilyakaznacheev/cleanenv#clean-env)   
**База данных** [PostgreSQL Driver](https://github.com/jackc/pgx)
**Routing** [HttpRouter](https://github.com/julienschmidt/httprouter)