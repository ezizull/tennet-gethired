package assets

import (
	"encoding/json"

	domainAsset "tennet/gethired/domain/assets"
	domainErrors "tennet/gethired/domain/errors"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all asset data
func (r *Repository) GetAll(page int64, limit int64) (*PaginationResultAsset, error) {
	var assets []Asset
	var total int64

	err := r.DB.Model(&Asset{}).Count(&total).Error
	if err != nil {
		return &PaginationResultAsset{}, err
	}
	offset := (page - 1) * limit
	err = r.DB.Limit(int(limit)).Offset(int(offset)).Find(&assets).Error

	if err != nil {
		return &PaginationResultAsset{}, err
	}

	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &PaginationResultAsset{
		Data:       arrayToDomainMapper(&assets),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil
}

// Create ... Insert New data
func (r *Repository) Create(newAsset *domainAsset.Asset) (createdAsset *domainAsset.Asset, err error) {
	asset := fromDomainMapper(newAsset)

	tx := r.DB.Create(asset)

	if tx.Error != nil {
		byteErr, _ := json.Marshal(tx.Error)
		var newError domainErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return
		}
		switch newError.Number {
		case 1062:
			err = domainErrors.NewAppErrorWithType(domainErrors.ResourceAlreadyExists)
		default:
			err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		}
		return
	}

	createdAsset = asset.toDomainMapper()
	return
}

// GetByID ... Fetch only one asset by Id
func (r *Repository) GetByID(id int) (*domainAsset.Asset, error) {
	var asset Asset
	err := r.DB.Where("id = ?", id).First(&asset).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		default:
			err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		}
		return &domainAsset.Asset{}, err
	}

	return asset.toDomainMapper(), nil
}

// GetOneByMap ... Fetch only one asset by Map
func (r *Repository) GetOneByMap(updateAsset map[string]interface{}) (*domainAsset.Asset, error) {
	var asset Asset

	err := r.DB.Where(updateAsset).Limit(1).Find(&asset).Error
	if err != nil {
		err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		return nil, err
	}
	return asset.toDomainMapper(), err
}

// Update ... Update asset
func (r *Repository) Update(id int64, updateAsset *domainAsset.Asset) (*domainAsset.Asset, error) {
	var asset Asset

	asset.ID = id
	err := r.DB.Model(&asset).
		Updates(updateAsset).Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError domainErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &domainAsset.Asset{}, err
		}
		switch newError.Number {
		case 1062:
			err = domainErrors.NewAppErrorWithType(domainErrors.ResourceAlreadyExists)
			return &domainAsset.Asset{}, err

		default:
			err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
			return &domainAsset.Asset{}, err
		}
	}

	err = r.DB.Where("id = ?", id).First(&asset).Error

	return asset.toDomainMapper(), err
}

// Delete ... Delete asset
func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Delete(&Asset{}, id)
	if tx.Error != nil {
		err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = domainErrors.NewAppErrorWithType(domainErrors.NotFound)
	}

	return
}
