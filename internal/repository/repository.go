package repository

type TasksAdmin interface {
	Create()    //Создать такс
	Delete()    //Удалить таск
	Open()      //Открыть таски для решения
	Close()     //Закрыть таски для решения
	OpenView()  //Сделать таски доступными для просмотра
	CloseView() //закрыть от просмота
	Disable()   //Отключить и открепить уже решенные таски
	Enable()    //вернуть назад к решению
	Backup()    //Сделать копию тасков
	Update()    //Обновить
	Import()    //Импортировать
}

type TasksUser interface {
	ShowTasks() //Просмотреть таски
	//TrySolve()  //Попытаться решить
}

type TasksSubmission interface {
	TrySolve(TaskID int, TeamID int) (a bool /*, time.Time*/)
}

type Authors interface {
	Create()
	Delete()
	Change()
}

type repository interface {
	TasksAdmin
	TasksUser
	TasksSubmission
	Authors
}
