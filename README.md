# Math App

## Требования

Наличие на машине утилит `docker` и `docker-compose`

## Настройка
В `docker-compose.yml` в значении `MATH_APP_API_PORT` можно изменить порт, на котором будет запущен сервер.

## Запуск через `docker-compose`
Выполнить в консоли команды:  
`docker-compose build`  
`docker-compose up`  

## Запуск без `docker-compose`
Выполнить следующие команды. Во второй команде можно изменить порт.  
`docker build -t math_app .`  
`docker run --env MATH_APP_API_PORT=8080 -p 8080:8080 -it math_app
`

