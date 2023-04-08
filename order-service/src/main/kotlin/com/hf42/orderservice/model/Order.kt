package com.hf42.orderservice.model

import io.quarkus.arc.impl.Identified
import org.bson.codecs.pojo.annotations.BsonCreator
import org.bson.codecs.pojo.annotations.BsonProperty
import javax.validation.constraints.NotBlank
import javax.validation.constraints.NotNull
import javax.validation.constraints.Null
import javax.validation.constraints.Pattern
import javax.validation.constraints.Positive
import javax.validation.constraints.Size

data class Order @BsonCreator constructor(
    @Identified("orderID")
    @field:Null(message = "Order ID must be null")
    @field:Size(min = 36, max = 36, message = "Order ID must be 36 characters")
    @param:BsonProperty("orderID")
    var orderID: String? = null,

    @field:NotBlank(message = "Customer ID is required")
    @field:Size(min = 24, max = 24, message = "Customer ID must be 24 characters")
    @param:BsonProperty("customerID")
    val customerID: String,

    @field:NotBlank(message = "Order date is required")
    @field:Size(min = 16, max = 16, message = "Order date must be 16 characters")
    @field:Pattern(
        regexp = "^[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}\$",
        message = "Order date must be in yyyy-MM-dd HH:mm format"
    )
    @param:BsonProperty("orderDate")
    val orderDate: String,

    @field:NotNull(message = "Books are required")
    @field:Size(min = 1, message = "Books must contain at least one book")
    @param:BsonProperty("books")
    val books: List<Book>,

    @field:NotNull(message = "Total price is required")
    @field:Positive(message = "Total price must be greater than 0")
    @param:BsonProperty("totalPrice")
    val totalPrice: Double,

    @field:NotNull(message = "Status is required")
    @field:Pattern(regexp = "^(pending|shipped|delivered)\$", message = "Status must be pending, shipped, or delivered")
    @field:Size(min = 7, max = 9, message = "Status must be 7 or 9 characters")
    @param:BsonProperty("status")
    var status: String = "pending",
)