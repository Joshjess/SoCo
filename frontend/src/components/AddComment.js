import { Button, Form, Input } from 'antd';
import { CommentOutlined } from '@ant-design/icons';
import axios from 'axios';

const { TextArea } = Input;

const headers = {
  'Content-Type': 'text/plain'
};



function sendPost(title, text) {
  axios.post('http://localhost:8080/v1/posts/create/', 
    {
      Title: title, 
      Text: text,
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

// const onFinish = (values) => {
//   console.log('Success:', values);
//   sendPost(values.title, values.text)
// }

function AddComment(post_id) {
  
  const onFinish = (values) => {
    console.log('Success:', values);
    console.log('Success:', post_id);
    // sendPost(values.title, values.text)
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
