# Final Project


## Final Project API Documentation
Run: go run cmd/server/main.go  
URL: 127.0.0.1:8000  
All API require a bearer token. (Except Register & Login)

-------------------------

### Users API
URL: 127.0.0.1:8000/users

-------------------------

#### Register User
| Description | URL |
| --- | -------- |
| Register User | 127.0.0.1:8000/users/register |
| Method | POST |
##### Input
```json
{
	"username": "user1",
	"email": "user@email.com",
	"password": "password1",
	"age": 22
}
```

##### Output
```json
{
    "username": "user1",
	"email": "user@email.com",
	"password": "password1",
	"age": 22
}
```

#### User Login
| Description | URL |
| --- | -------- |
| User Login | 127.0.0.1:8000/users/login |
| Method | POST |

##### Input
```json
{
    "email": "user@email.com",
 	"password": "password"
}
```

##### Output
```json
{
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NTg5MTEwNjExNTksImV4cCI6MTY1ODkxMTEyMTE1OSwidWlkIjoxfQ.-sZ1McGjIC-VNFOAAFutI621CPFlHRH5-ABRnraelTk"
}
```

#### Update User
| Description | URL |
| --- | -------- |
| Update User | 127.0.0.1:8000/users |
| Method | PUT |

##### Input
```json
{
	"username": "user1",
	"email": "user@email.com"
}
```

##### Output
```json
{
    "username": "user1",
	"email": "user@email.com",
	"password": "password1",
	"age": 22,
    "updated_at":"2022-08-03T11:44:56.25Z"
}
```

#### Delete User
| Description | URL |
| --- | -------- |
| Delete User | 127.0.0.1:8000/users |
| Method | DELETE |

##### Output
```json
{
    "message": "Your account has been successfully deleted"
}
```

-------------------------

### Photos API
URL: 127.0.0.1:8000/photos

-------------------------

#### Post Photo
| Description | URL |
| --- | -------- |
| Post Photo | 127.0.0.1:8000/photos |
| Method | POST |

##### Input
```json
{
	"title": "title photo",
 	"caption": "caption photo",
 	"photo_url": "https://photo.domain.com"
}
```

##### Output
```json
{
    "id": 1,
	"title": "title photo",
 	"caption": "caption photo",
 	"photo_url": "https://photo.domain.com",
    "user_id": 1,
    "created_at":"2022-08-03T11:44:56.25Z"
}
```

#### Get Photos
| Description | URL |
| --- | -------- |
| Get Photos | 127.0.0.1:8000/photos |
| Method | GET |

##### Output
```json
[
    {
        "id": 1,
        "title": "title photo",
        "caption": "caption photo",
        "photo_url": "https://photo.domain.com",
        "user_id": 1,
        "created_at":"2022-08-03T11:44:56.25Z",
        "updated_at":"2022-08-03T11:44:56.25Z",
        "user":{
            "username": "user1",
	        "email": "user@email.com"
        }
    },
    {
        "id": 2,
        "title": "title photo",
        "caption": "caption photo",
        "photo_url": "https://photo.domain.com",
        "user_id": 2,
        "created_at":"2022-08-03T11:44:56.25Z",
        "updated_at":"2022-08-03T11:44:56.25Z",
        "user":{
            "username": "user2",
	        "email": "user2@email.com"
        }
    }
]
```

#### Update Photo
| Description | URL |
| --- | -------- |
| Update Photo | 127.0.0.1:8000/photos/{id} |
| Method | PUT |

##### Input
```json
{
	"title": "title photo",
 	"caption": "caption photo",
 	"photo_url": "https://photo.domain.com"
}
```

##### Output
```json
{
    "id": 1,
	"title": "title photo",
 	"caption": "caption photo",
 	"photo_url": "https://photo.domain.com",
    "user_id": 1,
    "updated_at":"2022-08-03T11:44:56.25Z"
}
```

#### Delete Photo
| Description | URL |
| --- | -------- |
| Delete Photo | 127.0.0.1:8000/photos/{id} |
| Method | DELETE |

##### Output
```json
{
    "message": "Your photo has been successfully deleted"
}
```

-------------------------

### Comments API
URL: 127.0.0.1:8000/comments

-------------------------

#### Post Comment
| Description | URL |
| --- | -------- |
| Post Comment | 127.0.0.1:8000/comments |
| Method | POST |

##### Input
```json
{
	"message": "comment message",
 	"photo_id": 1
}
```

