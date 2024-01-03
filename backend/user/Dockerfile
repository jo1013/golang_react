# Go 애플리케이션을 위한 베이스 이미지 사용
FROM golang:latest AS builder


# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY ./main /app/main


# Build the Beego application
RUN go build -o main .

# Use a smaller base image for the final runtime image
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the built executable from the build image
# COPY --from=build /app/main /app/main

# Expose port 8080 for the application
EXPOSE 8080

COPY . .

# Run the Beego application
CMD ["./main"]