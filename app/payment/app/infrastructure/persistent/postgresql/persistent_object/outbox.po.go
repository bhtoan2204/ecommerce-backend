package persistentobject

import "time"

type Outbox struct {
	BasePO
	EventType   string `gorm:"size:50;not null"`
	EventData   string `gorm:"type:text;not null"`
	Processed   bool   `gorm:"default:false;not null"`
	RetryCount  int    `gorm:"default:0;not null"`
	ProcessedAt *time.Time
}

func (o *Outbox) TableName() string {
	return "outboxs"
}
