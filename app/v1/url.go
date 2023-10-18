package v1

import (
	"errors"

	"github.com/baaj2109/shorturl/common"
	"github.com/baaj2109/shorturl/dao"
	"k8s.io/utils/lru"
)

type Server interface {
	GetShortUrl(string, int) (string, error)
	ResotreUrl(string) (string, error)
}

type service struct {
	cCache   *lru.Cache
	uCache   *lru.Cache
	urlModel *dao.UrlCodeDAO
}

func NewService() *service {
	return &service{
		cCache:   lru.New(100),
		uCache:   lru.New(100),
		urlModel: dao.NewUrlCodeDAO(),
	}
}

// encode url
func (s *service) GetShortUrl(url string, userId int) (string, error) {
	shortCode := ""
	encodedUrl := common.MD5(url)
	if uc, ok := s.uCache.Get(encodedUrl); ok {
		shortCode = uc.(string)
	} else {
		uc := s.urlModel.GetByUrl(url)
		if uc.Code != "" {
			shortCode = uc.Code
		} else {
			id := 0
			if uc.Id != 0 {
				id = uc.Id
			} else {
				_id, err := s.urlModel.AddUrl(url, userId)
				id = _id
				if err != nil {
					return "", err
				}
			}
			if id == 0 {
				return "", errors.New("failed to get id")
			}
			shortCode = TransToCode(id)
			if shortCode == "" {
				return "", errors.New("gen code failed")
			}
			common.SugarLogger.Info("add new short url, code:%s", shortCode)
			if err := s.urlModel.UpdateCode(id, shortCode); err != nil {
				return "", err
			}
		}
	}
	// add to cache
	go func() {
		s.cCache.Add(shortCode, url)
		s.uCache.Add(encodedUrl, shortCode)
	}()
	return common.Config.GetString("app.url") + "/" + shortCode, nil
}

// restore url
func (s *service) ResotreUrl(code string) (string, error) {
	// get from cache
	var originalUrl string
	if uc, ok := s.cCache.Get(code); ok {
		// return uc.(string), nil
		originalUrl = uc.(string)
	} else {
		uc := s.urlModel.GetByCode(code)
		if uc.Url == "" {
			return "", errors.New("invalid code")
		}
		originalUrl = uc.Url
		go func() {
			s.cCache.Add(code, originalUrl)
		}()
	}
	// click increase
	go func() {
		// addClick <= code
	}()

	return originalUrl, nil
}

func TransToCode(id int) string {
	bytes := []byte("0lv12NUJ3789qazwegbyhnujmipQAZWsxSXEDCR4kt56FVTGBYHMIodcrfKLOP")

	var code string
	for m := id; m > 0; m = m / 62 {
		n := m % 62
		code += string(bytes[n])
		if m < 62 {
			break
		}
	}
	return code
}
