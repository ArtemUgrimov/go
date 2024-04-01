package school

type Teacher struct {
	Name       string
	Password   string
	FullAccess bool
	token      string
}

func NewTeacher(name string, pass string, fullAccess bool) Teacher {
	t := Teacher{
		Name:       name,
		Password:   pass,
		FullAccess: fullAccess,
	}
	return t
}
