package task

type UseCaseTask interface {
	FindAll() ([]Task, error)
	FindById(Id int) (Task, error)
	Create(task TaskRequest) (Task, error)
	Update(Id int, task TaskRequest) (Task, error)
	Delete(Id int) (Task, error)
}

type useCaseTask struct {
	repository Repository
}

func TaskUseCase(repository Repository) *useCaseTask {
	return &useCaseTask{repository}
}

func (u *useCaseTask) FindAll() ([]Task, error) {
	task, err := u.repository.FindAll()
	return task, err
}

func (u *useCaseTask) FindById(id int) (Task, error) {
	task, err := u.repository.FindById(id)
	return task, err
}

func (u *useCaseTask) Delete(id int) (Task, error) {
	task, err := u.repository.Delete(id)
	return task, err
}

func (u *useCaseTask) Create(taskRequest TaskRequest) (Task, error) {
	task := Task{
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
		Doing:       taskRequest.Doing,
	}

	task, err := u.repository.Create(task)
	return task, err
}

func (u *useCaseTask) Update(Id int, taskRequest TaskRequest) (Task, error) {
	task, err := u.repository.FindById(Id)

	task.Title = taskRequest.Title
	task.Description = taskRequest.Description
	task.Doing = taskRequest.Doing

	newTask, err := u.repository.Update(task)
	return newTask, err
}
