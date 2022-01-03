package repository

// Tasks
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
	ShowTasks() //Просмотреть таски             !    НУЖНЫ ЛИ ЭТИ МЕТОДЫ В ТАСКАХ 	  !
	ShowHint()  //                              ! ИЛИ ИХ ЛУЧШЕ В КОМАНДУ?? и наоборот !
	//TrySolve()  //Попытаться решить
}

type TasksSubmission interface {
	TrySolve(TaskID int, TeamID int) (a bool /*, time.Time*/)
}

type TasksAuthors interface {
	Create()
	Delete()
	Change()
}

//users
type Roles interface {
	Create()
}

type Persons interface {
	Create()
	Find()
	ChangePassword()
}

type Teams interface {
	Create()
	Join()
	ViewTeam()
	ViewScoreboard()
	ChangeAdmin()
	Kick()
}

type Repository interface {
	TasksAdmin
	TasksUser
	TasksSubmission
	TasksAuthors
	Roles
	Persons
	Teams
}
