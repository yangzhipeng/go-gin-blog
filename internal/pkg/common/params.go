package common

type CommonParams struct {
	Page  int   `form:"page,default=1"`   // 当前页码
	Limit int   `form:"limit,default=10"` // 每页大小
	Id    int   `form:"id"`               // id
	Ids   []int `form:"ids[]"`            // ids
}
