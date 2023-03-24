package com.hf42.userservice.model

import jakarta.validation.constraints.Email
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.Size

data class Login(
    @field:NotBlank(message = "Email is mandatory")
    @field:Email(message = "Email should be valid") val email: String = "",

    @field:NotBlank(message = "Password is mandatory")
    @field:Size(min = 8, message = "Password should be greater than 8 characters") val password: String = "",
)