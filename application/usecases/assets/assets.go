package assets

import (
	assetDomain "tennet/gethired/domain/assets"
	assetsRepository "tennet/gethired/infrastructure/repository/mysql/assets"
)

// Service is a struct that contains the repository implementation for assets use case
type Service struct {
	AssetRepository assetsRepository.Repository
}

// GetAll is a function that returns all assets
func (s *Service) GetAll(page int64, limit int64) (*PaginationResultAsset, error) {

	all, err := s.AssetRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return &PaginationResultAsset{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// GetByID is a function that returns a asset by id
func (s *Service) GetByID(id int) (*assetDomain.Asset, error) {
	return s.AssetRepository.GetByID(id)
}

// Create is a function that creates a asset
func (s *Service) Create(asset *NewAsset) (*assetDomain.Asset, error) {
	assetModel := asset.toDomainMapper()
	return s.AssetRepository.Create(assetModel)
}

// GetByMap is a function that returns a asset by map
func (s *Service) GetByMap(assetMap map[string]interface{}) (*assetDomain.Asset, error) {
	return s.AssetRepository.GetOneByMap(assetMap)
}

// Delete is a function that deletes a asset by id
func (s *Service) Delete(id int) error {
	return s.AssetRepository.Delete(id)
}

// Update is a function that updates a asset by id
func (s *Service) Update(id int64, asset UpdateAsset) (*assetDomain.Asset, error) {
	assetModel := asset.toDomainMapper()
	return s.AssetRepository.Update(id, &assetModel)
}
