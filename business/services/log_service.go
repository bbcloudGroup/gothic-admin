package services

import (
	"gothic-admin/business/models"
	"gothic-admin/business/repositories"
)

type LogService interface {
	Delete(params models.Form) error
	GetPage(current int, pageSize int, params map[string]string) models.Page
	LogOperation(userID uint, menuID uint, methodID uint, data string)
}

func NewLogService(repo repositories.AdminRepo) LogService {
	return &logService{repo:repo}
}

type logService struct {
	repo repositories.AdminRepo
}

func (s *logService) Delete(params models.Form) (err error) {
	if len(params.IDS) == 0 {
		return
	}
	return s.repo.DeleteLog(params.IDS)
}

func (s *logService) GetPage(current int, pageSize int, params map[string]string) models.Page {
	var data []interface{}

	logs, count :=  s.repo.GetLogPage(current, pageSize, params)

	for _, log := range logs {
		data = append(data, log)
	}

	return models.Page{
		Success:  true,
		Current:  current,
		PageSize: pageSize,
		Total:    count,
		Data:     data,
	}
}



func (s *logService) LogOperation(userID uint, menuID uint, methodID uint, data string) {

	_, _ = s.repo.AddLog(userID, menuID, methodID, data)

}
