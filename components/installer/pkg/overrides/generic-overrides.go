package overrides

import (
	"strings"

	"github.com/ghodss/yaml"
)

//Map is a map of overrides. Values in the map can be nested maps (of the same type) or strings
type Map map[string]interface{}

//ToMap converts yaml to Map. Supports only map-like yamls (no lists!)
func ToMap(value string) (Map, error) {
	target := Map{}

	if value == "" {
		//Otherwise, nil Map is returned by yaml.Unmarshal
		return target, nil
	}

	err := yaml.Unmarshal([]byte(value), &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

//ToYaml converts Map to YAML
func ToYaml(oMap Map) (string, error) {
	if len(oMap) == 0 {
		return "", nil
	}

	res, err := yaml.Marshal(oMap)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

//FlattenMap flattens an Map into a map of aggregated keys and value (to entries like, for example: "istio.ingress.service.gateway: xyz")
func FlattenMap(oMap Map) map[string]string {
	res := map[string]string{}
	flattenMap(oMap, "", res)
	return res
}

//UnflattenMap converts external "flat" overrides into Map. Opposite of FlattenMap function.
func UnflattenMap(sourceMap map[string]string) Map {
	mergedMap := Map{}
	if len(sourceMap) == 0 {
		return mergedMap
	}

	for key, value := range sourceMap {
		keys := strings.Split(key, ".")
		mergeIntoMap(keys, value, mergedMap)
	}

	return mergedMap
}

//MergeMaps merges all values from overridesMap map into baseMap, adding and/or overwriting final keys (string values) if both maps contain such entries.
//baseMap WILL be modified during merge.
//overridesMap won't be modified by future merges, since a deep-copy of it's nested maps are used for merging such nested maps.
func MergeMaps(baseMap, overridesMap Map) {

	//Helper function to deep-copy nested maps
	putValueToMap := func(baseMap map[string]interface{}, key string, overrideVal interface{}) {
		overrideValMap, overrideIsMap := overrideVal.(map[string]interface{})
		if overrideIsMap {
			baseMap[key] = deepCopyMap(overrideValMap)
		} else {
			baseMap[key] = overrideVal
		}
	}

	for key, overrideVal := range overridesMap {
		//Can be nil
		baseVal := baseMap[key]

		baseMapVal, baseIsMap := baseVal.(map[string]interface{})
		ovrrMapVal, newIsMap := overrideVal.(map[string]interface{})

		if baseIsMap && newIsMap {
			//Two maps case! Reccursion happens here!
			MergeMaps(baseMapVal, ovrrMapVal)
		} else {
			//All other cases, even "pathological" one, when baseMap[key] is a map and overrideVal is a string.
			putValueToMap(baseMap, key, overrideVal)
		}
	}

}

//Recursively copies the map. Used to ensure immutability of input maps when merging.
func deepCopyMap(src map[string]interface{}) map[string]interface{} {
	dst := map[string]interface{}{}

	//Helper recursive function
	var cp func(src map[string]interface{}, dst map[string]interface{})

	cp = func(src map[string]interface{}, dst map[string]interface{}) {
		for key, value := range src {

			nestedMap, isMap := value.(map[string]interface{})
			if isMap {
				//Nested map!
				nestedCopy := map[string]interface{}{}
				cp(nestedMap, nestedCopy)
				dst[key] = nestedCopy
			} else {
				dst[key] = value
			}
		}
	}
	cp(src, dst)

	return dst
}

// Flattens given Map. The keys in result map will contain all intermediate keys joined with dots, e.g.: "istio.ingress.service.gateway: xyz"
func flattenMap(oMap Map, keys string, result map[string]string) {

	var prefix string

	if len(keys) == 0 {
		prefix = ""
	} else {
		prefix = keys + "."
	}

	for key, value := range oMap {

		aString, isString := value.(string)
		if isString {
			result[prefix+key] = aString
		} else {
			//Nested map!
			nestedMap := value.(map[string]interface{})
			flattenMap(nestedMap, prefix+key, result)
		}
	}
}

//Merges value into given map, introducing intermediate "nested" maps for every intermediate key.
func mergeIntoMap(keys []string, value string, dstMap Map) {
	currentKey := keys[0]
	//Last key points directly to string value
	if len(keys) == 1 {
		dstMap[currentKey] = value
		return
	}

	//All keys but the last one should point to a nested map
	nestedMap, isMap := dstMap[currentKey].(Map)

	if !isMap {
		nestedMap = Map{}
		dstMap[currentKey] = nestedMap
	}

	mergeIntoMap(keys[1:], value, nestedMap)
}
