package entity

// Department is a struct that represent department entity
type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// DepartmentUpdate is a struct that represent department update entity
type DepartmentUpdate struct {
	ID   int     `json:"id"`
	Name *string `json:"name"`
}
