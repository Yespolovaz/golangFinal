#🌟 Project Name: Sigma Bank 🚀
Welcome to the **Sigma bank**! 🎉

##📝 Project Description
The Sigma bank is a comprehensive web application that combines online banking, shopping 🛒, insurance, and money transfers. It provides users with a seamless experience to manage their finances, purchase products, obtain insurance coverage, and transfer money between accounts. 💰💳

1. [Project Description](#project-description)
2. [API Structure](#api-structure)
3. [Database Entity Structure](#database-entity-structure)
4. [Team Members](#team-members)
5. [CRUD Functionality](#crud-functionality)
6. [Code Structure](#code-structure)
7. [Getting Started](#getting-started)

#Getting Started

##🚀 API Structure
The API is designed to provide endpoints for various functionalities, including user management, banking operations, shopping, insurance, and money transfers. Each endpoint is carefully crafted to ensure security, efficiency, and ease of use for the end-users. 🌐

### User Management API

- `/api/users` - Endpoint to manage users (CRUD operations).

### Banking API

- `/api/accounts` - Endpoint to manage bank accounts (CRUD operations).
- `/api/transactions` - Endpoint to perform transactions (e.g., deposit, withdrawal, transfer).

### Shopping API

- `/api/products` - Endpoint to manage products (CRUD operations).
- `/api/cart` - Endpoint to manage shopping carts.

### Insurance API

- `/api/policies` - Endpoint to manage insurance policies (CRUD operations).

### Money Transfer API

- `/api/transfers` - Endpoint to initiate money transfers between accounts.

##🏦 Database Entity Structure
The database of the project consists of several entities interconnected to support the functionalities of the application. The main entities include users, bank accounts, products, insurance policies, and transactions. These entities are designed to store relevant information and maintain relationships to facilitate seamless operations within the application. 🗃️

### User Entity

- `id` - Unique identifier for the user.
- `username` - Username of the user.
- `email` - Email address of the user.
- `password` - Encrypted password of the user.
- ...

### Bank Account Entity

- `id` - Unique identifier for the bank account.
- `userId` - Foreign key referencing the user who owns the account.
- `accountNumber` - Account number.
- `balance` - Current balance of the account.
- ...

### Product Entity

- `id` - Unique identifier for the product.
- `name` - Name of the product.
- `description` - Description of the product.
- `price` - Price of the product.
- ...

### Insurance Policy Entity

- `id` - Unique identifier for the insurance policy.
- `userId` - Foreign key referencing the user who owns the policy.
- `type` - Type of insurance.
- `coverage` - Coverage details.
- ...

### Transaction Entity

- `id` - Unique identifier for the transaction.
- `accountId` - Foreign key referencing the bank account involved in the transaction.
- `type` - Type of transaction (e.g., deposit, withdrawal, transfer).
- `amount` - Amount of the transaction.
- ...

👨‍💻 Team Members
- [Mombekov Dias](https://github.com/Dias21B030874) (Student ID: 21B030874) 🧑‍💼
- [Yespolova Zharkynay](https://github.com/Dias21B030874) (Student ID: 21B030666) 👩‍💼
- [Kidirmaganbetov Nurken](https://github.com/Dias21B030874) (Student ID: 21B030) 🧑‍💻

##🔨 CRUD Functionality
The project implements CRUD (Create, Read, Update, Delete) functionality for managing various entities in the system. For example, users can create, view, update, and delete their bank accounts, products, insurance policies, and transactions. The CRUD operations ensure that users have full control over their data and can perform necessary actions with ease. 🔄

##📁 Code Structure
The codebase of the project follows the Standard Layout convention to maintain consistency and readability. The project is organized into modules, each focusing on specific functionalities such as user management, banking operations, shopping, insurance, and money transfers. The codebase is well-documented and adheres to best practices to ensure maintainability and scalability. 📦

In this structure:

- `cmd`: Contains the main application entry point.
- `pkg`: Contains the internal packages of the application.
  - `api`: Handles API endpoints.
  - `models`: Defines the data models.
  - `repositories`: Implements data access logic.
  - `services`: Implements business logic.
- `migrations`: Contains database migration scripts.
- `config`: Contains configuration files.
- `go.mod`: Go module file.


##🚀 Getting Started
To get started with the *Sigma bank*, follow these steps:

Thank you for choosing the **Sigma Bank**! We hope you have a great experience using our application. 🌟
