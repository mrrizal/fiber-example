package book

import (
	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm"
	"strconv"
)

func getIDFromURLQuery(c *fiber.Ctx, next, previous string, previousPage *bool) (int, error) {
	span, _ := apm.StartSpan(c.Context(), "getIDFromURLQuery", "utils")
	defer span.End()
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
