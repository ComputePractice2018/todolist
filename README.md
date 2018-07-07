# todolist

## Usecases

1. Как пользователь, я хочу иметь возможность просмотреть все мои задачи.
2. Как пользователь, я хочу иметь возможность добавить задачу (Название), чтобы пополнять список задач.
3. Как пользователь, я хочу иметь возможность добавлять подзачи (1),изменять подзадачу (2) ,удалять подзадачи(3).
4. Как пользователь, я хочу иметь возможность  отмечать  выполненые подзадачи.
5. Как пользователь, я хочу иметь возможность удалить задачу, чтобы не хранить неактуальную информацию.

## REST API

### №1 GET api/task/getList 

Ответ: 200 ОК
```json
    [{
        "id" : 1,
        "name" : "Задача 1",
        "success" : true|false 
    }]
```

### №2 POST api/task/add

Тело запроса:
```json
    {
        "name" : "Задача 1",
        "success" : false 
    }
 ```   
Ответ: 201 Created
Location: api/task/add/1

### №3.1 POST api/field/add/{idRef}

Тело запроса:
```json
    {
        "nameField" : "Подзадача 1",
        "successField" : false,
         "idRef": {idRef}
    }
 ```   
Ответ: 201 Created
Location: api/field/add/1



### №3.2 PUT api/edit/{idRef}/{idField}

Тело запроса:
```json
    {   
        "idField" : {idField},
        "nameField" : "Подзадача 1",
        "successField" : false|true,
         "idRef": {idRef}
    }
 ```  
Ответ: 202 Accepted

### №3.3 DELETE api/field/del/{idField}

Ответ: 204  No Content


### №4 PUT api/success/{idRef}/{idField}

Тело запроса:
```json
    {   
        "idField" : {idField},
        "successField" : false|true,
         "idRef": {idRef}
    }
 ```  
Ответ: 202 Accepted

### №5 DELETE api/task/del/{id}

Ответ: 204  No Content

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