# Basic Booking CLI Application

This is a simple **Command-Line Interface (CLI)** application built using **Golang** as part of my practice with the Go programming language. The application allows users to simulate a basic seat booking system.

## Overview
The Booking CLI Application demonstrates fundamental concepts in Golang, such as:
- Input and output handling in the terminal
- Working with loops, conditions, and functions
- Managing data using slices and maps
- Concurrency for handling multiple bookings (if implemented)

The application is beginner-friendly and serves as a stepping stone for learning how to build CLI tools with Go.

## Features
1. **User Input**: Users can input their name, email, and the number of seats to book.
2. **Validation**: Basic input validation is performed (e.g., no negative seat numbers).
3. **Booking Management**: Bookings are stored in a slice, and relevant information is displayed after each booking.
4. **Summary**: The application shows a summary of all current bookings.

## Prerequisites
To run this application, you need to have:
- **Golang** installed on your machine ([Download here](https://golang.org/dl/))
- A terminal or command-line tool

## How to Run
1. Clone or download this repository:
   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```

2. Run the Go file:
   ```bash
   go run main.go
   ```

3. Follow the on-screen instructions to make bookings.

## Example Workflow
```plaintext
Welcome to the CLI Booking Application!
Enter your name: Ahan
Enter your email: ahan@example.com
Enter the number of seats you want to book: 3

Thank you, Ahan! You have successfully booked 3 seats.

Current Bookings:
- Ahan (ahan@example.com) booked 3 seats
```

## Code Structure
- `main.go`: Contains the core logic of the booking system, including input validation, booking storage, and summary display.

## Key Concepts Practiced
- Working with slices, maps, and basic data structures
- Handling user input/output
- Writing functions for modular code
- Basic error checking and validation

## Future Improvements
- Add a database for persistent storage
- Implement concurrency to handle simultaneous bookings
- Include a cancellation feature
- Add seat availability tracking

## License
This project is for educational purposes and is released under the MIT License.

---
**Author**: [Ahan Bandyopadhyay](#)
