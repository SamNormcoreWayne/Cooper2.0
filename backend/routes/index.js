'use strict'

const login = require('./users/loginUser')
const registration = require('./users/createUser')
const getUser = require('./users/getUser')
const getCourses = require('./courses/getCourses')
const registerCourse = require('./courses/registerCourse')
const getUsersCourses = require('./users/getCoursesForUser')
const unregisterUserFromCourse = require('./users/unregisterFromCourse')

const router = require('express-promise-router')()
const category = require('./users/association/associateCategory')
const major = require('./users/association/associateMajor')
const range = require('./users/association/associateRange')
const institute = require('./users/association/associateInstitute')
const checkAccessToken = require('../middleware/checkAccessToken')

// User Routes
router.get('/api/users/getUser', checkAccessToken, getUser)
router.post('/api/users/register', registration)
router.post('/api/users/login', login)
router.get('/api/users/courses', checkAccessToken, getUsersCourses)
router.get('/api/users/courses/unregister', checkAccessToken, unregisterUserFromCourse)

router.get('/api/courses', checkAccessToken, getCourses)
router.post('/api/courses/register', checkAccessToken, registerCourse)

router.post('/api/users/association/range', range.postRange)
router.post('/api/users/association/category', category.postCategory)
router.post('/api/users/association/major', major.postMajor)
router.post('/api/users/association/institute', institute.postInstitute)

module.exports = router
