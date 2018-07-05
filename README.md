# todolist

## Как собрать и звпустить

Backend:

```bat
cd backend 
docker build -f Dockerfile -t todolist:<имя ветки> .
docker run --rm --name todolist -e NAME=<параметр приложения> todolist:<имя ветки>
```