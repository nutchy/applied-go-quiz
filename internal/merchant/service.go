package merchant

type Service struct {
	repo repo
}

type repo interface {
	GetMerchantByMID(mid string) Merchant
}

func NewService() *Service {
	return &Service{}
}

func (s Service) GetMerchantByMID(mid string) Merchant {
	return s.repo.GetMerchantByMID("23")
}
