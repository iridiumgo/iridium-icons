package icon

func DeepCopyMap(m map[string]any) map[string]any {
	if m == nil {
		return nil
	}

	copy := make(map[string]any, len(m))
	for k, v := range m {
		copy[k] = deepCopyValue(v)
	}
	return copy
}

func deepCopyValue(v any) any {
	switch x := v.(type) {
	case map[string]any:
		return DeepCopyMap(x)
	case []any:
		sliceCopy := make([]any, len(x))
		for i, e := range x {
			sliceCopy[i] = deepCopyValue(e)
		}
		return sliceCopy
	default:
		return x // primitives, structs, etc. are copied by value
	}
}
