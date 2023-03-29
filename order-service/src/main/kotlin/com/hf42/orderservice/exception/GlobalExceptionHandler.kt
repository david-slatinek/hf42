package com.hf42.orderservice.exception

import com.fasterxml.jackson.databind.exc.ValueInstantiationException
import javax.ws.rs.core.Response
import javax.ws.rs.ext.ExceptionMapper
import javax.ws.rs.ext.Provider

@Provider
class GlobalExceptionHandler : ExceptionMapper<ValueInstantiationException> {
    override fun toResponse(exception: ValueInstantiationException?): Response {
        exception.let { e ->
            return Response.status(Response.Status.BAD_REQUEST)
                .entity(mapOf("error" to e?.message)).build()
        }
    }
}