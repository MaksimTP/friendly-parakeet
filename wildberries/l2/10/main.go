package main

type Event struct {
	Body      string
	CreatedAt string
	ExpiresAt string
}

type Calendar struct {
	Events []Event
}

func (c *Calendar) CreateEvent(body string, date string) error {
	err := validateDate(date)
	if err != nil {
		return err
	}

}

func validateDate(date string) error {
	return nil
}

func main() {
}
