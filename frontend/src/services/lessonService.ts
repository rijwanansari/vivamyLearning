import api from './api';
import { Lesson, ApiResponse, PaginatedResponse } from '../types';

export interface CreateLessonData {
  title: string;
  description: string;
  video_url: string;
  video_id: string;
  script: string;
  duration: number;
  sequence: number;
  is_published: boolean;
  is_free: boolean;
}

export interface LessonProgress {
  lesson_id: number;
  watch_time: number;
  is_completed: boolean;
}

class LessonService {
  async getCourseLessons(courseId: number): Promise<PaginatedResponse<Lesson>> {
    const response = await api.get<PaginatedResponse<Lesson>>(`/courses/${courseId}/lessons`);
    return response.data;
  }

  async getFreeLessons(courseId: number): Promise<PaginatedResponse<Lesson>> {
    const response = await api.get<PaginatedResponse<Lesson>>(`/courses/${courseId}/lessons/free`);
    return response.data;
  }

  async getLesson(lessonId: number): Promise<ApiResponse<Lesson>> {
    const response = await api.get<ApiResponse<Lesson>>(`/lessons/${lessonId}`);
    return response.data;
  }

  async createLesson(courseId: number, lessonData: CreateLessonData): Promise<ApiResponse<Lesson>> {
    const response = await api.post<ApiResponse<Lesson>>(`/courses/${courseId}/lessons`, lessonData);
    return response.data;
  }

  async updateLesson(lessonId: number, lessonData: Partial<CreateLessonData>): Promise<ApiResponse<Lesson>> {
    const response = await api.put<ApiResponse<Lesson>>(`/lessons/${lessonId}`, lessonData);
    return response.data;
  }

  async deleteLesson(lessonId: number): Promise<ApiResponse<null>> {
    const response = await api.delete<ApiResponse<null>>(`/lessons/${lessonId}`);
    return response.data;
  }

  async updateProgress(progress: LessonProgress): Promise<ApiResponse<null>> {
    const response = await api.post<ApiResponse<null>>('/lessons/progress', progress);
    return response.data;
  }

  async markCompleted(lessonId: number, watchTime: number): Promise<ApiResponse<null>> {
    const response = await api.post<ApiResponse<null>>(`/lessons/${lessonId}/complete`, {
      watch_time: watchTime
    });
    return response.data;
  }

  async getLessonProgress(courseId: number): Promise<ApiResponse<any[]>> {
    const response = await api.get<ApiResponse<any[]>>(`/courses/${courseId}/lessons/progress`);
    return response.data;
  }

  async reorderLessons(courseId: number, lessons: Array<{lesson_id: number, sequence: number}>): Promise<ApiResponse<null>> {
    const response = await api.put<ApiResponse<null>>(`/courses/${courseId}/lessons/reorder`, lessons);
    return response.data;
  }
}

export default new LessonService();
