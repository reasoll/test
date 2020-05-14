package routers

import (
	"github.com/astaxie/beego"
	"mvc/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/homeAction", &controllers.MainController{}, "*:Home")

	beego.Router("/loginAction", &controllers.MainController{}, "*:Login")

	beego.Router("/treeview", &controllers.SysController{}, "*:Tree")

	//添加主目录，既左侧目录
	beego.Router("/addtree", &controllers.SysController{}, "*:AddTree")

	//添加知识库一级目录
	beego.Router("/addnewtree", &controllers.SysController{}, "*:AddTrees")

	//添加知识库二级目录
	beego.Router("/addmorenewtree", &controllers.KnowController{}, "*:AddMoreTrees")

	beego.Router("/gettree", &controllers.SysController{}, "*:GetTree")

	beego.Router("/frontPage", &controllers.SysController{}, "*:GetPage")

	beego.Router("/admin/:id:string", &controllers.SysController{}, "*:Admin")

	//beego.Router("/adddirectory",&controllers.KnowController{},"*:AddDirectory")

	// beego.Router("/searchknowledge",&controllers.KnowController{},"*:SearchKnowledge")

	//显示二级目录
	//beego.Router("/getmulu",&controllers.KnowController{},"*:Getdirectory")

	beego.Router("/getmoremulu", &controllers.KnowController{}, "*:GetMoreDirectory")

	//显示知识库页面
	beego.Router("/getknowledge", &controllers.KnowController{}, "*:GetKnowledge")

	//跳转文章总页面
	beego.Router("/jump", &controllers.KnowController{}, "*:JumpPage")

	//跳转知识库管理页面
	beego.Router("/knowledgepage", &controllers.KnowController{}, "*:JumpKnowledgePages")

	//编辑知识库一级目录名
	beego.Router("/editknowledge", &controllers.KnowController{}, "*:EditKnowledge")

	//删除一级目录
	beego.Router("/deleteknowledge", &controllers.KnowController{}, "*:DeleteKnowledge")

	//跳转文章详细页面
	beego.Router("/articlePage", &controllers.KnowController{}, "*:JumpArticlePage")

	//关注知识库页面
	beego.Router("/getguanzhu", &controllers.GuanzhuController{}, "*:GetUserGuanzhu")

	//收藏知识库页面
	beego.Router("/getcollection", &controllers.CollectionController{}, "*:GetUserCollection")

	//添加关注
	beego.Router("/addGuanzhu", &controllers.GuanzhuController{}, "*:AddGuanzhuInformation")

	//取消关注
	beego.Router("/deleteGuanzhu", &controllers.GuanzhuController{}, "*:DeleteGuanzhu")

	//取消收藏
	beego.Router("/deleteCollection", &controllers.CollectionController{}, "*:DeleteCollection")

	//跳转添加文章页面
	beego.Router("/jumpAddArticle", &controllers.KnowController{}, "*:JumpAddArticle")

	//添加文章
	beego.Router("/addArticle", &controllers.KnowController{}, "*:AddArticle")

	//显示文章H5代码
	beego.Router("/getArticle", &controllers.KnowController{}, "*:GetArticle")

	//test页面
	beego.Router("/jumptest", &controllers.KnowController{}, "*:Jumptest")

	//添加收藏
	beego.Router("/addCollection", &controllers.CollectionController{}, "*:AddCollection")

	//初始化文章列表
	beego.Router("/getArticleList", &controllers.KnowController{}, "*:GetArticleList")

	//编辑主目录
	beego.Router("/editTree", &controllers.SysController{}, "*:EditTree")

	//删除主目录
	beego.Router("/deleteTree", &controllers.SysController{}, "*:DeleteTree")
}
