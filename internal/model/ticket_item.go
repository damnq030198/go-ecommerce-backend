package model

// VO: Get ticketItems returns
type TicketItemsOutput struct {
	TicketId       int    `json:"ID"`             // Sửa tag và thêm field
	TicketName     string `json:"Name"`           // Sửa tag
	StockAvailable int    `json:"StockAvailable"` // Sửa tag
	StockInitial   int    `json:"StockInitial"`   // Sửa tag
}

// DTO
type TicketItemRequest struct {
	TicketId string `json:"ticket_Id"`
}
