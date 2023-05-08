package service

type assetService interface {
	GetAssets() ([]string, error)
}

func (srvc service) GetAssets() ([]string, error) {
	assets := []string{"ASSET11", "ASSET212", "ASSET465"}
	return assets, nil
}
