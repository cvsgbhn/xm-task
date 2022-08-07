package dtl

import (
	"fmt"
	"strconv"
	"strings"
	"xm-task/packages/dbmodels"
	"xm-task/packages/entities"
)

func CompanyFromDB(c dbmodels.Company) entities.Company {
	h := fmt.Sprintf("%x", c.ID)

	return entities.Company{
		Name:      c.Name,
		Code:      h,
		Country:   c.Country,
		Website:   c.Website,
		Phone:     c.Phone,
		UpdatedAt: c.UpdatedAt,
	}
}

func CompaniesFromDB(c []dbmodels.Company) entities.Companies {
	dc := make([]entities.Company, len(c))

	for i, v := range c {
		dc[i] = CompanyFromDB(v)
	}

	return dc
}

func CompanyToDB(c entities.Company) dbmodels.Company {
	dStr := strings.Replace(c.Code, "0x", "", -1)
	dStr = strings.Replace(dStr, "0X", "", -1)
	d, _ := strconv.ParseInt(dStr, 16, 64)

	return dbmodels.Company{
		ID:        d,
		Name:      c.Name,
		Country:   c.Country,
		Website:   c.Website,
		Phone:     c.Phone,
		UpdatedAt: c.UpdatedAt,
	}
}
