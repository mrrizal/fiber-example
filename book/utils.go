package book

import "strconv"

func getIDFromURLQuery(next, previous string, previousPage *bool) (int, error) {
	if previous != "" {
		*previousPage = true
		tempID, err := strconv.ParseInt(previous, 10, 32)
		if err != nil {
			return 0, err
		}
		return int(tempID), nil
	} else if next != "" {
		tempID, err := strconv.ParseInt(next, 10, 32)
		if err != nil {
			return 0, err
		}
		return int(tempID), nil
	}
	return 0, nil
}