import axios from "axios";
import React, { useState, useEffect } from "react";
import { Button, Space } from 'antd';
import { LikeOutlined, DislikeOutlined } from '@ant-design/icons';
import { useCookies } from 'react-cookie';


function postCommentLike(id, vote, cookie) {
  // console.log(id, vote, cookie.token)

  let data = {
    comment_id: id,
    vote: vote,
  };

  let headers = {
    'Content-Type': 'text/plain',
    'Authorization': 'Bearer ' + cookie.token,
  };

  axios.post(`${process.env.REACT_APP_BACKEND_URL}/v1/comments/vote`, data, {headers: headers})
  .catch(function (error) {
    console.log(error);
  })
  .then(function (response) {
    console.log(response);
  })
}

function handleLike(id, cookies) {
  // console.log(id)
  postCommentLike(id, true, cookies)
}

function handleDislike(id, cookie) {
  // console.log(id)
  postCommentLike(id, false, cookie)
}

function Comments(post_id) {

  const [comments, setComments] = useState([]);
  const [cookies, setCookie] = useCookies(['token']);

  const fetchData = () => { 
    axios.get(`${process.env.REACT_APP_BACKEND_URL}/v1/comments/post/` + post_id.post_id)
    .then(response => {
      setComments(response.data)
    })
    .catch(function (error) {
      console.log(error);
    })
  }
  
  useEffect(() => {
    fetchData()
  }, [])

  if (comments.length > 0) {
    return (
      <>
        <h4>Comments</h4>
        {comments.map((comment) => (
          <>
            <p>{comment.text} by {comment.username}</p>

            <Button key={comment.id} type="primary" onClick={() => handleLike(comment.id, cookies)}>
              <LikeOutlined />Like
            </Button>

            <Button key={comment.id} danger onClick={() => handleDislike(comment.id, cookies)}>
              <DislikeOutlined />Dislike
            </Button>

            <p>{comment.likes} likes</p>
            <p>{comment.dislikes} dislikes</p>

            <Space />
          </>
        ))}
      </>
    );
  } else {   
    return (
      <>
        <h4>Comments</h4>
        <p>Looks quite empty</p>
      </>
    );
  }
}

export default Comments;
