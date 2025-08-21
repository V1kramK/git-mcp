package main

import (
	"github.com/git-mcp-api/mcp-server/config"
	"github.com/git-mcp-api/mcp-server/models"
	tools_repokey_timestamp "github.com/git-mcp-api/mcp-server/tools/repokey_timestamp"
	tools_counts "github.com/git-mcp-api/mcp-server/tools/counts"
	tools_file_bucketpath "github.com/git-mcp-api/mcp-server/tools/file_bucketpath"
	tools_key "github.com/git-mcp-api/mcp-server/tools/key"
	tools_owner "github.com/git-mcp-api/mcp-server/tools/owner"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_repokey_timestamp.CreateGet_repokey_timestampTool(cfg),
		tools_counts.CreatePut_countsTool(cfg),
		tools_file_bucketpath.CreatePut_file_bucketpathTool(cfg),
		tools_key.CreateGet_keyTool(cfg),
		tools_key.CreatePut_keyTool(cfg),
		tools_owner.CreateGet_owner_repo_filenameTool(cfg),
	}
}
