package controllers

import (
	"dashboard_beego/models"
	"dashboard_beego/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

//MainController the main
type MainController struct {
	beego.Controller
}

// Get func
func (c *MainController) Get() {

	subtopo := c.GetString(":subtopo")
	m := sql.QuerySubtopo(subtopo)
	outgoingJSON, error := json.MarshalIndent(m, "", "    ")
	if error != nil {
		log.Println(error.Error())
		return
	}
	var runList []models.RunDB
	err := json.Unmarshal(outgoingJSON, &runList)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("test passed: %d", runList[0].TcPassed)
	//c.Data["runs"] = runList[0].TcPassed
	c.Data["Name"] = "naveen"
	c.TplName = "test.tpl"
}
