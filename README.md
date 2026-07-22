# tada-list-API

A basic todo list API written in GoLang using the GIN framework.
NOTE: Data is stored in memory and will be erased on terminating the program.

Schema:
```
ID STRING
TITLE STRING
PRIORITY STRING
STATUS STRING
DESCRIPTION STRING
```

Available Routes:
- `GET /tasks` - Gets all tasks present 
- `GET /tasks/:id` - Finds a task given the ID (Eg: /tasks/2 returns task with ID as 2)
- `POST /tasks` - Adss a task to the slice given the json structure
- `DELETE /tasks/:id` - Deletes a task given the task's ID
- `PATCH /tasks/:id/complete` - Sets a task as complete given the ID
- `PATCH /tasks/:id/update` - Modifies a task's details such as Title, Status, Priority, Description given the json
Eg:
```json
{
"title" : "Commit to Github"
}
```

