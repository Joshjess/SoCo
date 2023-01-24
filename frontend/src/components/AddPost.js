import { Button, Form, Input } from 'antd';
import axios from 'axios';

const { TextArea } = Input;

const headers = {
  'Content-Type': 'text/plain'
};



function sendPost(title, text) {
  axios.post('http://localhost:8080/v1/posts/create', 
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

function submitPost(title, text) {
  console.log('Submit:', title)
  console.log('Submit:', text);
}

const onFinish = (values) => {
  console.log('Success:', values);
  sendPost(values.title, values.text)
}

const AddPost = () => (
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
        // onFinishFailed={onFinishFailed}
    >
      <Form.Item
          name="title"
          rules={[
            {
              required: true,
              message: 'Please input your title!',
            },
          ]}
      >
        <Input placeholder="Title" showCount maxLength={20}/>
      </Form.Item>

      <Form.Item
          name="text"
          rules={[
            {
              required: true,
              message: 'Please input your content!',
            },
          ]}
      >
        <TextArea placeholder="Content" showCount maxLength={100}/>
      </Form.Item>

      <Form.Item>
        <Button type="primary" htmlType="submit">Post</Button>
      </Form.Item>
    </Form>
  </>
);

export default AddPost;
