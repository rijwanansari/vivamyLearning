import React from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import { User, BookOpen, PlusCircle, LogOut, Home } from 'lucide-react';
import authService from '../services/authService';

const Navbar: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const user = authService.getCurrentUser();

  const handleLogout = () => {
    authService.logout();
    navigate('/login');
  };

  const isActive = (path: string) => {
    return location.pathname === path ? 'bg-primary-700 text-white' : 'text-gray-300 hover:bg-primary-600 hover:text-white';
  };

  return (
    <nav className="bg-primary-500 shadow-lg fixed top-0 left-0 right-0 z-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center">
            <Link to="/dashboard" className="flex-shrink-0 flex items-center">
              <BookOpen className="h-8 w-8 text-white mr-2" />
              <span className="text-white text-xl font-bold">VivaLearning</span>
            </Link>
            
            <div className="hidden md:block ml-10">
              <div className="flex items-baseline space-x-4">
                <Link
                  to="/dashboard"
                  className={`px-3 py-2 rounded-md text-sm font-medium transition-colors ${isActive('/dashboard')}`}
                >
                  <Home className="inline-block w-4 h-4 mr-1" />
                  Dashboard
                </Link>
                
                <Link
                  to="/courses"
                  className={`px-3 py-2 rounded-md text-sm font-medium transition-colors ${isActive('/courses')}`}
                >
                  <BookOpen className="inline-block w-4 h-4 mr-1" />
                  Courses
                </Link>
                
                <Link
                  to="/my-courses"
                  className={`px-3 py-2 rounded-md text-sm font-medium transition-colors ${isActive('/my-courses')}`}
                >
                  <User className="inline-block w-4 h-4 mr-1" />
                  My Courses
                </Link>
                
                <Link
                  to="/create-course"
                  className={`px-3 py-2 rounded-md text-sm font-medium transition-colors ${isActive('/create-course')}`}
                >
                  <PlusCircle className="inline-block w-4 h-4 mr-1" />
                  Create Course
                </Link>
              </div>
            </div>
          </div>

          <div className="flex items-center">
            <div className="flex items-center space-x-4">
              <span className="text-white text-sm">
                Welcome, {user?.name}
              </span>
              
              <button
                onClick={handleLogout}
                className="bg-primary-600 text-white px-3 py-2 rounded-md text-sm font-medium hover:bg-primary-700 transition-colors flex items-center"
              >
                <LogOut className="w-4 h-4 mr-1" />
                Logout
              </button>
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
