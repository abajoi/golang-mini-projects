package main

import "errors"

func main() {

	resp, err := c.SetTimer(cont.B(), &set)
	if err != nil || resp.Result != OK {
		if resp.Result != OK {
			err = errors.New(resp.Cause)
		}
	}

	resp, err := c.SetTimer(cont.B(), &set)
	if err == nil {
		if resp != nil {
			if resp.Result == OK {
				return IDs, nil
			} else {
				err1 = errors.New(resp.Cause)
			}
		} else {
			err1 = errors.New("ERROR")
		}
	} else {
		err1 = err
	}

}
