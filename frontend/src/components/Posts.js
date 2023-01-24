import { Button, Card } from 'antd';
import { LikeOutlined, DislikeOutlined, CommentOutlined } from '@ant-design/icons';
import axios from 'axios';
import AddComment from './AddComment';
import Comments from './Comments';
import { useCookies } from 'react-cookie';

// const axios = require('axios');

function Posts(posts) {

  function postLike(id, vote, cookie) {
    console.log(id, vote, cookie.token)
  
    let data = {
      post_id: id,
      vote: vote,
    };
  
    let headers = {
      'Content-Type': 'text/plain',
      'Authorization': 'Bearer ' + cookie.token,
    };
  
    axios.post('http://localhost:8080/v1/posts/vote/', data, {headers: headers})
    .catch(function (error) {
      console.log(error);
    })
    .then(function (response) {
      console.log(response);
    })
  }
  
  function requestComments(id) {
    
    let headers = {
      'Content-Type': 'text/plain',
    };
    
    axios.get('http://localhost:8080/v1/comments/post/' + id, {headers: headers})
    .catch(function (error) {
      console.log(error);
    })
    .then(function (response) {
      console.log(response);
    })
  }
  
  function handleLike(id, cookies) {
    console.log(id)
    postLike(id, true, cookies)
  }
  
  function handleDislike(id, cookie) {
    console.log(id)
    postLike(id, false, cookie)
  }
  
  const [cookies, setCookie] = useCookies(['token'])
  
  console.log(cookies)

  return (
    <>
      <h1>Posts</h1>
      {posts.posts.map((post, index) => (
        <Card key={post.id} title={post.title + " from " + post.username} size="small">
          <p>{post.text}</p>

          <Button key={post.id} type="primary" onClick={() => handleLike(post.id, cookies)}>
            <LikeOutlined />Like
          </Button>

          <Button key={post.id} danger onClick={() => handleDislike(post.id, cookies)}>
            <DislikeOutlined />Dislike
          </Button>

          <p>{post.upvote} likes</p>
          <p>{post.downvote} dislikes</p>

          <AddComment post_id={post.id} />

          {/* <Comments comments={requestComments(post.id).data}/> */}

        </Card>
      ))}
    </>
  );
};

export default Posts;
