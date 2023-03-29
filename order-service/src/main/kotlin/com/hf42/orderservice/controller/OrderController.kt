package com.hf42.orderservice.controller

import com.hf42.orderservice.model.Order
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

        val id = orderService.insertOrder(order)
        return if (id == null) {
            Response.status(Response.Status.INTERNAL_SERVER_ERROR)
                .entity(mapOf("error" to "Order was not created")).build()
        } else {
            Response.status(Response.Status.CREATED).entity(mapOf("id" to id)).build()
        }
    }

    @GET
    @Path("/order/{id}")
    fun get(@PathParam("id") id: String): Response {
        if (id.length != 36) {
            return Response.status(Response.Status.BAD_REQUEST)
                .entity(mapOf("error" to "Order ID must be 36 characters")).build()
        }

        val order = orderService.getOrder(id)

        return if (order == null) {
            Response.status(Response.Status.NOT_FOUND)
                .entity(mapOf("error" to "Order not found")).build()
        } else {
            Response.status(Response.Status.OK).entity(order).build()
        }
    }
}