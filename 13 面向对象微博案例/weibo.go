package main

import (
	"fmt"
	"time"
)

// 基类
type Base struct {
	id   int
	name string
}

// 博主子类
type Blogger struct {
	Base
	weibos  []*PostContent         // 存微博，使用切片结构体指针
	comment map[int][]*PostContent // 存微博的留言信息, int代表哪条微博, map指针
	fans    []FansInterfacer       // 存fans对象。
}

// 粉丝子类

type Fan struct {
	Base
}

type FriendFan struct {
	Fan
}

type BadFan struct {
	Fan
}

// 微博内容结构体

type PostContent struct {
	id          int
	content     string    //内容
	commentTime time.Time //时间
	wbType      int       // 心情类型
	postMan     string    // 发布人
	to          string    // 留言给谁。 给评论功能使用的。

}

// 接口

type FansInterfacer interface {
	Update(b BloggerInterfacer, wbId int)
	Active(b BloggerInterfacer, wbId int)
}

type BloggerInterfacer interface {
	Attach(fan FansInterfacer)
	Detach()
	Notify(wbId int)
}

// 功能函数

// 粉丝功能
func (f *FriendFan) Update(b BloggerInterfacer, wbId int) {
	blogger, ok := b.(*Blogger) // 断言及类型转换
	if ok {
		fmt.Printf("%s你好! 你关注的博主%s刚更新了一条微博.\n", f.name, blogger.name)
	}

	// 评论微博
	f.Active(b, wbId)
}

func (f *FriendFan) Active(b BloggerInterfacer, wbId int) {
	blogger, ok := b.(*Blogger)
	message := ""
	if ok {
		weibo := blogger.Getweibo(wbId)
		switch weibo.wbType {
		case 1:
			message = "愿你开心每一天"
		case 2:
			message = "加油，一起努力"
		default:
			message = ""
		}

		comment := PostContent{0, message, time.Now(), weibo.wbType, f.name, blogger.name}
		blogger.PostComment(comment, wbId)
		blogger.ShowComment(wbId)

	}
}

func (b *BadFan) Update() {
}

func (b *BadFan) Active() {
}

// 博主功能
func (b *Blogger) Attach(fan FansInterfacer) { // 添加粉丝
	b.fans = append(b.fans, fan)
}

func (b *Blogger) Detach() { // 取关粉丝
}

func (b *Blogger) Notify(wbId int) { // 发布微博通知粉丝
	for _, fan := range b.fans {
		fan.Update(b, wbId)
	}
}

func (b *Blogger) getId() int { // 获取微博ID
	var Id int = 0
	if len(b.weibos) == 0 {
		return Id
	} else {
		return Id + 1
	}
}

func (b *Blogger) PostWeibo(message string, wbType int) { // 发布微博
	content := new(PostContent)
	content.id = b.getId()
	content.content = message
	content.wbType = wbType
	content.commentTime = time.Now()
	content.postMan = b.name

	// 存储微博数据
	b.weibos = append(b.weibos, content)

	// 通知粉丝
	b.Notify(content.id)
}

func (b *Blogger) Getweibo(wbId int) *PostContent { // 获取微博ID
	for _, weibo := range b.weibos {
		if weibo.id == wbId {
			return weibo
		}
	}

	return nil

}

func (b *Blogger) PostComment(comment PostContent, wbId int) { // 发布评论
	b.comment[wbId] = append(b.comment[wbId], &comment)

}

func (b *Blogger) ShowComment(wbId int)  {
	weibo := b.Getweibo(wbId)
	fmt.Printf("博主: %s\n", b.name)
	fmt.Printf("微博内容: %s\n", weibo.content)
	for _, comment := range b.comment[wbId] {
		fmt.Printf("粉丝: %s\n", comment.postMan)
		fmt.Printf("评论内容: %s\n", comment.content)
	}

}



// 博主功能函数

// 1. 创建博主对象函数
func NewBlogger(name string) *Blogger {
	blogger := new(Blogger)
	blogger.id = 1
	blogger.name = name
	blogger.weibos = make([]*PostContent, 0)
	blogger.comment = make(map[int][]*PostContent)
	return blogger
}

func main() {
	blogger := NewBlogger("zhenping")
	friendfan := new(FriendFan)
	friendfan.id = 1
	friendfan.name = "weizi"
	blogger.Attach(friendfan)

	blogger.PostWeibo("今天天气真不错，心情很好", 1)

}
