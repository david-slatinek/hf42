package com.hf42.orderservice.controller

import com.hf42.orderservice.model.Order
import com.hf42.orderservice.service.OrderService
import javax.inject.Inject
import javax.ws.rs.Consumes
import javax.ws.rs.POST
import javax.ws.rs.Path
import javax.ws.rs.Produces
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

        val id = orderService.insertOrder(order)
        return if (id == null) {
            Response.status(Response.Status.INTERNAL_SERVER_ERROR)
                .entity(mapOf("error" to "Order was not created")).build()
        } else {
            Response.status(Response.Status.CREATED).entity(mapOf("id" to id)).build()
        }
    }
}