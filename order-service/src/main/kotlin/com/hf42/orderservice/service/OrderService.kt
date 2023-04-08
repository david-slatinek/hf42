package com.hf42.orderservice.service

import com.hf42.orderservice.model.Order
import com.mongodb.client.MongoClient
import com.mongodb.client.MongoCollection
import com.mongodb.client.model.Filters.eq
import javax.enterprise.context.ApplicationScoped
import javax.inject.Inject


@ApplicationScoped
class OrderService {
    @Inject
    lateinit var mongoClient: MongoClient

    private fun getCollection(): MongoCollection<Order> {
        return mongoClient.getDatabase("order-service").getCollection("orders", Order::class.java)
    }

    fun insertOrder(order: Order): Boolean {
        val id = getCollection().insertOne(order)
        return id.insertedId != null
    }

    fun getOrder(id: String): Order? {
        return getCollection().find(eq("orderID", id)).first()
    }

    fun updateOrder(order: Order): Boolean {
        return getCollection().replaceOne(eq("orderID", order.orderID), order).matchedCount == 1L
    }

    fun deleteOrder(id: String): Boolean {
        return getCollection().deleteOne(eq("orderID", id)).deletedCount == 1L
    }
}