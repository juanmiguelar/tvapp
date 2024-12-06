# TVApp Project

TVApp is a full-stack web application designed for managing and displaying news articles. The project consists of three main components:

1. **Frontend**: Built with Vue 3, TypeScript, and Vite, and served using Nginx.
2. **Backend**: Built with Go using Gin and GQLGen for a GraphQL API, with MongoDB as the database.
3. **Database**: MongoDB for storing news articles and their associated authors.

---

## Features

- **Frontend**: 
  - Responsive user interface for listing, creating, and deleting news articles.
  - Built using Vue 3 and TypeScript with state management provided by Pinia.
  
- **Backend**:
  - GraphQL API for querying, creating, updating, and deleting news articles.
  - MongoDB integration for persistent storage.
  - Built using Go and the GQLGen framework.

- **Database**:
  - MongoDB for data persistence, with containerized support for easy deployment.

---

## Setup Instructions

### Prerequisites

- **Docker** and **Podman Compose** installed on your system.
- **Node.js** (for local frontend development).

---

### Running Locally with Docker/Podman

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-repo/tvapp.git
   cd tvapp
   ```

2. **Ensure `.env` files are present**:
   - Create a `.env` file in the root and `database` directories with the required environment variables. Example:
     ```env
     MONGO_INITDB_ROOT_USERNAME=admin
     MONGO_INITDB_ROOT_PASSWORD=password
     MONGO_INITDB_DATABASE=tvapp_db
     ```

3. **Build and run the services**:
   ```bash
   podman-compose up --build
   ```

4. **Access the application**:
   - Frontend: [http://localhost:5173](http://localhost:5173)
   - Backend GraphQL Playground: [http://localhost:8080](http://localhost:8080)

---

### Directory Structure

```plaintext
tvapp/
├── backend/
│   ├── database/       # MongoDB connection setup
│   ├── graph/          # GQLGen schema and resolvers
│   ├── server.go       # Main entry point for the backend
├── database/
│   ├── Dockerfile      # MongoDB Dockerfile
│   ├── data/           # Persistent storage for MongoDB
│   ├── .env            # MongoDB environment variables
├── frontend/
│   ├── src/            # Vue 3 components and stores
│   ├── public/         # Static assets
│   ├── Dockerfile      # Frontend Dockerfile
├── docker-compose.yml  # Compose file to orchestrate services
```

---

### Development

#### Frontend

1. **Navigate to the frontend directory**:
   ```bash
   cd frontend/tvapp-frontend
   ```

2. **Install dependencies**:
   ```bash
   npm install
   ```

3. **Run the development server**:
   ```bash
   npm run dev
   ```

4. **Access the application**:
   - Frontend: [http://localhost:5173](http://localhost:5173)

---

### Backend

1. **Navigate to the backend directory**:
   ```bash
   cd backend
   ```

2. **Run the backend locally**:
   ```bash
   go run server.go
   ```

3. **Access the GraphQL Playground**:
   [http://localhost:8080](http://localhost:8080)

---

## API Overview

### GraphQL Schema

```graphql
type Author {
  name: String!
  email: String!
}

type News {
  id: ID!
  title: String!
  content: String!
  author: Author!
}

type Query {
  getNews: [News!]!
}

type Mutation {
  createNews(
    title: String!
    content: String!
    authorName: String!
    authorEmail: String!
  ): News!

  updateNews(
    id: ID!
    title: String
    content: String
    authorName: String
    authorEmail: String
  ): News!

  deleteNews(id: ID!): Boolean!
}
```

---

## Docker Setup

### Frontend

The frontend is built using a multi-stage Dockerfile. It uses `node:18-alpine` for the build stage and `nginx:alpine` for serving static files.

#### Example Dockerfile:

```dockerfile
# Build Stage
FROM node:18-alpine as builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Serve Stage
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

---

## Contribution Guidelines

1. Fork the repository.
2. Create a new branch for your feature or fix:
   ```bash
   git checkout -b feature-name
   ```
3. Commit and push your changes.
4. Submit a pull request.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Authors

- [Your Name](https://github.com/your-username)
- Contributions by [Other Developers]
