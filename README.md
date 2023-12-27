## VK Community Wrapped

Bunch of script for data calculation. 

Source of calculation expected to be raw json from [VK API](https://dev.vk.com/en/method) (wall.get, wall.getComments etc.)

## Usage

1. `./bash/fetch_posts.sh ./posts.txt`
2. `cat ./posts.txt | jq ".[].id" >> postsIds.txt`
3. `./bash/fetch_comments.sh ./postsIds.txt`
4. `./bash/fetch_likes.sh ./postIds.txt`
5. `go run main.go ~/comments ~/likes ~/posts.txt`
```
Directory with comments: ../wrapped2023/comments
Directory with likes: ../wrapped2023/likes
File with posts: ../wrapped2023/posts.txt
Output directory: ./collected
Start running tasks.
- [0] *tasks.MostLikedPostsTask has finished.
- [1] *tasks.MostLikingUsersTask has finished.
- [2] *tasks.MostPostingUsersTask has finished.
- [3] *tasks.MostCommentingUsersTask has finished.
- [4] *tasks.MostLikedSignedPostsTask has finished.
- [5] *tasks.MostRepostedPostsTask has finished.
- [6] *tasks.MostLikedBandsTask has finished.
- [7] *tasks.MostPostedBandsTask has finished.
Tasks have successfully finished.
```