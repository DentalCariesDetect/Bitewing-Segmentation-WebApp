package usecases

func stringIsNil(value string) *string {
	if value != "" {
		return &value
	}
	return nil
}

func uintIsNil(value uint64) *uint64 {
	if value != 0 {
		return &value
	}
	return nil
}
