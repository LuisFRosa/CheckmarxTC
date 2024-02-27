# CheckmarxTC
Checkmarx Technical Challenge using ReactJS and GoLang

## Prerequisites

- Docker
- Docker Compose

## Getting Started

1. **Clone the repository:**

    ```bash
    git clone https://github.com/LuisFRosa/CheckmarxTC.git
    ```

2. **Build the Docker containers and start the services:**

    ```bash
    docker-compose up --build
    ```

3. **Access the application:**

    Once the services are up and running, you can access your application at `http://localhost:3000`, where `3000` is the port specified in your `docker-compose.yml` file.

## Stopping the Application

To stop the application and remove the containers, run:

```bash
docker-compose down