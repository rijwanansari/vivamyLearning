package dto

// Course DTOs
type CreateCourseRequest struct {
	Title            string  `json:"title" validate:"required,min=3,max=200"`
	Description      string  `json:"description" validate:"max=2000"`
	ShortDescription string  `json:"short_description" validate:"max=500"`
	Thumbnail        string  `json:"thumbnail" validate:"url"`
	Level            string  `json:"level" validate:"oneof=beginner intermediate advanced"`
	Category         string  `json:"category" validate:"required,min=2,max=100"`
	Tags             string  `json:"tags"`
	Price            float64 `json:"price" validate:"min=0"`
	IsPublished      bool    `json:"is_published"`
}

type UpdateCourseRequest struct {
	Title            *string  `json:"title,omitempty" validate:"omitempty,min=3,max=200"`
	Description      *string  `json:"description,omitempty" validate:"omitempty,max=2000"`
	ShortDescription *string  `json:"short_description,omitempty" validate:"omitempty,max=500"`
	Thumbnail        *string  `json:"thumbnail,omitempty" validate:"omitempty,url"`
	Level            *string  `json:"level,omitempty" validate:"omitempty,oneof=beginner intermediate advanced"`
	Category         *string  `json:"category,omitempty" validate:"omitempty,min=2,max=100"`
	Tags             *string  `json:"tags,omitempty"`
	Price            *float64 `json:"price,omitempty" validate:"omitempty,min=0"`
	IsPublished      *bool    `json:"is_published,omitempty"`
}

type CourseResponse struct {
	ID               uint                  `json:"id"`
	Title            string                `json:"title"`
	Description      string                `json:"description"`
	ShortDescription string                `json:"short_description"`
	Thumbnail        string                `json:"thumbnail"`
	Level            string                `json:"level"`
	Category         string                `json:"category"`
	Tags             []string              `json:"tags"`
	Duration         int                   `json:"duration"`
	Price            float64               `json:"price"`
	IsPublished      bool                  `json:"is_published"`
	CreatedBy        uint                  `json:"created_by"`
	CreatedAt        string                `json:"created_at"`
	UpdatedAt        string                `json:"updated_at"`
	LessonCount      int                   `json:"lesson_count"`
	EnrolledCount    int                   `json:"enrolled_count"`
	CompletionRate   float64               `json:"completion_rate"`
	Lessons          []LessonResponse      `json:"lessons,omitempty"`
	IsEnrolled       bool                  `json:"is_enrolled,omitempty"`
	UserProgress     *UserProgressResponse `json:"user_progress,omitempty"`
}

type CourseListResponse struct {
	ID               uint     `json:"id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"short_description"`
	Thumbnail        string   `json:"thumbnail"`
	Level            string   `json:"level"`
	Category         string   `json:"category"`
	Tags             []string `json:"tags"`
	Duration         int      `json:"duration"`
	Price            float64  `json:"price"`
	LessonCount      int      `json:"lesson_count"`
	EnrolledCount    int      `json:"enrolled_count"`
	CompletionRate   float64  `json:"completion_rate"`
	IsEnrolled       bool     `json:"is_enrolled,omitempty"`
}

type UserProgressResponse struct {
	Progress     float64 `json:"progress"`
	LastLessonID uint    `json:"last_lesson_id"`
	IsCompleted  bool    `json:"is_completed"`
	EnrolledAt   string  `json:"enrolled_at"`
	CompletedAt  *string `json:"completed_at,omitempty"`
}
