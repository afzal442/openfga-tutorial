package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"

    "github.com/openfga/go-sdk/client"
)

func main() {
    // Initialize FGA client
    fgaClient, err := client.NewSdkClient(&client.ClientConfiguration{
        ApiUrl:               os.Getenv("FGA_API_URL"),   // required, e.g. https://api.fga.example
        StoreId:              os.Getenv("FGA_STORE_ID"),  // optional, not needed for `CreateStore` and `ListStores`, required before calling all other methods
        AuthorizationModelId: os.Getenv("FGA_MODEL_ID"),  // optional, can be overridden per request
    })

    if err != nil {
        fmt.Println("Error creating FGA client:", err)
        return
    }

    // Create a new store
    resp, err := fgaClient.CreateStore(context.Background()).Body(client.ClientCreateStoreRequest{Name: "FGA Demo"}).Execute()
    if err != nil {
        fmt.Println("Error creating store:", err)
        return
    }
    fmt.Println("Store created successfully:", resp)

    // Define and unmarshal the WriteAuthorizationModelRequest JSON
    var writeAuthorizationModelRequestString = "{\"schema_version\":\"1.1\",\"type_definitions\":[{\"type\":\"user\"},{\"type\":\"document\",\"relations\":{\"reader\":{\"this\":{}},\"writer\":{\"this\":{}},\"owner\":{\"this\":{}}},\"metadata\":{\"relations\":{\"reader\":{\"directly_related_user_types\":[{\"type\":\"user\"}]},\"writer\":{\"directly_related_user_types\":[{\"type\":\"user\"}]},\"owner\":{\"directly_related_user_types\":[{\"type\":\"user\"}]}}}}]}"
    var body client.ClientWriteAuthorizationModelRequest
    if err := json.Unmarshal([]byte(writeAuthorizationModelRequestString), &body); err != nil {
        fmt.Println("Error unmarshalling authorization model request:", err)
        return
    }

    // Write the authorization model
    data, err := fgaClient.WriteAuthorizationModel(context.Background()).
        Body(body).
        Execute()

    if err != nil {
        fmt.Println("Error writing authorization model:", err)
        return
    }

    fmt.Println("Authorization Model ID:", data.AuthorizationModelId)
}
