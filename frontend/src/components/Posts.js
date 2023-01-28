import { Button, Card } from 'antd';
import { LikeOutlined, DislikeOutlined, CommentOutlined } from '@ant-design/icons';
import axios from 'axios';
import AddComment from './AddComment';
import Comments from './Comments';

const headers = {
  'Content-Type': 'text/plain'
};

function requestLike(id, vote) {
  axios.post(`${process.env.REACT_APP_BACKEND_URL}/v1/posts/vote`, {
      post_id: id,
      vote: vote,
    },
    {headers}
  )
  .catch(function (error) {
    console.log(error);
  })
  .then(function (response) {
    console.log(response);
  })
}

function requestComments(id) {
  axios.get(`${process.env.REACT_APP_BACKEND_URL}/v1/comments/post/` + id, 
  {}, {headers} )
  .catch(function (error) {
    console.log(error);
  })
  .then(function (response) {
    console.log(response);
  })
}

function handleLike(id) {
  console.log(id)
  requestLike(id, true)
}

function handleDislike(id) {
  console.log(id)
  requestLike(id, false)
}

function Posts(posts) {
  return (
    <>
      <h1>Posts</h1>
      {posts.posts.map((post, index) => (
        <Card key={index} title={post.title} size="small">
          <p>{post.text}</p>

          <Button key={index} type="primary" onClick={() => handleLike(index)}>
            <LikeOutlined />Like
          </Button>

          <Button key={index} danger onClick={() => handleDislike(index)}>
            <DislikeOutlined />Dislike
          </Button>

          <AddComment post_id={index} />

          {/* <Comments comments={requestComments(index).data}/> */}

        </Card>
      ))}
    </>
  );
};

export default Posts;
