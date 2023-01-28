import { Button, Card } from 'antd';
import { LikeOutlined, DislikeOutlined, CommentOutlined } from '@ant-design/icons';
import axios from 'axios';

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

function handleLike(id) {
  console.log(id)
  requestLike(id, true)
}

function handleDislike(id) {
  console.log(id)
  requestLike(id, false)
}

function Comments(comments) {
  return (
    <>
      <h4>Comments</h4>
      {comments.comments.map((comment, index) => {
        if (comment.message === "No comments found") {
          return <p>{comment.message}</p>
        }
        else (
        <Card key={index} title={comment.title} size="small">
          <p>{comment.text}</p>

          <Button key={index} type="primary" onClick={() => handleLike(index)}>
            <LikeOutlined />Like
          </Button>

          <Button key={index} danger onClick={() =>handleDislike(index)}>
            <DislikeOutlined />Dislike
          </Button>

        </Card>
      )})}
    </>
  );
};

export default Comments;
