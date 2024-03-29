/*
 * Generado por gqlgenerate.
 *
 * Este archivo puede contener errores, de ser asi, coloca el issue en el repositorio de github
 * https://github.com/pjmd89/gogql
 *
 * Estos arvhivos corren riesgo de sobreescritura, por ese motivo gqlgnerate crea una carpeta llamada generate, asi que,
 * copia todas las carpetas que estan dentro de la carpeta generate y pegalas en la carpeta raiz de tu proyecto.
 *
 * gqlgenerate no creara archivos en la carpeta raiz de tu modulo porque puedes sufrir perdida de informacion.
 */
package main

import (
	"embed"
	"negocio_backend/lib"
	"negocio_backend/resolvers/directives"
	"negocio_backend/resolvers/objectTypes/changelog"
	"negocio_backend/resolvers/objectTypes/costumer"
	"negocio_backend/resolvers/objectTypes/edgecostumer"
	"negocio_backend/resolvers/objectTypes/pageinfo"
	"negocio_backend/resolvers/objectTypes/systeminfo"
	"negocio_backend/resolvers/scalars"

	"github.com/pjmd89/gogql/lib/gql"
	"github.com/pjmd89/gogql/lib/http"
	"github.com/pjmd89/gogql/lib/rest"
	"github.com/pjmd89/goutils/systemutils"
	"github.com/pjmd89/goutils/systemutils/debugmode"
	"github.com/pjmd89/mongomodel/mongomodel"
)

var (
	//go:embed "schema"
	embedFS embed.FS
)

var myConfig = lib.Config()
var systemLog = systemutils.NewLog(myConfig.SystemLog)
var accessLog = systemutils.NewLog(myConfig.AccessLog)
var logs = systemutils.Logs{System: systemLog, Access: accessLog}
var db = mongomodel.NewConn(&myConfig.DBConfigFile)
var schema = gql.Init(embedFS, "schema")

var restfull = rest.Init()
var myHttp = http.Init(logs, myConfig.HTTPConfigFile).SetGql(schema).SetRest(restfull)

func main() {
	lib.MyConfig = myConfig
	lib.Logs = logs
	systemLog.Info().Println("debugmode: ", debugmode.Enabled)
	myHttp.Start()
}
func OnDB(currentDB string, currentCollection string) (r string) {
	r = currentDB
	return
}
func httpOnSession() (r interface{}) {
	return
}
func httpCheckOrigin(url http.URL) (r bool, info interface{}) {
	r = true
	return
}
func httpOnBegin(url http.URL, httpPath *http.Path, originData interface{}) (r bool) {
	return true
}
func httpOnFinish() {}
func init() {
	db.(*mongomodel.MongoDBConn).OnDatabase = OnDB
	myHttp.OnSession = httpOnSession
	myHttp.CheckOrigin = httpCheckOrigin
	myHttp.OnBegin = httpOnBegin
	myHttp.OnFinish = httpOnFinish
	schema.ObjectType("SystemInfo", systeminfo.NewSystemInfo(db))
	schema.ObjectType("ChangeLog", changelog.NewChangeLog(db))
	schema.ObjectType("PageInfo", pageinfo.NewPageInfo(db))
	schema.ObjectType("Costumer", costumer.NewCostumer(db))
	schema.ObjectType("EdgeCostumer", edgecostumer.NewEdgeCostumer(db))
	schema.Scalar("DateTime", scalars.NewDateTimeScalar())
	schema.Scalar("DNI", scalars.NewDNIScalar())
	schema.Scalar("ID", scalars.NewIDScalar())
	schema.Directive("paginate", directives.NewPaginate())
	schema.Directive("sort", directives.NewSort())
	schema.Directive("search", directives.NewSearch(schema.GetScalars()))

}
