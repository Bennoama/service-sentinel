# Service Sentinel

**Service Sentinel** is a monitoring application written in Go that allows users to set up and manage various monitors (e.g., HTTP and ICMP) to track the health and performance of services.
It features a user-friendly UI for creating and managing monitors, as well as a backend that stores monitoring results in a database.

## Features

### Current Functionality
1. **Database-Driven Initialization**:
   - On startup, the service retrieves monitor configurations from the database and starts monitoring them automatically.
2. **Dynamic Monitoring**:
   - Each new monitor added via the UI immediately begins monitoring the configured service.

### Planned Functionality
1. **Add New Monitors**:
   - Add HTTP or ICMP monitors via the UI (feature under development).
2. **Save Monitoring Results**:
   - Responses from each monitor will be stored in the database for analysis (feature under development).
3. **View Monitoring Results**:
   - The UI will display the responses from all monitors for easy review (feature under development).

## Getting Started

### Prerequisites
- **Go**: Ensure Go is installed (minimum version: 1.20).
- **Database**: The application requires a database connection. Supported databases include PostgreSQL (preferred).

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/Bennoama/service-sentinel.git
   cd service-sentinel

2. Install dependencies:
   go mod tidy
3.Set up the database:
  Create a database and apply migrations using the provided schema (to be included in the repository).
4. Run the application:
  go run .\cmd\main.go

5. Access the UI:
  The UI will be available at http://localhost:8080 (default port).

Contributing
Contributions are welcome! Please fork the repository, create a feature branch, and submit a pull request.
     
