package handler

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"html/template"
	"net/http"

	"gitlab.com/bawi/task-queue/controller"
)

func Home(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	taskController := controller.NewTaskController(c)

	if task := r.FormValue("task"); task != "" {
		log.Infof(c, "Criando uma task")
		taskController.CreateTask()
	}

	if taskETA5min := r.FormValue("taskETA5min"); taskETA5min != "" {
		log.Infof(c, "Criando uma task que executa daqui 5 minutos")
		taskController.CreateTaskETA()
	}

	if _delete := r.FormValue("delete"); _delete != "" {
		log.Infof(c, "Deletando uma task")
		taskController.DeleteTask(_delete)
	}


	if err := handlerTemplate.Execute(w, taskController.GetAllTasks()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// OK
}

var handlerTemplate = template.Must(template.New("handler").Parse(handlerHTML))

const handlerHTML = `
{{range .}}
<p>{{.Name}}</p>
{{end}}

<p>Cria uma task:</p>
<form action="/" method="POST">
<input type="text" name="task">
<input type="submit" value="Add">
</form>

<p>Cria uma task que executa daqui 5 minutos:</p>
<form action="/" method="POST">
<input type="text" name="taskETA5min">
<input type="submit" value="Add">
</form>

<p>Deleta uma task:</p>
<form action="/" method="POST">
<input type="text" name="delete">
<input type="submit" value="Add">
</form>
`
