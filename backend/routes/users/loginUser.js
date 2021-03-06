'use strict'
const axios = require('axios')
const jwt = require('jsonwebtoken')
const bcrypt = require('bcryptjs')
const urlJoin = require('url-join')

const { HttpError } = require('../../errors')
const env = require('../../env')
const DB_BASE_URL = env('DB_BASE_URL')

/**
 * @typedef ErrorResponse
 * @property {[integer]} statusCode
 * @property {[string]} message
 */

/**
 * @typedef LoginRequest
 * @property {{tester1@test.com}} username - username (email)
 * @property {{Test1234}} password - user's password
 */

/**
 * @route POST /api/users/login
 * @group users - Login existing user
 * @param {LoginRequest.model} LoginRequest.body.required - the new login request
 * @returns {LoginResponse.model} 200 - The User's information and session
 * @returns {ErrorResponse.model}  default - HttpError - User not found
 */
module.exports = function (req, res, next) {
  const userRequest = {
    'username': req.body.username,
    'password': req.body.password
  }

  let newRequest = {
    'username': userRequest.username,
    'password': userRequest.password
  }

  getUserObjectIfExists(res, newRequest, next)
}

function getUserObjectIfExists (res, user, next) {
  return axios.get(urlJoin(DB_BASE_URL, 'users?username=' + user.username))
    .then(function (response) {
      return hashUserData(res, response.data, user.password, next)
    })
    .catch(() => {
      next(new HttpError(400, 'Failed to login: bad username or password'))
    })
}

function hashUserData (res, user, pwAttempted, next) {
  const options = {}
  const userToken = {
    'accessToken': jwt.sign({ userId: user.userId }, 'cooper', Object.assign(options, { expiresIn: '2 hours' })),
    'tokenType': 'bearer',
    'expiresIn': 60 * 60 * 24,
    'refreshToken': jwt.sign({ userId: user.userId }, 'cooper', Object.assign(options, { expiresIn: '2 days' }))
  }

  // check pw against db stored pw
  if (bcrypt.compareSync(pwAttempted, user.password)) {
    user.accessToken = userToken.accessToken
    user.refreshToken = userToken.refreshToken
    user.tokenType = userToken.tokenType
    user.expiresIn = userToken.expiresIn
    return sendLoginUser(res, user, next)
  } else {
    next(new HttpError(400, 'Failed to login: bad username or password'))
  }
}

function sendLoginUser (res, user, next) {
  // Save this user to the database
  return axios.post(urlJoin(DB_BASE_URL, 'users', 'newSession'), user)
    .then(function (response) {
      return res.status(201).json({
        username: response.data.username,
        userId: response.data.userId,
        name: response.data.name,
        role: response.data.role,
        accessToken: response.data.accessToken,
        tokenType: response.data.tokenType,
        expiresIn: response.data.expiresIn,
        refreshToken: response.data.refreshToken
      })
    })
    .catch(function () {
      next(new HttpError(400, 'Failed to create new user session'))
    })
}
