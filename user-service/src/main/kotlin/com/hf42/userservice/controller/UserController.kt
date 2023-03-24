package com.hf42.userservice.controller

import com.hf42.userservice.model.Login
import com.hf42.userservice.model.User
import com.hf42.userservice.repository.UserRepository
import io.swagger.v3.oas.annotations.Operation
import io.swagger.v3.oas.annotations.media.Content
import io.swagger.v3.oas.annotations.responses.ApiResponse
import io.swagger.v3.oas.annotations.responses.ApiResponses
import jakarta.validation.Valid
import org.mindrot.jbcrypt.BCrypt
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.*

@RestController
@RequestMapping(
    "/",
    consumes = ["application/json"],
    produces = ["application/json"],
    headers = ["Content-Type=application/json"]
)

class UserController(private val userRepository: UserRepository) {
    @Operation(summary = "Register a new user")
    @ApiResponses(
        value = [
            ApiResponse(
                responseCode = "201",
                description = "Created",
                content = [Content(
                    mediaType = "application/json",
                    schema = io.swagger.v3.oas.annotations.media.Schema(implementation = Map::class)
                )]
            ),
            ApiResponse(
                responseCode = "400", description = "Bad Request", content = [Content(
                    mediaType = "application/json",
                    schema = io.swagger.v3.oas.annotations.media.Schema(implementation = Map::class)
                )]
            ),
        ]
    )
    @PostMapping("/register")
    fun register(@Valid @RequestBody user: User): ResponseEntity<Map<String, String>> {
        if (userRepository.findByEmail(user.email) != null) {
            return ResponseEntity(mapOf("error" to "User already exists"), HttpStatus.BAD_REQUEST);
        }

        user.password = BCrypt.hashpw(user.password, BCrypt.gensalt(15));
        val obj = userRepository.save(user)
        return ResponseEntity(mapOf("id" to obj.id.toString()), HttpStatus.CREATED);
    }

    @Operation(summary = "Login a user")
    @ApiResponses(
        value = [
            ApiResponse(
                responseCode = "200", description = "OK", content = [Content(
                    mediaType = "application/json",
                    schema = io.swagger.v3.oas.annotations.media.Schema(implementation = Map::class)
                )]
            ),
            ApiResponse(
                responseCode = "400", description = "Bad Request", content = [Content(
                    mediaType = "application/json",
                    schema = io.swagger.v3.oas.annotations.media.Schema(implementation = Map::class)
                )]
            ),
        ]
    )
    @PostMapping("/login")
    fun login(@Valid @RequestBody login: Login): ResponseEntity<Map<String, String>> {
        val user = userRepository.findByEmail(login.email)
            ?: return ResponseEntity(mapOf("error" to "invalid email or password"), HttpStatus.BAD_REQUEST)

        if (!BCrypt.checkpw(login.password, user.password)) {
            return ResponseEntity(mapOf("error" to "invalid email or password"), HttpStatus.BAD_REQUEST)
        }

        return ResponseEntity(mapOf("message" to "login successful"), HttpStatus.OK);
    }

    @Operation(summary = "Delete a user by id")
    @ApiResponses(
        value = [
            ApiResponse(
                responseCode = "200", description = "OK", content = [Content(
                    mediaType = "application/json",
                    schema = io.swagger.v3.oas.annotations.media.Schema(implementation = Map::class)
                )]
            ),
        ]
    )
    @DeleteMapping("/{userId}")
    fun delete(@PathVariable userId: String): ResponseEntity<Map<String, String>> {
        userRepository.deleteById(userId)
        return ResponseEntity(mapOf("message" to "user deleted"), HttpStatus.OK);
    }
}