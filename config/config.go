package config

func GetInt(key string) int {
    return globalConfig[key].IntValue
}

func GetString(key string) string {
    return globalConfig[key].StringValue
}

func GetBool(key string) bool {
    return globalConfig[key].BoolValue
}

func Set(key string, value interface{}) {
    cval := _ConfigValue{Key: key}

    switch v := value.(type) {
    case string:
        cval.StringValue = v
    case int:
        cval.IntValue = v
    case bool:
        cval.BoolValue = v
    default:
        return
    }

    globalConfig[key] = cval
}

type _ConfigValue struct {
    Key string
    IntValue int
    StringValue string
    BoolValue bool
}

var globalConfig map[string]_ConfigValue

func init() {
    globalConfig = make(map[string]_ConfigValue)
}
