# Go Rest API 🏦💱
Go Rest API is a Event creation API ( CRUD ), developed in Go.

## ✨ Technologies
- Go (Gin Framework)
- SQLite
- JWT
- Rest API

## 📖 How to Run Locally
1. Clone the repository:
   ```sh
   git clone https://github.com/JoaoVictorLiz/Go_Rest_API.git
   cd Go_Rest_API
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Start the database and run the project:
   ```sh
   go run main.go
   ```

## 🔧 Available Endpoints
| Method | Route             | Description                    |
|--------|------------------|--------------------------------|
| `GET` | `/events`       | Fetches all events |
| `GET` | `/events/:id`       | Fetches a unique event|
| `POST` | `/events`       | Create a new event |
| `PUT`  | `/events/:id`  | Update a unique existing event by event ID|
| `DELETE`  | `/events/:id`      | Delete a unique event |
| `POST`  | `/events/:id/register`      | Register for a specific event |
| `DELETE`  | `/events/:id/register`      | unregister for a specific event|
| `POST`  | `/signup`      | Register a new user |
| `POST`  | `/login`      | Login |

## 📢 Contributing
Feel free to submit issues and pull requests!

## 📣 Contact
- LinkedIn: [My LinkedIn](https://www.linkedin.com/in/joão-victor-liz-da-silveira-b347a71b5/)
- Email: joaovictorlizsilveira@gmail.com
