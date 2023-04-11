package assets

import domainAsset "tennet/gethired/domain/assets"

func (n *NewAsset) toDomainMapper() *domainAsset.Asset {
	return &domainAsset.Asset{
		Name: n.Name,
	}
}

func (n *UpdateAsset) toDomainMapper() domainAsset.Asset {
	domainAsset := domainAsset.Asset{}

	if n.Name != nil {
		domainAsset.Name = *n.Name
	}

	return domainAsset
}
