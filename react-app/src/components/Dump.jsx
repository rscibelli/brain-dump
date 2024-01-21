import React, { useEffect, useState } from 'react';
import { getPosts, createPost } from '../functions/dbCalls';

const Dump = () => {
  const [newPost, setNewPost] = useState('');
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    let allposts = getPosts()
    console.log("posts " + allposts)
    setPosts(allposts)
  }, []);

  const handlePostChange = (e) => {
    setNewPost(e.target.value);
  };

  const handlePostSubmit = () => {
    if (newPost.trim() !== '') {
      setPosts([...posts, newPost]);
      createPost(newPost)
      setNewPost('');
    }
  };

  return (
    <div className="blog p-2 mt-2">
      <div className='mx-5'>
        <h3 className='text-start'>What's on your mind?</h3>
        <textarea className="form-control bg-light" rows="3" onChange={handlePostChange}></textarea>
        <div className='text-end'>
            <button className="btn btn-primary mt-1" onClick={handlePostSubmit}>Submit</button>
        </div>
      </div>

      <hr className='w-75 mx-auto' />

      <div className="previous-posts">
        <h4>Previous Blog Posts</h4>
        {posts.map((post, index) => (
          <div key={index} className="blog-post">
            <p className="blog-post-content">{post}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Dump;