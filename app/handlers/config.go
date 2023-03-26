package handlers

import (
	"github.com/gin-gonic/gin"
	"market/app/model"
	"market/app/response"
	"market/app/utils"
	"market/app/validator/v_data"
	"market/app/vars"
	"strconv"
	"strings"
)

type Config struct{}

func (h *Config) Configs(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VConfigList)
	offset := utils.GetPages(params.Page, params.PageSize)
	configs, total, err := model.NewConfig(vars.DBMysql).List(params.Key, params.Desc, params.State, offset, params.PageSize)
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, gin.H{"total": total, "list": configs})
}

func (h *Config) Config(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	config, err := model.NewConfig(vars.DBMysql).FindOneById(id)
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, config)
}

func (h *Config) ConfigCreate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VConfigCreate)
	err := model.NewConfig(vars.DBMysql).CreateConfig(model.Config{
		Key:    params.Key,
		Val:    params.Val,
		Desc:   params.Desc,
		State:  1,
		Bak1:   params.Bak1,
		Bak2:   params.Bak2,
		Remark: params.Remark,
	})
	if err != nil {
		response.Fail(ctx, "创建失败: "+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (h *Config) ConfigUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VConfigUpdate)
	err := model.NewConfig(vars.DBMysql).UpdateConfig(params.Id, map[string]interface{}{
		"_k":     params.Key,
		"_v":     params.Val,
		"_desc":  params.Desc,
		"state":  params.State,
		"bak1":   params.Bak1,
		"bak2":   params.Bak2,
		"remark": params.Remark,
	})
	if err != nil {
		response.Fail(ctx, "修改失败: "+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (h *Config) ApiFindKey(ctx *gin.Context) {
	key := ctx.Query("key")
	config, err := model.NewConfig(vars.DBMysql).ApiFindOneByKey(key)
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, config.Val)
}

func (h *Config) ApiFindKeys(ctx *gin.Context) {
	keys := ctx.Query("keys")
	splits := strings.Split(keys, ",")
	configs := make(map[string]string)
	for _, split := range splits {
		config, _ := model.NewConfig(vars.DBMysql).ApiFindOneByKey(split)
		if config == nil {
			configs[split] = ""
		} else {
			configs[split] = config.Val
		}
	}
	response.Success(ctx, configs)
}
