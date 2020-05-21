package repositories

import (
	"github.com/go-redis/redis/v7"
	"gothic-admin/datasource"
	"time"
)

const (
	CaptchaPrefix = "captcha:"
)


type CaptchaRepo interface {
	Exist(mobile string) (bool, string)
	Set(mobile string, captcha string) error
	Get(mobile string) string
	Delete(mobile string)
}

type captchaRepo struct {
	cache datasource.Cache
}

func NewCaptchaRepo(cache datasource.Cache) CaptchaRepo {
	return &captchaRepo{cache:cache}
}

func (r *captchaRepo) buildKey(mobile string) string {
	return CaptchaPrefix + mobile
}

func (r *captchaRepo) Exist(mobile string) (bool, string) {
	key := r.buildKey(mobile)
	v, err := r.cache.Get(key).Result()
	if err == redis.Nil {
		return false, ""
	}
	return true, v
}

func (r *captchaRepo) Set(mobile string, captcha string) error {
	key := r.buildKey(mobile)
	err := r.cache.SetNX(key, captcha, 5 * time.Minute).Err()
	return err
}

func (r *captchaRepo) Get(mobile string) string {
	key := r.buildKey(mobile)
	v, err := r.cache.Get(key).Result()
	if err != nil {
		return ""
	}
	return v

}


func (r *captchaRepo) Delete(mobile string) {
	key := r.buildKey(mobile)
	_, err := r.cache.Del(key).Result()
	if err != nil {
		panic(err)
	}
}
