package models

import (
	"fmt"
	"net/http"
)

type Balance struct {
	UserId  string  `json:"UserId"`
	Balance float64 `json:"Balance"`
}

type GetBalance struct {
	UserId string `json:"UserId"`
}

type ReserveBalance struct {
	UserId    string  `json:"UserId"`
	ServiceId string  `json:"ServiceId"`
	OrderId   string  `json:"OrderId"`
	Price     float64 `json:"Price"`
}

type Revenue struct {
	UserId    string  `json:"UserId"`
	ServiceId string  `json:"ServiceId"`
	OrderId   string  `json:"OrderId"`
	Price     float64 `json:"Price"`
}

func (*Revenue) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (i *Revenue) Bind(r *http.Request) error {
	if i.UserId == "" {
		return fmt.Errorf("UserId is a required field")
	}
	if i.OrderId == "" {
		return fmt.Errorf("OrderId is a required field")
	}
	if i.ServiceId == "" {
		return fmt.Errorf("ServiceId is a required filed")
	}
	if i.Price < 0 {
		return fmt.Errorf("postive price is a required")
	}
	return nil
}

func (*ReserveBalance) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (i *ReserveBalance) Bind(r *http.Request) error {
	if i.UserId == "" {
		return fmt.Errorf("UserId is a required field")
	}
	if i.OrderId == "" {
		return fmt.Errorf("OrderId is a required field")
	}
	if i.ServiceId == "" {
		return fmt.Errorf("ServiceId is a required filed")
	}
	if i.Price < 0 {
		return fmt.Errorf("postive price is a required")
	}
	return nil
}

func (*GetBalance) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (i *GetBalance) Bind(r *http.Request) error {
	if i.UserId == "" {
		return fmt.Errorf("UserId is a required field")
	}
	return nil
}

func (*Balance) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (i *Balance) Bind(r *http.Request) error {
	if i.UserId == "" {
		return fmt.Errorf("UserId is a required field")
	}
	if i.Balance < 0 {
		return fmt.Errorf("positive balance is a required")
	}
	return nil
}
