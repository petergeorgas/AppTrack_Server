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

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	firestoreClient, err := datastore.NewFirestoreClient(os.Getenv("APP_TRACK_FIREBASE"))

	if err != nil {
		log.Fatalln("Failed to connect to Firebase/Firestore!" + err.Error())
	}
	defer firestoreClient.Close()
	resv := graph.NewResolver(firestoreClient)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resv}))
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://192.168.1.239", "http://192.168.1.239:3000", "http://localhost", "http://localhost:3000", "https://apptrack-178utrkn2-petergeorgas.vercel.app"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "OPTIONS"},
		AllowCredentials: true,
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cors.Handler(srv))

	log.Printf("SERVER STARTED ON PORT: %s" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
