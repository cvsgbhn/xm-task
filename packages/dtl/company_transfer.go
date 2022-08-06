package dtl

import (
	"fmt"
	"strconv"
	"strings"
	"xm-task/packages/domain"
	"xm-task/packages/storage"
)

func CompanyFromDB(c storage.Company) domain.Company {
	h := fmt.Sprintf("%x", c.ID)

	return domain.Company{
		Name:      c.Name,
		Code:      h,
		Country:   c.Country,
		Website:   c.Website,
		Phone:     c.Phone,
		UpdatedAt: c.UpdatedAt,
	}
}

func CompanyToDB(c domain.Company) storage.Company {
	dStr := strings.Replace(c.Code, "0x", "", -1)
	dStr = strings.Replace(dStr, "0X", "", -1)
	d, _ := strconv.ParseInt(dStr, 16, 64)

	return storage.Company{
		ID:        d,
		Name:      c.Name,
		Country:   c.Country,
		Website:   c.Website,
		Phone:     c.Phone,
		UpdatedAt: c.UpdatedAt,
	}
}
