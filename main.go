package main

import (
    "log"
    "github.com/antihax/optional"
    petstore "github.com/goprog0248/petstore-client" // Adjust the import path
)

func main() {
    // Set up configuration
    configuration := petstore.NewConfiguration()
    apiClient := petstore.NewAPIClient(configuration)

    // Create a new pet
    pet := petstore.Pet{
        Id:   1,
        Name: "Buddy",
    }

    _, _, err := apiClient.PetApi.AddPet(pet)
    if err != nil {
        log.Fatalf("Error adding pet: %v", err)
    }

    log.Println("Pet added successfully")
}