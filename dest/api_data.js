define({ "api": [
  {
    "type": "get",
    "url": "/index/:user_id",
    "title": "Request User Information",
    "version": "0.1.0",
    "name": "Index",
    "group": "Index",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "description": "<p>API to get the user information.</p>",
    "examples": [
      {
        "title": "Example usage:",
        "content": "curl -i http://localhost:5000/users/2",
        "type": "json"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "user_id",
            "description": "<p>The user's unique ID.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name of the User.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"name\": \"Tom\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UserNotFound",
            "description": "<p>The <code>user_id&lt; /code&gt; of the User was not found.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n    \"error\": \"UserNotFound\",\n    \"message\": \"User {user_id} doesn't exist\"\n}",
          "type": "json"
        }
      ]
    },
    "sampleRequest": [
      {
        "url": "http://localhost:5000/users/:user_id"
      }
    ],
    "filename": "src/index.js",
    "groupTitle": "Index"
  },
  {
    "type": "get",
    "url": "/users/:user_id",
    "title": "Request User Information",
    "version": "0.1.0",
    "name": "GetUser",
    "group": "User",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "description": "<p>API to get the user information.</p>",
    "examples": [
      {
        "title": "Example usage:",
        "content": "curl -i http://localhost:5000/users/2",
        "type": "json"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "user_id",
            "description": "<p>The user's unique ID.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name of the User.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"name\": \"Tom\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UserNotFound",
            "description": "<p>The <code>user_id&lt; /code&gt; of the User was not found.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n    \"error\": \"UserNotFound\",\n    \"message\": \"User {user_id} doesn't exist\"\n}",
          "type": "json"
        }
      ]
    },
    "sampleRequest": [
      {
        "url": "http://localhost:5000/users/:user_id"
      }
    ],
    "filename": "src/user.js",
    "groupTitle": "User"
  }
] });