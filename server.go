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
	"github.com/rs/cors"
)

const defaultPort = "8008"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	firestoreClient, err := datastore.NewFirestoreClient()

	if err != nil {
		log.Fatalln("Failed to connect to Firebase/Firestore!\n" + err.Error())
	}
	defer firestoreClient.Close()
	resv := graph.NewResolver(firestoreClient)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resv}))
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost*", "https://apptrack*.vercel.app"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "OPTIONS"},
		AllowCredentials: true,
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cors.Handler(srv))
	//
	log.Printf("SERVER STARTED ON PORT: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
