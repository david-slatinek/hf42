package com.hf42.userservice.model

data class User(
    private val id: String,
    private var firstName: String,
    private var lastName: String,
    private val email: String,
    private val password: String,
    private var streetAddress: String,
    private var city: String,
    private var postOfficeNumber: Int
)
