package models

type Employee struct {
	FirstName       string  `json:"first_name,omitempty"`
	LastName        string  `json:"last_name,omitempty"`
	ID              int     `json:"ID,omitempty"`
	HireDate        string  `json:"hire_date,omitempty"`
	TerminationDate string  `json:"termination_date,omitempty"`
	Salary          float32 `json:"salary,omitempty"`
}

type NewEmployee struct {
	FirstName       string  `json:"first_name,omitempty"`
	LastName        string  `json:"last_name,omitempty"`
	ID              int     `json:"ID,omitempty"`
	HireDate        string  `json:"hire_date,omitempty"`
	TerminationDate string  `json:"termination_date,omitempty"`
	Salary          float32 `json:"salary,omitempty"`
	PerkiraanGaji   float32 `json:"perkiraan_gaji_2016,omitempty"`
	TotalUlasan     int     `json:"total_ulasan"`
}

type File struct {
	Data     interface{} `json:"data"`
	FileName string      `json:"file_name"`
}

type FileOpen struct {
	FileName string `param:"file_name" validate:"required" required:"file_name is required"`
}

type InputCity struct {
	City string `json:"city"`
}

type InputMulti struct {
	Data string `json:"data"`
}
type AnnualReviews struct {
	ID         int    `json:"ID,omitempty"`
	EmpID      int    `json:"emp_id,omitempty"`
	ReviewDate string `json:"review_date,omitempty"`
}

type TempDataArray struct {
	SortArray            []int    `json:"sort_array,omitempty"`
	TotalDuolicate       []string `json:"total_duplicate,omitempty"`
	DeleteValueFromInput []int    `json:"delete_value_from_input,omitempty"`
	SumValue             []int    `json:"sum_value,omitempty"`
}
