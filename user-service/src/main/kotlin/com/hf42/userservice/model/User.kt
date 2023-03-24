package com.hf42.userservice.model

import com.fasterxml.jackson.annotation.JsonIgnoreProperties
import jakarta.validation.constraints.*
import org.springframework.data.annotation.Id
import org.springframework.data.mongodb.core.mapping.Document

@Document("users")
data class User(
    @Id
    @field:Null(message = "Id should not be specified")
    val id: String?,

    @field:NotBlank(message = "First name is mandatory")
    @field:Size(min = 2, max = 20, message = "First name should be between 2 and 20 characters")
    private var firstName: String,

    @field:NotBlank(message = "Last name is mandatory")
    @field:Size(min = 2, max = 20, message = "Last name should be between 2 and 20 characters")
    private var lastName: String,

    @field:NotBlank(message = "Email is mandatory")
    @field:Email(message = "Email should be valid")
    @field:Size(min = 2, max = 30, message = "Email should be between 2 and 20 characters")
    val email: String,

    @field:NotBlank(message = "Password is mandatory")
    @field:Size(min = 8, message = "Password should be greater than 8 characters")
    var password: String,

    @field:NotBlank(message = "Street address is mandatory")
    @field:Size(min = 2, max = 30, message = "Street address should be between 2 and 30 characters")
    private var streetAddress: String,

    @field:NotBlank(message = "City is mandatory")
    @Size(min = 2, max = 20, message = "City should be between 2 and 20 characters")
    private var city: String,

    @field:Min(value = 1000, message = "Zip code should be greater than 1000")
    private var postOfficeNumber: Int
)
