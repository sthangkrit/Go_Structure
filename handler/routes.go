package handler

import (
	"net/http"


	"Go_Structure/service/ping"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	//Validation  gin.HandlerFunc
}

type Routes struct {
	transaction []route
}

func (r Routes) InitTransactionRoute() http.Handler {

	ping := ping.NewEndpoint()


	r.transaction = []route{
		{
			Name:        "Ping Pong : GET",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/ping",
			Endpoint:    ping.PingGetEndpoint,
		},
		{
			Name:        "Ping Pong : GET Prams",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/ping/:name",
			Endpoint:    ping.PingGetParamsEndpoint,
		},
		{
			Name:        "Ping Pong : POST Prams+Body",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodPost,
			Pattern:     "/ping/:name",
			Endpoint:    ping.PingPostParamsAndBodyEndpoint,
		},
	}

	

	// txDocument := []route{
	// 	{
	// 		Name:     "Ping Pong : GET",
	// 		Method:   http.MethodGet,
	// 		Pattern:  "/ping",
	// 		Endpoint: documents.PingGetEndpoint,
	// 	},
	// 	{
	// 		Name:        "Ping Pong : GET Prams",
	// 		Description: "Ping Pong : Data",
	// 		Method:      http.MethodGet,
	// 		Pattern:     "/ping/:name",
	// 		Endpoint:    documents.PingGetParamsEndpoint,
	// 	},
	// 	{
	// 		Name:        "Ping Pong : POST Prams+Body",
	// 		Description: "Ping Pong : Data",
	// 		Method:      http.MethodPost,
	// 		Pattern:     "/ping/:name",
	// 		Endpoint:    documents.PingPostParamsAndBodyEndpoint,
	// 	},
	// 	{
	// 		Name:        "GET DocumentDetail Endpoint",
	// 		Description: "GET DocumentDetail Endpoint",
	// 		Method:      http.MethodPost,
	// 		Pattern:     "/getDocumentDetail",
	// 		Endpoint:    documents.GetDocumentDetailEndpoint,
	// 	},
	// 	{
	// 		Name:        "GET DocumentListPO Endpoint",
	// 		Description: "GET DocumentListPO Endpoint",
	// 		Method:      http.MethodPost,
	// 		Pattern:     "/getDocumentList",
	// 		Endpoint:    documents.GettPurchesOrderDocListEndpoint,
	// 	},
	// }

	

	ro := gin.New()

	

	store := ro.Group("/app")
	for _, e := range r.transaction {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}
