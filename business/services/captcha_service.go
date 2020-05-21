package services

import (
	"gothic-admin/business/repositories"
	"math/rand"
	"regexp"
	"strconv"
)

type CaptchaService interface {
	IsPhoneNumber(mobile string) bool
	SendCaptcha(mobile string) (captcha string)
	Check(mobile string, captcha string) bool
}

func NewCaptchaService(repo repositories.CaptchaRepo) CaptchaService {
	return &captchaService{repo:repo}
}

type captchaService struct {
	repo repositories.CaptchaRepo
}

func send2mobile(mobile string, captcha string) {
	// TODO: 发送验证码
}

func (s *captchaService) SendCaptcha(mobile string) (captcha string) {
	// 已存在验证码
	if b, v := s.repo.Exist(mobile); b == true {
		return v
	}
	// 生成验证码
	captcha = strconv.Itoa(1000 + rand.Intn(8999))
	err := s.repo.Set(mobile, captcha)
	if err != nil {
		panic(err)
	}
	// 发送验证码
	send2mobile(mobile, captcha)
	return
}

func (s *captchaService) Check(mobile string, captcha string) bool {

	if s.repo.Get(mobile) == captcha {
		s.repo.Delete(mobile)
		return true
	}

	return false
}

func (s *captchaService) IsPhoneNumber(mobile string) (is bool) {
	// TODO: 手机号验证
	is, _ = regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	return true
}