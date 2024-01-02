package tasks

type taskData struct {
	TaskId        int    `json:"task_id"`
	Deadline      uint   `json:"deadline"`
	Stage         int    `json:"stage"`
	Title         string `json:"title"`
	Text          string `json:"text"`
	CompanyId     int    `json:"company_id"`
	CreatedId     int    `json:"created_id"`
	ResponsibleId int    `json:"responsible_id"`
	CreatedAt     uint   `json:"created_at"`
}

type stageResponse struct {
	Stages map[int]string `json:"data"`
	Err    *string        `json:"error"`
}

type taskResponse struct {
	Data *taskData `json:"data"`
	Err  *string   `json:"error"`
}

type taskListResponse struct {
	Data []taskData `json:"data"`
	Err  *string    `json:"error"`
}
