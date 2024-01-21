import { Schema, model } from "mongoose";

const postSchema = new Schema({
  name: String,
  date: String,
  post: String,
});

export default model("Post", postSchema);

// db.posts.insert({name: "Robert Scibelli", date: "January 21st 2024", post: "Hello world"})