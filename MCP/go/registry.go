package main

import (
	"github.com/article-rewriter-and-article-extractor-api/mcp-server/config"
	"github.com/article-rewriter-and-article-extractor-api/mcp-server/models"
	tools_api "github.com/article-rewriter-and-article-extractor-api/mcp-server/tools/api"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_api.CreateGetinfoTool(cfg),
	}
}
