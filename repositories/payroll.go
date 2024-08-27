package repositories

import (
	"context"

	"github.com/26thavenue/payroll_go/models"
	"gorm.io/gorm"
)

type PayrollRepository struct {
	db *gorm.DB
}

func (r *PayrollRepository)GetAll (ctx context.Context)([]*models.Payroll, error) {

	payrolls := []*models.Payroll{}

	res := r.db.Model(&models.Payroll{}).Find(&payrolls)

	if res.Error !=nil{
		return nil, res.Error
	}

	
	return payrolls, nil
}

func (r *PayrollRepository)GetOne (ctx context.Context, payrollId string)(*models.Payroll, error) {
	return nil, nil
}

func (r *PayrollRepository)CreatePayroll(ctx context.Context, payroll models.Payroll)(*models.Payroll, error) {
	return nil, nil
}

func (r *PayrollRepository)DeletePayroll (ctx context.Context, payrollId string) error{
	return nil
}


func NewPayrollRepository(db *gorm.DB) models.PayrollRepository{
	return &PayrollRepository{
		db:db,
	}
}