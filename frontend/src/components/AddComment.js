import { Button, Form, Input } from 'antd';
import { CommentOutlined } from '@ant-design/icons';
import axios from 'axios';
import { useCookies } from 'react-cookie';


function postComment (text, id, cookie) {
  console.log(text, id, cookie.token)

  let data = {
    text: text,
    post_id: id,
  };

  let headers = {
    'Content-Type': 'text/plain',
    'Authorization': 'Bearer ' + cookie.token,
  };

  axios.post('http://localhost:8080/v1/comments/create/', data, {headers: headers})
  .catch(function (error) {
    console.log(error);
  })
  .then(function (response) {
    console.log(response);
  })
}

function AddComment(post_id) {
  
  const [cookies, setCookie] = useCookies(['token'])
  
  const onFinish = (values) => {
    console.log('Success:', values);
    console.log('Success:', post_id.post_id);
    postComment(values.comment, post_id.post_id, cookies)
  }

  return (
    <>
      <Form
          labelCol={{
            span: 4,
          }}
          wrapperCol={{
            span: 14,
          }}
          layout="horizontal"
          style={{
            maxWidth: 600,
          }}
          onFinish={onFinish}
      >
        <Form.Item
            name="comment"
            rules={[
              {
                required: true,
                message: 'Please add comment!',
              },
            ]}
        >
          <Input placeholder="comment" showCount maxLength={100}/>
        </Form.Item>

        <Form.Item>
          <Button htmlType="submit">
            <CommentOutlined />Comment
          </Button>
        </Form.Item>
      </Form>
    </>
  );
}

export default AddComment;
