package models

type Response struct {
	Status 	string 	`json:"status"`
	Message string 	`json:"message"`
}


type Page struct {
	Success		bool	`json:"success"`
	Current		int		`json:"current"`
	PageSize	int 	`json:"pageSize"`
	Total 		int 	`json:"total"`
	Data 		[]interface{}	`json:"data"`
}


type Tag struct {
	Key		string 	`json:"key"`
	Label 	string 	`json:"label"`
}

type AccountInfo struct {
	Profile		string 	`json:"profile"`
	Mail 		string 	`json:"email"`
	Name 		string 	`json:"name"`
	Avatar		string 	`json:"avatar"`
	ID 			string	`json:"userid"`
	Signature	string	`json:"signature"`
	Title 		string 	`json:"title"`
	Group 		string 	`json:"group"`
	Tag			[]Tag	`json:"tags"`
	NotifyCount	int		`json:"notifyCount"`
	UnreadCount	int		`json:"unreadCount"`
	Country		string 	`json:"country"`
	Geographic	struct {
		Province	Tag	`json:"province"`
		City		Tag `json:"city"`
	}					`json:"geographic"`
	Address 	string	`json:"address"`
	Phone 		string	`json:"phone"`
}

