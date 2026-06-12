# Flowchart Diagram — HRIS Payroll System

> Diagram dibuat berdasarkan `gohrispayrollsystem.go`

---

## 1. Struktur & Relasi (Class Diagram)

```mermaid
classDiagram
    class HRIS {
        +Employees map~string, string~
        +Payrolls map~string, PayrollCalculator~
        +RegisterEmployee(id, name, payroll) error
        +CalculateTotalPayout() float64
        +PrintPayrollReport()
    }

    class PayrollCalculator {
        <<interface>>
        +CalculateSalary() float64, error
        +GetEmployeeType() string
        +PrintPayroll()
    }

    class FullTimeEmployee {
        +BaseSalary float64
        +Allowance float64
        +TaxRate float64
        +CalculateSalary() float64, error
        +GetEmployeeType() string
        +PrintPayroll()
    }

    class ContractEmployee {
        +MonthlyRate float64
        +PerformanceBonus float64
        +CalculateSalary() float64, error
        +GetEmployeeType() string
        +PrintPayroll()
    }

    class Freelancer {
        +HourlyRate float64
        +HoursWorked int
        +CalculateSalary() float64, error
        +GetEmployeeType() string
        +PrintPayroll()
    }

    HRIS --> PayrollCalculator : uses
    PayrollCalculator <|.. FullTimeEmployee : implements
    PayrollCalculator <|.. ContractEmployee : implements
    PayrollCalculator <|.. Freelancer : implements
```

---

## 2. Flowchart: RegisterEmployee

```mermaid
flowchart TD
    A([Start: RegisterEmployee]) --> B{id == kosong?}

    B -- Ya --> C[/Return ErrEmptyId/]
    B -- Tidak --> D{name == kosong?}

    D -- Ya --> E[/Return ErrEmptyName/]
    D -- Tidak --> F{ID sudah terdaftar<br>di Employees?}

    F -- Ya --> G[/Return ErrEmployeeRegistered/]
    F -- Tidak --> H[Panggil validatePayroll]

    H --> I{validatePayroll<br>return error?}
    I -- Ya --> J[/Return error/]
    I -- Tidak --> K["Simpan ke Employees[id] = name"]
    K --> L["Simpan ke Payrolls[id] = payroll"]
    L --> M[/Return nil/]

    style C fill:#f44,color:#fff
    style E fill:#f44,color:#fff
    style G fill:#f44,color:#fff
    style J fill:#f44,color:#fff
    style M fill:#4caf50,color:#fff
```

---

## 3. Flowchart: validatePayroll

```mermaid
flowchart TD
    A([Start: validatePayroll]) --> B{"Type switch:<br>payroll.(type)"}

    B -- "*FullTimeEmployee" --> C{"BaseSalary < 0 OR<br>Allowance < 0 OR<br>TaxRate < 0?"}
    C -- Ya --> ERR[/Return ErrNegativeValue/]
    C -- Tidak --> OK

    B -- "*ContractEmployee" --> D{"MonthlyRate < 0 OR<br>PerformanceBonus < 0?"}
    D -- Ya --> ERR
    D -- Tidak --> OK

    B -- "*Freelancer" --> E{"HourlyRate < 0 OR<br>HoursWorked < 0?"}
    E -- Ya --> ERR
    E -- Tidak --> OK

    B -- "Default" --> OK[/Return nil/]

    style ERR fill:#f44,color:#fff
    style OK fill:#4caf50,color:#fff
```

---

## 4. Flowchart: CalculateSalary (Polymorphism)

```mermaid
flowchart TD
    A([CalculateSalary dipanggil]) --> B{"Tipe Employee?"}

    B -- FullTimeEmployee --> C["(BaseSalary + Allowance)<br>× (1.0 − TaxRate)"]
    B -- ContractEmployee --> D["MonthlyRate<br>+ PerformanceBonus"]
    B -- Freelancer --> E["HourlyRate<br>× HoursWorked"]

    C --> F[/Return salary, nil/]
    D --> F
    E --> F

    style C fill:#1e88e5,color:#fff
    style D fill:#7b1fa2,color:#fff
    style E fill:#e65100,color:#fff
    style F fill:#4caf50,color:#fff
```

---

## 5. Flowchart: CalculateTotalPayout

```mermaid
flowchart TD
    A([Start: CalculateTotalPayout]) --> B["totalSalary = 0.0"]
    B --> C{Masih ada payroll<br>di Payrolls?}

    C -- Ya --> D["salary, err = payroll.CalculateSalary()"]
    D --> E{err == nil?}
    E -- Ya --> F["totalSalary += salary"]
    E -- Tidak --> C
    F --> C

    C -- Tidak --> G[/Return totalSalary/]

    style G fill:#4caf50,color:#fff
```

---

## 6. Flowchart: PrintPayrollReport

```mermaid
flowchart TD
    A([Start: PrintPayrollReport]) --> B["Print Header:<br>=========<br>Payroll Report<br>========="]
    B --> C{Masih ada<br>payroll di map?}

    C -- Ya --> D["Print separator: ---------"]
    D --> E["Print id & nama karyawan"]
    E --> F["payroll.PrintPayroll()"]
    F --> C

    C -- Tidak --> G["Print Footer: ========="]
    G --> H["Panggil CalculateTotalPayout()"]
    H --> I["Print Total Payout: Rp ..."]
    I --> J["Print Footer: ========="]
    J --> K([End])

    style B fill:#1e88e5,color:#fff
    style F fill:#7b1fa2,color:#fff
    style I fill:#4caf50,color:#fff
```

---

## 7. Flowchart Keseluruhan Sistem (Overview)

```mermaid
flowchart LR
    subgraph INPUT["📥 Input"]
        ID["id: string"]
        NAME["name: string"]
        PAYROLL["payroll: PayrollCalculator"]
    end

    subgraph REGISTER["📋 RegisterEmployee"]
        V1{"Validasi ID"}
        V2{"Validasi Nama"}
        V3{"Cek Duplikat"}
        V4["validatePayroll()"]
        SAVE["Simpan ke Maps"]
    end

    subgraph TYPES["👥 Tipe Karyawan"]
        FT["FullTimeEmployee<br>Gaji = (Base+Allowance)×(1−Tax)"]
        CT["ContractEmployee<br>Gaji = Monthly + Bonus"]
        FL["Freelancer<br>Gaji = Hourly × Hours"]
    end

    subgraph OUTPUT["📊 Output"]
        CALC["CalculateTotalPayout()"]
        REPORT["PrintPayrollReport()"]
    end

    INPUT --> V1
    V1 --> V2
    V2 --> V3
    V3 --> V4
    V4 --> SAVE

    SAVE --> TYPES

    TYPES --> CALC
    TYPES --> REPORT

    style FT fill:#1e88e5,color:#fff
    style CT fill:#7b1fa2,color:#fff
    style FL fill:#e65100,color:#fff
    style SAVE fill:#4caf50,color:#fff
    style CALC fill:#00897b,color:#fff
    style REPORT fill:#00897b,color:#fff
```
