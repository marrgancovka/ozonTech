# Система для добавления и чтения постов с комментариями  
## Запуск приложения
Чтобы выбрать используемое хранилище нужно изменить переменную _STORAGE_ в файле _.env_ в корне проекта  
```makefile
# Для хранения данных в памяти (in-memory)
STORAGE=in-memory
# Для хранения данных в бд Postgres
STORAGE=postgres
```
Для запуска приложения нужно выполнить команду: 
```makefile
make start
```
Если выбрано использовать Postgres, то нужно выполнить миграции:
```makefile
make migrate-up
```
Остановить приложение можно командой:
```makefile
make stop
```

## Примеры использования приложения

### Регистрация и авторизация
Пример запроса:
```makefile
# Регистрация
mutation{
  register(name:"margarita", password: "1108")
}
# Авторизация
mutation{
  login(name:"margarita", password: "1108")
}
```
Пример ответа:
```makefile
{
  "data": {
    "register": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk4NzYxODEsImlkIjoxfQ.7zH5cuVyUWzr5Kha2IWq9DRnhWBt4_9dH-3TEcT62bY"
  }
}
```
Полученный токен нужно будет скопировать в headers для зпросов, которые требуют авторизации (создание постов и комментариев)
```makefile
{
  "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk4NTI2NTMsImlkIjoyfQ.NwhiBp-rgt61MKa-ZzHVgT_GhIWJIEsZJUEH49_QiY8"
}
```
### Создание поста
Пример запроса:
```makefile
mutation {
  createPost(title: "Мой первый пост", content: "Всем привет!", commentsAllowed: true) {
    id
    title
    content
    userID
  }
}
```
Пример ответа:
```makefile
{
  "data": {
    "createPost": {
      "id": "1",
      "title": " Мой первый пост",
      "content": "Всем привет!",
      "userID": "2"
    }
  }
}
```
### Создание комментария
Пример запроса:
```makefile
mutation {
  createComment(postID: 1, parentID: 0, content: "Спасибо за первый пост") {
    id
    content
  }
}
```
Пример ответа:
```makefile
{
  "data": {
    "createComment": {
      "id": "1",
      "content": "Спасибо за первый пост"
    }
  }
}
```

### Просмотр списка постов
Пример запроса:
```makefile
{
  posts{
    id
    title
    content
  }
}
```
Пример ответа:
```makefile
{
  "data": {
    "posts": [
      {
        "id": "1",
        "title": " Мой первый пост",
        "content": "Всем привет!"
      },
      {
        "id": "2",
        "title": " Мой второй пост",
        "content": "Хорошего настроения"
      },
      {
        "id": "3",
        "title": " Мой третий пост",
        "content": "Расскажите про GraphQL"
      }
    ]
  }
}
```

### Просмотр одного поста с комментариями
Пример запроса:
```makefile
{
  post(id:1){
    id
    title
    content
    comments {
      id
      content
      parentCommentID
      childComments{
        id
        content
        parentCommentID
        postID
        userID
      }
    }
  }
}
```
Пример ответа:
```makefile
{
  "data": {
    "post": {
      "id": "1",
      "title": " Мой первый пост",
      "content": "Всем привет!",
      "comments": [
        {
          "id": "1",
          "content": "Спасибо за первый пост",
          "parentCommentID": "",
          "childComments": [
            {
              "id": "3",
              "content": "Пожалуйста",
              "parentCommentID": "1",
              "postID": "1",
              "userID": "2"
            }
          ]
        },
        {
          "id": "2",
          "content": "Спасибо еще раз за первый пост",
          "parentCommentID": "",
          "childComments": null
        }
      ]
    }
  }
}
```