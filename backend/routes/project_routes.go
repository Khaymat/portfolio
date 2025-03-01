package controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "portfolio/config"
    "portfolio/models"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

var projectCollection *mongo.Collection = config.DB.Collection("projects")

func GetProjects(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var projects []models.Project
    cursor, err := projectCollection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var project models.Project
        cursor.Decode(&project)
        projects = append(projects, project)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(projects)
}
