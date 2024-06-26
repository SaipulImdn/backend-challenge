---

# Online Store

## Description

This project is a back-end application for an online store built using Go-lang and Fiber. The application provides APIs for product management, shopping cart, and transactions.

## Features

- View product list by category.
- Add products to the shopping cart.
- View the list of products in the shopping cart.
- Remove products from the shopping cart.
- Checkout and process payments.
- Customer registration and login.

## Project Structure

- `cmd`: Directory for `main.go`.
- `internal`: Directory for application code, containing:
  - `controllers`
  - `middleware`
  - `models`
  - `routes`
- `configs`: Directory for `config.go`.
- `.env.example`: Example environment configuration file.
- `docker-compose.yml`: File for setting up Docker compose.
- `Dockerfile`: File for building the Docker image.
- `go.mod` and `go.sum`: Files for Go dependency management.

## How to Run

1. Clone this repository.
2. Create a `.env` file based on `.env.example`.
3. Run `docker-compose up` to start the application.
4. The application will be available at `http://localhost:5000`.

## API Endpoints

- `GET /products/:category`: Get a list of products by category.
- `POST /cart`: Add products to the shopping cart.
- `GET /cart/:user_id`: View the list of products in the shopping cart.
- `DELETE /cart/:cart_id`: Remove products from the shopping cart.
- `POST /checkout`: Checkout and process payments.
- `POST /register`: Customer registration.
- `POST /login`: Customer login.

---
