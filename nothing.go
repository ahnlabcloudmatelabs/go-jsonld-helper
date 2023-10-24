package jsonld_helper

type Nothing struct {
	Error error
}

func (n Nothing) Value() any  { return nil }
func (n Nothing) Length() int { return 0 }

func (n Nothing) ReadKey(string) JsonLDReader { return n }
func (n Nothing) ReadIndex(int) JsonLDReader  { return n }

func (n Nothing) Get() any                       { return nil }
func (n Nothing) GetOrElse(defaultValue any) any { return defaultValue }
func (n Nothing) GetOrThrow(err error) (any, error) {
	if err != nil {
		return err, err
	}

	return n.Error, n.Error
}

func (n Nothing) StringOrElse(defaultValue string) string { return defaultValue }
func (n Nothing) StringOrThrow(err error) (string, error) {
	if err != nil {
		return "", err
	}

	return "", n.Error
}

func (n Nothing) BoolOrElse(defaultValue bool) bool { return defaultValue }
func (n Nothing) BoolOrThrow(err error) (bool, error) {
	if err != nil {
		return false, err
	}

	return false, n.Error
}

func (n Nothing) IntOrElse(defaultValue int) int { return defaultValue }
func (n Nothing) IntOrThrow(err error) (int, error) {
	if err != nil {
		return 0, err
	}

	return 0, n.Error
}

func (n Nothing) FloatOrElse(defaultValue float64) float64 { return defaultValue }
func (n Nothing) FloatOrThrow(err error) (float64, error) {
	if err != nil {
		return 0, err
	}

	return 0, n.Error
}
