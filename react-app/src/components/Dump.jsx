import React, { useEffect, useState } from 'react';
import { getPosts, createPost } from '../functions/dbCalls';
import 'bootstrap/dist/css/bootstrap.min.css';
import Card from './Card';


const Dump = () => {
  const [title, setTitle] = useState('');
  const [newPost, setNewPost] = useState('');
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    async function fetchData() {
      let allposts = await getPosts()
      console.log("posts " + allposts)
      setPosts(allposts)
    }
    fetchData()
  }, []);

  const handleTitleChange = (e) => {
    setTitle(e.target.value)
  }

  const handlePostChange = (e) => {
    setNewPost(e.target.value);
  };

  const handlePostSubmit = () => {
    if (newPost.trim() !== '') {
      let newCompletePost = {
        "name": "Robert Scibelli",
        "date": getDate(),
        "title": title,
        "post": newPost
      }
      setPosts([...posts, newCompletePost]);
      createPost(newCompletePost)
      setNewPost('');
      setTitle('');
    }
  };

  return (
    <div className="blog p-2 mt-2">
      <div className='mx-5'>
        <h3 className='text-start'>Create a Dump</h3>
        <div>
          <textarea id="postTitle" className="form-control bg-light" placeholder="Title" rows="1" onChange={handleTitleChange} value={title}></textarea>
        </div>
        <div>
          <textarea id="postContent" className="form-control bg-light" placeholder="Content" rows="3" onChange={handlePostChange} value={newPost}></textarea>
        </div>
        <div className='text-end'>
          <button className="btn btn-primary mt-1" onClick={handlePostSubmit}>Submit</button>
        </div>
      </div>

      <hr className='w-75 mx-auto' />

      <div className="px-5 mx-5">
        <h4 className="mb-4">Previous Dumps</h4>
        <div className="row row-cols-1">
          {posts.map((post, index) => (
            <div key={index} className="col mb-3">
              <Card post={post} />
            </div>
          ))}
        </div>
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