package com.hf42.orderservice.model

data class Book(
    val isbn: String,
    val title: String,
    val subtitle: String,
    val author: String,
    val year: Int,
    val description: String,
    val categories: List<String>,
    val originalTitle: String,
    val originalSubtitle: String,
    val originalYear: Int,
    val translator: String,
    val size: String,
    val weight: String,
    val pages: Int,
    val publisher: String,
    val language: String,
    val price: Double
) {
}