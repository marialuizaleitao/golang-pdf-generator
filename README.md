# Go PDF Generator

This is a simple Go application for generating PDF files from HTML templates using the wkhtmltopdf library.

## Technologies Used

- **Go**: Go is an open-source programming language developed by Google, known for its simplicity, performance, and concurrency support. We used Go for the backend logic of this application.

- **wkhtmltopdf**: wkhtmltopdf is an open-source command-line tool that converts HTML documents to PDF files using the WebKit rendering engine. We used the Go wrapper package `github.com/SebastiaanKlippert/go-wkhtmltopdf` to interface with wkhtmltopdf.

## Installation

To run this application, you need to have Go installed on your system. You also need to have wkhtmltopdf installed and accessible from the command line.

To install the required Go dependencies, you can use the following command:
```bash
go mod download
```

## Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/your-repository.git
   cd your-repository
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug fixes, feel free to open an issue or submit a pull request.
