# Практическая работа 11
## REST Api
### Выполнил Сотников М.Е. ЭФМО-01-25

Учебный RESTful API проект на Go для управления заметками. Проект построен с использованием принципов чистой архитектуры (слои Core, Service, Repository)

Современные Go-проекты делятся на слои:   

* *cmd/api/main.go*        — точка входа
* *internal/http/handlers*  — обработчики HTTP-запросов
* *internal/core/service*   — бизнес-логика
* *internal/repo*           — работа с данными (память или БД)
* *internal/platform*       — вспомогательные пакеты (логирование, валидация)
* *api/openapi.yaml*       — документация


Post (Create) (201 Created):
<img width="839" height="669" alt="POSTNote" src="https://github.com/user-attachments/assets/03ae748f-e6ba-4900-8982-494db14c3d47" />   

<img width="836" height="670" alt="POSTNote2" src="https://github.com/user-attachments/assets/4a0845ed-e5f1-4bf3-a5cb-15b4892c1dd0" />

Delete (204 No Content):
<img width="847" height="580" alt="DELETENote" src="https://github.com/user-attachments/assets/05dbd47d-6d6a-4d9f-83a8-8faaeb9cb6d7" />

Get (Read) несуществующей заметки (404):
<img width="831" height="626" alt="GETunexisting" src="https://github.com/user-attachments/assets/6c15f496-75e6-46e3-b624-86cf19ce329a" />    

Get by ID (200)
<img width="840" height="689" alt="GETexistingByID" src="https://github.com/user-attachments/assets/2db487a1-0c30-4931-a536-be1e63c5a485" />    

Get all (200)
<img width="849" height="851" alt="GETall" src="https://github.com/user-attachments/assets/9d1abc2d-080f-4f30-9945-2969992d03f4" />    

Коды ответов:
*	200 — успешное выполнение;
*	201 — создан новый ресурс;
*	204 — операция выполнена, контент не возвращается;
*	400 — некорректный запрос;
*	404 — ресурс не найден;
*	500 — внутренняя ошибка сервера.

## Выводы
* REST *(Representational State Transfer)* — это архитектурный стиль взаимодействия клиентских и серверных приложений через протокол HTTP.
* CRUD - Create (*POST*) Read (*GET*) Update (*PUT*, *PATCH*) Delete (*DELETE*)
* Handler (Ручка) - работает с протоколом, принимает HTTP-запрос, достает данные из JSON, проверяет заголовки и отправляет ответ
* Service - бизнес-логика приложения
* Repository - слой данных, общается с базой данных






