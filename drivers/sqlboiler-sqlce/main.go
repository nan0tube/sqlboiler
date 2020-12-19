package main

import (
	"github.com/volatiletech/sqlboiler/v4/drivers"
	"github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlce/driver"
)

func main() {
	drivers.DriverMain(&driver.SQLCEDriver{})
}
