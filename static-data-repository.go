package main

var todos = []TodoItem{
	{
		Title:   "Grocery Shopping",
		Details: "Buy milk, eggs, bread, spinach, chicken, and some snacks for the weekend.",
		Status:  false,
	}, {
		Title:   "Finish Project Report",
		Details: "Complete the final section of the report, review formatting, and send it to the team by 5 PM.",
		Status:  false,
	}, {
		Title:   "Book Doctor’s Appointment",
		Details: "Schedule a general check-up for next week, preferably Thursday morning.",
		Status:  false,
	}, {
		Title:   "Call Mom",
		Details: "Catch up and check in about her weekend plans and how she's feeling lately.",
		Status:  true,
	}, {
		Title:   "Clean Out Email Inbox",
		Details: "Sort through unread emails, unsubscribe from spam, and organize important ones into folders.",
		Status:  false,
	},
}

type staticRepository struct {
}

func (r *staticRepository) GetAll() []TodoItem {
	return todos
}

func (r *staticRepository) GetDone() []TodoItem {
	toReturn := []TodoItem{}
	for _, v := range todos {
		if v.Status {
			toReturn = append(toReturn, v)
		}
	}
	return toReturn
}

func (r *staticRepository) GetPending() []TodoItem {
	toReturn := []TodoItem{}
	for _, v := range todos {
		if !v.Status {
			toReturn = append(toReturn, v)
		}
	}
	return toReturn
}
