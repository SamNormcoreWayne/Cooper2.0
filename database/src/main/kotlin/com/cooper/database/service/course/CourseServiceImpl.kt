package com.cooper.database.service.course

import com.cooper.database.error.ObjectNotCreated
import com.cooper.database.error.ObjectNotDeleted
import com.cooper.database.error.ObjectNotFound
import com.cooper.database.model.Course
import com.cooper.database.repository.CourseRepository
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.stereotype.Service
import java.util.*


@Service
class CourseServiceImpl : CourseService {

    @Autowired
    private val courseRepository: CourseRepository? = null

    override fun listAllCourses(): List<Course> {
        val all = ArrayList<Course>()

        try {
            courseRepository?.let { repository ->
                for (session in repository.findAll()) {
                    all.add(session)
                }
            }
        } catch (err: Exception) {
            throw ObjectNotFound(message = "Could not find any courses")
        }

        return all
    }

    override fun findById(id: Long?): Course? {
        id?.let { courseId ->
            try {
                return this.courseRepository?.findById(courseId)?.get()
            } catch (err: Exception) {
                throw ObjectNotFound(message = "Could not find course with id: $id")
            }
        } ?: run {
            throw ObjectNotFound(message = "Could not find course with id: $id")
        }
    }

    override fun findByTitle(title: String?): List<Course>? {
        title?.let { courseTitle ->
            try {
                return this.courseRepository?.findByTitle(courseTitle)
            } catch (err: Exception) {
                throw ObjectNotFound(message = "Could not find any course with title: $title")
            }
        } ?: run {
            throw ObjectNotFound(message = "Could not find any course with id: $title")
        }
    }

    override fun findCoursesByUserId(id: String?): List<Course>? {
        id?.let { userId ->
            try {
                return this.courseRepository?.findByUserId(userId)
            } catch (err: Exception) {
                throw ObjectNotFound(message = "Could not find any courses for user: $userId")
            }
        } ?: run {
            throw ObjectNotFound(message = "Could not find any courses for user: $id")
        }
    }

    override fun findBySection(section: String?): List<Course>? {
        section?.let { courseSection ->
            try {
                return this.courseRepository?.findBySection(courseSection)
            } catch (err: Exception) {
                throw ObjectNotFound(message = "Could not find any course with section: $courseSection")
            }
        } ?: run {
            throw ObjectNotFound(message = "Could not find any course with section: $section")
        }
    }

    override fun unregister(courseId: Long) {
        try {
            this.courseRepository?.deleteById(courseId)
        } catch (err: Exception) {
            throw ObjectNotDeleted(
                    status = HttpStatus.BAD_REQUEST,
                    message = "Could not unregister from course: $courseId"
            )
        }
    }

    override fun create(course: Course): Course? {
        try {
            return this.courseRepository?.save(course)
        } catch (err: Exception) {
            throw ObjectNotCreated(
                    status = HttpStatus.BAD_REQUEST,
                    message = "Could not create course: ${course.title}"
            )
        }
    }
}
