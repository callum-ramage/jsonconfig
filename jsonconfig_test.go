package jsonconfig_test

import (
  "github.com/callum-ramage/jsonconfig"
  "fmt"
  "testing"
)

func TestLoadAbstract(test *testing.T) {
  config, err := jsonconfig.LoadAbstract("./TestConfig.json", `
  {
    "test_string": "try to overwrite",
    "test_array": [
      "try to overwrite array"
    ],
    "test_object": {
      "test_string": "try to overwrite",
      "test_object": {
        "even_deeper": "works"
      },
      "test_default": "works"
    },
    "test_default": "works"
  }
  `)
  if err != nil {
    test.Error()
    return
  }

  if config["test_string"].Str != "string value" {
    fmt.Println(config["test_string"].Str)
    test.Error()
  }

  if config["pl//ace"].Str != "valid json" {
    fmt.Println(config["pl//ace"].Str)
    test.Error()
  }

  if config["test_array"].Arr[0].Str != "array value 0" {
    fmt.Println(config["test_array"].Arr[0].Str)
    test.Error()
  }

  if config["test_object"].Obj["test_number"].Num != float64(5.3) {
    fmt.Println(config["test_object"].Obj["test_number"].Num)
    test.Error()
  }

  if config["test_object"].Obj["test_number"].Int != 5 {
    fmt.Println(config["test_object"].Obj["test_number"].Int)
    test.Error()
  }

  if config["test_object"].Obj["test_string"].Str != "wont be over written" {
    fmt.Println(config["test_object"].Obj["test_string"].Str)
    test.Error()
  }

  if config["test_object"].Obj["test_object"].Obj["even_deeper"].Str != "works" {
    fmt.Println(config["test_object"].Obj["test_object"].Obj["even_deeper"].Str)
    test.Error()
  }

  if config["test_object"].Obj["test_default"].Str != "works" {
    fmt.Println(config["test_object"].Obj["test_default"].Str)
    test.Error()
  }

  if !config["test_bool"].Bool {
    fmt.Println(config["test_bool"].Bool)
    test.Error()
  }

  if config["test_default"].Str != "works" {
    fmt.Println(config["test_default"].Str)
    test.Error()
  }
}

func ExampleLoadAbstract() {
  /*
  ./ExampleConfig.json is
  {
    "example_string": "string value",
    "example_array": [
      "array value 0"
    ],
    "example_object": {
      "example_number": 5.3
    }
  }
  */
  config, err := jsonconfig.LoadAbstract("./ExampleConfig.json", `{"example_default": 4}`)

  if err != nil {
    return
  }

  fmt.Println("example_string:", config["example_string"].Str)
  fmt.Println("example_array:", config["example_array"].Arr[0].Str)
  fmt.Println("example_object:", config["example_object"].Obj["example_number"].Num)
  fmt.Println("example_default:", config["example_default"].Int)

  // Output: example_string: string value
  // example_array: array value 0
  // example_object: 5.3
  // example_default: 4
}

type exampleObject struct {
  Example_number float64
}

type configuration struct {
  Example_string string
  Example_array []string
  Example_object *exampleObject
  Example_default int
}

func ExampleLoad() {
  /*
  ./ExampleConfig.json is
  {
    "example_string": "string value",
    "example_array": [
      "array value 0"
    ],
    "example_object": {
      "example_number": 5.3
    }
  }
  */
  config := configuration{Example_default: 4}
  err := jsonconfig.Load("./ExampleConfig.json", &config)

  if err != nil {
    return
  }

  fmt.Println("example_string:", config.Example_string)
  fmt.Println("example_array:", config.Example_array[0])
  fmt.Println("example_object:", config.Example_object.Example_number)
  fmt.Println("example_default:", config.Example_default)

  // Output: example_string: string value
  // example_array: array value 0
  // example_object: 5.3
  // example_default: 4
}
