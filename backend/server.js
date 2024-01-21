import express from "express";
import BlogPost from "../models/BlogPost.js";
import { connect } from "mongoose";
const app = express();
const PORT = 2081;

app.listen(PORT, () => console.log(`Server running on port ${PORT}`));

app.use(function(req, res, next) {
  res.setHeader('Access-Control-Allow-Origin', '*');
  res.setHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE');
  res.setHeader('Access-Control-Allow-Headers', 'Content-Type');
  res.setHeader('Access-Control-Allow-Credentials', true);
  next();
});

connect("mongodb://localhost:27017/brain-dump-db", {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

app.get("/api/posts", async (req, res) => {
    try {
        const posts = await BlogPost.find();
        res.json(posts);
    } catch (error) {
        console.error(error);
        res.status(500).send("Server Error");
    }
});

app.post('/api/post', async (req, res) => {
    try {
      const newPostData = req.body;
      const newPost = new BlogPost(newPostData);
  
      await newPost.save();
  
      res.status(201).json({ message: 'Post added successfully', newPost });
    } catch (error) {
      console.error('Error adding post:', error);
      res.status(500).json({ error: 'Internal Server Error' });
    }
  });