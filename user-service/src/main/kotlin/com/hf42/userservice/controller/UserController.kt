package com.hf42.userservice.controller

import com.hf42.userservice.model.Login
import com.hf42.userservice.model.User
import com.hf42.userservice.repository.UserRepository
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
    @PostMapping("/register")
    fun register(@Valid @RequestBody user: User): ResponseEntity<Map<String, String>> {
        if (userRepository.findByEmail(user.email) != null) {
            return ResponseEntity(mapOf("error" to "User already exists"), HttpStatus.BAD_REQUEST);
        }

        user.password = BCrypt.hashpw(user.password, BCrypt.gensalt(15));
        val obj = userRepository.save(user)
        return ResponseEntity(mapOf("id" to obj.id.toString()), HttpStatus.CREATED);
    }

    @PostMapping("/login")
    fun login(@Valid @RequestBody login: Login): ResponseEntity<Map<String, String>> {
        val user = userRepository.findByEmail(login.email)
            ?: return ResponseEntity(mapOf("error" to "invalid email or password"), HttpStatus.BAD_REQUEST)

        if (!BCrypt.checkpw(login.password, user.password)) {
            return ResponseEntity(mapOf("error" to "invalid email or password"), HttpStatus.BAD_REQUEST)
        }

        return ResponseEntity(mapOf("message" to "login successful"), HttpStatus.OK);
    }

    @DeleteMapping("/{userId}")
    fun delete(@PathVariable userId: String): ResponseEntity<Map<String, String>> {
        userRepository.deleteById(userId)
        return ResponseEntity(mapOf("message" to "user deleted"), HttpStatus.OK);
    }
}