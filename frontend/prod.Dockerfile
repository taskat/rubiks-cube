# Stage 1: Build the npm project
FROM node:18.15.0 AS builder
WORKDIR /app

# Copy package.json and package-lock.json separately to leverage Docker cache
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the npm project
RUN npm run build

# Stage 2: Create the production-ready nginx container
FROM nginx:latest

# Copy the built assets from the previous stage to nginx's default html directory
COPY --from=builder /app/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 8081 to the outside world
EXPOSE 8081
# The nginx container automatically starts serving content from /usr/share/nginx/html
