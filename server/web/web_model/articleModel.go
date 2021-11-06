package web_model

import "2021/yunsongcailu/yunsong_server/tools"

type ArticleModel struct {
	Id         int64
	State      int            `xorm:"int index" json:"state"`       // 是否激活 1 否 0 是
	MenuId     int64          `xorm:"int index" json:"menu_id"`     // 一级类别
	CategoryId int64          `xorm:"int index" json:"category_id"` // 二级类别
	IsIndex    int            `xorm:"int index" json:"is_index"`    // 首页展示
	IsTop      int            `xorm:"int index" json:"is_top"`      // 本类推荐
	IsBottom   int            `xorm:"int index" json:"is_bottom"`   // 本类底部展示
	IsOriginal int            `xorm:"int index" json:"is_original"` // 是否原创  0 原创 1 引用
	LinkUrl    string         `xorm:"varchar(255)" json:"link_url"` // 引用地址
	Author     string         `xorm:"varchar(50)" json:"author"`    // 作者
	AuthorId   int64          `xorm:"int index" json:"author_id"`   // 作者ID
	Title      string         `xorm:"varchar(150)" json:"title"`    // 标题
	Desc       string         `xorm:"varchar(255)" json:"desc"`     // 描述
	Banner     string         `xorm:"varchar(255)" json:"banner"`
	Cover      string         `xorm:"varchar(255)" json:"cover"`   // 封面
	Content    string         `xorm:"text" json:"content"`         // 内容
	Video      string         `xorm:"varchar(255)" json:"video"`   // 视频
	Star       int            `xorm:"int" json:"star"`             // 点赞
	Tread      int            `xorm:"int" json:"tread"`            // 踩
	ReadNum int	`xorm:"int" json:"read_num"`	// 阅读量
	CommentNum int `xorm:"int" json:"comment_num"` // 评论数
	Keyword    string         `xorm:"varchar(255)" json:"keyword"` // 关键字
	CreateTime tools.JsonTime `xorm:"created" json:"create_time"`  // 时间
	UpdateTime tools.JsonTime `xorm:"updated" json:"update_time"`
}
