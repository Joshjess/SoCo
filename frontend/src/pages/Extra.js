import axios from 'axios';
import React, { useState, useEffect } from 'react';
import { Space } from 'antd';
import Posts from '../components/Posts';
import AddPost from '../components/AddPost';


function Extra() {
  // console.log(cookies.token)
  const [posts, setPosts] = useState([]);


  const fetchData = () => { 
    axios.get('http://localhost:8080/v1/posts/')
    .then(response => {
      setPosts(response.data)
    })
    .catch(function (error) {
      console.log(error);
    })
  }
  
  useEffect(() => {
    fetchData()
  }, [])


  return (
    <>
      <h1>Extra</h1>

      <AddPost />

      <div>
        <Space
            direction="vertical"
            size="middle"
            style={{
              display: 'flex',
            }}
        > 
          <Posts posts={posts} />
        </Space>
      </div>
    </>
  );
};

export default Extra;
