ARG GO_VERSION=1.22.0
ARG APP_BASE_IMAGE=golang:${GO_VERSION}-bookworm

# Build stage
FROM ${APP_BASE_IMAGE} as dev

WORKDIR /usr/src/app

# Install Google Chrome and necessary dependencies
RUN curl -o /tmp/google-chrome.deb https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && \
    apt-get update && \
    apt-get install -y /tmp/google-chrome.deb

# Copy Go modules files
COPY ./app/go.mod ./app/go.sum ./

# Download and verify Go modules
RUN go mod download && go mod verify

# Copy application source code
COPY ./app .

# Build the application
RUN go build -v -o /usr/local/bin/run-app ./main.go

# run-app stage
FROM debian:bookworm as run-app

# Install necessary dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    fonts-liberation \
    libappindicator3-1 \
    libasound2 \
    libatk-bridge2.0-0 \
    libatk1.0-0 \
    libcups2 \
    libdbus-1-3 \
    libgdk-pixbuf2.0-0 \
    libnspr4 \
    libnss3 \
    libxcomposite1 \
    libxdamage1 \
    libxrandr2 \
    xdg-utils \
    wget && \
    update-ca-certificates

# Copy the built application from the dev stage
COPY --from=dev /usr/local/bin/run-app /usr/local/bin/run-app

# Copy Google Chrome and its dependencies from the dev stage
COPY --from=dev /opt/google/chrome /opt/google/chrome

# Manually create the symbolic links for Google Chrome
RUN ln -sf /opt/google/chrome/google-chrome /usr/bin/google-chrome && \
    ln -sf /opt/google/chrome/google-chrome /usr/bin/google-chrome-stable && \
    ln -sf /opt/google/chrome/google-chrome /etc/alternatives/google-chrome && \
    ln -sf /opt/google/chrome/google-chrome /etc/alternatives/google-chrome-stable && \
    ln -sf /opt/google/chrome/google-chrome /etc/alternatives/gnome-www-browser && \
    ln -sf /opt/google/chrome/google-chrome /etc/alternatives/x-www-browser

# Copy the entrypoint script
COPY entrypoint.sh /usr/local/bin/entrypoint.sh

# Make the entrypoint script executable
RUN chmod +x /usr/local/bin/entrypoint.sh

# Set the entrypoint
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]