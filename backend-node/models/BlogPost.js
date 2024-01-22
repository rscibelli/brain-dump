import { Schema, model } from "mongoose";

const postSchema = new Schema({
  name: String,
  date: String,
  post: String,
});

export default model("BlogPost", postSchema, "posts");

// db.posts.insert({name: "Robert Scibelli", date: "January 21st 2024", post: "Hello world"})