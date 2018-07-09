# todolist

## Usecases

1. Как пользователь, я хочу иметь возможность просмотреть все мои задачи.
2. Как пользователь, я хочу иметь возможность добавить задачу (Название), чтобы пополнять список задач.
3. Как пользователь, я хочу иметь возможность удалить задачу, чтобы не хранить неактуальную информацию.
4. Как пользователь, я хочу иметь возможность изменить задачу.
5. Как пользователь, я хочу иметь возможность отметить выполненность или снимать выполненность задачи 

## REST API

### №1 GET api/todolist/task 

Ответ: 200 ОК
```json
    [{
        "id" : 1,
        "name" : "Задача 1",
        "success" : true|false 
    }]
```


### №2 POST api/todolist/task

Тело запроса:
```json
    {
        "name" : "Задача 1",
        "success" : false 
    }
 ```
Ответ: 201 Created
Location: api/todolist/task/



### №3 DELETE api/todolist/task/{id}

Ответ: 204  No Content


### №4 PUT api/todolist/task/{id}

Тело запроса:
```json
    {
        "id": 1,
        "name" : "Задача 1",
        "success" : false 
    }
 ```

Ответ: 202 Accepted Location: /api/todolist/task/1


### №5 PUT api/todolist/task/complete/{id}

Тело запроса:
```json
    {
        "id": 1,
        "name" : "Задача 1",
        "success" : true 
    }
 ```

Ответ: 202 Accepted Location: /api/todolist/task/complete/1


======

## Как собрать и запускать

Backend:

```bat
cd backend 
docker build -f Dockerfile -t todolistbackend:<имя ветки> .
docker run --rm --name todolistbackend -e NAME=<параметр приложения> todolistbackend:<имя ветки>
```

Frontend:

```bat
cd frontend 
docker build -f Dockerfile -t todolistfrontend:<имя ветки> .
docker run -d --rm --name todolistfrontend -p 80:80 todolistfrontend:<имя ветки>
```