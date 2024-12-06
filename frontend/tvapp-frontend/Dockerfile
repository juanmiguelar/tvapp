# Stage 1: Build the Vue application
FROM node:18-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Install dependencies
COPY package*.json ./
RUN npm install --production=false

# Copy application source code
COPY . .

# Build the application
RUN npm run build

# Stage 2: Serve with Nginx
FROM nginx:alpine

# Remove default Nginx configuration
RUN rm /etc/nginx/conf.d/default.conf

# Copy custom Nginx configuration (optional)
# Uncomment the following lines if you want a custom Nginx configuration
# COPY nginx.conf /etc/nginx/conf.d/

# Copy built files from the previous stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose the port for the Nginx server
EXPOSE 80

# Start Nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]
