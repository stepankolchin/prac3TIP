## Практическое занятие №3. Колчин Степан Сергеевич, ЭФМО-02-25. Реализация простого HTTP-сервера на стандартной библиотеке net/http. Обработка запросов GET/POST


### Структура проекта

```
Prac3/
│   README.md
│   go.mod
│
├───cmd
│     main.go
│
└───internal
        ├───api
        │       handlers.go         # Обработчики HTTP запросов
        │       handlers_test.go    # Юнит-тесты
        │       middleware.go       # Middleware (CORS, логирование)
        │       responses.go        # Вспомогательные функции ответов
        │
        └───storage
                memory.go            # In-memory хранилище задач
```

1. Health check

<img width="820" height="76" alt="изображение" src="https://github.com/user-attachments/assets/d529d75a-fcae-4ebb-80ee-638af8297703" />

2. Создание и просмотр задач

<img width="820" height="293" alt="изображение" src="https://github.com/user-attachments/assets/74b3e483-33e0-4140-9568-aa77d2a304a9" />

<img width="820" height="59" alt="изображение" src="https://github.com/user-attachments/assets/492515ef-a1da-4f3e-b6f7-3cfb8f044658" />



4. Получение задачи по `ID`

<img width="820" height="97" alt="изображение" src="https://github.com/user-attachments/assets/8a4b87cc-54c4-48f9-9eb2-7ee98a56534d" />


### Функциональность

- `GET /health` - проверка работоспособности сервера

- `GET /tasks` - список всех задач (с фильтрацией ?q=)

- `POST /tasks` - создание новой задачи

- `GET /tasks/{id}` - получение задачи по ID

- `PATCH /tasks/{id}` - обновление задачи

- `DELETE /tasks/{id}` - удаление задачи

### Дополнительно

1. `CORS (минимально)`: добавить заголовки Access-Control-Allow-Origin: * для GET/POST (в отдельной middleware).

<img width="820" height="208" alt="изображение" src="https://github.com/user-attachments/assets/1e466a1f-6154-4a9c-b02e-7cf40153f300" />

2. Валидация длины `title` (например, 1…140 символов).

<img width="820" height="335" alt="изображение" src="https://github.com/user-attachments/assets/adabc995-6d01-4853-9595-d474e013b805" />

3. Метод `PATCH /tasks/{id}`для отметки Done=true.

<img width="806" height="218" alt="изображение" src="https://github.com/user-attachments/assets/07b5b28f-b6bd-4082-bb6f-7870323c593a" />

4. Метод `DELETE /tasks/{id}`

    <img width="806" height="258" alt="изображение" src="https://github.com/user-attachments/assets/47edba91-696b-4e6e-a7c8-2c886d10faec" />


5. `Graceful shutdown` через http.Server и контекст.

<img width="806" height="377" alt="изображение" src="https://github.com/user-attachments/assets/f293e3a1-d7b2-4282-9654-afa1969226fa" />
