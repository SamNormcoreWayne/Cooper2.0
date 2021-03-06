package com.cooper.database.error

import org.springframework.http.HttpStatus
import org.springframework.web.server.ResponseStatusException

class ObjectNotDeleted(status: HttpStatus, override val message: String) : ResponseStatusException(status, message)
