package schemas

type BaseResp struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
	Data    any    `json:"data"`
}

type ResponseUUID struct {
	UUID string `json:"uuid"`
}
