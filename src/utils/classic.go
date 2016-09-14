package utils

import (
	"github.com/codegangsta/negroni"

)

func Classic() *negroni.Negroni {
	return negroni.New(NewLogger(),SessionMiddleware(), Instrumentation())
}
