package holidays

import "time"

// Book is represents cal
type Book struct {
	events []event
	index  map[string]event
}

func newBookfromEvents(events []event) (Book, error) {
	index := make(map[string]event)
	for _, e := range events {
		days, err := e.days()
		if err != nil {
			return Book{}, err
		}

		for _, day := range days {
			index[_key(day)] = e
		}
	}
	return Book{events, index}, nil
}

func (b Book) isHoliday(d time.Time) bool {
	e := b.findEvent(d)

	if e == nil {
		return isWeekend(d)
	}

	return e.isHoliday()
}

func (b Book) isWorkingday(d time.Time) bool {
	e := b.findEvent(d)

	if e == nil {
		return !isWeekend(d)
	}

	return e.isWorkingday()
}

func (b Book) findEvent(d time.Time) *event {
	e, ok := b.index[_key(d)]
	if !ok {
		return nil
	}

	return &e
}

func isWeekend(d time.Time) bool {
	day := d.Weekday()
	return day == 6 || day == 0
}

func _key(d time.Time) string {
	return d.Format("2006-01-02")
}
