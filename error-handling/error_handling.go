package erratum

import (
	"errors"
)

func Use(o ResourceOpener, input string) (ret error) {
	var res Resource
	var err error

	for {
		res, err = o()
		if err != nil {
			if errors.As(err, &TransientError{}) {
				continue
			} else {
				return err
			}
		} else {
			break
		}
	}

	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(FrobError); ok {
				res.Defrob(err.defrobTag)
				ret = err
			} else {
				ret = r.(error)
			}
		}
	}()

	res.Frob(input)
	return nil
}
