package controller

import (
	"github.com/qedus/nds"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
	"time"
	"fmt"
	"golang.org/x/net/context"
	"github.com/douglasfsti/task-queue/struct"
	"github.com/satori/go.uuid"
)

type TaskInterface interface {
	CreateTask()
	CreateTaskETA()
	DeleteTask(name string)
	GetAllTasks() []_struct.TaskStruct
}

type TaskController struct {
	ctx context.Context
}

func NewTaskController(ctx context.Context) TaskInterface {
	return &TaskController{ctx: ctx}
}

func (c *TaskController) createTask() *taskqueue.Task {
	task := taskqueue.NewPOSTTask("/worker",
		map[string][]string{"createAt": {fmt.Sprintf("%v", time.Now())}})
	task.Name = uuid.NewV4().String()
	return task
}

func (c *TaskController) insertTaskDatastore(name string) (*datastore.Key, error) {
	return nds.Put(c.ctx, datastore.NewIncompleteKey(c.ctx, "Task", nil),
		&_struct.TaskStruct{Name:name})
}

func (c *TaskController) CreateTask() {
	if _, err := taskqueue.Add(c.ctx, c.createTask(), ""); err != nil {
		log.Errorf(c.ctx, "Erro ao criar task: %v", err)
	}
}

func (c *TaskController) CreateTaskETA() {
	task := c.createTask()
	task.ETA = time.Now().Add(time.Duration(5 * time.Minute))
	if _, err := taskqueue.Add(c.ctx, task, ""); err != nil {
		log.Errorf(c.ctx, "Erro ao criar task: %v", err)
	}

	if _, err := c.insertTaskDatastore(task.Name); err != nil {
		log.Errorf(c.ctx, "Erro ao adicionar datastore: %v", err)
	}
}

func (c *TaskController) GetAllTasks() []_struct.TaskStruct {
	//q := datastore.NewQuery("Task")
	//var tasks []_struct.TaskStruct
	//if _, err := q.GetAll(c.ctx, &tasks); err != nil {
	//	log.Errorf(c.ctx, "Erro ao buscar no datastore: %v", err)
	//	return []_struct.TaskStruct{}
	//}

	return []_struct.TaskStruct{}
}

func (c *TaskController) DeleteTask(name string) {
	if err := c.deleteTask(name); err != nil {
		log.Errorf(c.ctx, "Erro ao deletar task: %v", err)
	}

	if err := c.deleteTaskDatastore(name); err != nil {
		log.Errorf(c.ctx, "Erro ao remover key datastore: %v", err)
	}
}

func (c *TaskController) deleteTask(name string) error {
	task := c.createTask()
	task.Name = name
	return taskqueue.Delete(c.ctx, task, "")
}

func (c *TaskController) deleteTaskDatastore(name string) error {
	return nds.Delete(c.ctx, datastore.NewKey(c.ctx, "Task", name, 0, nil))
}
