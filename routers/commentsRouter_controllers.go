package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:DocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:SubtipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:SubtipoDocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:SubtipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:SubtipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:SubtipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:SubtipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:SubtipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:SubtipoDocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:SubtipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:SubtipoDocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/documentos_crud/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