##### Output
```json
{
    "id": 1,
	"message": "message comment",
 	"photo_id": 1,
 	"user_id": 1,
    "created_at":"2022-08-03T11:44:56.25Z"
}
```

#### Get Comments
| Description | URL |
| --- | -------- |
| Get Comments | 127.0.0.1:8000/comments |
| Method | GET |

##### Output
```json
[
    {
        "id": 1,
        "message": "message comment",
        "photo_id": 1,
        "user_id": 1,
        "updated_at":"2022-08-03T11:44:56.25Z",
        "created_at":"2022-08-03T11:44:56.25Z",
        "user":{
            "id":1,
            "username": "user1",
            "email": "user@email.com"
        },
        "photo":{
            "id": 1,
            "title": "title photo",
            "caption": "caption photo",
            "photo_url": "https://photo.domain.com",
            "user_id": 1
        }
    },
    {
        "id": 2,
        "message": "message comment",
        "photo_id": 2,
        "user_id": 2,
        "updated_at":"2022-08-03T11:44:56.25Z",
        "created_at":"2022-08-03T11:44:56.25Z",
        "user":{
            "id":2,
            "username": "user2",
            "email": "user2@email.com"
        },
        "photo":{
            "id": 2,
            "title": "title photo",
            "caption": "caption photo",
            "photo_url": "https://photo.domain.com",
            "user_id": 2
        }
    }
]
```

#### Update Comment
| Description | URL |
| --- | -------- |
| Update Comment | 127.0.0.1:8000/comments/{id} |
| Method | PUT |

##### Input
```json
{
	"message": "comment message"
}
```

##### Output
```json
{
    "id": 1,
	"message": "message comment",
 	"photo_id": 1,
 	"user_id": 1,
    "updated_at":"2022-08-03T11:44:56.25Z"
}
```

#### Delete Comment
| Description | URL |
| --- | -------- |
| Delete Comment | 127.0.0.1:8000/comments/{id} |
| Method | DELETE |

##### Output
```json
{
	"message": "Your photo has been successfully deleted"
}
```

-------------------------

### Social Medias API
URL: 127.0.0.1:8000/comments

-------------------------

#### Post Social Media
| Description | URL |
| --- | -------- |
| Post Social Media | 127.0.0.1:8000/socialmedias |
| Method | POST |

##### Input
```json
{
	"name": "social media name",
 	"social_media_url": "https://domainsocialmedia.com/user",
 	"profile_image_url": "https://domainsocialmedia.com/userimage.jpg"
}
```

##### Output
```json
{
    "id": 1,
	"name": "social media name",
 	"social_media_url": "https://domainsocialmedia.com/user",
 	"user_id": 1,
    "created_at":"2022-08-03T11:44:56.25Z"
}
```

#### Get Social Medias
| Description | URL |
| --- | -------- |
| Get Social Medias | 127.0.0.1:8000/socialmedias |
| Method | GET |

##### Output
```json
[
    {
         "id": 1,
	    "name": "social media name",
 	    "social_media_url": "https://domainsocialmedia.com/user",
        "user_id": 1,
        "updated_at":"2022-08-03T11:44:56.25Z",
        "created_at":"2022-08-03T11:44:56.25Z",
        "user":{
            "id":1,
            "username": "user1",
            "profile_image_url": "https://domainsocialmedia.com/userimage.jpg"
        }
    },
    {
        "id": 1,
	    "name": "social media name",
 	    "social_media_url": "https://domainsocialmedia.com/user",
        "user_id": 2,
        "updated_at":"2022-08-03T11:44:56.25Z",
        "created_at":"2022-08-03T11:44:56.25Z",
        "user":{
            "id":2,
            "username": "user2",
            "profile_image_url": "https://domainsocialmedia.com/userimage.jpg"
        }
    }
]
```

#### Update Social Media
| Description | URL |
| --- | -------- |
| Update Social Media | 127.0.0.1:8000/socialmedias/{id} |
| Method | PUT |

##### Input
```json
{
	"name": "social media name",
 	"social_media_url": "https://domainsocialmedia.com/user"
}
```

##### Output
```json
{
    "id": 1,
	"name": "social media name",
 	"social_media_url": "https://domainsocialmedia.com/user",
 	"user_id": 1,
    "updated_at":"2022-08-03T11:44:56.25Z"
}
```

#### Delete Social media
| Description | URL |
| --- | -------- |
| Delete Social media | 127.0.0.1:8000/socialmedias/{id} |
| Method | DELETE |

##### Output
```json
{
	"message": "Your social media has been successfully deleted"
}
```