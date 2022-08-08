package dtl

import (
	"fmt"
	"strconv"
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
	d, _ := strconv.ParseInt(c.Code, 16, 64)

	return dbmodels.Company{
		ID:        d,
		Name:      c.Name,
		Country:   c.Country,
		Website:   c.Website,
		Phone:     c.Phone,
		UpdatedAt: c.UpdatedAt,
	}
}
