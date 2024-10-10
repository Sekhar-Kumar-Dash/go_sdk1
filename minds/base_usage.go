package minds

import (
	"fmt"
	"go_sdk/minds"
	"os"
	// Adjust import path
)

func main() {
	// --- Connect ---
	apiKey := os.Getenv("MINDSDB_API_KEY") // Assuming you've set the API key as an env variable
	if apiKey == "" {
		panic("MINDSDB_API_KEY environment variable is not set")
	}

	client := minds.NewClient(apiKey, "") // Use default base URL

	// Or use a custom server
	// baseURL := "https://custom_cloud.mdb.ai/"
	// client := minds.NewClient(apiKey, baseURL)

	// --- Create Datasource ---
	postgresConfig := &minds.DatabaseConfig{
		Name:        "my_datasource",
		Description: "<DESCRIPTION-OF-YOUR-DATA>",
		Engine:      "postgres",
		ConnectionData: map[string]string{
			"user":     "demo_user",
			"password": "demo_password",
			"host":     "samples.mindsdb.com",
			"port":     "5432", // Note: port should be a string
			"database": "demo",
			"schema":   "demo_data",
		},
		Tables: []string{"<TABLE-1>", "<TABLE-2>"},
	}

	// --- Create Mind ---
	// With datasource at the same time
	createOpts := &minds.CreateMindOptions{
		Datasources: []interface{}{*postgresConfig},
	}
	mind, err := client.Minds.Create("mind_name", createOpts, false)
	if err != nil {
		panic(err)
	}

	// // Or separately
	// datasource, err := client.Datasources.Create(postgresConfig, false)
	// if err != nil {
	// 	panic(err)
	// }
	// mind, err = client.Minds.Create("mind_name", &minds.CreateMindOptions{
	// 	Datasources: []interface{}{datasource},
	// }, false)
	// if err != nil {
	// 	panic(err)
	// }

	// With prompt template
	createOpts = &minds.CreateMindOptions{
		PromptTemplate: &[]string{"You are a coding assistant."}[0],
	}
	mind, err = client.Minds.Create("mind_name_with_template", createOpts, false)
	if err != nil {
		panic(err)
	}

	// Or add to an existing mind
	mind, err = client.Minds.Create("mind_name_to_update", nil, false)
	if err != nil {
		panic(err)
	}
	// By config
	if err := mind.AddDatasource(*postgresConfig); err != nil {
		panic(err)
	}
	// Or by datasource
	// if err := mind.AddDatasource(datasource); err != nil {
	// 	panic(err)
	// }

	// --- Managing Minds ---
	// Create or replace
	mind, err = client.Minds.Create("mind_to_replace", createOpts, true)
	if err != nil {
		panic(err)
	}

	// Update
	newName := "updated_mind_name"
	updateOpts := &minds.UpdateMindOptions{
		Name:        &newName,
		Datasources: []interface{}{*postgresConfig}, // Replace current datasources
	}
	if err := mind.Update(updateOpts); err != nil {
		panic(err)
	}

	// List
	minds, err := client.Minds.List()
	if err != nil {
		panic(err)
	}
	fmt.Println("Minds:", minds)

	// Get by name
	mind, err = client.Minds.Get("updated_mind_name")
	if err != nil {
		panic(err)
	}
	fmt.Println("Retrieved Mind:", mind)

	// Removing datasource
	if err := mind.DelDatasource("my_datasource"); err != nil {
		panic(err)
	}

	// Remove Mind
	if err := client.Minds.Drop("updated_mind_name"); err != nil {
		panic(err)
	}

	// Call completion (replace "mind_name_with_template" with an existing Mind name)
	completion, err := mind.Completion("2+3", false)
	if err != nil {
		panic(err)
	}
	fmt.Println("Completion:", completion)

	// Stream completion (replace "mind_name_with_template" with an existing Mind name)
	fmt.Println("Streaming Completion:")
	streamCompletion, err := mind.Completion("What is the capital of France?", true)
	if err != nil {
		panic(err)
	}
	fmt.Println(streamCompletion)

	// --- Managing Datasources ---
	// Create or replace
	datasource, err = client.Datasources.Create(postgresConfig, true)
	if err != nil {
		panic(err)
	}

	// List
	datasources, err := client.Datasources.List()
	if err != nil {
		panic(err)
	}
	fmt.Println("Datasources:", datasources)

	// Get
	datasource, err = client.Datasources.Get("my_datasource")
	if err != nil {
		panic(err)
	}
	fmt.Println("Retrieved Datasource:", datasource)

	// Remove
	if err := client.Datasources.Drop("my_datasource"); err != nil {
		panic(err)
	}
}
