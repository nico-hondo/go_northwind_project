// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package models

import (
	"database/sql"
)

type Category struct {
	CategoryID   int16          `db:"category_id" json:"categoryId"`
	CategoryName string `db:"category_name" json:"categoryName"`
	Description  string `db:"description" json:"description"`
	Picture      []byte         `db:"picture" json:"picture"`
}

type Customer struct {
	CustomerID   string         `db:"customer_id" json:"customerId"`
	CompanyName  string `db:"company_name" json:"companyName"`
	ContactName  string `db:"contact_name" json:"contactName"`
	ContactTitle string `db:"contact_title" json:"contactTitle"`
	Address      string `db:"address" json:"address"`
	City         string `db:"city" json:"city"`
	Region       string `db:"region" json:"region"`
	PostalCode   string `db:"postal_code" json:"postalCode"`
	Country      string `db:"country" json:"country"`
	Phone        string `db:"phone" json:"phone"`
	Fax          string `db:"fax" json:"fax"`
}

type Employee struct {
	EmployeeID      int16          `db:"employee_id" json:"employeeId"`
	LastName        string `db:"last_name" json:"lastName"`
	FirstName       string `db:"first_name" json:"firstName"`
	Title           string `db:"title" json:"title"`
	TitleOfCourtesy string `db:"title_of_courtesy" json:"titleOfCourtesy"`
	BirthDate       sql.NullTime   `db:"birth_date" json:"birthDate"`
	HireDate        sql.NullTime   `db:"hire_date" json:"hireDate"`
	Address         string `db:"address" json:"address"`
	City            string `db:"city" json:"city"`
	Region          string `db:"region" json:"region"`
	PostalCode      string `db:"postal_code" json:"postalCode"`
	Country         string `db:"country" json:"country"`
	HomePhone       string `db:"home_phone" json:"homePhone"`
	Extension       string `db:"extension" json:"extension"`
	Photo           []byte         `db:"photo" json:"photo"`
	Notes           string `db:"notes" json:"notes"`
	ReportTo        sql.NullInt16  `db:"report_to" json:"reportTo"`
	PhotoPath       string `db:"photo_path" json:"photoPath"`
}

type Order struct {
	OrderID        int16           `db:"order_id" json:"orderId"`
	OrderDate      sql.NullTime    `db:"order_date" json:"orderDate"`
	RequiredDate   sql.NullTime    `db:"required_date" json:"requiredDate"`
	ShippedDate    sql.NullTime    `db:"shipped_date" json:"shippedDate"`
	Freight        sql.NullFloat64 `db:"freight" json:"freight"`
	ShipName       string  `db:"ship_name" json:"shipName"`
	ShipAddress    string  `db:"ship_address" json:"shipAddress"`
	ShipCity       string  `db:"ship_city" json:"shipCity"`
	ShipRegion     string  `db:"ship_region" json:"shipRegion"`
	ShipPostalCode string  `db:"ship_postal_code" json:"shipPostalCode"`
	ShipCountry    string  `db:"ship_country" json:"shipCountry"`
	EmployeeID     sql.NullInt16   `db:"employee_id" json:"employeeId"`
	CustomerID     string  `db:"customer_id" json:"customerId"`
	ShipperID      sql.NullInt16   `db:"shipper_id" json:"shipperId"`
}

type OrderDetail struct {
	OrderID   int16           `db:"order_id" json:"orderId"`
	ProductID int16           `db:"product_id" json:"productId"`
	UnitPrice sql.NullFloat64 `db:"unit_price" json:"unitPrice"`
	Quantity  sql.NullInt16   `db:"quantity" json:"quantity"`
	Discount  sql.NullFloat64 `db:"discount" json:"discount"`
}

type Product struct {
	ProductID       int16           `db:"product_id" json:"productId"`
	ProductName     string  `db:"product_name" json:"productName"`
	QuantityPerUnit string  `db:"quantity_per_unit" json:"quantityPerUnit"`
	UnitPrice       sql.NullFloat64 `db:"unit_price" json:"unitPrice"`
	UnitsInStock    sql.NullInt16   `db:"units_in_stock" json:"unitsInStock"`
	UnitsInOrder    sql.NullInt16   `db:"units_in_order" json:"unitsInOrder"`
	ReorderLevel    sql.NullInt16   `db:"reorder_level" json:"reorderLevel"`
	Discontinued    sql.NullInt32   `db:"discontinued" json:"discontinued"`
	SupplierID      sql.NullInt16   `db:"supplier_id" json:"supplierId"`
	CategoryID      sql.NullInt16   `db:"category_id" json:"categoryId"`
}

type Shipper struct {
	ShipperID   int16          `db:"shipper_id" json:"shipperId"`
	CompanyName string `db:"company_name" json:"companyName"`
	Phone       string `db:"phone" json:"phone"`
}

type Supplier struct {
	SupplierID   int16          `db:"supplier_id" json:"supplierId"`
	CompanyName  string `db:"company_name" json:"companyName"`
	ContactName  string `db:"contact_name" json:"contactName"`
	ContactTitle string `db:"contact_title" json:"contactTitle"`
	Address      string `db:"address" json:"address"`
	City         string `db:"city" json:"city"`
	Region       string `db:"region" json:"region"`
	PostalCode   string `db:"postal_code" json:"postalCode"`
	Country      string `db:"country" json:"country"`
	Phone        string `db:"phone" json:"phone"`
	Fax          string `db:"fax" json:"fax"`
	Homepage     string `db:"homepage" json:"homepage"`
}
