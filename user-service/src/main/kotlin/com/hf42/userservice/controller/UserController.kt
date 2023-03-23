package com.hf42.userservice.controller

import com.hf42.userservice.model.User
import com.hf42.userservice.repository.UserRepository
import jakarta.validation.Valid
import org.mindrot.jbcrypt.BCrypt
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController


@RestController
@RequestMapping("/user")
class UserController(private val userRepository: UserRepository) {
    @PostMapping("/register")
    fun register(@Valid @RequestBody user: User): ResponseEntity<Map<String, String>> {
        user.password = BCrypt.hashpw(user.password, BCrypt.gensalt(15));
        val obj = userRepository.save(user)
        return ResponseEntity(mapOf("id" to obj.id.toString()), HttpStatus.CREATED);
    }
}