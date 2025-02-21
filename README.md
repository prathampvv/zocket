Zocket - Task Dashboard with AI-Powered Chat

Zocket is a Next.js (App Router) + TypeScript + Tailwind CSS application designed to provide a task dashboard with real-time updates and an AI-powered chat for task recommendations. It includes JWT authentication for secure access and uses a PostgreSQL database with Prisma ORM for data persistence. The project is built using a scalable architecture, ensuring smooth performance and easy maintenance.

The key features of Zocket include a task dashboard that allows users to manage tasks with real-time updates, an AI-powered chat that provides intelligent task recommendations, and JWT-based authentication for secure login and access control. The responsive UI is built with Tailwind CSS, making the application accessible across different devices.

The project follows a structured file organization, with separate folders for the frontend (built with Next.js), the backend (powered by Node.js and Express.js), and the database schema (managed using Prisma ORM). It also includes API routes, reusable UI components, and configuration files.

To set up the project, users need to clone the repository from GitHub, install the necessary dependencies, and configure the environment variables for both the frontend and backend. The backend requires a PostgreSQL database connection and a secret key for JWT authentication, while the frontend needs an API URL and an OpenAI API key for AI-powered recommendations.

After setting up the environment, the application can be started by running the backend and frontend servers separately. The backend server runs on port 5000, while the frontend server operates on port 3000, making the application accessible via http://localhost:3000.

Contributions to Zocket are welcome, and developers can submit issues and pull requests by following the Git branching workflow. New features should be developed in separate branches, committed with appropriate messages, and pushed to the repository before merging.

The project is licensed under the MIT License, allowing for open-source contributions and modifications. Whether you're looking for a smart task management solution or an AI-powered assistant for productivity, Zocket provides an intuitive and scalable platform to meet your needs
