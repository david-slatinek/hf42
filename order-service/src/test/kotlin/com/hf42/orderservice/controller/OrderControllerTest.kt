package com.hf42.orderservice.controller

import com.fasterxml.jackson.databind.ObjectMapper
import io.quarkus.test.junit.QuarkusTest
import io.restassured.RestAssured.given
import org.hamcrest.core.StringContains.containsString
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Assertions.assertTrue
import org.junit.jupiter.api.Test
import java.time.LocalDateTime
import javax.ws.rs.core.MediaType.APPLICATION_JSON


@QuarkusTest
class OrderControllerTest {
    @Test
    fun testAdd() {
        val body: MutableMap<String, Any> = HashMap()
        body["customerID"] = "641d4e3fea22771984f10689"
        body["orderDate"] = "2023-03-29 12:00"

        body["books"] = MutableList(1) {
            val book: MutableMap<String, Any> = HashMap()
            book["isbn"] = "978-961-7083-13-2"
            book["title"] = "Test"
            book["author"] = "Test"
            book["year"] = LocalDateTime.now().year
            book["description"] = "Test"
            book["categories"] = listOf("Test1", "Test2")
            book["size"] = "Test"
            book["weight"] = "Test"
            book["pages"] = 42
            book["publisher"] = "Test"
            book["language"] = "Test"
            book["quantity"] = 2
            book["price"] = 10.0
            book["totalPrice"] = 20
            book
        }
        body["totalPrice"] = 20
        body["status"] = "test"

        val addResponse = given()
            .`when`()
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            .body(body)
            .post("/order")
            .then()
            .log().all()
            .statusCode(201)
            .contentType(APPLICATION_JSON)
            .body(containsString("orderID"))
            .extract().response()

        val registerResponse =
            ObjectMapper().readValue(addResponse.asInputStream(), mutableMapOf<String, Any>()::class.java)

        assertTrue(registerResponse.containsKey("orderID"))
        assertEquals(36, registerResponse["orderID"].toString().length)

        given()
            .`when`()
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            .delete("/order/${registerResponse["orderID"]}").then().log().all()
    }

    @Test
    fun testGet() {
        val body: MutableMap<String, Any> = HashMap()
        body["customerID"] = "641d4e3fea22771984f10689"
        body["orderDate"] = "2023-03-29 12:00"

        body["books"] = MutableList(1) {
            val book: MutableMap<String, Any> = HashMap()
            book["isbn"] = "978-961-7083-13-2"
            book["title"] = "Test"
            book["author"] = "Test"
            book["year"] = LocalDateTime.now().year
            book["description"] = "Test"
            book["categories"] = listOf("Test1", "Test2")
            book["size"] = "Test"
            book["weight"] = "Test"
            book["pages"] = 42
            book["publisher"] = "Test"
            book["language"] = "Test"
            book["quantity"] = 2
            book["price"] = 10.0
            book["totalPrice"] = 20
            book
        }
        body["totalPrice"] = 20
        body["status"] = "test"

        val addResponse = given()
            .`when`()
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            .body(body)
            .post("/order")
            .then()
            .log().all()
            .statusCode(201)
            .contentType(APPLICATION_JSON)
            .body(containsString("orderID"))
            .extract().response()

        val registerResponse =
            ObjectMapper().readValue(addResponse.asInputStream(), mutableMapOf<String, Any>()::class.java)

        assertTrue(registerResponse.containsKey("orderID"))
        assertEquals(36, registerResponse["orderID"].toString().length)

        given()
            .`when`()
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            .get("/order/${registerResponse["orderID"]}")
            .then()
            .log().all()
            .statusCode(200)
            .contentType(APPLICATION_JSON)
            .body(containsString("orderID"))
            .body(containsString("customerID"))
            .body(containsString("orderDate"))
            .body(containsString("books"))
            .body(containsString("totalPrice"))

        given()
            .`when`()
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            .delete("/order/${registerResponse["orderID"]}").then().log().all()
    }

