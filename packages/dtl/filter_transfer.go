package dtl

import (
	"net/url"
	"xm-task/packages/domain"
)

func FilterToStruct(f url.Values) domain.Filter {
	return domain.Filter{
		Name:    f.Get("name"),
		Country: f.Get("country"),
		Phone:   f.Get("phone"),
		Website: f.Get("website"),
	}
}
