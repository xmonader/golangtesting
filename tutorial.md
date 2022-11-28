- `go mod init github.com/xmonader/goclienttesting`
- `go get github.com/gin-gonic/gin`
- `touch main.go`
- scaffold the application with a memory store for todos

```go
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type TodoStore struct {
	todos map[int]Todo
}

type App struct {
	store TodoStore
}

func (a *App) listHandler(c *gin.Context) {
	fmt.Printf("list all todos request\n")
	c.JSON(http.StatusOK, a.store.todos)
}
func (a *App) getHandler(c *gin.Context) {
	fmt.Printf("get one todo request\n")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}

	if todo, ok := a.store.todos[id]; ok {
		c.JSON(http.StatusOK, todo)
	}
}
func (a *App) postHandler(c *gin.Context) {
	fmt.Printf("creating todo request\n")
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		newId := len(a.store.todos) + 1
		todo.ID = newId
		a.store.todos[newId] = todo
		c.String(http.StatusOK, "Success")
	}
}

func (a *App) putHandler(c *gin.Context) {
	fmt.Printf("updating todo request\n")
	var todo Todo
	todoIdStr := c.Param("id")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})

	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		fmt.Println(todo, a.store.todos)
		if _, ok := a.store.todos[todoId]; ok {
			todo.ID = todoId
			a.store.todos[todoId] = todo
			c.String(http.StatusOK, "Success")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})

		}
	}

}
func NewApp() *App {
	return &App{store: TodoStore{todos: map[int]Todo{}}}
}

func main() {
	a := NewApp()
	r := gin.Default()
	r.GET("/todos/", a.listHandler)
	r.GET("/todos/:id", a.getHandler)
	r.POST("/todos/", a.postHandler)
	r.PUT("/todos/:id", a.putHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


```

### requests

#### Get All
curl http://localhost:8080/todos

#### Get one
curl http://localhost:8080/todos/1

#### Create one
curl -X POST -H "Content-Type: application/json" -d '{"title": "todo" }' http://localhost:8080/todos
#### Update one
curl -X PUT -H "Content-Type: application/json" -d '{"title": "todo"}'' http://localhost:8080/todos