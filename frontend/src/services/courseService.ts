import api from './api';
import { Course, PaginatedResponse, ApiResponse, CourseAnalytics } from '../types';

export interface CreateCourseData {
  title: string;
  description: string;
  short_description: string;
  thumbnail: string;
  level: 'beginner' | 'intermediate' | 'advanced';
  category: string;
  tags: string;
  price: number;
  is_published: boolean;
}

export interface CourseFilters {
  category?: string;
  level?: string;
  page?: number;
  limit?: number;
  search?: string;
}

class CourseService {
  async getCourses(filters: CourseFilters = {}): Promise<PaginatedResponse<Course>> {
    const params = new URLSearchParams();
    Object.entries(filters).forEach(([key, value]) => {
      if (value) params.append(key, value.toString());
    });
    
    const response = await api.get<PaginatedResponse<Course>>(`/courses?${params}`);
    return response.data;
  }

  async searchCourses(filters: CourseFilters): Promise<PaginatedResponse<Course>> {
    const params = new URLSearchParams();
    Object.entries(filters).forEach(([key, value]) => {
      if (value) params.append(key, value.toString());
    });
    
    const response = await api.get<PaginatedResponse<Course>>(`/courses/search?${params}`);
    return response.data;
  }

  async getCourse(id: number): Promise<ApiResponse<Course>> {
    const response = await api.get<ApiResponse<Course>>(`/courses/${id}`);
    return response.data;
  }

  async createCourse(courseData: CreateCourseData): Promise<ApiResponse<Course>> {
    const response = await api.post<ApiResponse<Course>>('/courses', courseData);
    return response.data;
  }

  async updateCourse(id: number, courseData: Partial<CreateCourseData>): Promise<ApiResponse<Course>> {
    const response = await api.put<ApiResponse<Course>>(`/courses/${id}`, courseData);
    return response.data;
  }

  async deleteCourse(id: number): Promise<ApiResponse<null>> {
    const response = await api.delete<ApiResponse<null>>(`/courses/${id}`);
    return response.data;
  }

  async getMyCourses(): Promise<PaginatedResponse<Course>> {
    const response = await api.get<PaginatedResponse<Course>>('/my/courses');
    return response.data;
  }

  async enrollInCourse(courseId: number): Promise<ApiResponse<null>> {
    const response = await api.post<ApiResponse<null>>(`/courses/${courseId}/enroll`);
    return response.data;
  }

  async unenrollFromCourse(courseId: number): Promise<ApiResponse<null>> {
    const response = await api.delete<ApiResponse<null>>(`/courses/${courseId}/enroll`);
    return response.data;
  }

  async getEnrolledCourses(): Promise<PaginatedResponse<Course>> {
    const response = await api.get<PaginatedResponse<Course>>('/my/enrolled-courses');
    return response.data;
  }

  async getCourseProgress(courseId: number): Promise<ApiResponse<any>> {
    const response = await api.get<ApiResponse<any>>(`/courses/${courseId}/progress`);
    return response.data;
  }

  async getCourseAnalytics(courseId: number): Promise<ApiResponse<CourseAnalytics>> {
    const response = await api.get<ApiResponse<CourseAnalytics>>(`/courses/${courseId}/analytics`);
    return response.data;
  }
}

export default new CourseService();
