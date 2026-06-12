package gohrispayrollsystem

import (
	"errors"
	"fmt"
)

var (
	ErrEmployeeRegistered = errors.New("invalid : karyawan dengan ID tersebut sudah terdaftar")
	ErrEmptyId            = errors.New("invalid : id karyawan tidak boleh kosong")
	ErrEmptyName          = errors.New("invalid : nama karyawan tidak boleh kosong")
	ErrNegativeValue      = errors.New("invalid : input tidak boleh bernilai negatif")
)

type HRIS struct {
	Employees map[string]string            // id_employee, nama
	Payrolls  map[string]PayrollCalculator // id_employee, kontrak
}

func validatePayroll(payroll PayrollCalculator) error {
	switch v := payroll.(type) {
	case *FullTimeEmployee:
		if v.BaseSalary < 0 || v.Allowance < 0 || v.TaxRate < 0 {
			return ErrNegativeValue
		}
	case *ContractEmployee:
		if v.MonthlyRate < 0 || v.PerformanceBonus < 0 {
			return ErrNegativeValue
		}
	case *Freelancer:
		if v.HourlyRate < 0 || v.HoursWorked < 0 {
			return ErrNegativeValue
		}
	}
	return nil
}

func (hris *HRIS) RegisterEmployee(id string, name string, payroll PayrollCalculator) error {
	if id == "" {
		return ErrEmptyId
	}
	if name == "" {
		return ErrEmptyName
	}
	_, exists := hris.Employees[id]

	if exists {
		return ErrEmployeeRegistered
	}

	if err := validatePayroll(payroll); err != nil {
		return err
	}

	hris.Employees[id] = name
	hris.Payrolls[id] = payroll

	return nil
}

func (hris *HRIS) CalculateTotalPayout() float64 {
	totalSalary := 0.0
	for _, payroll := range hris.Payrolls {
		if salary, err := payroll.CalculateSalary(); err == nil {
			totalSalary += salary
		}
	}
	return totalSalary
}

func (hris *HRIS) PrintPayrollReport() {
	fmt.Println("=========================================================")
	fmt.Println("Payroll Report")
	fmt.Println("=========================================================")
	for id, payroll := range hris.Payrolls {
		fmt.Println("---------------------------------------------------------")
		fmt.Println("id: ", id, "Nama: ", hris.Employees[id])
		payroll.PrintPayroll()
	}
	fmt.Println("=========================================================")
	fmt.Printf("Total Payout: Rp %.2f\n", hris.CalculateTotalPayout())
	fmt.Println("=========================================================")
}

type PayrollCalculator interface {
	CalculateSalary() (float64, error)
	GetEmployeeType() string // FullTime, Contract, Freelance
	PrintPayroll()
}

type FullTimeEmployee struct {
	BaseSalary float64
	Allowance  float64
	TaxRate    float64
}

func (emp *FullTimeEmployee) CalculateSalary() (float64, error) {
	return (emp.BaseSalary + emp.Allowance) * (1.0 - emp.TaxRate), nil
}

func (emp *FullTimeEmployee) GetEmployeeType() string {
	return "FullTime"
}

func (emp *FullTimeEmployee) PrintPayroll() {
	total_salary, _ := emp.CalculateSalary()
	fmt.Printf(
		"Tipe : %s\nGaji Pokok : Rp %.2f\nTunjangan : Rp %.2f\nPajak : Rp %.2f\nTotal Gaji: Rp %.2f\n",
		emp.GetEmployeeType(),
		emp.BaseSalary,
		emp.Allowance,
		emp.TaxRate,
		total_salary,
	)
}

type ContractEmployee struct {
	MonthlyRate      float64
	PerformanceBonus float64
}

func (emp *ContractEmployee) CalculateSalary() (float64, error) {
	return emp.MonthlyRate + emp.PerformanceBonus, nil
}

func (emp *ContractEmployee) GetEmployeeType() string {
	return "Contract"
}

func (emp *ContractEmployee) PrintPayroll() {
	total_salary, _ := emp.CalculateSalary()
	fmt.Printf(
		"Tipe : %s\nGaji per Bulan : Rp %.2f\nBonus : Rp %.2f\nTotal Gaji: Rp %.2f\n",
		emp.GetEmployeeType(),
		emp.MonthlyRate,
		emp.PerformanceBonus,
		total_salary,
	)
}

type Freelancer struct {
	HourlyRate  float64
	HoursWorked int
}

func (emp *Freelancer) CalculateSalary() (float64, error) {
	return emp.HourlyRate * float64(emp.HoursWorked), nil
}

func (emp *Freelancer) GetEmployeeType() string {
	return "Freelance"
}

func (emp *Freelancer) PrintPayroll() {
	total_salary, _ := emp.CalculateSalary()
	fmt.Printf(
		"Tipe : %s\nGaji per Jam : Rp %.2f\nJam Kerja : %d jam\nTotal Gaji: Rp %.2f\n",
		emp.GetEmployeeType(),
		emp.HourlyRate,
		emp.HoursWorked,
		total_salary,
	)
}
