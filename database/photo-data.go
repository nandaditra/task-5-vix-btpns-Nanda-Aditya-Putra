package database

type PhotoUpdateData struct {
	ID       uint64 `json:"id" form:"id" `
	Title    string `json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photourl" form:"photourl"`
	UserID   uint64 `json:"user_id,omitempty" form:"user_id, omiempty"`
}

type PhotoCreateData struct {
	Title    string `json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photourl" form:"photourl"`
	UserID   uint16 `json:"user_id,omitempty" form:"user_id"`
}
