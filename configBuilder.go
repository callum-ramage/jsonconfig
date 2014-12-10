package jsonconfig

import (
  "encoding/json"
  "os"
  "strings"
)

//Is a convenience struct that makes working with abstract JSON data more tolerable.
//The internal values Arr and Obj can be nil, so should not be assumed to be safe.
type JSONValue struct{
  Value interface{}
  Arr []JSONValue
  Str string
  Int int
  Num float64
  Bool bool
  Obj map[string]JSONValue
}

//Creates a JSONValue from the interface provided. It attempts to fill the values Arr, Str, Int, Num, and Obj
//by checking against the type of the value provided.
func NewJSONValue(value interface{}) JSONValue {
  outputValue := JSONValue{Value: value}
  outputValue.Arr = outputValue.Array()
  outputValue.Str = outputValue.String()
  outputValue.Num = outputValue.Number()
  outputValue.Int = outputValue.Integer()
  outputValue.Bool = outputValue.Boolean()
  outputValue.Obj = outputValue.Object()
  return outputValue
}

//Checks if the type of the json value is an array and if appropriate, casts it into
//an array of JSONValue.
func (key JSONValue) Array() []JSONValue {
  switch typedValue := key.Value.(type) {
    case []interface{}:
      typedArray := make([]JSONValue, len(typedValue))
      for i := range typedValue {
        typedArray[i] = NewJSONValue(typedValue[i])
      }
      return typedArray
    default:
      return nil
  }
}

//Checks if the type of the json value is a string and if appropriate, casts it into a string.
func (key JSONValue) String() string {
  switch typedValue := key.Value.(type) {
    case string:
      return typedValue
    default:
      return ""
  }
}

//Checks if the type of the json value is a float64 and if appropriate, casts it into an int.
func (key JSONValue) Integer() int {
  switch typedValue := key.Value.(type) {
    case float64:
      return int(typedValue)
    default:
      return 0
  }
}

//Checks if the type of the json value is a float64 and if appropriate, casts it into a float64.
func (key JSONValue) Number() float64 {
  switch typedValue := key.Value.(type) {
    case float64:
      return typedValue
    default:
      return 0
  }
}

//Checks if the type of the json value is a bool and if appropriate, casts it into a bool.
func (key JSONValue) Boolean() bool {
  switch typedValue := key.Value.(type) {
    case bool:
      return typedValue
    default:
      return false
  }
}

//Checks if the type of the json value is an object and if appropriate, casts it into a map of JSONValue.
func (key JSONValue) Object() map[string]JSONValue {
  switch typedValue := key.Value.(type) {
    case map[string]interface{}:
      return convertMap(typedValue)
    default:
      return nil
  }
}

//Converts an abstract map of json data into a map of JSONValue.
func convertMap(from map[string]interface{}) map[string]JSONValue {
  output := map[string]JSONValue{}
  for mapKey, mapValue := range from {
    output[mapKey] = NewJSONValue(mapValue)
  }
  return output
}

//Attempts to parse the file as a json object, removing any //comments in the process.
func loadFileAsJSON(filename string) (map[string]JSONValue, error) {
  file, err := os.Open(filename)
  if err != nil {
    return map[string]JSONValue{}, err
  }

  untypedMap := map[string]interface{}{}
  dec := json.NewDecoder(NewJsonCommentStripper(file))
  if err = dec.Decode(&untypedMap); err != nil {
    return map[string]JSONValue{}, err
  }

  return convertMap(untypedMap), nil
}

//Attempts to parse the string as a json object, removing any //comments in the process.
func loadStringAsJSON(jsonstr string) (map[string]JSONValue, error) {
  untypedMap := map[string]interface{}{}
  dec := json.NewDecoder(NewJsonCommentStripper(strings.NewReader(jsonstr)))
  if err := dec.Decode(&untypedMap); err != nil {
    return map[string]JSONValue{}, err
  }

  return convertMap(untypedMap), nil
}

//Carefully copies default values into the loaded config file.
//If the key already exists in the loaded config file then the one in the default config
//is ignored unless the value in both the default config and the loaded config is an object.
//If the value is an object then the process is repeated, treating this key as a config in both
//the loaded config and default config.
func applyDefaults(config, defaults map[string]JSONValue) {
  for key, value := range defaults {
    if _, exists := config[key]; !exists {
      config[key] = value
    } else {
      switch value.Value.(type) {
        case map[string]interface{}:
          switch config[key].Value.(type) {
            case map[string]interface{}:
              applyDefaults(config[key].Obj, defaults[key].Obj)
          }
      }
    }
  }
}

//Loads the file containing a json object into an abstract map of JSONValue valueType.
//You can provide a default configuration by providing a partial example of the config
//file as a string.
func LoadAbstract(filename string, defaults string) (config map[string]JSONValue, err error) {
  config, err = loadFileAsJSON(filename)
  if err != nil {
    return
  }

  if len(defaults) > 0 {
    defaultValues, err := loadStringAsJSON(defaults)
    if err != nil {
      return map[string]JSONValue{}, err
    }
    applyDefaults(config, defaultValues)
  }
  return
}

//Loads the file containing a json object into the provided data structure. You can
//provide default values by defining them in the provided data structure before handing
//it to this func.
func Load(filename string, config interface{}) error {
  file, err := os.Open(filename)
  if err != nil {
    return err
  }

  dec := json.NewDecoder(NewJsonCommentStripper(file))
  if err = dec.Decode(config); err != nil {
    return err
  }

  return nil
}
