package service

type BaseService interface {
	HealthCheck() bool
}

type baseService struct {
}

func NewBaseService() BaseService {
	return &baseService{}
}

func (bs baseService) HealthCheck() bool {
	return true
}
