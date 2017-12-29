package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	_ "github.com/nslcn/db"
	"github.com/nslcn/models"
)

// Nodes represent lights
type Nodes struct {
	beego.Controller
}

// Get returns list of lights
func (c *Nodes) Get() {
	c.Data["id"] = "ksjksjks"

	id := c.Ctx.Input.Param(":id")

	var data interface{}
	if id == "" {
		data = getAllNodes()
	} else {
		data = getNode(id)
	}

	log.Println("id is ", id)
	c.Data["json"] = data
	c.ServeJSON()
	//c.TplName = "index.tpl"
}

//Post adds a light node
func (c *Nodes) Post() {
	node := new(models.Node)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, node); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		beego.Error(err, c.Ctx.Input.RequestBody)
		return
	}

	o := orm.NewOrm()
	beego.Info(o.Insert(node))
	c.ServeJSON()
}

//Put updates a light node
func (c *Nodes) Put() {
	node := new(models.Node)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, node); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte(err.Error()))
		beego.Error(err, c.Ctx.Input.RequestBody)
		return
	}

	id := c.Ctx.Input.Param(":id")
	if id != node.ID {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("mismatch in node id"))
		beego.Error("mismatch in node id. Param:", id, ",body id:", node.ID)
		return
	}

	o := orm.NewOrm()
	if num, err := o.Update(node); err != nil {
		beego.Error("update error:", err)
		if err == orm.ErrMissPK {
			c.Ctx.Output.SetStatus(404)
			c.Ctx.Output.Body([]byte(err.Error()))
			return
		}
	} else if num == 0 {
		beego.Error(id, "update failed")
		c.Ctx.Output.SetStatus(404)
	}

	c.ServeJSON()
}

//Delete deletes a light node
func (c *Nodes) Delete() {
	node := new(models.Node)
	node.ID = c.Ctx.Input.Param(":id")

	o := orm.NewOrm()
	beego.Info(o.Delete(node))
	c.ServeJSON()
}

func getAllNodes() (nodes []*models.Node) {
	o := orm.NewOrm()

	_, err := o.QueryTable(new(models.Node)).All(&nodes)
	handleDBError(err)
	return
}

func getNode(id string) (node *models.Node) {
	o := orm.NewOrm()
	node = &models.Node{ID: id}

	err := o.Read(node)
	handleDBError(err)
	return
}
func handleDBError(err error) {
	if err == orm.ErrNoRows {
		log.Println("no nodes found")
	} else if err == orm.ErrMissPK {
		log.Println("no primary key found")
	}
}
