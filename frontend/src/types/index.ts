export interface User {
  id: number;
  name: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface Course {
  id: number;
  title: string;
  description: string;
  short_description: string;
  thumbnail: string;
  level: 'beginner' | 'intermediate' | 'advanced';
  category: string;
  tags: string;
  price: number;
  is_published: boolean;
  creator_id: number;
  created_at: string;
  updated_at: string;
  lessons_count?: number;
  total_duration?: number;
  enrollment_count?: number;
}

export interface Lesson {
  id: number;
  course_id: number;
  title: string;
  description: string;
  video_url: string;
  video_id: string;
  script: string;
  duration: number;
  sequence: number;
  is_published: boolean;
  is_free: boolean;
  created_at: string;
  updated_at: string;
}

export interface UserCourse {
  id: number;
  user_id: number;
  course_id: number;
  enrolled_at: string;
  progress_percentage: number;
  completed_lessons: number;
  total_lessons: number;
  last_accessed_at: string;
  course?: Course;
}

export interface UserLesson {
  id: number;
  user_id: number;
  lesson_id: number;
  watch_time: number;
  is_completed: boolean;
  completed_at?: string;
  last_watched_at: string;
  lesson?: Lesson;
}

export interface AuthResponse {
  success: boolean;
  message: string;
  data: {
    user: User;
    access_token: string;
    refresh_token: string;
    expires_at: string;
  };
}

export interface ApiResponse<T> {
  success: boolean;
  message: string;
  data: T;
}

export interface PaginatedResponse<T> {
  success: boolean;
  message: string;
  data: T[];
  pagination: {
    page: number;
    limit: number;
    total: number;
    total_pages: number;
  };
}

export interface CourseAnalytics {
  total_enrollments: number;
  completion_rate: number;
  average_progress: number;
  total_watch_time: number;
  most_popular_lesson: string;
  enrollment_trend: Array<{
    date: string;
    count: number;
  }>;
}
