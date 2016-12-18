package routes

type BaseRoute struct {
	Pattern string
}

func (this *BaseRoute) GetPattern() string {
	return this.Pattern
}

