# After writing a telegram bot, I started taking the Udemy course
Course: [Backend Master Class [Golang + Postgres + Kubernetes + gRPC]](https://www.udemy.com/share/105PNI3@vgqgMszJTBS_PimvIiwd2I1n0p9H2f5nGWh1DrDvUUMWqXTTEnR6b9sAp31jfWQmkQ==/)

## Working with database [Postgres + SQLC]
 
1. [x] **Design DB schema and generate SQL code with dbdiagram.io**

4. [x] **Use Docker + Postgres + TablePlus to create DB schema**

5. [x] **How to write & run database migration in Golang**

6. [x] **Generate CRUD Golang code from SQL | Compare db/sql, gorm, sqlx & sqlc**
    - Practice Test 1: Write SQL queries for transfers and entries table and generate Go code

7. [x] **Write unit tests for database CRUD with random data in Golang**
    - Write tests for the CRUD operations of Entry and Transfer tables
    - Practice Test 2: Configure the right sql_package in your sqlc.yaml file

8. [x] **A clean way to implement database transaction in Golang**

9. [x] **DB transaction lock & How to handle deadlock in Golang**

10. [x] **How to avoid deadlock in DB transaction? Queries order matters!**

11. [x] **Deeply understand transaction isolation levels & read phenomena**

12. [x] **Setup Github Actions for Golang + Postgres to run automated tests**

## Section 2: Building RESTful HTTP JSON API [Gin + JWT + PASETO]

1. [x] **Implement RESTful HTTP API in Go using Gin**

2. [x] **Load config from file & environment variables in Go with Viper**

3. [x] **Mock DB for testing HTTP API in Go and achieve 100% coverage**

4. [x] **Implement transfer money API with a custom params validator**

5. [x] **Add users table with unique & foreign key constraints in PostgreSQL**

6. [x] **How to handle DB errors in Golang correctly**

7. [x] **How to securely store passwords? Hash password in Go with Bcrypt!**

8. [x] **How to write stronger unit tests with a custom gomock matcher**

9. [x] **Why PASETO is better than JWT for token-based authentication?**

10. [x] **How to create and verify JWT & PASETO token in Golang**

11. [x] **Implement login user API that returns PASETO or JWT access token in Go**

12. [x] **Implement authentication middleware and authorization rules in Golang using Gin**

## Section 3: Deploying the application to production [Kubernetes + AWS]

1. [x] **Build a minimal Golang Docker image with a multistage Dockerfile**

2. [x] **How to use docker network to connect 2 stand-alone containers**

3. [x] **How to write docker-compose file and control service start-up orders with wait-for.sh**

4. [x] **How to create a free tier AWS account**

5. [x] **Auto build & push docker image to AWS ECR with Github Actions**

6. [x] **How to create a production DB on AWS RDS**

7. [x] **Store & retrieve production secrets with AWS secrets manager**

8. [x] **Kubernetes architecture & How to create an EKS cluster on AWS**

9. [x] **How to use kubectl & k9s to connect to a kubernetes cluster on AWS EKS**

10. [x] **How to deploy a web app to Kubernetes cluster on AWS EKS**

11. [x] **Register a domain name & set up A-record using Route53**

12. [x] **How to use Ingress to route traffics to different services in Kubernetes**

13. [x] **Automatic issue TLS certificates in Kubernetes with Let's Encrypt**

14. [x] **Automatic deploy to Kubernetes with Github Action**

## Section 4: Advanced Backend Topics [Sessions + gRPC]

1. [x] **How to manage user session with refresh token - Golang**

2. [x] **Generate DB documentation page and schema SQL dump from DBML**

3. [x] **Introduction to gRPC**

4. [x] **Define gRPC API and generate Go code with protobuf**

5. [x] **How to run a golang gRPC server and call its API**

6. [x] **Implement gRPC API to create and login users in Go**

7. [x] **gRPC gateway: write code once, serve both gRPC & HTTP requests**

8. [x] **How to extract info from gRPC metadata**

9. [x] **Automatic generate & serve Swagger docs from Go server**

10. [x] **Embed static frontend files inside Golang backend server's binary**

11. [x] **Validate gRPC parameters and send human/machine friendly response**

12. [x] **Run DB migrations directly inside Golang code**
 
13. [x] **Partial update DB record with SQLC nullable parameters**

14. [x] **Build gRPC update API with optional parameters**
 
15. [x] **Add authorization to protect gRPC API**
 
16. [ ] **Write structured logs for gRPC APIs**
 
17. [ ] **How to write HTTP logger middleware in Go**

## Section 5: Asynchronous processing with background workers [Asynq + Redis]

1. [ ] **Implement background worker in Go with Redis and Asynq**

2. [ ] **Integrate async worker to Go web server**

3. [ ] **Send async tasks to Redis within a DB transaction**

4. [ ] **How to handle errors and print logs for Go Asynq workers**

5. [ ] **A bit of delay might be good for your async tasks**

6. [ ] **How to send emails in Go via Gmail**

7. [ ] **How to skip test in Go and config test flag in vscode**

8. [ ] **Email verification in Go: design DB and send email**

9. [ ] **Implement email verification API in Go**

10. [ ] **Unit test gRPC API with mock DB & Redis**

11. [ ] **How to test a gRPC API that requires authentication**

## Section 6: Improve the stability and security of the server

1. [ ] **Config sqlc version 2 for Go and Postgres**

2. [ ] **Switch DB driver from lib/pq to pgx**

3. [ ] **How to handle DB errors with PGX driver**

4. [ ] **Docker compose: port + volume mapping**

5. [ ] **How to install & use binary packages in Go**

6. [ ] **Implement role-based access control (RBAC) in Go**

7. [ ] **Grant AWS EKS cluster access to Postgres and Redis using security group**

8. [ ] **Deploy gRPC + HTTP server to AWS EKS cluster**

9. [ ] **Don't lose money on AWS**

10. [ ] **Go 1.22 fixed the most common for-loop trap**
