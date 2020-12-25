package searchengine

import (
	"encoding/json"
	"go-server/middleware"
	"go-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
)

var (
	// searcher is coroutine safe
	searcher           = riot.New()
	repElement         models.Representative
	searchJSONResponse []models.Representative
)

func initEngine() {
	// Add documents to the index
	for _, rep := range middleware.RepMap {
		repString, _ := json.Marshal(middleware.RepMap[rep.GUID])
		// fmt.Println(rep.GUID)
		searcher.Index(rep.GUID, types.DocData{Content: string(repString), Fields: middleware.RepMap[rep.GUID].TotalVotes})
	}

	// Wait for the index to refresh
	searcher.Flush()
}

func init() {
	initEngine()
}

// ExecuteSearch executes the search query from http request
func ExecuteSearch(c *gin.Context) {
	searchTerm := c.Query("search-term")
	sea := searcher.Search(types.SearchReq{
		Text: searchTerm,
		RankOpts: &types.RankOpts{
			OutputOffset: 0,
			MaxOutputs:   100,
		}})

	searchResponse := sea.Docs.(types.ScoredDocs)
	for i := range searchResponse {
		searchElements := searchResponse[i].Content
		json.Unmarshal([]byte(searchElements), &repElement)
		// JSONElement, _ := json.Marshal(searchResponse[i].Content)
		searchJSONResponse = append(searchJSONResponse, repElement)
		// fmt.Println(searchResponse[i].Content)
		// fmt.Println("")
	}
	msg := map[string]interface{}{"Status": "Ok", "search-results": searchJSONResponse}
	c.JSON(http.StatusOK, msg)
	searchJSONResponse = []models.Representative{}
}
