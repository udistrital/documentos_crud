package controllers

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/documentos_crud/models"
	"github.com/udistrital/utils_oas/time_bogota"
)

// FirmaElectronicaController operations for FirmaElectronica
type FirmaElectronicaController struct {
	beego.Controller
}

// URLMapping ...
func (c *FirmaElectronicaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create FirmaElectronica
// @Param	body		body 	models.FirmaElectronica	true		"body for FirmaElectronica content"
// @Success 201 {int} models.FirmaElectronica
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *FirmaElectronicaController) Post() {
	var v models.FirmaElectronica
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddFirmaElectronica(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get FirmaElectronica by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.FirmaElectronica
// @Failure 404 not found resource
// @router /:id [get]
func (c *FirmaElectronicaController) GetOne() {
	id := c.Ctx.Input.Param(":id")
	v, err1 := models.GetFirmaElectronicaById(id)
	//Inicio manipulación repuesta
	//Creación mapa
	var data map[string]interface{}
	//Conversion de string a mapa
	err := json.Unmarshal([]byte(v.DocumentoId.Metadatos), &data)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	}
	//Elimino y transformo a string de nuevo
	if llaves, ok := data["llaves"].(map[string]interface{}); ok {
		delete(llaves, "llave_privada")
	}
	modifiedJson, err := json.Marshal(data)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	}
	//Reasigno al campo de la estructura
	v.DocumentoId.Metadatos = string(modifiedJson)
	//Fin manipulación respuesta
	if err1 != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get FirmaElectronica
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.FirmaElectronica
// @Failure 404 not found resource
// @router / [get]
func (c *FirmaElectronicaController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = models.Alert{Type: "error", Code: "E_400", Body: "Error: invalid query key/value pair"}
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllFirmaElectronica(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = []interface{}{}
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the FirmaElectronica
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *FirmaElectronicaController) Delete() {
	id := c.Ctx.Input.Param(":id")
	if err := models.DeleteFirmaElectronica(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "404", "Body": err.Error(), "Type": "error"}
		c.Data["System"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the FirmaElectronica
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.FirmaElectronica	true		"body for FirmaElectronica content"
// @Success 200 {object} models.FirmaElectronica
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *FirmaElectronicaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.Atoi(idStr)
	v := models.FirmaElectronica{Id: idStr}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//v.FechaCreacion = time_bogota.TiempoCorreccionFormato(v.FechaCreacion)
		v.FechaModificacion = time_bogota.TiempoBogotaFormato()

		if err := models.UpdateFirmaElectronicaById(&v); err == nil {
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			c.Ctx.Output.SetStatus(400)
			//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
			c.Data["System"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		c.Ctx.Output.SetStatus(400)
		//c.Data["development"] = map[string]interface{}{"Code": "400", "Body": err.Error(), "Type": "error"}
		c.Data["System"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}