    @Test
    fun testUpdate() {
        val body: MutableMap<String, Any> = HashMap()
        body["customerID"] = "641d4e3fea22771984f10689"
        body["orderDate"] = "2023-03-29 12:00"

        body["books"] = MutableList(1) {
            val book: MutableMap<String, Any> = HashMap()
            book["isbn"] = "978-961-7083-13-2"
            book["title"] = "Test"
            book["author"] = "Test"
            book["year"] = LocalDateTime.now().year
            book["description"] = "Test"
            book["categories"] = listOf("Test1", "Test2")
            book["size"] = "Test"
            book["weight"] = "Test"
            book["pages"] = 42
            book["publisher"] = "Test"
            book["language"] = "Test"
            book["quantity"] = 2
            book["price"] = 10.0
            book["totalPrice"] = 20
            book
        }
        body["totalPrice"] = 20
        body["status"] = "test"

        val addResponse = given()
            .`when`()
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            .body(body)
            .post("/order")
            .then()
            .log().all()
            .statusCode(201)
            .contentType(APPLICATION_JSON)
            .body(containsString("orderID"))
            .extract().response()

        val registerResponse =
            ObjectMapper().readValue(addResponse.asInputStream(), mutableMapOf<String, Any>()::class.java)

        assertTrue(registerResponse.containsKey("orderID"))
        assertEquals(36, registerResponse["orderID"].toString().length)

        val updateBody: MutableMap<String, Any> = HashMap()
        updateBody["customerID"] = "641d4e3fea22771984f10689"
        updateBody["orderDate"] = "2023-03-29 12:00"
        updateBody["orderID"] = registerResponse["orderID"].toString()

        updateBody["books"] = MutableList(1) {
            val book: MutableMap<String, Any> = HashMap()
            book["isbn"] = "978-961-7083-13-2"
            book["title"] = "Test"
            book["author"] = "Test"
            book["year"] = LocalDateTime.now().year
            book["description"] = "Test - updated"
            book["categories"] = listOf("Test1", "Test2")
            book["size"] = "Test"
            book["weight"] = "Test"
            book["pages"] = 42
            book["publisher"] = "Test"
            book["language"] = "Test"
            book["quantity"] = 2
            book["price"] = 10.0
            book["totalPrice"] = 20
            book
        }
        updateBody["totalPrice"] = 20
        updateBody["status"] = "test"

        given()
            .`when`()
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            .body(updateBody)
            .put("/order/${registerResponse["orderID"]}")
            .then()
            .log().all()
            .statusCode(204)

        given()
            .`when`()
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            .delete("/order/${registerResponse["orderID"]}").then().log().all()
    }

    @Test
    fun testDelete() {
        val body: MutableMap<String, Any> = HashMap()
        body["customerID"] = "641d4e3fea22771984f10689"
        body["orderDate"] = "2023-03-29 12:00"

        body["books"] = MutableList(1) {
            val book: MutableMap<String, Any> = HashMap()
            book["isbn"] = "978-961-7083-13-2"
            book["title"] = "Test"
            book["author"] = "Test"
            book["year"] = LocalDateTime.now().year
            book["description"] = "Test"
            book["categories"] = listOf("Test1", "Test2")
            book["size"] = "Test"
            book["weight"] = "Test"
            book["pages"] = 42
            book["publisher"] = "Test"
            book["language"] = "Test"
            book["quantity"] = 2
            book["price"] = 10.0
            book["totalPrice"] = 20
            book
        }
        body["totalPrice"] = 20
        body["status"] = "test"

        val addResponse = given()
            .`when`()
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            .body(body)
            .post("/order")
            .then()
            .log().all()
            .statusCode(201)
            .contentType(APPLICATION_JSON)
            .body(containsString("orderID"))
            .extract().response()

        val registerResponse =
            ObjectMapper().readValue(addResponse.asInputStream(), mutableMapOf<String, Any>()::class.java)

        assertTrue(registerResponse.containsKey("orderID"))
        assertEquals(36, registerResponse["orderID"].toString().length)

        given()
            .`when`()
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            .delete("/order/${registerResponse["orderID"]}")
            .then()
            .log().all()
            .statusCode(204)

        given()
            .`when`()
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            .get("/order/${registerResponse["orderID"]}")
            .then()
            .log().all()
            .statusCode(404)
    }
}