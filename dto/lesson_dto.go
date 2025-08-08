package dto

// Lesson DTOs
type CreateLessonRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=200"`
	Description string `json:"description" validate:"max=1000"`
	VideoURL    string `json:"video_url" validate:"url"`
	VideoID     string `json:"video_id"`
	Script      string `json:"script"`
	Duration    int    `json:"duration" validate:"min=0"`
	Sequence    int    `json:"sequence" validate:"required,min=1"`
	IsPublished bool   `json:"is_published"`
	IsFree      bool   `json:"is_free"`
}

type UpdateLessonRequest struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=3,max=200"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=1000"`
	VideoURL    *string `json:"video_url,omitempty" validate:"omitempty,url"`
	VideoID     *string `json:"video_id,omitempty"`
	Script      *string `json:"script,omitempty"`
	Duration    *int    `json:"duration,omitempty" validate:"omitempty,min=0"`
	Sequence    *int    `json:"sequence,omitempty" validate:"omitempty,min=1"`
	IsPublished *bool   `json:"is_published,omitempty"`
	IsFree      *bool   `json:"is_free,omitempty"`
}

type LessonResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoURL    string `json:"video_url"`
	VideoID     string `json:"video_id"`
	Script      string `json:"script,omitempty"` // May be hidden for non-enrolled users
	Duration    int    `json:"duration"`
	CourseID    uint   `json:"course_id"`
	Sequence    int    `json:"sequence"`
	IsPublished bool   `json:"is_published"`
	IsFree      bool   `json:"is_free"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	IsCompleted bool   `json:"is_completed,omitempty"` // For enrolled users
}

type LessonListResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoID     string `json:"video_id"`
	Duration    int    `json:"duration"`
	Sequence    int    `json:"sequence"`
	IsFree      bool   `json:"is_free"`
	IsCompleted bool   `json:"is_completed,omitempty"`
}

// Progress tracking DTOs
type UpdateProgressRequest struct {
	LessonID    uint `json:"lesson_id" validate:"required"`
	WatchTime   int  `json:"watch_time" validate:"min=0"`
	IsCompleted bool `json:"is_completed"`
}

type EnrollCourseRequest struct {
	CourseID uint `json:"course_id" validate:"required"`
}

// Search and filter DTOs
type CourseFilterRequest struct {
	Category    string   `query:"category"`
	Level       string   `query:"level" validate:"omitempty,oneof=beginner intermediate advanced"`
	MinPrice    *float64 `query:"min_price" validate:"omitempty,min=0"`
	MaxPrice    *float64 `query:"max_price" validate:"omitempty,min=0"`
	Tags        string   `query:"tags"` // comma-separated
	Search      string   `query:"search"`
	IsPublished *bool    `query:"is_published"`
	Page        int      `query:"page" validate:"min=1"`
	Limit       int      `query:"limit" validate:"min=1,max=100"`
	SortBy      string   `query:"sort_by" validate:"omitempty,oneof=title created_at price duration enrolled_count"`
	SortOrder   string   `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}

// Response wrapper for paginated results
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalPages int         `json:"total_pages"`
}

// API Response wrapper
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
