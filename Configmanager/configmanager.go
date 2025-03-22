package configmanager

type configmanager struct {
	config map[string]interface{}
}

var (
	isInitialized bool = false
	instance      configmanager
)

func getstring(key string) error {

}
