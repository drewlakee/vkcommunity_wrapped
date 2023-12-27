package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"vkcommunity_wrapped/internal/models"
	"vkcommunity_wrapped/internal/tasks"
)

func main() {
	if len(os.Args) < 4 {
		printExpectedParameters("")
		return
	}

	commentsDirPath := os.Args[1]
	likesDirPath := os.Args[2]
	postsFilePath := os.Args[3]
	validate(commentsDirPath, likesDirPath, postsFilePath)

	outputDirPath := "./collected"
	if len(os.Args) >= 5 && strings.TrimSpace(os.Args[4]) != "" {
		outputDirPath = os.Args[4]
	}

	fmt.Println("Directory with comments: " + commentsDirPath)
	fmt.Println("Directory with likes: " + likesDirPath)
	fmt.Println("File with posts: " + postsFilePath)
	fmt.Println("Output directory: " + outputDirPath)

	commentsFiles, err := os.ReadDir(commentsDirPath)
	if err != nil {
		log.Fatal(err)
	}

	likesFiles, err := os.ReadDir(likesDirPath)
	if err != nil {
		log.Fatal(err)
	}

	postsFile, err := os.Open(postsFilePath)
	defer postsFile.Close()

	context := models.Context{
		UserLikes:    parseUserLikes(likesDirPath, likesFiles),
		UserComments: parseUserComments(commentsDirPath, commentsFiles),
		WallPosts:    parseWallPosts(postsFile),
		OutputDir:    outputDirPath,
	}

	fmt.Println("Start running tasks.")
	executableTasks := []tasks.Task{
		&tasks.MostLikedPostsTask{},
		&tasks.MostLikingUsersTask{},
		&tasks.MostPostingUsersTask{},
		&tasks.MostCommentingUsersTask{},
		&tasks.MostLikedSignedPostsTask{},
		&tasks.MostRepostedPostsTask{},
		&tasks.MostLikedBandsTask{},
		&tasks.MostPostedBandsTask{},
	}

	counter := 0
	for _, task := range executableTasks {
		task.Run(context)
		fmt.Println(fmt.Sprintf("- [%d] %s has finished.", counter, reflect.TypeOf(task).String()))
		counter++
	}

	fmt.Println("Tasks have successfully finished.")
}

func printExpectedParameters(warning string) {
	fmt.Println("Expected parameters:")
	fmt.Println("\t[0] - source directory for comments")
	fmt.Println("\t[1] - source directory for likes")
	fmt.Println("\t[2] - source file for posts")
	fmt.Println("\t[3] - (optional) destination output directory. Default './collected'")
	fmt.Println()
	if strings.TrimSpace(warning) != "" {
		log.Fatal(warning)
	}
}

func validate(commentsDirPath string, likesDirPath string, postsFilePath string) {
	if strings.TrimSpace(commentsDirPath) == "" {
		log.Fatal("Arg [0] (commentsDirPath) must be not empty")
	} else if !exists(commentsDirPath) {
		printExpectedParameters("Arg [0] (commentsDirPath) must exist")
	}

	if strings.TrimSpace(likesDirPath) == "" {
		log.Fatal("Arg [1] (likesDirPath) must be not empty")
	} else if !exists(likesDirPath) {
		printExpectedParameters("Arg [1] (likesDirPath) must exist")
	}

	if strings.TrimSpace(postsFilePath) == "" {
		log.Fatal("Arg [2] (postsFilePath) must be not empty")
	} else if !exists(postsFilePath) {
		printExpectedParameters("Arg [2] (postsFilePath) must exist")
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func parseWallPosts(file *os.File) []models.WallPost {
	var wallPosts []models.WallPost
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var wallPostsLine []models.WallPost
		err := json.Unmarshal(scanner.Bytes(), &wallPostsLine)
		if err != nil {
			log.Fatal(err)
		}

		for _, wallPost := range wallPostsLine {
			wallPosts = append(wallPosts, wallPost)
		}
	}
	return wallPosts
}

func parseUserComments(root string, dirFiles []os.DirEntry) models.UserComments {
	postComments := make(map[int32][]models.Comment)
	for _, file := range dirFiles {
		bytes, err := os.ReadFile(root + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		var comments []models.Comment
		err = json.Unmarshal(bytes, &comments)
		if err != nil {
			log.Fatal(err)
		}

		postId, err := strconv.ParseInt(strings.Replace(file.Name(), ".json", "", -1), 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		postComments[int32(postId)] = comments
	}

	return models.UserComments{postComments}
}

func parseUserLikes(root string, dirFiles []os.DirEntry) models.UserLikes {
	postLike := make(map[int32][]int64)
	for _, file := range dirFiles {
		bytes, err := os.ReadFile(root + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		var users []int64
		err = json.Unmarshal(bytes, &users)
		if err != nil {
			log.Fatal(err)
		}

		postId, err := strconv.ParseInt(strings.Replace(file.Name(), ".json", "", -1), 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		postLike[int32(postId)] = users
	}

	return models.UserLikes{postLike}
}
