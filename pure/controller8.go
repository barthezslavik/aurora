package main

import (
	"errors"
	"fmt"
)

// User represents a user in the system.
type User struct {
	ID    int
	Age   int
	Posts []Post
}

// Post represents a post by a user.
type Post struct {
	UserID  int
	Content string
}

// UserProfileController provides methods to manipulate and retrieve user profiles.
type UserProfileController struct {
	users map[int]User
	posts []Post
}

// NewUserProfileController creates a new instance of UserProfileController.
func NewUserProfileController() *UserProfileController {
	return &UserProfileController{
		users: make(map[int]User),
		posts: make([]Post, 0),
	}
}

// GetProfile retrieves the profile of a user.
func (c *UserProfileController) GetProfile(userID int) (User, error) {
	user, err := c.GetUser(userID)
	if err != nil {
		return User{}, err
	}

	postStats, err := c.GetPostStats(userID)
	if err != nil {
		return User{}, err
	}

	fmt.Println("Post Stats: ", postStats) // Or handle the stats as needed

	return user, nil
}

// UpdateAge updates the age of a user.
func (c *UserProfileController) UpdateAge(userID int, newAge int) error {
	if !ValidateAge(newAge) {
		return errors.New("invalid age")
	}

	return c.UpdateUserAge(userID, newAge)
}

// GetPostStats calculates statistics for a user's posts.
func (c *UserProfileController) GetPostStats(userID int) (map[string]float64, error) {
	userPosts := c.FindAllPostsByUserId(userID)
	count := len(userPosts)
	avgLength := AvgLength(userPosts)

	return map[string]float64{"Count": float64(count), "AverageLength": avgLength}, nil
}

// GetUser retrieves a user by ID.
func (c *UserProfileController) GetUser(userID int) (User, error) {
	user, exists := c.users[userID]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

// ValidateAge checks if the given age is valid.
func ValidateAge(age int) bool {
	return age >= 0 && age <= 150
}

// UpdateUserAge updates the age of a user.
func (c *UserProfileController) UpdateUserAge(userID int, age int) error {
	user, exists := c.users[userID]
	if !exists {
		return errors.New("user not found")
	}
	user.Age = age
	c.users[userID] = user
	return nil
}

// AvgLength calculates the average length of posts.
func AvgLength(posts []Post) float64 {
	if len(posts) == 0 {
		return 0
	}
	totalLength := 0
	for _, post := range posts {
		totalLength += len(post.Content)
	}
	return float64(totalLength) / float64(len(posts))
}

// FindAllPostsByUserId finds all posts by a given user ID.
func (c *UserProfileController) FindAllPostsByUserId(userID int) []Post {
	var userPosts []Post
	for _, post := range c.posts {
		if post.UserID == userID {
			userPosts = append(userPosts, post)
		}
	}
	return userPosts
}

func main() {
	// Example usage
	controller := NewUserProfileController()
	// Add your code here to populate the controller with users and posts,
	// and to call its methods as needed.
	controller.users[1] = User{ID: 1, Age: 20}
	controller.posts = append(controller.posts, Post{UserID: 1, Content: "Hello World"})

	// GetProfile
	user, err := controller.GetProfile(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
