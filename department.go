/*
Get department data

We are presently getting JSON and then returning JSON, which seems dumb,
however the retreival will be replaced with some other data soure in the
future, this is just to get the package up and running
*/
package department

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

const departmentsFile string = "static-data/departments.json"

type DepartmentList struct {
	Departments []Department // this variable name needs to begin with an upper case?
}

type Department struct {
	Full_name string // this variable name needs to begin with an upper case?
}

func main() {
	departmentData, err := ioutil.ReadFile(departmentsFile)
	if err != nil { panic(err) }
	//fmt.Printf("%s\n", string(departmentData))

    var departmentsJson DepartmentList
    json.Unmarshal(departmentData, &departmentsJson)
    //fmt.Printf("Results: %v\n", departmentsJson)

    reJson, _ := json.Marshal(departmentsJson)
    fmt.Printf("%v\n", string(reJson))
}
