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

@Path("/api")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
class OrderController {
    @Inject
    lateinit var orderService: OrderService

    @POST
    @Path("/order")
    fun add(order: Order): Response {
        val id = orderService.getCollection()?.insertOne(order)
        return Response.status(Response.Status.CREATED).build()
    }
}