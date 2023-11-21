package models

import (
	"fmt"
	"log"
	"reakgo/utility"

	"github.com/jmoiron/sqlx"
)

type Company struct {
	Id            int64  `json:"companyId" db:"id"`
	Name          string `json:"companyName" db:"name"`
	Gstin         string `json:"gstin" db:"gstin"`
	Address       string `json:"companyaddress" db:"address"`
	ContactNumber string `json:"contactNumber" db:"contactNumber"`
	ContactEmail  string `json:"contactEmail" db:"contactEmail"`
}

func (comp Company) PutCompany(company Company, tx *sqlx.Tx) (int64, error) {
	var insterid int64
	rows, err := tx.NamedExec("INSERT INTO `company` (id,name,gstin,address,contactNumber,contactEmail) VALUES ( :Id,:Name,:Gstin,:Address,:ContactNumber,:ContactEmail)", map[string]interface{}{"Id": company.Id, "Name": company.Name, "Gstin": company.Gstin, "Address": company.Address, "ContactNumber": company.ContactNumber, "ContactEmail": company.ContactEmail})
	// Check error
	if err != nil {
		log.Println(err)
		//logger remove for duplicate entry that means duplicate error message not send at email.
		istrue, _ := utility.CheckSqlError(err, "Duplicate entry")
		if !istrue {
			utility.Logger(err)
		}
		return insterid, err
	}
	return rows.LastInsertId()
}

func (comp Company) PostCompany(company Company, tx *sqlx.Tx) (bool, error) {
	userData, err := tx.NamedExec("UPDATE company SET name=:Name,gstin=:Gstin,address =:Address,contactNumber=:ContactNumber,contactEmail=:ContactEmail WHERE id=:Id", map[string]interface{}{"Id": company.Id, "Name": company.Name, "Gstin": company.Gstin, "Address": company.Address, "ContactNumber": company.ContactNumber, "ContactEmail": company.ContactEmail})
	// Check error
	if err != nil {
		log.Println("error: ", err)
		utility.Logger(err)
	} else {
		Rowefffect, _ := userData.RowsAffected()
		if Rowefffect == 0 {
			log.Println("input value is not change with previous one or id= " + fmt.Sprint(company.Id) + "is not valid")
		}
		return Rowefffect > 0, err
	}
	return false, err
}
