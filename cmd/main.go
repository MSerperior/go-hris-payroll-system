package main

import (
	"fmt"

	hris "go-hris-payroll-system"
)

func main() {
	fmt.Println("=========================================================")
	fmt.Println("       HRIS & PAYROLL SYSTEM - Demo & Testing")
	fmt.Println("=========================================================")

	// Inisialisasi sistem HRIS
	system := &hris.HRIS{
		Employees: make(map[string]string),
		Payrolls:  make(map[string]hris.PayrollCalculator),
	}

	// =====================================================
	// 1. Registrasi Karyawan (Happy Path)
	// =====================================================
	fmt.Println("\n--- [1] Registrasi Karyawan ---")

	// Karyawan Full-Time
	err := system.RegisterEmployee("FT-001", "Budi Santoso", &hris.FullTimeEmployee{
		BaseSalary: 10_000_000,
		Allowance:  2_000_000,
		TaxRate:    0.10,
	})
	printResult("Register Budi Santoso (FullTime)", err)

	// Karyawan Full-Time kedua
	err = system.RegisterEmployee("FT-002", "Siti Aminah", &hris.FullTimeEmployee{
		BaseSalary: 12_000_000,
		Allowance:  3_000_000,
		TaxRate:    0.15,
	})
	printResult("Register Siti Aminah (FullTime)", err)

	// Karyawan Kontrak
	err = system.RegisterEmployee("CT-001", "Andi Wijaya", &hris.ContractEmployee{
		MonthlyRate:      7_000_000,
		PerformanceBonus: 1_500_000,
	})
	printResult("Register Andi Wijaya (Contract)", err)

	// Freelancer
	err = system.RegisterEmployee("FL-001", "Dewi Lestari", &hris.Freelancer{
		HourlyRate:  150_000,
		HoursWorked: 120,
	})
	printResult("Register Dewi Lestari (Freelance)", err)

	// =====================================================
	// 2. Pengujian Validasi & Error Handling
	// =====================================================
	fmt.Println("\n--- [2] Pengujian Validasi & Error Handling ---")

	// Test: ID kosong
	err = system.RegisterEmployee("", "Tanpa ID", &hris.FullTimeEmployee{
		BaseSalary: 5_000_000,
		Allowance:  1_000_000,
		TaxRate:    0.05,
	})
	printResult("Register dengan ID kosong", err)

	// Test: Nama kosong
	err = system.RegisterEmployee("XX-001", "", &hris.FullTimeEmployee{
		BaseSalary: 5_000_000,
		Allowance:  1_000_000,
		TaxRate:    0.05,
	})
	printResult("Register dengan Nama kosong", err)

	// Test: ID duplikat
	err = system.RegisterEmployee("FT-001", "Duplikat Budi", &hris.FullTimeEmployee{
		BaseSalary: 8_000_000,
		Allowance:  1_000_000,
		TaxRate:    0.10,
	})
	printResult("Register dengan ID duplikat (FT-001)", err)

	// Test: Gaji negatif (FullTime)
	err = system.RegisterEmployee("FT-099", "Negatif FullTime", &hris.FullTimeEmployee{
		BaseSalary: -5_000_000,
		Allowance:  1_000_000,
		TaxRate:    0.10,
	})
	printResult("Register FullTime dengan gaji negatif", err)

	// Test: Gaji negatif (Contract)
	err = system.RegisterEmployee("CT-099", "Negatif Contract", &hris.ContractEmployee{
		MonthlyRate:      -3_000_000,
		PerformanceBonus: 500_000,
	})
	printResult("Register Contract dengan gaji negatif", err)

	// Test: Gaji negatif (Freelancer)
	err = system.RegisterEmployee("FL-099", "Negatif Freelancer", &hris.Freelancer{
		HourlyRate:  -100_000,
		HoursWorked: 80,
	})
	printResult("Register Freelancer dengan gaji negatif", err)

	// =====================================================
	// 3. Hitung Gaji per Karyawan
	// =====================================================
	fmt.Println("\n--- [3] Hitung Gaji per Karyawan ---")

	for id, payroll := range system.Payrolls {
		salary, err := payroll.CalculateSalary()
		if err != nil {
			fmt.Printf("  [%s] %s => ERROR: %v\n", id, system.Employees[id], err)
		} else {
			fmt.Printf("  [%s] %s (%s) => Rp %.2f\n",
				id,
				system.Employees[id],
				payroll.GetEmployeeType(),
				salary,
			)
		}
	}

	// =====================================================
	// 4. Total Payout
	// =====================================================
	fmt.Println("\n--- [4] Total Payout ---")
	totalPayout := system.CalculateTotalPayout()
	fmt.Printf("  Total seluruh gaji: Rp %.2f\n", totalPayout)

	// =====================================================
	// 5. Cetak Laporan Payroll Lengkap
	// =====================================================
	fmt.Println("\n--- [5] Laporan Payroll Lengkap ---")
	system.PrintPayrollReport()

	fmt.Println("\n\n=========================================================")
	fmt.Println("       Semua pengujian selesai!")
	fmt.Println("=========================================================")
}

// printResult mencetak hasil operasi: sukses atau error
func printResult(operation string, err error) {
	if err != nil {
		fmt.Printf("  ✗ %s => ERROR: %v\n", operation, err)
	} else {
		fmt.Printf("  ✓ %s => OK\n", operation)
	}
}
