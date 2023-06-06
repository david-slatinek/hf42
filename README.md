<div align="center">
    <img alt="Go" src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"/>
    <img alt="AWS" src="https://img.shields.io/badge/Amazon_AWS-232F3E?style=for-the-badge&logo=amazonaws&logoColor=white"/>
    <img alt="Spring Boot" src="https://img.shields.io/badge/Spring_Boot-F2F4F9?style=for-the-badge&logo=spring-boot"/>
    <img alt="Kotlin" src="https://img.shields.io/badge/Kotlin-0095D5?&style=for-the-badge&logo=kotlin&logoColor=white"/>
    <img alt="Quarkus" src="https://img.shields.io/badge/Quarkus-4695EB?style=for-the-badge&logo=quarkus&logoColor=white"/>
    <img alt="RabbitMQ" src="https://img.shields.io/badge/rabbitmq-%23FF6600.svg?&style=for-the-badge&logo=rabbitmq&logoColor=white"/>
    <img alt="Amazon S3" src="https://img.shields.io/badge/amazons3-569A31?style=for-the-badge&logo=amazons3&logoColor=white"/>
    <img alt="AWS Lambda" src="https://img.shields.io/badge/aws_lambda-FF9900?style=for-the-badge&logo=awslambda&logoColor=white"/>
    <img alt="MongoDB" src="https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white"/>
    <img alt="MySQL" src="https://img.shields.io/badge/MySQL-005C84?style=for-the-badge&logo=mysql&logoColor=white"/>
    <img alt="Docker" src="https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white"/>
    <img alt="React" src="https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB"/>
    <img alt="Github Actions" src="https://img.shields.io/badge/Github%20Actions-282a2e?style=for-the-badge&logo=githubactions&logoColor=367cfe"/>
</div>

![Logo](payment-service/logo/logo.png)

Book store project with support for managing books, users, orders, and payments. The [book-service](book-service) is a
REST API that provides CRUD operations for books and was made with Go. The [user-service](user-service)
provides CRUD operations for users and was made with Spring Boot and Kotlin.

The [order-service](order-service) provides CRUD operations for orders and was made with Quarkus. After an order is
created, it is sent to a [RabbitMQ](messaging) queue. The [payment-service](payment-service) reads orders from the queue
and, if a given user is valid, which is validated by the *user-service*, simulates a payment. It creates a
payment record in the database, generates a PDF invoice, and uploads it to AWS S3. At this point,
AWS [Lambda](email-lambda) is triggered, which sends an email with the invoice to the user.

<div align="center">
  <img src="images/diagram.png" alt="Invoicing">
  <br/>
  <i>Invoicing.</i>
</div>

<br/>

<div align="center">
  <img src="images/email.png" alt="Email with a link to the invoice">
  <br/>
  <i>Email with a link to the invoice.</i>
</div>

<br/>

<div align="center">
  <img src="images/invoice.png" alt="An example of an invoice">
  <br/>
  <i>An example of an invoice.</i>
</div>

<br/>

The [gateway](gateway) is a reverse proxy that uses Kong to route requests to the other services. The [website](website)
is a React application that uses micro frontends to display the books.

Payment service uses MySQL for storage, other services use MongoDB.
