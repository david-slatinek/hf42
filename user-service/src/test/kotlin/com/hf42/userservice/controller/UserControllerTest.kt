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

        val userAsJson: String = ObjectMapper().writeValueAsString(body)

        val result = mockMvc.perform(
            post("/register")
                .contentType(APPLICATION_JSON)
                .content(userAsJson)
                .accept(APPLICATION_JSON)
        ).andDo(MockMvcResultHandlers.print())
            .andExpect(status().isCreated).andReturn()

        val parsedResponse =
            ObjectMapper().readValue(result.response.contentAsByteArray, mutableMapOf<String, Any>()::class.java)

        mockMvc.perform(delete("/" + parsedResponse["id"]).contentType(APPLICATION_JSON).accept(APPLICATION_JSON))
            .andDo(MockMvcResultHandlers.print())

    }

    @Test
    fun login() {
    }

    @Test
    fun delete() {
    }
}