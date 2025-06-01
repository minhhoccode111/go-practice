package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource

	for {
		resource, err = opener()

		if err != nil {
			if _, ok := err.(TransientError); ok {
				continue
			}

			return
		}

		break
	}

	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case FrobError:
				resource.Defrob(v.defrobTag)
				err = v
			case error:
				err = v
			default:
			}
		}
	}()

	resource.Frob(input)

	return
}
