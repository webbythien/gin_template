package schema

type Chat struct {
	ID       string `bson:"_id"`
	Message  int32  `json:"message"`
	RoomID   string `json:"room_id"`
	CreateAt string `json:"createAt"`
	UserId   string `bson:"user_id"`
}
