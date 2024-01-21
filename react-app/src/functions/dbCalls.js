import axios from 'axios';

const API_URL = 'http://10.0.0.251:2081';

// Example function to fetch data from the server
export const getPosts = async () => {
  try {
    const response = await axios.get(`${API_URL}/api/posts`);
    return response.data;
  } catch (error) {
    console.error('Error fetching data:', error);
  }
};

export const createPost = async (postData) => {
    try {
      const response = await axios.post(`${API_URL}/api/post`, postData);
      return response.data;
    } catch (error) {
      console.error('Error fetching data:', error);
    }
};