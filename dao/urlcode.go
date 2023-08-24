package dao

import (
	"time"

	"github.com/baaj2109/shorturl/common"
	"github.com/baaj2109/shorturl/model"
	"gorm.io/gorm"
)

type UrlCodeDAO struct {
}

func NewUrlCodeDAO() *UrlCodeDAO {
	return &UrlCodeDAO{}
}

func (u UrlCodeDAO) AddUrl(url string, id int) (int, error) {
	uc := model.UrlCode{
		Url:       url,
		Code:      "",
		MD5:       common.MD5(url),
		UserId:    id,
		CreatedAt: time.Now(),
	}
	if err := common.DB.Create(&uc).Error; err != nil {
		common.SugarLogger.Info("Failed to add to db, err:%s", err)
		return 0, err
	}
	common.SugarLogger.Info("Add to db success")
	return uc.Id, nil
}

func (u UrlCodeDAO) GetByUrl(url string) model.UrlCode {
	var uc model.UrlCode
	common.DB.Where("md5 = ?", common.MD5(url)).Find(&uc)
	return uc
}

func (u UrlCodeDAO) GetByCode(code string) model.UrlCode {
	var uc model.UrlCode
	common.DB.Where("code = ?", code).Find(&uc)
	return uc
}

func (u UrlCodeDAO) UpdateCode(id int, code string) (err error) {
	err = common.DB.Table("url_code").Where("id = ?", id).Update("code", code).Error
	if err != nil {
		common.SugarLogger.Info("Failed to update code, err:%s", err)
		return err
	}
	return nil
}

func (u UrlCodeDAO) SaveClicks(clicks map[string]int) {
	for code, count := range clicks {
		var uc model.UrlCode
		common.DB.Where("code = ?", code).Find(&uc).UpdateColumn("click", gorm.Expr("click + ?", count))
		common.SugarLogger.Infof("add %d click on %s", count, code)
	}
}
