package models

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"time"
)

type BlogPost struct {
	BPID         string
	Author       *User
	Title        string
	Content      string
	PostedDate   time.Time
	Commentaries []*Commentary
}

// CreateBlogPost: creates and add it to redis
func CreateBlogPost(user *User, title string, content string) (*BlogPost, error) {
	if title != "" && content != "" && user != nil {
		bpid, err := uuid.NewV4()
		if err == nil {
			blogPost := &BlogPost{
				BPID:         bpid,
				Author:       user,
				Title:        title,
				Content:      content,
				PostedDate:   time.Now(),
				Commentaries: nil,
			}

			conn := RedisPool.Get()
			defer c.Close()

			_, err := conn.Do("SADD", "BlogPost", blogPost.BPID)
			if err == nil {
				_, err := conn.Do("HMSET", bloPost.BPID,
					"Title", blogPost.Title,
					"Author", blogPost.Author,
					"Content", blogPost.Content,
					"PostedDate", blogPost.CreationDate,
					"Commentaries", blogPost.Commentaries)
				if err == nil {
					return blogPost, nil
				}
				return nil, fmt.Errorf("Cannot set the blogPost")
			}
			return nil, fmt.Errorf("Cannot add the blogPost")
		}
		return nil, fmt.Errorf("Cannot create BPID")
	}
	return nil, fmt.Errorf("Title or content or user missing")
}

// GetBlogPost: return blogPost if exists
func GetBlogPost(bpid string) (*BlogPost, error) {
	if bpid != "" {
		conn := RedisPool.Get()
		defer conn.Close()

		ifExist, err := redis.Int(conn.Do("SISMEMBER", "BlogPost", bpid))
		if err == nil && ifExist == 1 {
			title, author, content, postedDate, commentaries, err := redis.String(conn.Do("HMGET", bpid, "Title", "Author", "Content", "PostedDate", "Commentaries"))
			if err == nil {
				blogPost := &BlogPos{
					BPID:         bpid,
					Title:        title,
					Author:       author,
					Content:      content,
					PostedDate:   postedDate,
					Commentaries: commentaries,
				}
				return blogPost, nil
			}
			return nil, fmt.Errorf("Cannot get the information")
		}
		return nil, fmt.Errorf("blogPost not found")
	}
	return nil, fmt.Errorf("id missing")
}

// GetAllBlogPost: return all blogPost
func GetAllBlogPost() ([]*BlogPost, error) {}

func EditBlogPost() {}

func DeleteBlogPost() {}
