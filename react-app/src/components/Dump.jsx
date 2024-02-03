import React, { useEffect, useState } from 'react';
import { getPosts, createPost } from '../functions/dbCalls';

const Dump = () => {
  const [title, setTitle] = useState('');
  const [newPost, setNewPost] = useState('');
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    let allposts = getPosts()
    console.log("posts " + allposts)
    setPosts(allposts)
  }, []);

  const handleTitleChange = (e) => {
    setTitle(e.target.value)
  }

  const handlePostChange = (e) => {
    setNewPost(e.target.value);
  };

  const handlePostSubmit = () => {
    if (newPost.trim() !== '') {
      setPosts([...posts, newPost]);
      let post = {
        "name": "Robert Scibelli",
        "date": getDate(),
        "title": title,
        "post": newPost
      }
      createPost(post)
      setNewPost('');
      setTitle('');
    }
  };

  return (
    <div className="blog p-2 mt-2">
      <div className='mx-5'>
        <h3 className='text-start'>Create a Post</h3>
        <div>
          <textarea id="postTitle" className="form-control bg-light" placeholder="Title" rows="1" onChange={handleTitleChange}></textarea>
        </div>
        <div>
          <textarea id="postContent" className="form-control bg-light" placeholder="Content" rows="3" onChange={handlePostChange}></textarea>
        </div>
        <div className='text-end'>
          <button className="btn btn-primary mt-1" onClick={handlePostSubmit}>Submit</button>
        </div>
      </div>

      <hr className='w-75 mx-auto' />

      <div className="previous-posts">
        <h4>Previous Blog Posts</h4>
        {posts.map((post, index) => (
          <div key={index} className="blog-post">
            <p className="blog-post-content">{post.post}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

const getDate = () => {
  const timeElapsed = Date.now();
  const today = new Date(timeElapsed);
  return today.toLocaleDateString();
}

export default Dump;