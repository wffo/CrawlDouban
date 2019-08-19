package model

type Urls struct {
	Url string
}
type Profile struct {
	//标题
	Title string
	//发布日期
	Data string
	//描述详情
	Text string
	//合租2 整租1
	Method string
	//小区1 城中村2 公寓3 未分类4
	Type string
	//求租
	Qiu string
	//唯一的ID
	Id string
	//图片url
	Image_url []Urls
	//帖子url
	Url string


	Xiaoqu string
	Daqu string
}
type Jy struct {
	Proxy int
}