package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"strings"
)

// StringArray is a custom type for handling string arrays
type StringArray []string

// Value implements the driver.Valuer interface
func (a StringArray) Value() (driver.Value, error) {
	if a == nil {
		return "{}", nil
	}
	return "{" + strings.Join(a, ",") + "}", nil
}

// Scan implements the sql.Scanner interface
func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = StringArray{}
		return nil
	}

	if str, ok := value.(string); ok {
		str = strings.Trim(str, "{}")
		*a = strings.Split(str, ",")
		return nil
	}

	return errors.New("failed to scan StringArray")
}

type Experience struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(255);not null"`
	Company     string `json:"company" gorm:"type:varchar(255);not null"`
	Period      string `json:"period" gorm:"type:varchar(100);not null"`
	Description string `json:"description" gorm:"type:text;not null"`
}

// SetDescription converts string array to JSON string
func (e *Experience) SetDescription(desc []string) error {
	jsonStr, err := json.Marshal(desc)
	if err != nil {
		return err
	}
	e.Description = string(jsonStr)
	return nil
}

// GetDescription converts JSON string to string array
func (e *Experience) GetDescription() ([]string, error) {
	var desc []string
	err := json.Unmarshal([]byte(e.Description), &desc)
	return desc, err
}

func (Experience) TableName() string {
	return "experiences"
}

type Project struct {
	gorm.Model
	Title        string `json:"title" gorm:"type:varchar(255);not null"`
	Description  string `json:"description" gorm:"type:text;not null"`
	Technologies string `json:"technologies" gorm:"type:text;not null"`
	Link         string `json:"link" gorm:"type:varchar(255)"`
}

// SetTechnologies converts string array to JSON string
func (p *Project) SetTechnologies(tech []string) error {
	jsonStr, err := json.Marshal(tech)
	if err != nil {
		return err
	}
	p.Technologies = string(jsonStr)
	return nil
}

// GetTechnologies converts JSON string to string array
func (p *Project) GetTechnologies() ([]string, error) {
	var tech []string
	err := json.Unmarshal([]byte(p.Technologies), &tech)
	return tech, err
}

func (Project) TableName() string {
	return "projects"
}

type SkillCategory struct {
	gorm.Model
	Title  string `json:"title" gorm:"type:varchar(255);not null"`
	Skills string `json:"skills" gorm:"type:text;not null"`
}

// SetSkills converts string array to JSON string
func (s *SkillCategory) SetSkills(skills []string) error {
	jsonStr, err := json.Marshal(skills)
	if err != nil {
		return err
	}
	s.Skills = string(jsonStr)
	return nil
}

// GetSkills converts JSON string to string array
func (s *SkillCategory) GetSkills() ([]string, error) {
	var skills []string
	err := json.Unmarshal([]byte(s.Skills), &skills)
	return skills, err
}

func (SkillCategory) TableName() string {
	return "skill_categories"
} 