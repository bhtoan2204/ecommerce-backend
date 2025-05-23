package persistentobject

import (
	"time"
)

type WebhookLog struct {
	BasePO
	PaymentID  *string                `gorm:"type:uuid"`
	Provider   string                 `gorm:"size:50"`
	EventType  string                 `gorm:"size:100"`
	Payload    map[string]interface{} `gorm:"type:jsonb"`
	ReceivedAt time.Time              `gorm:"autoCreateTime"`
}

func (w *WebhookLog) TableName() string {
	return "webhook_logs"
}
