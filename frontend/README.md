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