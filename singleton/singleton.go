package singleton

var uniqueInstance *boiler

type boiler struct {
	IsEmpty bool
}

func (b *boiler) Fill() {
	b.IsEmpty = false
}

func GetInstance() *boiler {
	if uniqueInstance == nil {
		uniqueInstance = &boiler{true}
	}
	return uniqueInstance
}
