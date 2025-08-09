import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { PlayIcon, ClockIcon, CheckCircleIcon } from '../components/Icons';
import api from '../services/api';
import { UserCourse } from '../types';

const MyCourses: React.FC = () => {
  const [enrolledCourses, setEnrolledCourses] = useState<UserCourse[]>([]);
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState<'enrolled' | 'completed'>('enrolled');

  useEffect(() => {
    fetchMyCourses();
  }, []);

  const fetchMyCourses = async () => {
    try {
      const response = await api.get('/user/courses');
      setEnrolledCourses(response.data.data || []);
    } catch (error) {
      console.error('Error fetching my courses:', error);
      // For demo purposes, use mock data
      setEnrolledCourses([
        {
          id: 1,
          user_id: 1,
          course_id: 1,
          enrolled_at: new Date().toISOString(),
          progress_percentage: 65,
          completed_lessons: 8,
          total_lessons: 12,
          last_accessed_at: new Date().toISOString(),
          course: {
            id: 1,
            title: 'Introduction to React',
            description: 'Learn the fundamentals of React development',
            short_description: 'Learn React basics with hands-on projects',
            thumbnail: 'https://via.placeholder.com/300x200?text=React+Course',
            level: 'beginner',
            category: 'Web Development',
            tags: 'react,javascript,frontend',
            price: 49.99,
            is_published: true,
            creator_id: 1,
            created_at: new Date().toISOString(),
            updated_at: new Date().toISOString(),
            lessons_count: 12,
            total_duration: 480,
            enrollment_count: 145
          }
        },
        {
          id: 2,
          user_id: 1,
          course_id: 3,
          enrolled_at: new Date().toISOString(),
          progress_percentage: 100,
          completed_lessons: 16,
          total_lessons: 16,
          last_accessed_at: new Date().toISOString(),
          course: {
            id: 3,
            title: 'Go Web Development',
            description: 'Build modern web applications with Go',
            short_description: 'Complete guide to Go web development',
            thumbnail: 'https://via.placeholder.com/300x200?text=Go+Course',
            level: 'intermediate',
            category: 'Backend Development',
            tags: 'go,backend,web',
            price: 59.99,
            is_published: true,
            creator_id: 1,
            created_at: new Date().toISOString(),
            updated_at: new Date().toISOString(),
            lessons_count: 16,
            total_duration: 600,
            enrollment_count: 67
          }
        }
      ]);
    } finally {
      setLoading(false);
    }
  };

  const enrolledFilteredCourses = enrolledCourses.filter(uc => uc.progress_percentage < 100);
  const completedFilteredCourses = enrolledCourses.filter(uc => uc.progress_percentage === 100);

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric'
    });
  };

  if (loading) {
    return (
      <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div className="flex items-center justify-center h-64">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
        </div>
      </div>
    );
  }

  return (
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div className="px-4 py-6 sm:px-0">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">My Courses</h1>
          <p className="text-gray-600">Continue your learning journey</p>
        </div>

        {/* Tabs */}
        <div className="border-b border-gray-200 mb-8">
          <nav className="-mb-px flex space-x-8">
            <button
              onClick={() => setActiveTab('enrolled')}
              className={`py-2 px-1 border-b-2 font-medium text-sm ${
                activeTab === 'enrolled'
                  ? 'border-blue-500 text-blue-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
              }`}
            >
              In Progress ({enrolledFilteredCourses.length})
            </button>
            <button
              onClick={() => setActiveTab('completed')}
              className={`py-2 px-1 border-b-2 font-medium text-sm ${
                activeTab === 'completed'
                  ? 'border-blue-500 text-blue-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
              }`}
            >
              Completed ({completedFilteredCourses.length})
            </button>
          </nav>
        </div>

        {/* Course List */}
        <div className="space-y-6">
          {(activeTab === 'enrolled' ? enrolledFilteredCourses : completedFilteredCourses).map((userCourse) => (
            <div key={userCourse.id} className="bg-white rounded-lg shadow hover:shadow-md transition-shadow duration-200">
              <div className="p-6">
                <div className="flex items-start space-x-4">
                  {/* Course Thumbnail */}
                  <div className="flex-shrink-0">
                    <img
                      src={userCourse.course?.thumbnail}
                      alt={userCourse.course?.title}
                      className="w-32 h-20 object-cover rounded-lg"
                    />
                  </div>

                  {/* Course Details */}
                  <div className="flex-1 min-w-0">
                    <div className="flex items-start justify-between">
                      <div className="flex-1">
                        <h3 className="text-lg font-semibold text-gray-900 mb-2">
                          {userCourse.course?.title}
                        </h3>
                        <p className="text-gray-600 text-sm mb-3">
                          {userCourse.course?.short_description}
                        </p>

                        {/* Progress Bar */}
                        <div className="mb-3">
                          <div className="flex items-center justify-between text-sm text-gray-600 mb-1">
                            <span>Progress: {userCourse.progress_percentage}%</span>
                            <span>{userCourse.completed_lessons} of {userCourse.total_lessons} lessons</span>
                          </div>
                          <div className="w-full bg-gray-200 rounded-full h-2">
                            <div
                              className="bg-blue-600 h-2 rounded-full transition-all duration-300"
                              style={{ width: `${userCourse.progress_percentage}%` }}
                            ></div>
                          </div>
                        </div>

                        {/* Course Stats */}
                        <div className="flex items-center space-x-4 text-sm text-gray-500">
                          <div className="flex items-center">
                            <ClockIcon className="h-4 w-4 mr-1" />
                            <span>Last accessed: {formatDate(userCourse.last_accessed_at)}</span>
                          </div>
                          <div className="flex items-center">
                            <span className={`px-2 py-1 rounded-full text-xs ${
                              userCourse.course?.level === 'beginner' ? 'bg-green-100 text-green-800' :
                              userCourse.course?.level === 'intermediate' ? 'bg-yellow-100 text-yellow-800' :
                              'bg-red-100 text-red-800'
                            }`}>
                              {userCourse.course?.level}
                            </span>
                          </div>
                        </div>
                      </div>

                      {/* Action Buttons */}
                      <div className="flex flex-col space-y-2 ml-4">
                        {userCourse.progress_percentage === 100 ? (
                          <div className="flex items-center text-green-600">
                            <CheckCircleIcon className="h-5 w-5 mr-2" />
                            <span className="text-sm font-medium">Completed</span>
                          </div>
                        ) : (
                          <Link
                            to={`/courses/${userCourse.course?.id}`}
                            className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                          >
                            <PlayIcon className="h-4 w-4 mr-2" />
                            Continue
                          </Link>
                        )}
                        <Link
                          to={`/courses/${userCourse.course?.id}`}
                          className="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                        >
                          View Details
                        </Link>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          ))}
        </div>

        {/* Empty State */}
        {(activeTab === 'enrolled' ? enrolledFilteredCourses : completedFilteredCourses).length === 0 && (
          <div className="text-center py-12">
            <div className="text-gray-400 mb-4">
              {activeTab === 'enrolled' ? (
                <PlayIcon className="mx-auto h-12 w-12" />
              ) : (
                <CheckCircleIcon className="mx-auto h-12 w-12" />
              )}
            </div>
            <h3 className="text-lg font-medium text-gray-900 mb-2">
              {activeTab === 'enrolled' ? 'No courses in progress' : 'No completed courses'}
            </h3>
            <p className="text-gray-500 mb-4">
              {activeTab === 'enrolled' 
                ? 'Start learning by enrolling in a course' 
                : 'Complete a course to see it here'
              }
            </p>
            <Link
              to="/courses"
              className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
              Browse Courses
            </Link>
          </div>
        )}
      </div>
    </div>
  );
};

export default MyCourses;
