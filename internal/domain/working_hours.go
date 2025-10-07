package domain

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// WorkingHours представляет рабочие часы компании
type WorkingHours struct {
	Monday    DaySchedule
	Tuesday   DaySchedule
	Wednesday DaySchedule
	Thursday  DaySchedule
	Friday    DaySchedule
	Saturday  DaySchedule
	Sunday    DaySchedule
}

// DaySchedule представляет расписание на один день
type DaySchedule struct {
	IsOpen    bool
	OpenTime  *TimeString // Формат "HH:MM"
	CloseTime *TimeString // Формат "HH:MM"
}

// TimeString кастомный тип для TIME полей PostgreSQL, сериализуется как "HH:MM"
type TimeString string

// Scan implements sql.Scanner interface
func (t *TimeString) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*t = TimeString(v.Format("15:04"))
		return nil
	case []byte:
		*t = TimeString(v)
		return nil
	case string:
		*t = TimeString(v)
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into TimeString", value)
	}
}

// Value implements driver.Valuer interface
func (t TimeString) Value() (driver.Value, error) {
	if t == "" {
		return nil, nil
	}
	return string(t), nil
}
