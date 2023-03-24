package com.hf42.userservice.controller

import com.fasterxml.jackson.databind.ObjectMapper
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.http.MediaType.APPLICATION_JSON
import org.springframework.test.web.servlet.MockMvc
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders.delete
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post
import org.springframework.test.web.servlet.result.MockMvcResultHandlers
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.status
import java.io.IOException

@SpringBootTest
@AutoConfigureMockMvc
class UserControllerTest {
    @Autowired
    lateinit var mockMvc: MockMvc

    @Throws(IOException::class)
    @Test
    fun register() {
        val body: MutableMap<String, Any> = HashMap()
        body["firstName"] = "test"
        body["lastName"] = "test"
        body["email"] = "test@email.com"
        body["password"] = "test_test_test"
        body["streetAddress"] = "test"
        body["city"] = "test"
        body["postOfficeNumber"] = 1000

        val registerResult = mockMvc.perform(
            post("/register")
                .contentType(APPLICATION_JSON)
                .content(ObjectMapper().writeValueAsString(body))
                .accept(APPLICATION_JSON)
        ).andDo(MockMvcResultHandlers.print())
            .andExpect(status().isCreated).andReturn()

        val registerResponse =
            ObjectMapper().readValue(
                registerResult.response.contentAsByteArray,
                mutableMapOf<String, Any>()::class.java
            )

        assertTrue(registerResponse.containsKey("id"))
        assertEquals(24, registerResponse["id"].toString().length)

        mockMvc.perform(delete("/" + registerResponse["id"]).contentType(APPLICATION_JSON).accept(APPLICATION_JSON))
            .andDo(MockMvcResultHandlers.print())
    }

    @Test
    fun login() {
        val body: MutableMap<String, Any> = HashMap()
        body["firstName"] = "test"
        body["lastName"] = "test"
        body["email"] = "test@email.com"
        body["password"] = "test_test_test"
        body["streetAddress"] = "test"
        body["city"] = "test"
        body["postOfficeNumber"] = 1000

        val registerResult = mockMvc.perform(
            post("/register")
                .contentType(APPLICATION_JSON)
                .content(ObjectMapper().writeValueAsString(body))
                .accept(APPLICATION_JSON)
        ).andDo(MockMvcResultHandlers.print()).andReturn()

        val registerResponse =
            ObjectMapper().readValue(
                registerResult.response.contentAsByteArray,
                mutableMapOf<String, Any>()::class.java
            )

        val loginBody: MutableMap<String, Any> = HashMap()
        loginBody["email"] = body["email"] as Any
        loginBody["password"] = body["password"] as Any

        val loginResult = mockMvc.perform(
            post("/login")
                .contentType(APPLICATION_JSON)
                .content(ObjectMapper().writeValueAsString(loginBody))
                .accept(APPLICATION_JSON)
        ).andDo(MockMvcResultHandlers.print())
            .andExpect(status().isOk).andReturn()

        val loginResponse =
            ObjectMapper().readValue(loginResult.response.contentAsByteArray, mutableMapOf<String, Any>()::class.java)

        assertTrue(loginResponse.containsKey("message"))
        assertEquals("login successful", loginResponse["message"])

        mockMvc.perform(delete("/" + registerResponse["id"]).contentType(APPLICATION_JSON).accept(APPLICATION_JSON))
    }

    @Test
    fun delete() {

    }
}