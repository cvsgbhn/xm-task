package dtl

import (
	"net/url"
	"reflect"
	"testing"
	"time"
	"xm-task/packages/dbmodels"
	"xm-task/packages/entities"
)

func TestCompaniesFromDB(t *testing.T) {
	testTime := time.Now().UTC()
	type args struct {
		c []dbmodels.Company
	}
	tests := []struct {
		name string
		args args
		want entities.Companies
	}{
		{
			name: "Simple test - companies from db",
			args: args{
				c: []dbmodels.Company{
					{
						ID:        1,
						Name:      "a",
						Country:   "b",
						Website:   "c",
						Phone:     "d",
						CreatedAt: testTime,
						UpdatedAt: testTime,
						DeletedAt: time.Time{},
					}, {
						ID:        125,
						Name:      "q",
						Country:   "w",
						Website:   "e",
						Phone:     "r",
						CreatedAt: testTime,
						UpdatedAt: testTime,
						DeletedAt: time.Time{},
					},
				},
			},
			want: []entities.Company{
				{
					Code:      "1",
					Name:      "a",
					Country:   "b",
					Website:   "c",
					Phone:     "d",
					UpdatedAt: testTime,
				}, {
					Code:      "7d",
					Name:      "q",
					Country:   "w",
					Website:   "e",
					Phone:     "r",
					UpdatedAt: testTime,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompaniesFromDB(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompaniesFromDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyToDB(t *testing.T) {
	testTime := time.Now().UTC()

	type args struct {
		c entities.Company
	}
	tests := []struct {
		name string
		args args
		want dbmodels.Company
	}{
		{
			name: "Simple test - company to db",
			args: args{
				entities.Company{
					Code:      "a",
					Name:      "test",
					Country:   "qwerty",
					Website:   "asdf",
					Phone:     "134234324",
					UpdatedAt: testTime,
				},
			},
			want: dbmodels.Company{
				ID:        10,
				Name:      "test",
				Country:   "qwerty",
				Website:   "asdf",
				Phone:     "134234324",
				UpdatedAt: testTime,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompanyToDB(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyToDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterToStruct(t *testing.T) {
	type args struct {
		f url.Values
	}
	tests := []struct {
		name string
		args args
		want entities.Filter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterToStruct(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterToStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
