import { Button, Card } from 'antd';
import { LikeOutlined, DislikeOutlined } from '@ant-design/icons';
import axios from 'axios';


function Posts(posts) {
  
  function handleLike() {
    axios.post('http://localhost:8080/v1/users/signup', {
      post_id: values.username,
      vote: values.password,
    },
    {headers}
    )
    .then(function (response) {
      console.log(response);
    })


  }

  function handleDislike() {
    console.log('dislike')
  }
  
  return (
    <>
      <h1>Posts</h1>
      {posts.posts.map((post) => (
        <Card key={post.id} title={post.title} size="small">
          <p>{post.text}</p>

          <Button type="primary" onClick={handleLike}>
            <LikeOutlined />Like
          </Button>
          <Button danger onClick={handleDislike}>
            <DislikeOutlined />Dislike
          </Button>

        </Card>
      ))}
    </>
  );
};

export default Posts;
