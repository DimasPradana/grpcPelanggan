package models

type StructPelanggan struct {
	Unit        string `json:"unit,omitempty"`
	Alamat      string `json:"alamat,omitempty"`
	Namapelang  string `json:"namapelang,omitempty"`
	WkbGeometry string `json:"wkb_geometry,omitempty"`
	NoLanggan   string `json:"no_langgan,omitempty"`
	NoSambung   string `json:"no_sambung,omitempty"`
}
