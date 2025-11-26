# Frontend

React + TypeScript + Vite setup for the ticketing app.

## Setup

```bash
npm install
npm run dev
```

Opens on http://localhost:5173

## What's in here

- React 18 with TypeScript
- Vite for dev server and builds
- React Router for navigation
- Axios for API calls

## Structure

```
src/
├── pages/       - page components
├── components/  - shared components
├── contexts/    - React context (auth, etc)
├── services/    - API calls
├── utils/       - helpers
└── types/       - TS types
```

## Authentication

Login and protected routes are working with backend integration:

- JWT tokens stored in localStorage (`auth_token` key)
- Protected routes auto-redirect to login if unauthenticated
- Token injection via Axios interceptors (automatic headers)
- Logout clears token and session

### Development Login

Backend is currently using stub auth:
- Email: any valid email format (e.g., `user@example.com`)
- Password: anything works (validation not enforced yet)
- User data is mocked by backend, JWT generation is real

The `AuthContext` manages auth state. Use `ProtectedRoute` wrapper for pages that need auth.

## Config

Copy `.env.example` to `.env` if you need to change the API URL. Defaults should work fine for local dev.

The Vite config proxies `/api/*` to the backend on port 8080, so CORS isn't an issue locally.

## Building

```bash
npm run build    # outputs to dist/
npm run preview  # test the build locally
```

Static files go in `dist/` - can be deployed to S3, Vercel, anywhere really.

## Docker

There's a Dockerfile if you want to run this in a container. Or just use docker-compose from the root.