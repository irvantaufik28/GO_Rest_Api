package task

type UseCaseTask interface {
	FindAll() ([]Task, error)
	FindById(Id int) (Task, error)
	Create(task TaskRequest) (Task, error)
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
func (u *useCaseTask) Create(taskRequest TaskRequest) (Task, error) {
	task := Task{
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
		Doing:       taskRequest.Doing,
	}

	task, err := u.repository.Create(task)
	return task, err
}
