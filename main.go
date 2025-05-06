package main

import (
	"context"
	_ "embed"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type ManifestCode struct {
	Description string `yaml:"description"`
	File        string `yaml:"file"`
}
type ManifestDocument struct {
	Name       string                    `yaml:"name"`
	MDFileName string                    `yaml:"md_file_name"`
	Code       map[string][]ManifestCode `yaml:"code"`
}
type ManifestRoot struct {
	Documents []ManifestDocument `yaml:"documents"`
}

func main() {
	loader := NewLoader()
	m := loader.Manifest

	s := server.NewMCPServer(
		"ch32v003-development-guide-book",
		"0.0.1",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	itemNameList := []string{}
	for _, d := range m.Documents {
		itemNameList = append(itemNameList, d.Name)
	}

	itemNameListString := strings.Join(itemNameList, ", ")

	readGudeBookTool := mcp.NewTool("read_ch32v003_guide_book",
		mcp.WithDescription("CH32V003の実装方法を説明したドキュメントを返します。CH32V003の開発の際には参照してください。CH32V003の開発フレームワークには、WCH SDKとch32funがあります。利用する場合にはしていしてください。作りたいものに応じて、"+itemNameListString+"を指定してください。"),
		mcp.WithString("framework",
			mcp.Required(),
			mcp.Description("フレームワーク(WCH SDK, ch32fun)"),
			mcp.Enum("WCH SDK", "ch32fun"),
		),
		mcp.WithString("item",
			mcp.Required(),
			mcp.Description("内容("+itemNameListString+")"),
			mcp.Enum(itemNameList...),
		),
	)

	s.AddTool(readGudeBookTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

		frameworkName := request.Params.Arguments["framework"].(string)
		frameworkKeyName := frameworkName
		itemName := request.Params.Arguments["item"].(string)

		mainMD, err := loader.GetMainMD(frameworkName)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Invalid framework: %s", frameworkName)), nil
		}

		doc, err := loader.SelectManifestItem(itemName)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Invalid item: %s", itemName)), nil
		}

		itemMD, err := loader.GetItemMD(frameworkKeyName, itemName)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("item document not found: %s", itemName)), nil
		}

		b := strings.Builder{}
		b.WriteString(fmt.Sprintf(`
%s

%s

`, mainMD, itemMD))

		codeDocuments, err := loader.GetCodeDocument(frameworkKeyName, itemName)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("code documents not found for item: %s", itemName)), nil
		}

		for i, code := range codeDocuments {
			b.WriteString(fmt.Sprintf(`
# サンプルコード %s

`, doc.Code[frameworkKeyName][i].Description))
			b.WriteString("```\r\n")
			b.WriteString(code)
			b.WriteString("\r\n```\r\n")
		}

		return mcp.NewToolResultText(b.String()), nil
	})

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
