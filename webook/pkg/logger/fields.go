package logger

func String(key, val string) Field {
	return Field{
		Key:   key,
		Value: val,
	}
}
func Int32(key, val string) Field {
	return Field{
		Key:   key,
		Value: val,
	}
}
func Int64(key, val string) Field {
	return Field{
		Key:   key,
		Value: val,
	}
}
func Error(val error) Field {
	return Field{
		Key:   "error",
		Value: val,
	}
}
