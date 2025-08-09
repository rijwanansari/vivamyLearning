import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { MagnifyingGlassIcon, FunnelIcon } from '../components/Icons';
import api from '../services/api';
import { Course } from '../types';

const Courses: React.FC = () => {
  const [courses, setCourses] = useState<Course[]>([]);
  const [loading, setLoading] = useState(true);
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedLevel, setSelectedLevel] = useState<string>('all');
  const [selectedCategory, setSelectedCategory] = useState<string>('all');

  useEffect(() => {
    fetchCourses();
  }, []);

  const fetchCourses = async () => {
    try {
      const response = await api.get('/courses');
      setCourses(response.data.data || []);
    } catch (error) {
      console.error('Error fetching courses:', error);
      // For demo purposes, use mock data
      setCourses([
        {
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
        },
        {
          id: 2,
          title: 'Advanced TypeScript',
          description: 'Master TypeScript for large-scale applications',
          short_description: 'Advanced TypeScript patterns and best practices',
          thumbnail: 'https://via.placeholder.com/300x200?text=TypeScript+Course',
          level: 'advanced',
          category: 'Programming',
          tags: 'typescript,javascript,backend',
          price: 79.99,
          is_published: true,
          creator_id: 1,
          created_at: new Date().toISOString(),
          updated_at: new Date().toISOString(),
          lessons_count: 20,
          total_duration: 720,
          enrollment_count: 89
        },
        {
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
      ]);
    } finally {
      setLoading(false);
    }
  };

  const filteredCourses = courses.filter(course => {
    const matchesSearch = course.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
                         course.description.toLowerCase().includes(searchTerm.toLowerCase());
    const matchesLevel = selectedLevel === 'all' || course.level === selectedLevel;
    const matchesCategory = selectedCategory === 'all' || course.category === selectedCategory;
    
    return matchesSearch && matchesLevel && matchesCategory;
  });

  const formatDuration = (minutes: number) => {
    const hours = Math.floor(minutes / 60);
    const mins = minutes % 60;
    return hours > 0 ? `${hours}h ${mins}m` : `${mins}m`;
  };

  const getLevelBadgeColor = (level: string) => {
    switch (level) {
      case 'beginner': return 'bg-green-100 text-green-800';
      case 'intermediate': return 'bg-yellow-100 text-yellow-800';
      case 'advanced': return 'bg-red-100 text-red-800';
      default: return 'bg-gray-100 text-gray-800';
    }
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
          <h1 className="text-3xl font-bold text-gray-900 mb-2">Browse Courses</h1>
          <p className="text-gray-600">Discover new skills and advance your career</p>
        </div>

        {/* Search and Filters */}
        <div className="bg-white rounded-lg shadow p-6 mb-8">
          <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
            {/* Search */}
            <div className="relative">
              <MagnifyingGlassIcon className="h-5 w-5 absolute left-3 top-3 text-gray-400" />
              <input
                type="text"
                placeholder="Search courses..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className="pl-10 w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>

            {/* Level Filter */}
            <select
              value={selectedLevel}
              onChange={(e) => setSelectedLevel(e.target.value)}
              className="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="all">All Levels</option>
              <option value="beginner">Beginner</option>
              <option value="intermediate">Intermediate</option>
              <option value="advanced">Advanced</option>
            </select>

            {/* Category Filter */}
            <select
              value={selectedCategory}
              onChange={(e) => setSelectedCategory(e.target.value)}
              className="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="all">All Categories</option>
              <option value="Web Development">Web Development</option>
              <option value="Programming">Programming</option>
              <option value="Backend Development">Backend Development</option>
              <option value="Data Science">Data Science</option>
            </select>

            {/* Filter Button */}
            <button className="flex items-center justify-center px-4 py-2 bg-gray-100 text-gray-700 rounded-md hover:bg-gray-200 transition-colors">
              <FunnelIcon className="h-5 w-5 mr-2" />
              Filters
            </button>
          </div>
        </div>

        {/* Results Count */}
        <div className="mb-6">
          <p className="text-gray-600">
            Showing {filteredCourses.length} of {courses.length} courses
          </p>
        </div>

        {/* Course Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {filteredCourses.map((course) => (
            <Link
              key={course.id}
              to={`/courses/${course.id}`}
              className="bg-white rounded-lg shadow hover:shadow-lg transition-shadow duration-200 overflow-hidden group"
            >
              {/* Course Image */}
              <div className="aspect-video bg-gray-200 relative overflow-hidden">
                <img
                  src={course.thumbnail}
                  alt={course.title}
                  className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-200"
                />
                <div className="absolute top-4 left-4">
                  <span className={`px-2 py-1 text-xs font-medium rounded-full ${getLevelBadgeColor(course.level)}`}>
                    {course.level}
                  </span>
                </div>
                <div className="absolute top-4 right-4">
                  <span className="px-2 py-1 text-xs font-medium bg-black bg-opacity-70 text-white rounded">
                    ${course.price}
                  </span>
                </div>
              </div>

              {/* Course Content */}
              <div className="p-6">
                <div className="mb-2">
                  <h3 className="text-lg font-semibold text-gray-900 group-hover:text-blue-600 transition-colors line-clamp-2">
                    {course.title}
                  </h3>
                </div>
                
                <p className="text-gray-600 text-sm mb-4 line-clamp-2">
                  {course.short_description}
                </p>

                <div className="flex items-center justify-between text-sm text-gray-500 mb-4">
                  <span>{course.lessons_count} lessons</span>
                  <span>{formatDuration(course.total_duration || 0)}</span>
                  <span>{course.enrollment_count} students</span>
                </div>

                <div className="flex items-center justify-between">
                  <span className="text-xs bg-gray-100 text-gray-600 px-2 py-1 rounded">
                    {course.category}
                  </span>
                  <span className="text-blue-600 text-sm font-medium group-hover:underline">
                    View Course →
                  </span>
                </div>
              </div>
            </Link>
          ))}
        </div>

        {/* Empty State */}
        {filteredCourses.length === 0 && (
          <div className="text-center py-12">
            <div className="text-gray-400 mb-4">
              <svg className="mx-auto h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1} d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
              </svg>
            </div>
            <h3 className="text-lg font-medium text-gray-900 mb-2">No courses found</h3>
            <p className="text-gray-500">Try adjusting your search criteria or browse all courses.</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default Courses;
