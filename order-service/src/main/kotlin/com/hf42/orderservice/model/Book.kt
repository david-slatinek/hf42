package com.hf42.orderservice.model

import io.quarkus.arc.impl.Identified
import org.bson.codecs.pojo.annotations.BsonCreator
import org.bson.codecs.pojo.annotations.BsonProperty
import javax.validation.constraints.NotBlank
import javax.validation.constraints.NotNull
import javax.validation.constraints.Positive
import javax.validation.constraints.Size

data class Book @BsonCreator constructor(
    @Identified("isbn")
    @field:NotBlank(message = "ISBN is required")
    @field:Size(min = 10, max = 13, message = "ISBN must be between 10 and 13 characters")
    @param:BsonProperty("isbn")
    val isbn: String,

    @field:NotBlank(message = "Title is required")
    @field:Size(min = 1, max = 255, message = "Title must be between 1 and 255 characters")
    @param:BsonProperty("title")
    val title: String,

    @field:Size(min = 1, max = 255, message = "Subtitle must be between 1 and 255 characters")
    @param:BsonProperty("subtitle")
    val subtitle: String? = null,

    @field:NotBlank(message = "Author is required")
    @field:Size(min = 1, max = 255, message = "Author must be between 1 and 255 characters")
    @param:BsonProperty("author")
    val author: String,

    @field:NotNull(message = "Publication year is required")
    @field:Positive(message = "Publication year must be greater than 0")
    @param:BsonProperty("year")
    val year: Int,

    @field:NotBlank(message = "Description is required")
    @field:Size(min = 1, max = 255, message = "Description must be between 1 and 255 characters")
    @param:BsonProperty("description")
    val description: String,

    @field:NotBlank(message = "Categories are required")
    @field:NotNull(message = "Categories must be a list of strings")
    @field:Size(min = 1, message = "Categories must contain at least one category")
    @param:BsonProperty("categories")
    val categories: List<String>,

    @field:Size(min = 1, max = 255, message = "Original title must be between 1 and 255 characters")
    @param:BsonProperty("originalTitle")
    val originalTitle: String? = null,

    @field:Size(min = 1, max = 255, message = "Original subtitle must be between 1 and 255 characters")
    @param:BsonProperty("originalSubtitle")
    val originalSubtitle: String? = null,

    @field:Positive(message = "Original publication year must be greater than 0")
    @param:BsonProperty("originalYear")
    val originalYear: Int? = null,

    @field:Size(min = 1, max = 255, message = "Translator must be between 1 and 255 characters")
    @param:BsonProperty("translator")
    val translator: String? = null,

    @field:NotBlank(message = "Size is required")
    @field:Size(min = 1, max = 255, message = "Size must be between 1 and 255 characters")
    @param:BsonProperty("size")
    val size: String,

    @field:NotBlank(message = "Weight is required")
    @field:Size(min = 1, max = 255, message = "Weight must be between 1 and 255 characters")
    @param:BsonProperty("weight")
    val weight: String,

    @field:NotBlank(message = "Number of pages is required")
    @field:Positive(message = "Number of pages must be greater than 0")
    @param:BsonProperty("pages")
    val pages: Int,

    @field:NotBlank(message = "Publisher is required")
    @field:Size(min = 1, max = 255, message = "Publisher must be between 1 and 255 characters")
    @param:BsonProperty("publisher")
    val publisher: String,

    @field:NotBlank(message = "Language is required")
    @field:Size(min = 1, max = 255, message = "Language must be between 1 and 255 characters")
    @param:BsonProperty("language")
    val language: String,

    @field:NotNull(message = "Quantity is required")
    @field:Positive(message = "Quantity must be greater than 0")
    @param:BsonProperty("quantity")
    val quantity: Int,

    @field:NotBlank(message = "Price is required")
    @field:Positive(message = "Price must be greater than 0")
    @param:BsonProperty("price")
    val price: Double,

    @field:NotBlank(message = "Total price is required")
    @field:Positive(message = "Total price must be greater than 0")
    @param:BsonProperty("totalPrice")
    var totalPrice: Double,
)