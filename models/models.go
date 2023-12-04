package models

type QRCodeParams struct {
	AssetTag          string `json:"asset_tag" validate:"required"`
	AssetSerialNumber string `json:"asset_serial_number" validate:"required"`
	AssetType         string `json:"asset_type" validate:"required"`
	AssetCost         string `json:"asset_cost" validate:"min=1"`
	OsType            string `json:"os_type" validate:"required"`
	Vendor            string `json:"vendor" validate:"required"`
	PurchasedDate     string `json:"purchased_date" validate:"required"`
	EntityType        string `json:"entity" validate:"required"`
}
