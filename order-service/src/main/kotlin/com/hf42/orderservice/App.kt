package com.hf42.orderservice

import io.quarkus.runtime.Quarkus
import io.quarkus.runtime.QuarkusApplication

object App {

    @JvmStatic
    fun main(args: Array<String>) {
        Quarkus.run(Main::class.java, *args)
    }

    class Main : QuarkusApplication {
        @Throws(Exception::class)
        override fun run(vararg args: String): Int {
            println("Do startup logic here")
            Quarkus.waitForExit()
            return 0
        }
    }

}