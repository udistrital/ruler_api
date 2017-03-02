package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:DominioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:DominioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:DominioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:DominioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:DominioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:DominioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:DominioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:DominioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:DominioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:DominioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:PredicadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:PredicadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:PredicadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:PredicadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:PredicadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ruler_api/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
