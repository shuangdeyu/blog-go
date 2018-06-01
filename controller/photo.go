package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"jyj-go/model"
)

// 网站作品
func NetPage(c *gin.Context){
	c.HTML(http.StatusOK, "photo/net.html", nil)
}

// 摄影.山
func ShanPage(c *gin.Context){
	list, err := model.GetShanList()
	//fmt.Println(*list[0])
	if err != nil {
		Handle404(c)
	} else {
		c.HTML(http.StatusOK, "photo/shan.html", gin.H{
			"list": list,
		})
	}
}

// 摄影.人
func RenPage(c *gin.Context) {
	list, err := model.GetImgConList(1)
	if err != nil {
		Handle404(c)
	} else {
		c.HTML(http.StatusOK, "photo/ren.html", gin.H{
			"content": list.Content,
			"imgList": list.ImgList,
		})
	}
}

// 摄影.仙
func XianPage(c *gin.Context) {
	list, err := model.GetImgConList(2)
	if err != nil {
		Handle404(c)
	} else {
		c.HTML(http.StatusOK, "photo/xian.html", gin.H{
			"content": list.Content,
			"imgList": list.ImgList,
		})
	}
}

// 摄影.赏
func ViewPage(c *gin.Context) {
	list, err := model.GetImgList()
	if err != nil {
		Handle404(c)
	} else {
		c.HTML(http.StatusOK, "photo/view.html", gin.H{
			"list": list,
		})
	}
}