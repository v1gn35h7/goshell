package service

type transformationService interface {
	getEventsProto() (bool, error)
}

func (srvc service) getEventsProto() (bool, error) {
	// TODO: get proto for parsing
	return true, nil
}
