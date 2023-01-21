import axios from 'axios';
import React, { useState, useEffect } from 'react';
import { Button, Card, Space } from 'antd';
import { LikeOutlined, DislikeOutlined } from '@ant-design/icons';
import Posts from '../components/Posts';


const Extra = () => {
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

  // console.log(posts)

  // populate the page with the posts
  return (
    <>
      <h1>Extra</h1>

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
