package com.hf42.orderservice.service

import com.mongodb.client.MongoClient
import com.mongodb.client.MongoCollection
import com.hf42.orderservice.model.Order
import com.mongodb.client.result.InsertOneResult
import javax.enterprise.context.ApplicationScoped
import javax.inject.Inject

@ApplicationScoped
class OrderService {
    @Inject
    lateinit var mongoClient: MongoClient

    private fun getCollection(): MongoCollection<Order> {
        return mongoClient.getDatabase("order-service").getCollection("orders", Order::class.java)
    }

    fun insertOrder(order: Order): String? {
        val id = getCollection().insertOne(order)
        if (id.insertedId == null) {
            return null
        }
        return id.insertedId.asObjectId().value.toHexString()
    }
}