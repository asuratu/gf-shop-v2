package backend

type CommonPaginationReq struct {
	Page int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size int `json:"size" in:"query" d:"10" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
}

type CommonPaginationRes struct {
	List  interface{} `json:"list" dc:"列表数据"`
	Total int         `json:"total" dc:"总数"`
	Page  int         `json:"page" dc:"分页号码"`
	Size  int         `json:"size" dc:"分页数量"`
}
