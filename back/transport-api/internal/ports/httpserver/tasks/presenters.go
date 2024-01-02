package tasks

type addTaskRequest struct {
	Deadline      uint   `json:"deadline"`
	Title         string `json:"title"`
	Text          string `json:"text"`
	ResponsibleId int    `json:"responsible_id"`
}

type updateTaskRequest struct {
	Deadline      uint   `json:"deadline"`
	Stage         int    `json:"stage"`
	Text          string `json:"text"`
	ResponsibleId int    `json:"responsible_id"`
}
