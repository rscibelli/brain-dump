const mongoose = require("mongoose");

const postSchema = new mongoose.Schema({
  name: String,
  date: String,
  post: String,
});

module.exports = mongoose.model("Post", postSchema);

// db.posts.insert({name: "Robert Scibelli", date: "January 21st 2024", post: "Hello world"})