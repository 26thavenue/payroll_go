package repositories

import (
	"github.com/26thavenue/payroll_go/models"
	"context"
)

type PayrollRepository struct {
	db any
}

func (r *PayrollRepository)GetAll (ctx context.Context)([]*models.Payroll, error) {
	return nil, nil
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


func NewPayrollRepository(db any) models.PayrollRepository{
	return &PayrollRepository{
		db:db,
	}
}