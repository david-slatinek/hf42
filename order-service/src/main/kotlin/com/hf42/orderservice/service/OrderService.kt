package com.hf42.orderservice.service

import com.mongodb.client.MongoClient
import com.mongodb.client.MongoCollection
import com.hf42.orderservice.model.Order
import javax.enterprise.context.ApplicationScoped
import javax.inject.Inject

@ApplicationScoped
class OrderService {
    @Inject
    lateinit var mongoClient: MongoClient

    fun getCollection(): MongoCollection<Order>? {
        return mongoClient.getDatabase("order-service").getCollection("orders", Order::class.java)
    }
}