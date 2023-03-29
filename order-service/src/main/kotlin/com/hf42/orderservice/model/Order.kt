package com.hf42.orderservice.model

import io.quarkus.arc.impl.Identified
import javax.validation.constraints.NotBlank
import javax.validation.constraints.NotNull
import javax.validation.constraints.Null
import javax.validation.constraints.Pattern
import javax.validation.constraints.Positive
import javax.validation.constraints.Size

data class Order(
    @Identified("orderID")
    @field:Null(message = "Order ID must be null")
    val orderID: String?,

    @field:NotBlank(message = "Customer ID is required")
    @field:Size(min = 24, max = 24, message = "Customer ID must be 24 characters")
    val customerID: String,

    @field:NotBlank(message = "Order date is required")
    @field:Size(min = 16, max = 16, message = "Order date must be 16 characters")
    @field:Pattern(
        regexp = "^[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}\$",
        message = "Order date must be in yyyy-MM-dd HH:mm format"
    )
    val orderDate: String,

    @field:NotNull(message = "Books are required")
    @field:Size(min = 1, message = "Books must contain at least one book")
    val books: List<Book>,

    @field:NotNull(message = "Total price is required")
    @field:Positive(message = "Total price must be greater than 0")
    val totalPrice: Double
) {
}