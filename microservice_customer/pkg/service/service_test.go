package service

import (
	"microservice_customer/pkg/database"
	"microservice_customer/pkg/entity"
	"reflect"
	"testing"
)

func TestNewCostumerService(t *testing.T) {
	type args struct {
		dabase_pool database.DatabaseInterface
	}
	tests := []struct {
		name string
		args args
		want *customer_service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCostumerService(tt.args.dabase_pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCostumerService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_customer_service_GetCustomers(t *testing.T) {
	tests := []struct {
		name    string
		ps      *customer_service
		want    *entity.CustomerList
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetCustomers()
			if (err != nil) != tt.wantErr {
				t.Errorf("customer_service.GetCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customer_service.GetCustomers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_customer_service_GetCustomerByID(t *testing.T) {
	type args struct {
		ID *uint64
	}
	tests := []struct {
		name    string
		ps      *customer_service
		args    args
		want    *entity.Customer
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ps.GetCustomerByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("customer_service.GetCustomerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customer_service.GetCustomerByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_customer_service_CreateCustomer(t *testing.T) {
	type args struct {
		customer *entity.Customer
		logID    *int
	}
	tests := []struct {
		name    string
		ps      *customer_service
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ps.CreateCustomer(tt.args.customer, tt.args.logID); (err != nil) != tt.wantErr {
				t.Errorf("customer_service.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_customer_service_UpdateCustomer(t *testing.T) {
	type args struct {
		ID       *uint64
		customer *entity.Customer
		logID    *int
	}
	tests := []struct {
		name    string
		ps      *customer_service
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ps.UpdateCustomer(tt.args.ID, tt.args.customer, tt.args.logID); (err != nil) != tt.wantErr {
				t.Errorf("customer_service.UpdateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_customer_service_UpdateStatusCustomer(t *testing.T) {
	type args struct {
		ID    *uint64
		logID *int
	}
	tests := []struct {
		name    string
		ps      *customer_service
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ps.UpdateStatusCustomer(tt.args.ID, tt.args.logID); (err != nil) != tt.wantErr {
				t.Errorf("customer_service.UpdateStatusCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
