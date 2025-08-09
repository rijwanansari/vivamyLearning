# VivaLearning Frontend

A modern React TypeScript frontend for the VivaLearning online learning platform.

## 🚀 Features

- **User Authentication** - Login/Register with JWT tokens
- **Course Management** - Create, edit, and manage courses
- **Lesson Management** - Add lessons with video content and progress tracking
- **Progress Tracking** - Track learning progress and completion
- **Responsive Design** - Mobile-friendly interface with Tailwind CSS
- **Modern UI** - Clean, professional design with Lucide icons

## 🛠️ Tech Stack

- **React 18** - Modern React with hooks
- **TypeScript** - Type-safe development
- **Vite** - Fast build tool and dev server
- **Tailwind CSS** - Utility-first CSS framework
- **React Router** - Client-side routing
- **React Query** - Server state management
- **React Hook Form** - Form handling and validation
- **Axios** - HTTP client for API calls
- **Lucide React** - Beautiful icons
- **js-cookie** - Cookie management

## 📁 Project Structure

```
frontend/
├── src/
│   ├── components/           # Reusable UI components
│   │   ├── Navbar.tsx       # Navigation bar
│   │   └── ProtectedRoute.tsx # Route protection
│   ├── pages/               # Page components
│   │   ├── Login.tsx        # Login page
│   │   ├── Register.tsx     # Registration page
│   │   ├── Dashboard.tsx    # Main dashboard
│   │   ├── Courses.tsx      # Course listing
│   │   ├── CourseDetail.tsx # Course details
│   │   ├── LessonView.tsx   # Lesson viewer
│   │   ├── MyCourses.tsx    # User's courses
│   │   └── CreateCourse.tsx # Course creation
│   ├── services/            # API services
│   │   ├── api.ts          # Axios configuration
│   │   ├── authService.ts  # Authentication API
│   │   ├── courseService.ts # Course API
│   │   └── lessonService.ts # Lesson API
│   ├── types/              # TypeScript types
│   │   └── index.ts        # Type definitions
│   ├── App.tsx             # Main app component
│   ├── main.tsx            # Entry point
│   └── index.css           # Global styles
├── package.json            # Dependencies
├── vite.config.ts          # Vite configuration
├── tailwind.config.js      # Tailwind configuration
└── tsconfig.json           # TypeScript configuration
```

## 🚦 Getting Started

### Prerequisites

- **Node.js** (v18 or higher)
- **npm** or **yarn**
- **VivaLearning Backend** running on `http://localhost:8080`

### Installation

1. **Navigate to frontend directory:**
   ```bash
   cd frontend
   ```

2. **Install dependencies:**
   ```bash
   npm install
   ```

3. **Start development server:**
   ```bash
   npm run dev
   ```

4. **Open browser:**
   Navigate to `http://localhost:3000`

### Build for Production

```bash
npm run build
```

## 🔧 Configuration

### Environment Setup

The frontend is configured to proxy API requests to the backend:

```typescript
// vite.config.ts
server: {
  port: 3000,
  proxy: {
    '/api': {
      target: 'http://localhost:8080',
      changeOrigin: true,
    },
  },
}
```

### API Base URL

Update the API base URL in `src/services/api.ts` if needed:

```typescript
const API_BASE_URL = 'http://localhost:8080/api/v1';
```

## 📋 Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build
- `npm run lint` - Run ESLint

## 🎨 UI Components

### Authentication Flow
- **Login Page** - Email/password authentication
- **Register Page** - User registration
- **Protected Routes** - JWT token-based protection

### Course Management
- **Course Listing** - Browse all available courses
- **Course Creation** - Create new courses with rich metadata
- **Course Details** - View course information and lessons
- **My Courses** - Manage created and enrolled courses

### Lesson System
- **Lesson Viewer** - Watch lessons with progress tracking
- **Progress Tracking** - Automatic progress updates
- **Completion Status** - Mark lessons as completed

### Navigation
- **Responsive Navbar** - Mobile-friendly navigation
- **User Menu** - Account management and logout
- **Breadcrumbs** - Clear navigation context

## 🔐 Authentication

The app uses JWT tokens stored in cookies:

- **Access Token** - Stored in `access_token` cookie
- **User Info** - Stored in `user` cookie
- **Auto Logout** - Redirects to login on token expiry

## 📱 Responsive Design

The frontend is fully responsive with:

- **Mobile-first** approach using Tailwind CSS
- **Breakpoint system** - sm, md, lg, xl
- **Touch-friendly** interactions
- **Optimized layouts** for all screen sizes

## 🚀 Deployment Options

### Option 1: Static Hosting (Netlify, Vercel)
```bash
npm run build
# Deploy dist/ folder
```

### Option 2: Docker
```dockerfile
FROM node:18-alpine
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production
COPY . .
RUN npm run build
EXPOSE 3000
CMD ["npm", "run", "preview"]
```

### Option 3: Nginx
```nginx
server {
    listen 80;
    server_name your-domain.com;
    root /path/to/dist;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }

    location /api {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

## 🔄 API Integration

The frontend integrates with all backend APIs:

### Authentication APIs
- `POST /auth/login` - User login
- `POST /auth/register` - User registration

### Course APIs
- `GET /courses` - List published courses
- `GET /courses/search` - Search courses
- `POST /courses` - Create course
- `PUT /courses/:id` - Update course
- `DELETE /courses/:id` - Delete course
- `POST /courses/:id/enroll` - Enroll in course

### Lesson APIs
- `GET /courses/:id/lessons` - Get course lessons
- `POST /courses/:id/lessons` - Create lesson
- `PUT /lessons/:id` - Update lesson
- `POST /lessons/progress` - Update progress

## 🎯 Next Steps

To complete the frontend implementation:

1. **Complete Page Components** - Implement full functionality for all pages
2. **Form Validation** - Add comprehensive form validation
3. **Error Handling** - Implement global error handling
4. **Loading States** - Add loading spinners and skeletons
5. **Search & Filters** - Advanced course search and filtering
6. **Video Player** - Integrate video player for lessons
7. **Progress Visualization** - Charts and progress bars
8. **Notifications** - Toast notifications for user feedback
9. **Testing** - Add unit and integration tests
10. **PWA Features** - Add service worker for offline support

## 📞 Support

For issues or questions:
- Check the backend API documentation
- Review the Postman collection for API examples
- Ensure the backend server is running on port 8080

---

**Happy Learning! 🎓**
