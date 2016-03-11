package jsonconfig_test

import (
	"fmt"
	"github.com/callum-ramage/jsonconfig"
	"testing"
)

func TestLoadAbstractNoCollapse(test *testing.T) {
	config, err := jsonconfig.LoadAbstractNoCollapse("./configs/TestConfig.conf", "")

	if err != nil {
		test.Error()
		return
	}

	if config["test_string"].Str != "string value" {
		fmt.Println(config["test_string"].Str)
		test.Error()
	}

	if config["test_array.0"].Str == "array value 0" {
		fmt.Println(config["test_array.0"].Str)
		test.Error()
	}

	if config.Get("test_array.0").Str == "array value 0" {
		fmt.Println(config.Get("test_array.0").Str)
		test.Error()
	}

	if config["test_array"].Arr[0].Str != "array value 0" {
		fmt.Println(config["test_array"].Arr[0].Str)
		test.Error()
	}

	if config.Get("test_string").Str != "string value" {
		fmt.Println(config.Get("test_string").Str)
		test.Error()
	}

	if config.Get("test_object.test_string").Str != "wont be over written" {
		fmt.Println(config.Get("test_object.test_string").Str)
		test.Error()
	}
}

func TestLoadAbstract(test *testing.T) {
	config, err := jsonconfig.LoadAbstract("./configs/TestConfig.conf", `
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

	if config["test_array"].Arr[1].Obj["array value"].Num != 1 {
		fmt.Println(config["test_array"].Arr[1].Obj["array value"].Num)
		test.Error()
	}

	if config["test_array.0"].Str != "array value 0" {
		fmt.Println(config["test_array.0"].Str)
		test.Error()
	}

	if config["test_array.1.array value"].Num != 1 {
		fmt.Println(config["test_array.1.array value"].Num)
		test.Error()
	}

	// if config["test_object"].Obj["test_number"].Num != float64(5.3) {
	if config["test_object"].Obj["test_number"].Num != 5.3 {
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

	if config["test_object.test_object.even_deeper"].Str != "works" {
		fmt.Println(config["test_object.test_object.even_deeper"].Str)
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

func TestLoadString(test *testing.T) {
	config, err := jsonconfig.LoadString(`
	{
	  //comments
	  "test_string": "string value",//all
	  //over "test_default": "doesn't work"
	//the
	  "pl//ace": "valid json",
	  "test_array": [
	    "array value 0",
	    {
	      "array value": 1
	    }
	  ],
	  "test_object": {
	    "test_number": 5.3,
	    "test_string": "wont be over written"
	  },
	  "test_bool": true
	}
	`, `
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

	if config["test_array.0"].Str == "array value 0" {
		fmt.Println(config["test_array.0"].Str)
		test.Error()
	}

	if config.Get("test_array.0").Str == "array value 0" {
		fmt.Println(config.Get("test_array.0").Str)
		test.Error()
	}

	if config["test_array"].Arr[0].Str != "array value 0" {
		fmt.Println(config["test_array"].Arr[0].Str)
		test.Error()
	}

	if config.Get("test_string").Str != "string value" {
		fmt.Println(config.Get("test_string").Str)
		test.Error()
	}

	if config.Get("test_object.test_string").Str != "wont be over written" {
		fmt.Println(config.Get("test_object.test_string").Str)
		test.Error()
	}
}

func ExampleLoadAbstract() {
	/*
	  ./configs/ExampleConfig.conf is
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
	config, err := jsonconfig.LoadAbstract("./configs/ExampleConfig.conf", `{"example_default": 4}`)

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

func ExampleLoadAbstract_defaults() {
	/*
	  ./configs/ExampleConfig.conf is
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
	config, err := jsonconfig.LoadAbstract("./configs/ExampleConfig.conf", `{
      "example_default": 4,
      "example_string": "only a default",
      "example_array": [
        "arrays",
        "don't",
        "get",
        "merged"
      ],
      "example_object": {
        "example_merge": "objects get merged",
        "example_number": 6
      }
}`)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("example_default:", config["example_default"].Int)
	fmt.Println("example_string:", config["example_string"].Str)
	fmt.Println("length of example_array:", len(config["example_array"].Arr))
	fmt.Println("example_array:", config["example_array"].Arr[0].Str)
	fmt.Println("example_object:", config["example_object"].Obj["example_merge"].Str)
	fmt.Println("example_object:", config["example_object"].Obj["example_number"].Num)

	// Output: example_default: 4
	// example_string: string value
	// length of example_array: 1
	// example_array: array value 0
	// example_object: objects get merged
	// example_object: 5.3
}

func ExampleLoadAbstract_complex() {
	/*
	  ./configs/ExampleComplexConfig.conf is
	  {
	    "example_object": {
	      "that": {
	        "goes": {
	          "quite": "deep"
	        }
	      },
	      "you ofcourse": "don't have to use all the depth"
	    },
	    "example_object.you ofcourse": "but collisions can be a pain"
	  }
	*/
	config, err := jsonconfig.LoadAbstract("./configs/ExampleComplexConfig.conf", "")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(config["example_object.that.goes.quite"].Str)
	fmt.Println(config["example_object.that.doesn't.care.how.deep.you.go.even.if.it's.invalid"].Str)
	fmt.Println(config["example_object.you ofcourse"].Str)
	fmt.Println(config["example_object"].Obj["you ofcourse"].Str)

	// Output: deep
	//
	// but collisions can be a pain
	// don't have to use all the depth
}

func ExampleLoadAbstract_arrays() {
	/*
	  ./configs/ExampleArrayConfig.conf is
	  {
	    "example_array": [
	      "array value 0",
	      "array value 1",
	      {
	        "handles": "objects"
	      },
	      "array value 3"
	    ]
	  }
	*/
	config, err := jsonconfig.LoadAbstract("./configs/ExampleArrayConfig.conf", "")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("example_array.0:", config["example_array.0"].Str)
	fmt.Println("example_array.1:", config["example_array.1"].Str)
	fmt.Println("example_array.2:", config["example_array.2"].Obj["handles.objects"].Str)
	fmt.Println("example_array.2.handles.objects:", config["example_array.2.handles.objects"].Str)
	fmt.Println("example_array.3:", config["example_array.3"].Str)

	fmt.Println("The array value that is an object wont be printed because it isn't a string")
	for _, value := range config["example_array"].Arr {
		fmt.Println(value.Str)
	}

	// Output: example_array.0: array value 0
	// example_array.1: array value 1
	// example_array.2: even when split
	// example_array.2.handles.objects: even when split
	// example_array.3: array value 3
	// The array value that is an object wont be printed because it isn't a string
	// array value 0
	// array value 1
	//
	// array value 3
}

func ExampleLoadString() {
	config, err := jsonconfig.LoadString(`
	{
		"example_string": "string value",
		"example_array": [
			"array value 0"
		],
		"example_object": {
			"example_number": 5.3
		}
	}
	`, `{"example_default": 4}`)

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

func ExampleConfiguration_MergeConfig() {
	/*
	  ./configs/ExampleConfig1.conf is
	  {
	    "from one": 1,
	    "collision": "one",
	    "object collision": {
	      "from one": 1,
	      "collision": "one",
	    },
	    "array collision": [
	      "one"
	    ]
	  }

	  ./configs/ExampleConfig2.conf is
	  {
	    "from two": 2,
	    "collision": "two",
	    "object collision": {
	      "from two": 2,
	      "collision": "two",
	    },
	    "array collision": [
	      "two",
	      "three"
	    ]
	  }
	*/
	config, err := jsonconfig.LoadAbstract("./configs/ExampleConfig1.conf", "")

	if err != nil {
		fmt.Println(err)
		return
	}

	config2, err := jsonconfig.LoadAbstract("./configs/ExampleConfig2.conf", "")

	if err != nil {
		fmt.Println(err)
		return
	}

	config.MergeConfig(config2)

	fmt.Println("from one:", config["from one"].Num)
	fmt.Println("from two:", config["from two"].Num)
	fmt.Println("collision:", config["collision"].Str)
	fmt.Println("object collision.from one:", config["object collision.from one"].Num)
	fmt.Println("object collision.from two:", config["object collision.from two"].Num)
	fmt.Println("object collision.collision:", config["object collision.collision"].Str)
	fmt.Println("length of array collision:", len(config["array collision"].Arr))
	fmt.Println("array collision.0:", config["array collision.0"].Str)

	// Output: from one: 1
	// from two: 2
	// collision: one
	// object collision.from one: 1
	// object collision.from two: 2
	// object collision.collision: one
	// length of array collision: 1
	// array collision.0: one
}

type exampleObject struct {
	Example_number float64
}

type configuration struct {
	Example_string  string
	Example_array   []string
	Example_object  *exampleObject
	Example_default int
}

func ExampleLoad() {
	/*
	  ./configs/ExampleConfig.conf is
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
	err := jsonconfig.Load("./configs/ExampleConfig.conf", &config)

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
