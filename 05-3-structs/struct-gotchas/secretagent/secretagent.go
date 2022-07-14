package secretagent

type Agent struct {
	ID      int
	doubleO bool
}

func InitAgent(id int) *Agent {
	return &Agent{id, id == 7}
}
