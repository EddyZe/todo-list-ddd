package task

const (
	StatusPending    Status = "PENDING"     // Задача создана, но не назначена
	StatusInProgress Status = "IN_PROGRESS" // Задача в работе
	StatusReview     Status = "REVIEW"      // Задача на проверке/код-ревью
	StatusDone       Status = "DONE"        // Задача завершена
	StatusCancelled  Status = "CANCELLED"   // Задача отменена
)
