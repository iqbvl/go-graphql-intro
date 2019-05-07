package main

import (
	"context"
	"iqbvl/go-graphql-intro/resolvers"
	"iqbvl/go-graphql-intro/schema"
	"iqbvl/go-graphql-intro/services"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"

	repo_admin_user "iqbvl/go-graphql-intro/business-services/admin-user/repository"
)

func main() {
	conn, err := service.OpenDB()
	defer conn.Close()
	if err != nil {
		log.Panic(err)
	}
	ctx := context.Background()

	adminUser := repo_admin_user.NewSqlAdminUserRepository(conn)
	ctx = context.WithValue(ctx, "AdminUserRepository", adminUser)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		//AllowCredentials:   true,
		AllowedMethods: []string{"POST", "GET", "OPTIONS"},
		// OptionsPassthrough: true,
		AllowedHeaders: []string{"X-Requested-With",
			"Authorization",
			"Content-Type",
			"X-Auth-Token",
			"Origin",
			"Accept"},
		Debug: true,
	})
	Handle(ctx, c) // Login etc
	// s := `
	//         schema {
	//                 query: Query
	//         }
	//         type Query {
	//                 hello: String!
	//         }
	// `
	// schema := graphql.MustParseSchema(s, &query{})
	// http.Handle("/query", &relay.Handler{Schema: schema})
	log.Panic(http.ListenAndServe(":8080", nil))
}

func Handle(ctx context.Context, c *cors.Cors) {
	graphql.MaxDepth(5)

	resolverQuery := &resolvers.Resolver{}
	graphqlSchema := graphql.MustParseSchema(schema.GetSchema(), resolverQuery)
	http.Handle("/query",
		c.Handler(
			service.AddContext(ctx,
				&relay.Handler{Schema: graphqlSchema},
			),
		),
	)

	// resolverSystem := &resolvers.Resolver{}
	// graphqlSystemSchema := graphql.MustParseSchema(schema.GetSchema(), resolverSystem)
	// http.Handle("/system",
	// 	c.Handler(
	// 		service.AddContext(ctx,
	// 			&relay.Handler{Schema: graphqlSystemSchema},
	// 		),
	// 	),
	// )
}
