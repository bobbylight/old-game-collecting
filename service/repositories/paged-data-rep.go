package repositories

type PagedDataRep struct {

    Data []*GameRecord `json:"data,omitempty"`
    Start int          `json:"start"`
    Total int          `json:"total,omitempty"`
}
