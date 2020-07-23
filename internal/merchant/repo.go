package merchant

type Merchant struct {
	Name   string `json:"name"`
	MID    string `json:"mid"`
	Mobile string `json:"mobile"`
}

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

func (r Repo) GetMerchantByMID(mid string) Merchant {
	return Merchant{
		Name:   "Chayanon",
		MID:    mid,
		Mobile: "0812345678",
	}
}
