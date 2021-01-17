package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"go-server/models"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// collection object/instance
var (
	collection  *mongo.Collection
	RepMap      = make(map[string]models.Representative)
	userReps    = make(map[string][]string)
	tempRepList []models.Representative
)

// create connection with mongo db
func init() {
	loadTheEnv()
	createDBInstance()
	loadRepDB()
}

func loadTheEnv() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func createDBInstance() {
	// DB connection string
	connectionString := os.Getenv("DB_URI")

	// Database Name
	dbName := os.Getenv("DB_NAME")

	// Collection name
	collName := os.Getenv("DB_COLLECTION_NAME")

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

// New Functions

func loadRepDB() map[string]models.Representative {

	// create userReps map through csv file
	records := []*models.UserRepMap{}
	userFileData, err := os.Open("./test_data/user_favorite_reps.csv")
	if err != nil {
		panic(err)
	}
	defer userFileData.Close()
	if err := gocsv.UnmarshalFile(userFileData, &records); err != nil {
		panic(err)
	}
	for _, row := range records {
		// fmt.Println(rep.LastName)
		userReps[row.UserGUID] = append(userReps[row.UserGUID], row.RepGUID)
	}

	reps := []*models.Representative{}

	in, err := os.Open("./test_data/house_members.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()
	if err := gocsv.UnmarshalFile(in, &reps); err != nil {
		panic(err)
	}
	for _, rep := range reps {
		// fmt.Println(rep.LastName)
		RepMap[rep.GUID] = addRepToMap(rep)
	}

	in, err = os.Open("./test_data/senate_members.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()
	if err := gocsv.UnmarshalFile(in, &reps); err != nil {
		panic(err)
	}
	for _, rep := range reps {
		// fmt.Println(rep.LastName)
		RepMap[rep.GUID] = addRepToMap(rep)
	}

	return RepMap
}

func addRepToMap(rep *models.Representative) models.Representative {
	return models.Representative{
		GUID:                  rep.GUID,
		Office:                rep.Office,
		Name:                  rep.Name,
		LastName:              rep.LastName,
		Location:              rep.Location,
		Division:              rep.Division,
		GovWebsite:            rep.GovWebsite,
		PhotoURL:              rep.PhotoURL,
		Twitter:               rep.Twitter,
		TotalVotes:            rep.TotalVotes,
		MissedVotes:           rep.MissedVotes,
		PresentVotes:          rep.PresentVotes,
		PercentMissedVotes:    math.Round(10000*float64(rep.MissedVotes)/float64(rep.TotalVotes)) / 100,
		PercentPresentVotes:   math.Round(10000*float64(rep.PresentVotes)/float64(rep.TotalVotes)) / 100,
		PercentVotesWithParty: rep.PercentVotesWithParty,
	}
}

// LocalRepsHandler loads user information and is called on website homepage
func LocalRepsHandler(c *gin.Context) {
	//userGUID, _ := c.GetQuery("user_guid")
	userGUID := "55ee03f2dcd8c8e46b91cbb2e70d9e"
	//userGUID := "1234"
	c.Header("Content-Type", "application/json")
	var tempUserRepList []models.Representative
	for _, j := range userReps[userGUID] {
		// fmt.Println("j", j)
		// fmt.Println("Rep: ", repMap[j].Name)
		tempUserRepList = append(tempUserRepList, RepMap[j])
	}
	msg := map[string]interface{}{"Status": "Ok", "user_guid": userGUID, "users_rep_list": tempUserRepList}
	c.JSON(http.StatusOK, msg)
}

// EditLocalRep adds or removes a local rep in a user's feed
func EditLocalRep(c *gin.Context) {
	userGUID, _ := c.GetQuery("user_guid")
	repGUID, _ := c.GetQuery("rep_guid")
	editTask, _ := c.GetQuery("editTask")
	c.Header("Content-Type", "application/json")
	targetRepIndex := -1
	if editTask == "add" {
		// TODO: create map of maps
		userReps[userGUID] = append(userReps[userGUID], repGUID)
	} else if editTask == "remove" {
		tempUserRepList := userReps[userGUID]
		for i, value := range tempUserRepList {
			if value == repGUID {
				targetRepIndex = i
			}
		}
		if targetRepIndex != -1 {
			userReps[userGUID] = append(tempUserRepList[:targetRepIndex], tempUserRepList[targetRepIndex+1:]...)
		}
	} else {
		fmt.Println("edit Rep: provided invalid option")
		// log.Info("edit Rep: provided invalid option")
	}

	userRepUpdate := models.UserRepUpdate{
		UserGUID: userGUID,
		RepGUID:  repGUID,
		Action:   editTask,
	}

	userRepUpdateResponse, _ := json.Marshal(userRepUpdate)

	fmt.Println(string(userRepUpdateResponse))

	// if enableKafka {
	// 	err := writer.WriteMessages(context.Background(), kafka.Message{
	// 		//Key: []byte(repGUID),
	// 		Value: []byte(userRepUpdateResponse),
	// 	})
	// 	if err != nil {
	// 		panic("could not write message " + err.Error())
	// 	}
	// }

	msg := map[string]interface{}{"Status": "Ok", "user_guid": userGUID, "users_rep_list": userReps[userGUID]}
	c.JSON(http.StatusOK, msg)
}

// GetTopReps pulls in the trending representatives for a user to select
func GetTopReps(c *gin.Context) {
	reps := []*models.Representative{}

	in, err := os.Open("./test_data/top_house_reps.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()
	if err := gocsv.UnmarshalFile(in, &reps); err != nil {
		panic(err)
	}
	for _, rep := range reps {
		tempRepList = append(tempRepList, addRepToMap(rep))
	}

	in, err = os.Open("./test_data/top_senate_reps.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()
	if err := gocsv.UnmarshalFile(in, &reps); err != nil {
		panic(err)
	}
	for _, rep := range reps {
		tempRepList = append(tempRepList, addRepToMap(rep))
	}

	msg := map[string]interface{}{"Status": "Ok", "users_rep_list": tempRepList}
	c.JSON(http.StatusOK, msg)
}

// End of New Functions

// GetAllTask get all the task route
func GetAllTask(c *gin.Context) {
	payload := getAllTask()
	c.JSON(http.StatusOK, payload)
}

// CreateTask create task route
func CreateTask(c *gin.Context) {
	var task models.ToDoList
	_ = json.NewDecoder(c.Request.Body).Decode(&task)
	// fmt.Println(task, r.Body)
	insertOneTask(task)
	c.JSON(http.StatusOK, task)
}

// TaskComplete update task route
func TaskComplete(c *gin.Context) {
	task := c.Param("id")
	taskComplete(task)
	c.JSON(http.StatusOK, task)
	// params := mux.Vars(r)
	// taskComplete(params["id"])
	// json.NewEncoder(w).Encode(params["id"])
}

// UndoTask undo the complete task route
func UndoTask(c *gin.Context) {
	task := c.Param("id")
	fmt.Println(task)
	undoTask(task)
	c.JSON(http.StatusOK, task)
}

// DeleteTask delete one task route
func DeleteTask(c *gin.Context) {
	fmt.Println("deleteTask")
	task := c.Param("id")
	fmt.Println("task_id: ", task)
	deleteOneTask(task)
	c.JSON(http.StatusOK, task)
	// json.NewEncoder(w).Encode("Task not found")

}

// DeleteAllTask delete all tasks route
func DeleteAllTask(c *gin.Context) {
	count := deleteAllTask()
	c.JSON(http.StatusOK, count)
	// json.NewEncoder(w).Encode("Task not found")

}

// get all task from the DB and return it
func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// Insert one task in the DB
func insertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

// task complete method, update task's status to true
func taskComplete(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// task undo method, update task's status to false
func undoTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete one task from the DB, delete by ID
func deleteOneTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

// delete all the tasks from the DB
func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}
