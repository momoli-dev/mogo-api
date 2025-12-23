//nolint:forbidigo // Example uses fmt for output.
package main

import (
	"fmt"

	"github.com/momoli-dev/mogo-api/api"
	"github.com/momoli-dev/mogo-api/server"
)

func main() {
	myAPI := api.New(&api.NewParams{
		Title:      "Pet Store",
		Version:    "1.0.0",
		Origins:    []string{"http://*, https://*"},
		EnableDocs: true,
	})

	myAPI.Tag("pet")
	api.Put(myAPI, "/pet", api.NotImplemented, "Update an existing pet.")
	api.Post(myAPI, "/pet", api.NotImplemented, "Add a new pet to the store.")
	api.Get(myAPI, "/pet/findByStatus", api.NotImplemented, "Finds Pets by status.")
	api.Get(myAPI, "/pet/findByTags", api.NotImplemented, "Finds Pets by tags.")
	api.Get(myAPI, "/pet/{petId}", api.NotImplemented, "Find pet by ID.")
	api.Post(myAPI, "/pet/{petId}", api.NotImplemented, "Updates a pet in the store with form data.")
	api.Delete(myAPI, "/pet/{petId}", api.NotImplemented, "Deletes a pet.")
	api.Post(myAPI, "/pet/{petId}/uploadImage", api.NotImplemented, "Uploads an image.")

	server := server.NewServer(&server.Params{
		Addr:    ":8080",
		Handler: myAPI.GetHTTPHandler(),
	})

	<-server.StartGracefully()
	fmt.Println("Server shutting down...")
}
