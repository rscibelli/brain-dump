const express = require("express");
const Post = require("./models/Post");
const app = express();
const PORT = 2081;

app.listen(PORT, () => console.log(`Server running on port ${PORT}`));

const mongoose = require("mongoose");
mongoose.connect("mongodb://localhost:27017/brain-dump-db", {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

app.get("/api/posts", async (req, res) => {
    try {
        const posts = await Post.find();
        res.json(posts);
    } catch (error) {
        console.error(error);
        res.status(500).send("Server Error");
    }
});

app.post('/api/post', async (req, res) => {
    try {
      const newPostData = req.body;
      const newPost = new Post(newPostData);
  
      await newPost.save();
  
      res.status(201).json({ message: 'Post added successfully', newPost });
    } catch (error) {
      console.error('Error adding post:', error);
      res.status(500).json({ error: 'Internal Server Error' });
    }
  });