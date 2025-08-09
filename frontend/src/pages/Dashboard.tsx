import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { BookOpenIcon, PlayIcon, PlusIcon, UserGroupIcon } from '../components/Icons';
import authService from '../services/authService';

interface UserStats {
  enrolledCourses: number;
  completedCourses: number;
  totalWatchTime: number;
  createdCourses: number;
}

const Dashboard: React.FC = () => {
  const [user] = useState(authService.getCurrentUser());
  const [stats, setStats] = useState<UserStats>({
    enrolledCourses: 0,
    completedCourses: 0,
    totalWatchTime: 0,
    createdCourses: 0
  });

  useEffect(() => {
    // In a real app, you'd fetch these stats from the API
    setStats({
      enrolledCourses: 3,
      completedCourses: 1,
      totalWatchTime: 245, // minutes
      createdCourses: 0
    });
  }, []);

  const quickActions = [
    {
      title: 'Browse Courses',
      description: 'Discover new courses to learn',
      icon: BookOpenIcon,
      path: '/courses',
      color: 'bg-blue-500 hover:bg-blue-600',
      textColor: 'text-blue-600'
    },
    {
      title: 'My Courses',
      description: 'Continue your learning journey',
      icon: PlayIcon,
      path: '/my-courses',
      color: 'bg-green-500 hover:bg-green-600',
      textColor: 'text-green-600'
    },
    {
      title: 'Create Course',
      description: 'Share your knowledge with others',
      icon: PlusIcon,
      path: '/create-course',
      color: 'bg-purple-500 hover:bg-purple-600',
      textColor: 'text-purple-600'
    }
  ];

  const formatWatchTime = (minutes: number) => {
    const hours = Math.floor(minutes / 60);
    const mins = minutes % 60;
    return hours > 0 ? `${hours}h ${mins}m` : `${mins}m`;
  };

  return (
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div className="px-4 py-6 sm:px-0">
        {/* Welcome Header */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">
            Welcome back, {user?.name || 'Learner'}! 👋
          </h1>
          <p className="text-gray-600">Ready to continue your learning journey?</p>
        </div>

        {/* Stats Cards */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
          <div className="bg-white rounded-lg shadow p-6">
            <div className="flex items-center">
              <div className="flex-shrink-0">
                <BookOpenIcon className="h-8 w-8 text-blue-500" />
              </div>
              <div className="ml-4">
                <p className="text-sm font-medium text-gray-500">Enrolled Courses</p>
                <p className="text-2xl font-semibold text-gray-900">{stats.enrolledCourses}</p>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow p-6">
            <div className="flex items-center">
              <div className="flex-shrink-0">
                <PlayIcon className="h-8 w-8 text-green-500" />
              </div>
              <div className="ml-4">
                <p className="text-sm font-medium text-gray-500">Completed</p>
                <p className="text-2xl font-semibold text-gray-900">{stats.completedCourses}</p>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow p-6">
            <div className="flex items-center">
              <div className="flex-shrink-0">
                <UserGroupIcon className="h-8 w-8 text-purple-500" />
              </div>
              <div className="ml-4">
                <p className="text-sm font-medium text-gray-500">Watch Time</p>
                <p className="text-2xl font-semibold text-gray-900">{formatWatchTime(stats.totalWatchTime)}</p>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow p-6">
            <div className="flex items-center">
              <div className="flex-shrink-0">
                <PlusIcon className="h-8 w-8 text-orange-500" />
              </div>
              <div className="ml-4">
                <p className="text-sm font-medium text-gray-500">Created Courses</p>
                <p className="text-2xl font-semibold text-gray-900">{stats.createdCourses}</p>
              </div>
            </div>
          </div>
        </div>

        {/* Quick Actions */}
        <div className="mb-8">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">Quick Actions</h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            {quickActions.map((action, index) => (
              <Link
                key={index}
                to={action.path}
                className="group bg-white rounded-lg shadow hover:shadow-lg transition-shadow duration-200 p-6 border border-gray-200 hover:border-gray-300"
              >
                <div className="flex items-start">
                  <div className={`flex-shrink-0 p-3 rounded-lg bg-gray-50 group-hover:bg-gray-100`}>
                    <action.icon className={`h-6 w-6 ${action.textColor}`} />
                  </div>
                  <div className="ml-4">
                    <h3 className="text-lg font-semibold text-gray-900 group-hover:text-gray-700">
                      {action.title}
                    </h3>
                    <p className="text-gray-600 mt-1">{action.description}</p>
                    <div className="mt-3">
                      <span className={`text-sm font-medium ${action.textColor} group-hover:underline`}>
                        Get started →
                      </span>
                    </div>
                  </div>
                </div>
              </Link>
            ))}
          </div>
        </div>

        {/* Recent Activity Section */}
        <div className="bg-white rounded-lg shadow">
          <div className="px-6 py-4 border-b border-gray-200">
            <h2 className="text-lg font-semibold text-gray-900">Recent Activity</h2>
          </div>
          <div className="p-6">
            <div className="text-center py-8">
              <BookOpenIcon className="h-12 w-12 text-gray-400 mx-auto mb-4" />
              <p className="text-gray-500 mb-4">No recent activity yet</p>
              <Link
                to="/courses"
                className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              >
                Start Learning
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
