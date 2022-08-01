// Copyright 2016 The OPA Authors.  All rights reserved.
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package main

import (
/*
	"fmt"
	"os"

	"github.com/open-policy-agent/opa/cmd"
*/
	"context"
	"encoding/json"
	"fmt"
	"log"
	//"os"
    "strings"
	"github.com/open-policy-agent/opa/rego"

)

/*
func main() {
	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
*/

var input_json = `{
  "servers": [
    {
      "id": "app",
      "protocols": [ "https", "ssh" ],
      "ports": [ "p1", "p2", "p3" ]
    },
    {
      "id": "db",
      "protocols": [ "mysql" ],
      "ports": [ "p3" ]
    },
    {
      "id": "cache",
      "protocols": [ "memcache" ],
      "ports": [ "p3" ]
    },
    {
      "id": "ci",
      "protocols": [ "http" ],
      "ports": [ "p1", "p2" ]
    },
    {
      "id": "busybox",
      "protocols": [ "telnet" ],
      "ports": [ "p1" ]
    }
  ],
  "networks": [
    {
      "id": "net1",
      "public": false
    },
    {
      "id": "net2",
      "public": false
    },
    {
      "id": "net3",
      "public": true
    },
    {
      "id": "net4",
      "public": true
    }
  ],
  "ports": [
    {
      "id": "p1",
      "network": "net1"
    },
    {
      "id": "p2",
      "network": "net3"
    },
    {
      "id": "p3",
      "network": "net2"
    }
  ]
}`


func main() {

	ctx := context.Background()

	rego_file := "/home/ishtiyaque/example.rego"
	query_string := "data.example.violation";

	r := rego.New(
		rego.Query(query_string),
		rego.Load([]string{rego_file}, nil))

	// Construct a Rego object that can be prepared or evaluated.
/*
	r := rego.New(
		rego.Query(os.Args[2]),
		rego.Load([]string{os.Args[1]}, nil))
*/
	// Create a prepared query that can be evaluated.
	query, err := r.PrepareForEval(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Load the input document from stdin.
	var input interface{}
	//dec := json.NewDecoder(os.Stdin)
    dec := json.NewDecoder(strings.NewReader(input_json))

    
	dec.UseNumber()
	if err := dec.Decode(&input); err != nil {
		log.Fatal(err)
	}
    
     // json.Unmarshal([]byte(input_json), &input)
     //fmt.Println( input)

	// Execute the prepared query.
	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatal(err)
	}

    // Do something with the result.
	fmt.Println(rs)
}

// Capabilities + built-in metadata file generation:
//go:generate build/gen-run-go.sh internal/cmd/genopacapabilities/main.go capabilities.json
//go:generate build/gen-run-go.sh internal/cmd/genbuiltinmetadata/main.go builtin_metadata.json

// WASM base binary generation:
//go:generate build/gen-run-go.sh internal/cmd/genopawasm/main.go -o internal/compiler/wasm/opa/opa.go internal/compiler/wasm/opa/opa.wasm  internal/compiler/wasm/opa/callgraph.csv
