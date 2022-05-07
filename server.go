package main

import (
	"apptrack/datastore"
	"apptrack/graph"
	"apptrack/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	firestoreClient, err := datastore.NewFirestoreClient("application-tracker-5027c")

	if err != nil {
		log.Fatalln("Failed to connect to Firebase/Firestore!" + err.Error())
	}
	defer firestoreClient.Close()
	resv := graph.NewResolver(firestoreClient)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resv}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
