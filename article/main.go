package main

// gRPCサーバーの動作確認用
func main() {
	//c, _ := client.NewClient("localhost:8080")

	// 記事をCREATE
	/*
		input := &pb.ArticleInput{
			Author:  "gopher",
			Title:   "gRPC",
			Content: "gRPC is so cool!",
		}
		res, err := c.Service.CreateArticle(context.Background(), &pb.CreateArticleRequest{ArticleInput: input})
		if err != nil {
			log.Fatalf("Failed to CreateArticle: %v\n", err)
		}
		fmt.Printf("CreateArticle Response: %v\n", res)
	*/

	// 記事をREAD
	/*
		var id int64 = 1
		res, err := c.Service.ReadArticle(context.Background(), &pb.ReadArticleRequest{Id: id})
		if err != nil {
			log.Fatalf("Failed to ReadArticle: %v\n", err)
		}
		fmt.Printf("ReadArticle Response: %v\n", res)
	*/

	// 記事をUPDATE
	/*
		var id int64 = 1
		input := &pb.ArticleInput{
			Author:  "GraphQL master",
			Title:   "GraphQL",
			Content: "GraphQL is very smart!",
		}
		res, err := c.Service.UpdateArticle(context.Background(), &pb.UpdateArticleRequest{Id: id, ArticleInput: input})
		if err != nil {
			log.Fatalf("Failed to UpdateArticle: %v\n", err)
		}
		fmt.Printf("UpdateArticle Response: %v\n", res)
	*/

	// 記事をDELETE
	/*
		var id int64 = 1
		res, err := c.Service.DeleteArticle(context.Background(), &pb.DeleteArticleRequest{Id: id})
		if err != nil {
			log.Fatalf("Failed to UpdateArticle: %v\n", err)
		}
		fmt.Printf("The article has been deleted (%v)\n", res)
	*/

	// 記事を全取得
	/*
		stream, err := c.Service.ListArticle(context.Background(), &pb.ListArticleRequest{})
		if err != nil {
			log.Fatalf("Failed to ListArticle: %v\n", err)
		}
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to Server Streaming: %v\n", err)
			}
			fmt.Println(res)
		}
	*/
}
