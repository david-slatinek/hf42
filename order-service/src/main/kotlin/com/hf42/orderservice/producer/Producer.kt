package com.hf42.orderservice.producer

import com.rabbitmq.client.Channel
import com.rabbitmq.client.Connection
import com.rabbitmq.client.ConnectionFactory
import org.eclipse.microprofile.config.inject.ConfigProperty
import javax.enterprise.context.ApplicationScoped


@ApplicationScoped
class Producer(
    @ConfigProperty(name = "rabbitmq.host") host: String,
    @ConfigProperty(name = "rabbitmq.username") username: String,
    @ConfigProperty(name = "rabbitmq.password") password: String,
    @ConfigProperty(name = "rabbitmq.port") port: Int,
    @ConfigProperty(name = "rabbitmq.queue") queue: String,
    @ConfigProperty(name = "rabbitmq.exchange") var exchange: String
) {
    private final var factory: ConnectionFactory = ConnectionFactory()
    private final var connection: Connection
    private final var channel: Channel

    init {
        factory.host = host
        factory.username = username
        factory.password = password
        factory.port = port
        connection = factory.newConnection()
        channel = connection.createChannel()
        channel.exchangeDeclare(exchange, "fanout", true)
        channel.queueDeclare(queue, true, false, false, null)
    }

    fun produce(message: String) {
        channel.basicPublish(exchange, "", null, message.toByteArray())
    }
}