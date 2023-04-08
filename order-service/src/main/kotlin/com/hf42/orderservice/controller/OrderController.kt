package com.hf42.orderservice.controller

import com.hf42.orderservice.model.Order
import com.hf42.orderservice.producer.Producer
import com.hf42.orderservice.service.OrderService
import java.util.UUID
import javax.inject.Inject
import javax.ws.rs.*
import javax.ws.rs.core.MediaType
import javax.ws.rs.core.Response

@Path("/")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
class OrderController {
    @Inject
    lateinit var orderService: OrderService

    @Inject
    lateinit var producer: Producer

    @POST
    @Path("/order")
    fun add(order: Order): Response {
        if (order.totalPrice != order.books.sumOf { it.totalPrice }) {
            return Response.status(Response.Status.BAD_REQUEST)
                .entity(mapOf("error" to "Order price doesn't match sum of books prices")).build()
        }

        if (!order.books.all { it.price * it.quantity == it.totalPrice }) {
            return Response.status(Response.Status.BAD_REQUEST)
                .entity(mapOf("error" to "Books price times quantity doesn't match it's total price")).build()
        }

        order.orderID = UUID.randomUUID().toString()

        if (orderService.insertOrder(order)) {
            producer.produce(order)
            return Response.status(Response.Status.CREATED).entity(mapOf("orderID" to order.orderID)).build()
        }
        return Response.status(Response.Status.INTERNAL_SERVER_ERROR).entity(mapOf("error" to "Order was not created"))
            .build()
    }

    @GET
    @Path("/order/{id}")
    fun get(@PathParam("id") id: String): Response {
        if (id.length != 36) {
            return Response.status(Response.Status.BAD_REQUEST)
                .entity(mapOf("error" to "Order ID must be 36 characters")).build()
        }

        val order = orderService.getOrder(id)

        if (order == null) {
            return Response.status(Response.Status.NOT_FOUND).entity(mapOf("error" to "Order not found")).build()
        }
        return Response.status(Response.Status.OK).entity(order).build()
    }

    @PUT
    @Path("/order/{id}")
    fun update(@PathParam("id") id: String, order: Order): Response {
        if (order.orderID == null) {
            return Response.status(Response.Status.BAD_REQUEST)
                .entity(mapOf("error" to "Order ID must be present in body")).build()
        }

        if (id.length != 36 || order.orderID!!.length != 36) {
            return Response.status(Response.Status.BAD_REQUEST)
                .entity(mapOf("error" to "Order ID must be 36 characters")).build()
        }

        if (id != order.orderID) {
            return Response.status(Response.Status.BAD_REQUEST)
                .entity(mapOf("error" to "Order ID in path and body must match")).build()
        }

        if (order.totalPrice != order.books.sumOf { it.totalPrice }) {
            return Response.status(Response.Status.BAD_REQUEST)
                .entity(mapOf("error" to "Order price doesn't match sum of books prices")).build()
        }

        if (!order.books.all { it.price * it.quantity == it.totalPrice }) {
            return Response.status(Response.Status.BAD_REQUEST)
                .entity(mapOf("error" to "Books price times quantity doesn't match it's total price")).build()
        }

        if (orderService.updateOrder(order)) {
            return Response.status(Response.Status.NO_CONTENT).build()
        }
        return Response.status(Response.Status.NOT_FOUND).entity(mapOf("error" to "Order not found")).build()
    }

    @DELETE
    @Path("/order/{id}")
    fun delete(@PathParam("id") id: String): Response {
        if (id.length != 36) {
            return Response.status(Response.Status.BAD_REQUEST)
                .entity(mapOf("error" to "Order ID must be 36 characters")).build()
        }

        if (orderService.deleteOrder(id)) {
            return Response.status(Response.Status.NO_CONTENT).build()
        }
        return Response.status(Response.Status.NOT_FOUND).entity(mapOf("error" to "Order not found")).build()
    }
}