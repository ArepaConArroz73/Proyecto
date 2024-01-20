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
package pageinfo

import (
	"negocio_backend/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type PageInfo struct {
	model    models.PageInfo
	dbmodels map[string]PageInfoModel
}

type PageInfoModel struct {
	object dbutils.ModelInterface
	model  interface{}
}

func NewPageInfo(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	dbmodels := map[string]PageInfoModel{
		"costumer": PageInfoModel{&models.Costumer{}, models.Costumer{}},
	}

	for _, v := range dbmodels {
		v.object.Init(v.model, db)
	}
	o = &PageInfo{
		dbmodels: dbmodels,
	}
	return o
}
func (o *PageInfo) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		r, err = o.edgesQuery(info)

	case "mutation":

	}
	return
}
func (o *PageInfo) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
