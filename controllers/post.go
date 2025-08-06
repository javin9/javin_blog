package controllers

import (
	"javin_blog/system"
	"net/http"
	"strconv"
	"strings"

	"javin_blog/models"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func PostGet(c *gin.Context) {
	id, err := ParamUint(c, "id")
	if err != nil {
		HandleMessage(c, err.Error())
		return
	}
	post, err := models.GetPostById(id)
	if err != nil || !post.IsPublished {
		Handle404(c)
		return
	}
	post.View++
	post.UpdateView()
	post.Tags, _ = models.ListTagByPostId(id)
	post.Comments, _ = models.ListCommentByPostID(id)
	user, _ := c.Get(ContextUserKey)
	c.HTML(http.StatusOK, "post/display.html", gin.H{
		"post": post,
		"user": user,
		"cfg":  system.GetConfiguration(),
	})
}

func PostNew(c *gin.Context) {
	c.HTML(http.StatusOK, "post/new.html", gin.H{
		"user": c.MustGet(ContextUserKey),
		"cfg":  system.GetConfiguration(),
	})
}

func PostCreate(c *gin.Context) {
	tags := c.PostForm("tags")
	title := c.PostForm("title")
	body := c.PostForm("body")
	isPublished := c.PostForm("isPublished")
	published := "on" == isPublished

	post := &models.Post{
		Title:       title,
		Body:        body,
		IsPublished: published,
	}
	err := post.Insert()
	if err != nil {
		c.HTML(http.StatusOK, "post/new.html", gin.H{
			"post":    post,
			"message": err.Error(),
			"user":    c.MustGet(ContextUserKey),
			"cfg":     system.GetConfiguration(),
		})
		return
	}

	// add tag for post
	if len(tags) > 0 {
		tagArr := strings.Split(tags, ",")
		for _, tag := range tagArr {
			tagId, err := parseUint(tag)
			if err != nil {
				continue
			}
			pt := &models.PostTag{
				PostId: post.ID,
				TagId:  tagId,
			}
			pt.Insert()
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/post")
}

func PostEdit(c *gin.Context) {
	id, err := ParamUint(c, "id")
	if err != nil {
		HandleMessage(c, err.Error())
		return
	}
	post, err := models.GetPostById(id)
	if err != nil {
		Handle404(c)
		return
	}
	post.Tags, _ = models.ListTagByPostId(id)
	c.HTML(http.StatusOK, "post/modify.html", gin.H{
		"post": post,
		"user": c.MustGet(ContextUserKey),
		"cfg":  system.GetConfiguration(),
	})
}

func PostUpdate(c *gin.Context) {
	tags := c.PostForm("tags")
	title := c.PostForm("title")
	body := c.PostForm("body")
	isPublished := c.PostForm("isPublished")
	published := "on" == isPublished

	id, err := ParamUint(c, "id")
	if err != nil {
		HandleMessage(c, err.Error())
		return
	}

	post := &models.Post{
		Title:       title,
		Body:        body,
		IsPublished: published,
	}
	post.ID = id
	err = post.Update()
	if err != nil {
		c.HTML(http.StatusOK, "post/modify.html", gin.H{
			"post":    post,
			"message": err.Error(),
			"user":    c.MustGet(ContextUserKey),
			"cfg":     system.GetConfiguration(),
		})
		return
	}
	// 删除tag
	models.DeletePostTagByPostId(post.ID)
	// 添加tag
	if len(tags) > 0 {
		tagArr := strings.Split(tags, ",")
		for _, tag := range tagArr {
			tagId, err := parseUint(tag)
			if err != nil {
				continue
			}
			pt := &models.PostTag{
				PostId: post.ID,
				TagId:  tagId,
			}
			pt.Insert()
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/post")
}

func PostPublish(c *gin.Context) {
	var (
		err  error
		res  = gin.H{}
		post *models.Post
	)
	defer writeJSON(c, res)
	id, err := ParamUint(c, "id")
	if err != nil {
		res["message"] = err.Error()
		return
	}
	post, err = models.GetPostById(id)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	post.IsPublished = !post.IsPublished
	err = post.Update()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func PostDelete(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	id, err := ParamUint(c, "id")
	if err != nil {
		res["message"] = err.Error()
		return
	}
	post := &models.Post{}
	post.ID = id
	err = post.Delete()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	models.DeletePostTagByPostId(id)
	res["succeed"] = true
}

func PostIndex(c *gin.Context) {
	posts, _ := models.ListAllPost("")
	c.HTML(http.StatusOK, "admin/post.html", gin.H{
		"posts":    posts,
		"Active":   "posts",
		"user":     c.MustGet(ContextUserKey),
		"comments": models.MustListUnreadComment(),
		"cfg":      system.GetConfiguration(),
	})
}

// 搜索
func PostSearch(c *gin.Context) {
	var limitInt, pageInt, commentTotalInt int
	keyword := c.Query("keyword")
	isPublished := c.Query("isPublished")
	commentTotal := c.Query("commentTotal")

	if commentTotal == "" {
		commentTotalInt = 0
	} else {
		_, err := strconv.Atoi(commentTotal)
		if err != nil {
			commentTotalInt = 0
		}
	}

	limit := c.Query("limit")
	if limit == "" {
		limitInt = 10
	} else {
		_, err := strconv.Atoi(limit)
		if err != nil {
			limitInt = 10
		}
	}

	page := c.Query("page")
	if page == "" {
		pageInt = 1
	} else {
		_, err := strconv.Atoi(page)
		if err != nil {
			pageInt = 1
		}
	}

	postList, err := models.SearchPostListWithOptions(models.SearchOptions{
		Keyword:      keyword,
		IsPublish:    isPublished == "1",
		Limit:        limitInt,
		Offset:       pageInt,
		CommentTotal: commentTotalInt,
	})

	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    1,
			Data:    nil,
			Message: err.Error(),
		})
		return
	}
	// 返回 总数和当前页数 怎么改下面的返回
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Data:    postList,
		Message: "success",
	})
}

type SearchParams struct {
	Keyword      string `json:"keyword"`
	IsPublished  int    `json:"isPublished"`
	CommentTotal int    `json:"commentTotal"`
	Limit        int    `json:"limit"`
	Page         int    `json:"page"`
}

func PostSearch2(c *gin.Context) {
	var params SearchParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    1,
			Data:    nil,
			Message: err.Error(),
		})
		return
	}

	result, err := models.SearchPostListWithOptions(models.SearchOptions{
		Keyword:      params.Keyword,
		IsPublish:    params.IsPublished == 1,
		Limit:        params.Limit,
		Offset:       params.Page,
		CommentTotal: params.CommentTotal,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    1,
			Data:    nil,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    0,
		Data:    result,
		Message: "success",
	})
}
