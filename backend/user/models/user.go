package models

import (
    "time"
)

type User struct {
    ID           int64
    Username     string
    DisplayName  string
    Email        string
    Password     string
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
