package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"mvc/models"
	"strconv"
)

type KnowController struct {
	beego.Controller
}

//func (c *KnowController) AddDirectory(){
//	title := c.Input().Get("title")
//	id, _ := strconv.Atoi(c.Input().Get("code"))
//	err := models.InsertDirectory(title,id)
//	logs.Info("dataList :", err)
//	if err != nil {
//		c.ServeJSON()
//	}else{
//		c.ServeJSON()
//	}
//}

//func (c *KnowController) SearchKnowledge()  {
//	pid, _ := strconv.Atoi(c.Input().Get("pid"))
//	dataList2, err2 := models.SearchKnowledge(pid)
//	if err2 == nil {
//		c.Data["json"] =dataList2
//	}
//	logs.Info("dataList2 :", dataList2)
//	c.ServeJSON()
//}

//func (c *KnowController) Getdirectory() {
//	pid, _ := strconv.Atoi(c.Input().Get("pid"))
//	var datas []interface{}
//	switch pid {
//	case 1: datalist , _ := models.GetDirectorys()
//		datas = datalist
//	case 2:
//	case 3:
//	}
//	c.Data["json"] =datas
//	logs.Info("dataList2 :", datas)
//	c.ServeJSON()
//}

//后台页面点击一级知识库目录出现二级目录
func (c *KnowController) GetMoreDirectory() {
	title := c.GetString("title")
	datalist, err := models.GetMoreDirectorys(title)
	if err == nil {
		c.Data["json"] = datalist
	}
	logs.Info("dataList :", datalist)
	c.ServeJSON()
}

//拼写前端页面知识库画面
func (c *KnowController) GetKnowledge() {
	datalist, err := models.Knowledges()
	if err == nil {
		c.Data["json"] = datalist
	}
	logs.Info("dataList :", datalist)
	c.ServeJSON()
}

//跳转知识总页面
func (c *KnowController) JumpPage() {
	fid, _ := c.GetInt("id")
	//c.Data["gid"] = fid
	//c.TplName = "knowledge.tpl"
	datalist, err := models.JumpAllKnowPage(fid)
	if err == nil {
		c.Data["Fid"] = fid
		c.Data["List"] = datalist
		c.TplName = "knowledge.tpl"
	} else {
		c.Data["Fid"] = fid
		c.Data["List"] = datalist
		c.TplName = "knowledge.tpl"
	}
}

//跳转知识详细页面
func (c *KnowController) JumpArticlePage() {
	id, _ := c.GetInt("id")
	c.Data["articleId"] = id
	c.TplName = "knowledge_article.tpl"
}

//跳转知识库管理页面
func (c *KnowController) JumpKnowledgePages() {
	datalist, err := models.JumpToKnowledgePage()
	if err == nil {
		c.Data["List"] = datalist
	}
	c.TplName = "back_knowledgePage.tpl"
}

//添加知识库二级目录
func (c *KnowController) AddMoreTrees() {
	title := c.Input().Get("title")
	pid, _ := strconv.Atoi(c.Input().Get("code"))
	err := models.InsertAdminDepart3(title, pid)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

//编辑一级目录
func (c *KnowController) EditKnowledge() {
	title := c.Input().Get("title")
	pid, _ := strconv.Atoi(c.Input().Get("code"))
	err := models.EditKnowledge(title, pid)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

//删除一级目录
func (c *KnowController) DeleteKnowledge() {
	pid, _ := strconv.Atoi(c.Input().Get("code"))
	err := models.DeleteKnowledge(pid)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

//跳转添加知识文章页面
func (c *KnowController) JumpAddArticle() {
	gid, _ := c.GetInt("id")
	c.Data["gid"] = gid
	c.TplName = "addArticle.tpl"
}

//添加知识文章
func (c *KnowController) AddArticle() {
	texts := c.GetString("text")
	title := c.GetString("title")
	id, _ := c.GetInt("id")
	err := models.InsertArticle(texts, title, id)
	logs.Info("dataList :", err)
	if err != nil {
		c.ServeJSON()
	} else {
		c.ServeJSON()
	}
}

//显示知识文章
func (c *KnowController) GetArticle() {
	articleidId, _ := c.GetInt("articleid")
	datalist, err := models.GetArticles(articleidId)
	if err == nil {
		c.Data["json"] = datalist
	}
	//logs.Info("dataList :", datalist)
	c.ServeJSON()
}

func (c *KnowController) Jumptest() {
	c.TplName = "test.tpl"
}

//初始化文章列表
func (c *KnowController) GetArticleList() {
	datalist, err := models.GetArticleLists()
	if err == nil {
		c.Data["json"] = datalist
	}
	//logs.Info("dataList :", datalist)
	c.ServeJSON()
}
