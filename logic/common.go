package logic

// GetKeysFromFuelMap は map から key の配列を取り出すメソッド。key は string で value は Fuel
func GetKeysFromFuelMap(input map[string]Fuel) []string {
	item := make(map[string]interface{})
	for k, v := range input {
		item[k] = v
	}
	return GetKeysFromInterfaceMap(item)
}

// GetKeysFromStringSliceMap は（略）
func GetKeysFromStringSliceMap(input map[string][]string) []string {
	item := make(map[string]interface{})
	for k, v := range input {
		item[k] = v
	}
	return GetKeysFromInterfaceMap(item)
}

// GetKeysFromInterfaceMap は map から key の配列を取り出すメソッド。key は string で value は interface{}
func GetKeysFromInterfaceMap(input map[string]interface{}) []string {

	keys := make([]string, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	return keys
}
