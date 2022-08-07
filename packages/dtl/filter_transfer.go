package dtl

import (
	"net/url"
	"xm-task/packages/entities"
)

func FilterToStruct(f url.Values) entities.Filter {
	return entities.Filter{
		Name:    f.Get("name"),
		Country: f.Get("country"),
		Phone:   f.Get("phone"),
		Website: f.Get("website"),
	}
}
