package main

import (
	"fmt"

	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/jsonschema"
)

type Person struct {
	Name string `json:"name,omitempty"`
}

func main() {
	cfg := schemaregistry.NewConfig("http://localhost:8081")

	cli, err := schemaregistry.NewClient(
		cfg,
	)

	if err != nil {
		fmt.Printf("Error creating schema registry client: %v\n", err)
		return
	}

	conf := jsonschema.NewSerializerConfig()

	conf.EnableValidation = true
	conf.AutoRegisterSchemas = false
	conf.UseLatestVersion = true

	fmt.Printf("Validation enabled: %v\n", conf.EnableValidation)

	js, err := jsonschema.NewSerializer(cli, 2, conf)

	if err != nil {
		fmt.Printf("Error creating json schema serializer: %v\n", err)
		return
	}

	tony := Person{
		Name: "Tony",
	}

	jsPeek, err := json.Marshal(tony)

	if err != nil {
		fmt.Printf("Error marshalling js: %v\n", err)
		return
	}

	fmt.Printf("About to serialize: %s\n", jsPeek)

	res, err := js.Serialize("tony-silly-test-topic-value", tony)

	if err != nil {
		fmt.Printf("Error serializing: %v\n", err)
		return
	}

	fmt.Printf("From: %v\n", jsPeek)
	fmt.Printf("To:   %v\n", res)
	fmt.Printf("%s\n", jsPeek)
	fmt.Printf("%s\n", res)
}
